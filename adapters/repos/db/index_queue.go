//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package db

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/priorityqueue"
)

// IndexQueue is a persistent queue of vectors to index.
// It batches vectors together before sending them to the indexing worker.
// It persists the vectors on disk to ensure they are not lost in case of a crash.
// It is safe to use concurrently.
type IndexQueue struct {
	Index BatchIndexer

	logger logrus.FieldLogger

	maxQueueSize  int
	maxStaleTime  time.Duration
	retryInterval time.Duration

	// processCh is the channel used to send vectors to the indexing worker.
	processCh chan []vectorDescriptor

	// indexCh is the channel used to send vectors to the indexing worker.
	indexCh chan *chunk

	// if closed, prevents new vectors from being added to the queue.
	closed chan struct{}

	// tracks the workers
	wg sync.WaitGroup

	// tracks the background workers
	workerWg sync.WaitGroup

	// queue of not-yet-indexed vectors
	queue *vectorQueue

	// keeps track of the last time the queue was indexed
	staleTm *time.Timer

	pqMaxPool *pqMaxPool
}

type vectorDescriptor struct {
	id     uint64
	vector []float32
}

type IndexQueueOptions struct {
	// MaxQueueSize is the maximum number of vectors to queue
	// before sending them to the indexing worker.
	// It must be a multiple of BatchSize.
	MaxQueueSize int

	// BatchSize is the number of vectors to batch together
	// before sending them to the indexing worker.
	BatchSize int

	// MaxStaleTime is the maximum time to wait before sending
	// the pending vectors to the indexing worker, regardless
	// of the queue size.
	MaxStaleTime time.Duration

	// Logger is the logger used by the queue.
	Logger logrus.FieldLogger

	// RetryInterval is the interval between retries when
	// indexing fails.
	RetryInterval time.Duration

	// IndexWorkerCount is the number of workers used to index
	// the vectors.
	IndexWorkerCount int
}

type BatchIndexer interface {
	AddBatch(id []uint64, vector [][]float32) error
	SearchByVector(vector []float32, k int, allowList helpers.AllowList) ([]uint64, []float32, error)
	DistanceBetweenVectors(x, y []float32) (float32, bool, error)
}

func NewIndexQueue(
	index BatchIndexer,
	opts IndexQueueOptions,
) (*IndexQueue, error) {
	if opts.Logger == nil {
		opts.Logger = logrus.New()
	}

	if opts.MaxQueueSize == 0 {
		opts.MaxQueueSize = 1000
	}

	if opts.BatchSize == 0 {
		opts.BatchSize = 100
	}

	if opts.MaxStaleTime == 0 {
		opts.MaxStaleTime = 10 * time.Second
	}

	if opts.MaxQueueSize%opts.BatchSize != 0 {
		return nil, errors.New("maxQueueSize must be a multiple of batchSize")
	}

	if opts.RetryInterval == 0 {
		opts.RetryInterval = 1 * time.Second
	}

	if opts.IndexWorkerCount == 0 {
		// TODO: use the number of CPUs
		opts.IndexWorkerCount = 12
	}

	q := IndexQueue{
		Index:         index,
		maxQueueSize:  opts.MaxQueueSize,
		maxStaleTime:  opts.MaxStaleTime,
		retryInterval: opts.RetryInterval,
		queue:         newVectorQueue(opts.MaxQueueSize, opts.BatchSize),
		processCh:     make(chan []vectorDescriptor),
		indexCh:       make(chan *chunk),
		closed:        make(chan struct{}),
		logger:        opts.Logger.WithField("component", "index_queue"),
		pqMaxPool:     newPqMaxPool(0),
	}

	q.workerWg.Add(1)
	go func() {
		defer q.workerWg.Done()
		defer close(q.indexCh)

		q.processor()
	}()

	for i := 0; i < opts.IndexWorkerCount; i++ {
		q.workerWg.Add(1)
		go func() {
			defer q.workerWg.Done()

			q.indexer()
		}()
	}

	return &q, nil
}

// Close waits till the queue has ingested and persisted all pending vectors.
func (q *IndexQueue) Close() error {
	// check if the queue is closed
	select {
	case <-q.closed:
		return errors.New("index queue closed")
	default:
	}

	// prevent new jobs from being added
	close(q.closed)

	// wait for in-flight pushes to finish
	q.wg.Wait()

	// close the workers in cascade
	close(q.processCh)

	q.workerWg.Wait()

	return nil
}

// Push adds a vector to the persistent indexing queue.
// It waits until the vector is successfully persisted to the
// on-disk queue or sent to the indexing worker.
func (q *IndexQueue) Push(ctx context.Context, vectors ...vectorDescriptor) error {
	// check if the queue is closed
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-q.closed:
		return errors.New("index queue closed")
	default:
	}

	// count the number of in-flight pushes
	q.wg.Add(1)
	defer q.wg.Done()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case q.processCh <- vectors:
		return nil
	}
}

// Search defer to the index and brute force the unindexed data
func (q *IndexQueue) SearchByVector(vector []float32, k int, allowList helpers.AllowList) ([]uint64, []float32, error) {
	indexedResults, distances, err := q.Index.SearchByVector(vector, k, allowList)
	if err != nil {
		return nil, nil, err
	}
	results := q.pqMaxPool.GetMax(k)
	defer q.pqMaxPool.Put(results)
	for i := range indexedResults {
		results.Insert(indexedResults[i], distances[i])
	}

	snapshot := q.queue.AppendSnapshot(nil)
	ids := make([]uint64, len(snapshot))
	vectors := make([][]float32, len(snapshot))
	for i, v := range snapshot {
		ids[i] = v.id
		vectors[i] = v.vector
	}

	err = q.bruteForce(vector, k, ids, vectors, results, allowList)
	if err != nil {
		return nil, nil, err
	}

	if cap(ids) >= k {
		ids = ids[:k]
		distances = distances[:k]
	} else {
		ids = make([]uint64, k)
		distances = make([]float32, k)
	}

	for i := k - 1; i >= 0; i-- {
		element := results.Pop()
		ids[i] = element.ID
		distances[i] = element.Dist
	}
	return ids, distances, nil
}

func (q *IndexQueue) bruteForce(vector []float32, k int, ids []uint64, vectors [][]float32, results *priorityqueue.Queue, allowList helpers.AllowList) error {
	// actual brute force. Consider moving to a separate
	for i := range vectors {
		// skip filtered data
		if allowList != nil && allowList.Contains(ids[i]) {
			continue
		}

		dist, _, err := q.Index.DistanceBetweenVectors(vector, vectors[i])
		if err != nil {
			return err
		}

		if results.Len() < k || dist < results.Top().Dist {
			results.Insert(ids[i], dist)
			for results.Len() > k {
				results.Pop()
			}
		}
	}
	return nil
}

// This is the processor worker. Its job is to batch jobs together before
// sending them to the index.
// While the queue is not full or stale, it persists the jobs on disk.
// Once the queue is full or stale, it sends the jobs to the indexing worker.
// It batches concurrent jobs to reduce the number of disk writes.
func (q *IndexQueue) processor() {
	q.staleTm = time.NewTimer(q.maxStaleTime)

	for {
		select {
		case batch, ok := <-q.processCh:
			if !ok {
				// the channel is closed, abort.
				// drain the timer
				if !q.staleTm.Stop() {
					<-q.staleTm.C
				}
				return
			}

			var n int
			var err error

			// q.logger.WithField("batch size", len(batch)).Info("adding vectors to queue")
			for {
				n = q.queue.Add(batch...)
				if n == len(batch) {
					break
				}

				// q.logger.WithField("batch size", len(batch)).Info("queue full, indexing vectors")

				batch = batch[n:]

				err = q.index(true)
				if err != nil {
					q.logger.WithError(err).Error("failed to index vectors")
					return
				}
			}

			err = q.index(true)
			if err != nil {
				q.logger.WithError(err).Error("failed to index vectors")
				return
			}
		case <-q.staleTm.C:
			err := q.index(false)
			if err != nil {
				q.logger.WithError(err).Error("failed to index vectors")
				return
			}

			// reset the timer
			q.staleTm.Reset(q.maxStaleTime)
		case <-q.closed:
			// drain the timer
			if !q.staleTm.Stop() {
				<-q.staleTm.C
			}
			// if the queue is closed, do nothing.
			return
		}
	}
}

func (q *IndexQueue) index(fullOnly bool) error {
	for q.queue.HasChunksReady(fullOnly) {
		c := q.queue.BorrowNextChunk(fullOnly)
		// q.logger.WithField("index", c.id).Info("indexing chunk")
		select {
		case <-q.closed:
			return errors.New("index queue closed")
		case q.indexCh <- c:
		}
	}

	return nil
}

func (q *IndexQueue) indexer() {
	var ids []uint64
	var vectors [][]float32

	for b := range q.indexCh {
		for i := range b.chunk[:b.cursor] {
			ids = append(ids, b.chunk[i].id)
			vectors = append(vectors, b.chunk[i].vector)
		}

		q.indexVectors(ids, vectors)

		q.queue.ReleaseChunk(b)

		ids = ids[:0]
		vectors = vectors[:0]
	}
}

func (q *IndexQueue) indexVectors(ids []uint64, vectors [][]float32) error {
	for {
		err := q.Index.AddBatch(ids, vectors)
		if err == nil {
			break
		}

		q.logger.WithError(err).Infof("failed to index vectors, retrying in %s", q.retryInterval.String())

		t := time.NewTimer(q.retryInterval)
		select {
		case <-q.closed:
			// drain the timer
			if !t.Stop() {
				<-t.C
			}
			return errors.New("index queue closed")
		case <-t.C:
		}
	}

	return nil
}

type pqMaxPool struct {
	pool *sync.Pool
}

func newPqMaxPool(defaultCap int) *pqMaxPool {
	return &pqMaxPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return priorityqueue.NewMax(defaultCap)
			},
		},
	}
}

func (pqh *pqMaxPool) GetMax(capacity int) *priorityqueue.Queue {
	pq := pqh.pool.Get().(*priorityqueue.Queue)
	if pq.Cap() < capacity {
		pq.ResetCap(capacity)
	} else {
		pq.Reset()
	}

	return pq
}

func (pqh *pqMaxPool) Put(pq *priorityqueue.Queue) {
	pqh.pool.Put(pq)
}

type vectorQueue struct {
	sync.RWMutex

	vectors   []vectorDescriptor
	maxSize   int
	batchSize int

	curBatch *chunk
	chunks   []chunk
}

func newVectorQueue(maxSize, batchSize int) *vectorQueue {
	if maxSize%batchSize != 0 {
		panic("maxSize must be a multiple of batchSize")
	}

	q := vectorQueue{
		maxSize:   maxSize,
		batchSize: batchSize,
		vectors:   make([]vectorDescriptor, maxSize),
		chunks:    make([]chunk, maxSize/batchSize),
	}

	for i := range q.chunks {
		q.chunks[i].id = i
		q.chunks[i].free = true
		q.chunks[i].chunk = q.vectors[i*q.batchSize : (i+1)*q.batchSize]
	}

	q.curBatch = &q.chunks[0]
	q.curBatch.free = false

	return &q
}

func (q *vectorQueue) Add(vectors ...vectorDescriptor) int {
	q.Lock()
	defer q.Unlock()

	if !q.ensureHasSpace() {
		return 0
	}

	for i, v := range vectors {
		q.curBatch.chunk[q.curBatch.cursor] = v
		q.curBatch.cursor++

		if !q.ensureHasSpace() {
			return i + 1
		}
	}

	return len(vectors)
}

func (q *vectorQueue) ensureHasSpace() bool {
	if q.curBatch.cursor < q.batchSize {
		return true
	}

	// fmt.Println("batch full", q.curBatch.id)

	q.curBatch.full = true

	c := q.getFreeChunk()
	if c == nil {
		return false
	}

	q.curBatch = c
	q.curBatch.free = false
	q.curBatch.cursor = 0

	return true
}

func (q *vectorQueue) getFreeChunk() *chunk {
	for i := range q.chunks {
		if q.chunks[i].free {
			return &q.chunks[i]
		}
	}

	return nil
}

func (q *vectorQueue) HasChunksReady(fullOnly bool) bool {
	q.RLock()
	defer q.RUnlock()

	for i := range q.chunks {
		if !q.chunks[i].free && !q.chunks[i].borrowed {
			if fullOnly && !q.chunks[i].full {
				continue
			}

			if q.chunks[i].cursor == 0 {
				continue
			}

			return true
		}
	}

	return false
}

func (q *vectorQueue) BorrowNextChunk(fullOnly bool) *chunk {
	q.Lock()
	defer q.Unlock()

	for i := range q.chunks {
		if !q.chunks[i].free && !q.chunks[i].borrowed {
			if fullOnly && !q.chunks[i].full {
				continue
			}

			q.chunks[i].borrowed = true
			// fmt.Println("borrowed chunk", q.chunks[i].id)
			return &q.chunks[i]
		}
	}

	return nil
}

func (q *vectorQueue) ReleaseChunk(c *chunk) {
	q.Lock()
	defer q.Unlock()

	// fmt.Println("released chunk", c.id)
	c.full = false
	c.free = true
	c.borrowed = false
	c.cursor = 0
}

func (q *vectorQueue) AppendSnapshot(buf []vectorDescriptor) []vectorDescriptor {
	q.RLock()
	defer q.RUnlock()

	for i := range q.chunks {
		if q.chunks[i].free {
			continue
		}
		buf = append(buf, q.chunks[i].chunk[:q.chunks[i].cursor]...)
	}

	return buf
}

type chunk struct {
	id       int
	free     bool
	full     bool
	cursor   int
	borrowed bool
	chunk    []vectorDescriptor
}
