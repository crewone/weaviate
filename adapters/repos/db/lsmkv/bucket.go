//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package lsmkv

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/weaviate/sroar"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/inverted/terms"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv/segmentindex"
	"github.com/weaviate/weaviate/entities/cyclemanager"
	"github.com/weaviate/weaviate/entities/interval"
	"github.com/weaviate/weaviate/entities/lsmkv"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/entities/storagestate"
	"github.com/weaviate/weaviate/entities/storobj"
	"github.com/weaviate/weaviate/usecases/memwatch"
)

const FlushAfterDirtyDefault = 60 * time.Second

type BucketCreator interface {
	NewBucket(ctx context.Context, dir, rootDir string, logger logrus.FieldLogger,
		metrics *Metrics, compactionCallbacks, flushCallbacks cyclemanager.CycleCallbackGroup,
		opts ...BucketOption,
	) (*Bucket, error)
}

type Bucket struct {
	dir      string
	rootDir  string
	active   *Memtable
	flushing *Memtable
	disk     *SegmentGroup
	logger   logrus.FieldLogger

	// Lock() means a move from active to flushing is happening, RLock() is
	// normal operation
	flushLock        sync.RWMutex
	haltedFlushTimer *interval.BackoffTimer

	walThreshold      uint64
	flushDirtyAfter   time.Duration
	memtableThreshold uint64
	memtableResizer   *memtableSizeAdvisor
	strategy          string
	// Strategy inverted index is supposed to be created with, but existing
	// segment files were created with different one.
	// It can happen when new strategy were introduced to weaviate, but
	// files are already created using old implementation.
	// Example: RoaringSet strategy replaces CollectionSet strategy.
	// Field can be used for migration files of old strategy to newer one.
	desiredStrategy  string
	secondaryIndices uint16

	// Optional to avoid syscalls
	mmapContents bool

	// for backward compatibility
	legacyMapSortingBeforeCompaction bool

	flushCallbackCtrl cyclemanager.CycleCallbackCtrl

	status     storagestate.Status
	statusLock sync.RWMutex

	metrics *Metrics

	// all "replace" buckets support counting through net additions, but not all
	// produce a meaningful count. Typically, the only count we're interested in
	// is that of the bucket that holds objects
	monitorCount bool

	pauseTimer *prometheus.Timer // Times the pause

	// Whether tombstones (set/map/replace types) or deletions (roaringset type)
	// should be kept in root segment during compaction process.
	// Since segments are immutable, deletions are added as new entries with
	// tombstones. Tombstones are by default copied to merged segment, as they
	// can refer to keys/values present in previous segments.
	// Those tombstones can be removed entirely when merging with root (1st) segment,
	// due to lack of previous segments, tombstones may relate to.
	// As info about key/value being deleted (based on tombstone presence) may be important
	// for some use cases (e.g. replication needs to know if object(ObjectsBucketLSM) was deleted)
	// keeping tombstones on compaction is optional
	keepTombstones bool

	// Init and use bloom filter for getting key from bucket segments.
	// As some buckets can be accessed only with cursor (see flat index),
	// where bloom filter is not applicable, it can be disabled.
	// ON by default
	useBloomFilter bool

	// Net additions keep track of number of elements stored in bucket (of type replace).
	// As some buckets don't have to provide Count info (see flat index),
	// tracking additions can be disabled.
	// ON by default
	calcCountNetAdditions bool

	forceCompaction bool

	// optionally supplied to prevent starting memory-intensive
	// processes when memory pressure is high
	allocChecker memwatch.AllocChecker

	// optional segment size limit. If set, a compaction will skip segments that
	// sum to more than the specified value.
	maxSegmentSize int64

	// optional segments cleanup interval. If set, segments will be cleaned of
	// redundant obsolete data, that was deleted or updated in newer segments
	// (currently supported only in buckets of REPLACE strategy)
	segmentsCleanupInterval time.Duration

	// optional validation of segment file checksums. Enabling this option
	// introduces latency of segment availability, for the tradeoff of
	// ensuring segment files have integrity before reading them.
	enableChecksumValidation bool
}

func NewBucketCreator() *Bucket { return &Bucket{} }

// NewBucket initializes a new bucket. It either loads the state from disk if
// it exists, or initializes new state.
//
// You do not need to ever call NewBucket() yourself, if you are using a
// [Store]. In this case the [Store] can manage buckets for you, using methods
// such as CreateOrLoadBucket().
func (*Bucket) NewBucket(ctx context.Context, dir, rootDir string, logger logrus.FieldLogger,
	metrics *Metrics, compactionCallbacks, flushCallbacks cyclemanager.CycleCallbackGroup,
	opts ...BucketOption,
) (*Bucket, error) {
	beforeAll := time.Now()
	defaultMemTableThreshold := uint64(10 * 1024 * 1024)
	defaultWalThreshold := uint64(1024 * 1024 * 1024)
	defaultFlushAfterDirty := FlushAfterDirtyDefault
	defaultStrategy := StrategyReplace

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}

	b := &Bucket{
		dir:                   dir,
		rootDir:               rootDir,
		memtableThreshold:     defaultMemTableThreshold,
		walThreshold:          defaultWalThreshold,
		flushDirtyAfter:       defaultFlushAfterDirty,
		strategy:              defaultStrategy,
		mmapContents:          true,
		logger:                logger,
		metrics:               metrics,
		useBloomFilter:        true,
		calcCountNetAdditions: true,
		haltedFlushTimer:      interval.NewBackoffTimer(),
	}

	for _, opt := range opts {
		if err := opt(b); err != nil {
			return nil, err
		}
	}

	if b.memtableResizer != nil {
		b.memtableThreshold = uint64(b.memtableResizer.Initial())
	}

	sg, err := newSegmentGroup(logger, metrics, compactionCallbacks,
		sgConfig{
			dir:                      dir,
			strategy:                 b.strategy,
			mapRequiresSorting:       b.legacyMapSortingBeforeCompaction,
			monitorCount:             b.monitorCount,
			mmapContents:             b.mmapContents,
			keepTombstones:           b.keepTombstones,
			forceCompaction:          b.forceCompaction,
			useBloomFilter:           b.useBloomFilter,
			calcCountNetAdditions:    b.calcCountNetAdditions,
			maxSegmentSize:           b.maxSegmentSize,
			cleanupInterval:          b.segmentsCleanupInterval,
			enableChecksumValidation: b.enableChecksumValidation,
		}, b.allocChecker)
	if err != nil {
		return nil, fmt.Errorf("init disk segments: %w", err)
	}

	b.desiredStrategy = b.strategy
	// Actual strategy is stored in segment files. In case it is SetCollection,
	// while new implementation uses bitmaps and supposed to be RoaringSet,
	// bucket and segmentgroup strategy is changed back to SetCollection
	// (memtables will be created later on, with already modified strategy)
	// TODO what if only WAL files exists, and there is no segment to get actual strategy?
	if b.strategy == StrategyRoaringSet && len(sg.segments) > 0 &&
		sg.segments[0].strategy == segmentindex.StrategySetCollection {
		b.strategy = StrategySetCollection
		b.desiredStrategy = StrategyRoaringSet
		sg.strategy = StrategySetCollection
	}
	// As of v1.19 property's IndexInterval setting is replaced with
	// IndexFilterable (roaring set) + IndexSearchable (map) and enabled by default.
	// Buckets for text/text[] inverted indexes created before 1.19 have strategy
	// map and name that since 1.19 is used by filterable indeverted index.
	// Those buckets (roaring set by configuration, but in fact map) have to be
	// renamed on startup by migrator. Here actual strategy is set based on
	// data found in segment files
	if b.strategy == StrategyRoaringSet && len(sg.segments) > 0 &&
		sg.segments[0].strategy == segmentindex.StrategyMapCollection {
		b.strategy = StrategyMapCollection
		b.desiredStrategy = StrategyRoaringSet
		sg.strategy = StrategyMapCollection
	}

	// Inverted segments share a lot of their logic as the MapCollection,
	// and the main difference is in the way they store their data.
	// Setting the desired strategy to Inverted will make sure that we can
	// distinguish between the two strategies for search.
	// The changes only apply when we have segments on disk,
	// as the memtables will always be created with the MapCollection strategy.
	if b.strategy == StrategyMapCollection && len(sg.segments) > 0 &&
		sg.segments[0].strategy == segmentindex.StrategyInverted {
		b.desiredStrategy = StrategyInverted
	}

	b.disk = sg

	if err := b.mayRecoverFromCommitLogs(ctx); err != nil {
		return nil, err
	}

	err = b.setNewActiveMemtable()
	if err != nil {
		return nil, err
	}

	id := "bucket/flush/" + b.dir
	b.flushCallbackCtrl = flushCallbacks.Register(id, b.flushAndSwitchIfThresholdsMet)

	b.metrics.TrackStartupBucket(beforeAll)

	if err := GlobalBucketRegistry.TryAdd(dir); err != nil {
		// prevent accidentally trying to register the same bucket twice
		return nil, err
	}

	return b, nil
}

func (b *Bucket) GetDir() string {
	return b.dir
}

func (b *Bucket) GetRootDir() string {
	return b.rootDir
}

func (b *Bucket) GetStrategy() string {
	return b.strategy
}

func (b *Bucket) GetDesiredStrategy() string {
	return b.desiredStrategy
}

func (b *Bucket) GetSecondaryIndices() uint16 {
	return b.secondaryIndices
}

func (b *Bucket) GetStatus() storagestate.Status {
	b.statusLock.RLock()
	defer b.statusLock.RUnlock()

	return b.status
}

func (b *Bucket) GetMemtableThreshold() uint64 {
	return b.memtableThreshold
}

func (b *Bucket) GetWalThreshold() uint64 {
	return b.walThreshold
}

func (b *Bucket) GetFlushCallbackCtrl() cyclemanager.CycleCallbackCtrl {
	return b.flushCallbackCtrl
}

func (b *Bucket) IterateObjects(ctx context.Context, f func(object *storobj.Object) error) error {
	cursor := b.Cursor()
	defer cursor.Close()

	return b.iterateObjectsCursor(ctx, cursor, f)
}

func (b *Bucket) IterateObjectsWith(ctx context.Context, desiredSecondaryIndexCount int, f func(object *storobj.Object) error) error {
	cursor := b.CursorWith(desiredSecondaryIndexCount)
	defer cursor.Close()

	return b.iterateObjectsCursor(ctx, cursor, f)
}

func (b *Bucket) iterateObjectsCursor(ctx context.Context, cursor *CursorReplace, f func(object *storobj.Object) error) error {
	i := 0

	for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
		obj, err := storobj.FromBinary(v)
		if err != nil {
			return fmt.Errorf("cannot unmarshal object %d, %w", i, err)
		}
		if err := f(obj); err != nil {
			return fmt.Errorf("callback on object '%d' failed: %w", obj.DocID, err)
		}

		i++
	}

	return nil
}

func (b *Bucket) IterateMapObjects(ctx context.Context, f func([]byte, []byte, []byte, bool) error) error {
	cursor := b.MapCursor()
	defer cursor.Close()

	for kList, vList := cursor.First(ctx); kList != nil; kList, vList = cursor.Next(ctx) {
		for _, v := range vList {
			if err := f(kList, v.Key, v.Value, v.Tombstone); err != nil {
				return fmt.Errorf("callback on object '%v' failed: %w", v, err)
			}
		}
	}

	return nil
}

func (b *Bucket) SetMemtableThreshold(size uint64) {
	b.memtableThreshold = size
}

// Get retrieves the single value for the given key.
//
// Get is specific to ReplaceStrategy and cannot be used with any of the other
// strategies. Use [Bucket.SetList] or [Bucket.MapList] instead.
//
// Get uses the regular or "primary" key for an object. If a bucket has
// secondary indexes, use [Bucket.GetBySecondary] to retrieve an object using
// its secondary key
func (b *Bucket) Get(key []byte) ([]byte, error) {
	beforeFlushLock := time.Now()
	b.flushLock.RLock()
	if time.Since(beforeFlushLock) > 100*time.Millisecond {
		b.logger.WithField("duration", time.Since(beforeFlushLock)).
			WithField("action", "lsm_bucket_get_acquire_flush_lock").
			Debugf("Waited more than 100ms to obtain a flush lock during get")
	}
	defer b.flushLock.RUnlock()

	return b.get(key)
}

func (b *Bucket) get(key []byte) ([]byte, error) {
	beforeMemtable := time.Now()
	v, err := b.active.get(key)
	if time.Since(beforeMemtable) > 100*time.Millisecond {
		b.logger.WithField("duration", time.Since(beforeMemtable)).
			WithField("action", "lsm_bucket_get_active_memtable").
			Debugf("Waited more than 100ms to retrieve object from memtable")
	}
	if err == nil {
		// item found and no error, return and stop searching, since the strategy
		// is replace
		return v, nil
	}
	if errors.Is(err, lsmkv.Deleted) {
		// deleted in the mem-table (which is always the latest) means we don't
		// have to check the disk segments, return nil now
		return nil, nil
	}

	if !errors.Is(err, lsmkv.NotFound) {
		panic(fmt.Sprintf("unsupported error in bucket.Get: %v\n", err))
	}

	if b.flushing != nil {
		beforeFlushMemtable := time.Now()
		v, err := b.flushing.get(key)
		if time.Since(beforeFlushMemtable) > 100*time.Millisecond {
			b.logger.WithField("duration", time.Since(beforeFlushMemtable)).
				WithField("action", "lsm_bucket_get_flushing_memtable").
				Debugf("Waited over 100ms to retrieve object from flushing memtable")
		}
		if err == nil {
			// item found and no error, return and stop searching, since the strategy
			// is replace
			return v, nil
		}
		if errors.Is(err, lsmkv.Deleted) {
			// deleted in the now most recent memtable  means we don't have to check
			// the disk segments, return nil now
			return nil, nil
		}

		if !errors.Is(err, lsmkv.NotFound) {
			panic("unsupported error in bucket.Get")
		}
	}

	return b.disk.get(key)
}

func (b *Bucket) GetErrDeleted(key []byte) ([]byte, error) {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	v, err := b.active.get(key)
	if err == nil {
		// item found and no error, return and stop searching, since the strategy
		// is replace
		return v, nil
	}
	if errors.Is(err, lsmkv.Deleted) {
		// deleted in the mem-table (which is always the latest) means we don't
		// have to check the disk segments, return nil now
		return nil, err
	}

	if !errors.Is(err, lsmkv.NotFound) {
		panic(fmt.Sprintf("unsupported error in bucket.Get: %v\n", err))
	}

	if b.flushing != nil {
		v, err := b.flushing.get(key)
		if err == nil {
			// item found and no error, return and stop searching, since the strategy
			// is replace
			return v, nil
		}
		if errors.Is(err, lsmkv.Deleted) {
			// deleted in the now most recent memtable  means we don't have to check
			// the disk segments, return nil now
			return nil, err
		}

		if !errors.Is(err, lsmkv.NotFound) {
			panic("unsupported error in bucket.Get")
		}
	}

	return b.disk.getErrDeleted(key)
}

// GetBySecondary retrieves an object using one of its secondary keys. A bucket
// can have an infinite number of secondary keys. Specify the secondary key
// position as the first argument.
//
// A real-life example of secondary keys is the Weaviate object store. Objects
// are stored with the user-facing ID as their primary key and with the doc-id
// (an ever-increasing uint64) as the secondary key.
//
// Similar to [Bucket.Get], GetBySecondary is limited to ReplaceStrategy. No
// equivalent exists for Set and Map, as those do not support secondary
// indexes.
func (b *Bucket) GetBySecondary(pos int, key []byte) ([]byte, error) {
	bytes, _, err := b.GetBySecondaryIntoMemory(pos, key, nil)
	return bytes, err
}

// GetBySecondaryWithBuffer is like [Bucket.GetBySecondary], but also takes a
// buffer. It's in the response of the caller to pool the buffer, since the
// bucket does not know when the caller is done using it. The return bytes will
// likely point to the same memory that's part of the buffer. However, if the
// buffer is to small, a larger buffer may also be returned (second arg).
func (b *Bucket) GetBySecondaryWithBuffer(pos int, key []byte, buf []byte) ([]byte, []byte, error) {
	bytes, newBuf, err := b.GetBySecondaryIntoMemory(pos, key, buf)
	return bytes, newBuf, err
}

// GetBySecondaryIntoMemory copies into the specified memory, and retrieves
// an object using one of its secondary keys. A bucket
// can have an infinite number of secondary keys. Specify the secondary key
// position as the first argument.
//
// A real-life example of secondary keys is the Weaviate object store. Objects
// are stored with the user-facing ID as their primary key and with the doc-id
// (an ever-increasing uint64) as the secondary key.
//
// Similar to [Bucket.Get], GetBySecondary is limited to ReplaceStrategy. No
// equivalent exists for Set and Map, as those do not support secondary
// indexes.
func (b *Bucket) GetBySecondaryIntoMemory(pos int, key []byte, buffer []byte) ([]byte, []byte, error) {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	v, err := b.active.getBySecondary(pos, key)
	if err == nil {
		// item found and no error, return and stop searching, since the strategy
		// is replace
		return v, buffer, nil
	}
	if errors.Is(err, lsmkv.Deleted) {
		// deleted in the mem-table (which is always the latest) means we don't
		// have to check the disk segments, return nil now
		return nil, buffer, nil
	}

	if !errors.Is(err, lsmkv.NotFound) {
		panic("unsupported error in bucket.Get")
	}

	if b.flushing != nil {
		v, err := b.flushing.getBySecondary(pos, key)
		if err == nil {
			// item found and no error, return and stop searching, since the strategy
			// is replace
			return v, buffer, nil
		}
		if errors.Is(err, lsmkv.Deleted) {
			// deleted in the now most recent memtable  means we don't have to check
			// the disk segments, return nil now
			return nil, buffer, nil
		}

		if !errors.Is(err, lsmkv.NotFound) {
			panic("unsupported error in bucket.Get")
		}
	}

	k, v, buffer, err := b.disk.getBySecondaryIntoMemory(pos, key, buffer)
	if err != nil {
		return nil, nil, err
	}

	// additional validation to ensure the primary key has not been marked as deleted
	pkv, err := b.get(k)
	if err != nil {
		return nil, nil, err
	} else if pkv == nil {
		return nil, buffer, nil
	}

	return v, buffer, nil
}

// SetList returns all Set entries for a given key.
//
// SetList is specific to the Set Strategy, for Map use [Bucket.MapList], and
// for Replace use [Bucket.Get].
func (b *Bucket) SetList(key []byte) ([][]byte, error) {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	var out []value

	v, err := b.disk.getCollection(key)
	if err != nil && !errors.Is(err, lsmkv.NotFound) {
		return nil, err
	}
	out = v

	if b.flushing != nil {
		v, err = b.flushing.getCollection(key)
		if err != nil && !errors.Is(err, lsmkv.NotFound) {
			return nil, err
		}
		out = append(out, v...)

	}

	v, err = b.active.getCollection(key)
	if err != nil && !errors.Is(err, lsmkv.NotFound) {
		return nil, err
	}
	if len(v) > 0 {
		// skip the expensive append operation if there was no memtable
		out = append(out, v...)
	}

	return newSetDecoder().Do(out), nil
}

// Put creates or replaces a single value for a given key.
//
//	err := bucket.Put([]byte("my_key"), []byte("my_value"))
//	 if err != nil {
//		/* do something */
//	}
//
// If a bucket has a secondary index configured, you can also specify one or
// more secondary keys, like so:
//
//	err := bucket.Put([]byte("my_key"), []byte("my_value"),
//		WithSecondaryKey(0, []byte("my_alternative_key")),
//	)
//	 if err != nil {
//		/* do something */
//	}
//
// Put is limited to ReplaceStrategy, use [Bucket.SetAdd] for Set or
// [Bucket.MapSet] and [Bucket.MapSetMulti].
func (b *Bucket) Put(key, value []byte, opts ...SecondaryKeyOption) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	return b.active.put(key, value, opts...)
}

// SetAdd adds one or more Set-Entries to a Set for the given key. SetAdd is
// entirely agnostic of existing entries, it acts as append-only. This also
// makes it agnostic of whether the key already exists or not.
//
// Example to add two entries to a set:
//
//	err := bucket.SetAdd([]byte("my_key"), [][]byte{
//		[]byte("one-set-element"), []byte("another-set-element"),
//	})
//	if err != nil {
//		/* do something */
//	}
//
// SetAdd is specific to the Set strategy. For Replace, use [Bucket.Put], for
// Map use either [Bucket.MapSet] or [Bucket.MapSetMulti].
func (b *Bucket) SetAdd(key []byte, values [][]byte) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	return b.active.append(key, newSetEncoder().Do(values))
}

// SetDeleteSingle removes one Set element from the given key. Note that LSM
// stores are append only, thus internally this action appends a tombstone. The
// entry will not be removed until a compaction has run, and even then a
// compaction does not guarantee the removal of the data right away. This is
// because an entry could have been created in an older segment than those
// present in the compaction. This can be seen as an implementation detail,
// unless the caller expects to free disk space by calling this method. Such
// freeing is not guaranteed.
//
// SetDeleteSingle is specific to the Set Strategy. For Replace, you can use
// [Bucket.Delete] to delete the entire row, for Maps use [Bucket.MapDeleteKey]
// to delete a single map entry.
func (b *Bucket) SetDeleteSingle(key []byte, valueToDelete []byte) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	return b.active.append(key, []value{
		{
			value:     valueToDelete,
			tombstone: true,
		},
	})
}

// WasDeleted determines if an object used to exist in the LSM store
//
// There are 3 different locations that we need to check for the key
// in this order: active memtable, flushing memtable, and disk
// segment
func (b *Bucket) WasDeleted(key []byte) (bool, time.Time, error) {
	if !b.keepTombstones {
		return false, time.Time{}, fmt.Errorf("Bucket requires option `keepTombstones` set to check deleted keys")
	}

	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	_, err := b.active.get(key)
	if err == nil {
		return false, time.Time{}, nil
	}
	if errors.Is(err, lsmkv.Deleted) {
		var errDeleted *lsmkv.ErrDeleted
		if errors.As(err, &errDeleted) {
			return true, errDeleted.DeletionTime(), nil
		} else {
			return true, time.Time{}, nil
		}
	}
	if !errors.Is(err, lsmkv.NotFound) {
		return false, time.Time{}, fmt.Errorf("unsupported bucket error: %w", err)
	}

	// can still check flushing and disk

	if b.flushing != nil {
		_, err := b.flushing.get(key)
		if err == nil {
			return false, time.Time{}, nil
		}
		if errors.Is(err, lsmkv.Deleted) {
			var errDeleted *lsmkv.ErrDeleted
			if errors.As(err, &errDeleted) {
				return true, errDeleted.DeletionTime(), nil
			} else {
				return true, time.Time{}, nil
			}
		}
		if !errors.Is(err, lsmkv.NotFound) {
			return false, time.Time{}, fmt.Errorf("unsupported bucket error: %w", err)
		}

		// can still check disk
	}

	_, err = b.disk.getErrDeleted(key)
	if err == nil {
		return false, time.Time{}, nil
	}
	if errors.Is(err, lsmkv.Deleted) {
		var errDeleted *lsmkv.ErrDeleted
		if errors.As(err, &errDeleted) {
			return true, errDeleted.DeletionTime(), nil
		} else {
			return true, time.Time{}, nil
		}
	}
	if !errors.Is(err, lsmkv.NotFound) {
		return false, time.Time{}, fmt.Errorf("unsupported bucket error: %w", err)
	}

	return false, time.Time{}, nil
}

type MapListOptionConfig struct {
	acceptDuplicates           bool
	legacyRequireManualSorting bool
}

type MapListOption func(c *MapListOptionConfig)

func MapListAcceptDuplicates() MapListOption {
	return func(c *MapListOptionConfig) {
		c.acceptDuplicates = true
	}
}

func MapListLegacySortingRequired() MapListOption {
	return func(c *MapListOptionConfig) {
		c.legacyRequireManualSorting = true
	}
}

// MapList returns all map entries for a given row key. The order of map pairs
// has no specific meaning. For efficient merge operations, pair entries are
// stored sorted on disk, however that is an implementation detail and not a
// caller-facing guarantee.
//
// MapList is specific to the Map strategy, for Sets use [Bucket.SetList], for
// Replace use [Bucket.Get].
func (b *Bucket) MapList(ctx context.Context, key []byte, cfgs ...MapListOption) ([]MapPair, error) {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	c := MapListOptionConfig{}
	for _, cfg := range cfgs {
		cfg(&c)
	}

	segments := [][]MapPair{}
	// before := time.Now()
	disk, segmentsDisk, err := b.disk.getCollectionAndSegments(key)
	if err != nil && !errors.Is(err, lsmkv.NotFound) {
		return nil, err
	}

	allTombstones, err := b.loadAllTombstones(segmentsDisk)
	if err != nil {
		return nil, err
	}

	for i := range disk {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		propLengths := make(map[uint64]uint32)
		if segmentsDisk[i].strategy == segmentindex.StrategyInverted {
			propLengths, err = segmentsDisk[i].GetPropertyLengths()
			if err != nil {
				return nil, err
			}
		}

		segmentDecoded := make([]MapPair, len(disk[i]))
		for j, v := range disk[i] {
			// Inverted segments have a slightly different internal format
			// and separate property lengths that need to be read.
			if segmentsDisk[i].strategy == segmentindex.StrategyInverted {
				if err := segmentDecoded[j].FromBytesInverted(v.value, false); err != nil {
					return nil, err
				}
				docId := binary.BigEndian.Uint64(segmentDecoded[j].Key[:8])
				// check if there are any tombstones between the i and len(disk) segments
				for _, tombstones := range allTombstones[i+1:] {
					if tombstones != nil && tombstones.Contains(docId) {
						segmentDecoded[j].Tombstone = true
						break
					}
				}
				// put the property length in the value from the "external" property lengths
				binary.LittleEndian.PutUint32(segmentDecoded[j].Value[4:], math.Float32bits(float32(propLengths[docId])))

			} else {
				if err := segmentDecoded[j].FromBytes(v.value, false); err != nil {
					return nil, err
				}
				// Read "broken" tombstones with length 12 but a non-tombstone value
				// Related to Issue #4125
				// TODO: Remove the extra check, as it may interfere future in-disk format changes
				segmentDecoded[j].Tombstone = v.tombstone || len(v.value) == 12
			}
		}
		if len(segmentDecoded) > 0 {
			segments = append(segments, segmentDecoded)
		}
	}

	// fmt.Printf("--map-list: get all disk segments took %s\n", time.Since(before))

	// before = time.Now()
	// fmt.Printf("--map-list: append all disk segments took %s\n", time.Since(before))

	if b.flushing != nil {
		v, err := b.flushing.getMap(key)
		if err != nil && !errors.Is(err, lsmkv.NotFound) {
			return nil, err
		}
		if len(v) > 0 {
			segments = append(segments, v)
		}
	}

	// before = time.Now()
	v, err := b.active.getMap(key)
	if err != nil && !errors.Is(err, lsmkv.NotFound) {
		return nil, err
	}
	if len(v) > 0 {
		segments = append(segments, v)
	}
	// fmt.Printf("--map-list: get all active segments took %s\n", time.Since(before))

	// before = time.Now()
	// defer func() {
	// 	fmt.Printf("--map-list: run decoder took %s\n", time.Since(before))
	// }()

	if c.legacyRequireManualSorting {
		// Sort to support segments which were stored in an unsorted fashion
		for i := range segments {
			sort.Slice(segments[i], func(a, b int) bool {
				return bytes.Compare(segments[i][a].Key, segments[i][b].Key) == -1
			})
		}
	}

	return newSortedMapMerger().do(ctx, segments)
}

func (b *Bucket) loadAllTombstones(segmentsDisk []*segment) ([]*sroar.Bitmap, error) {
	hasTombstones := false
	allTombstones := make([]*sroar.Bitmap, len(segmentsDisk)+2)
	for i, segment := range segmentsDisk {
		if segment.strategy == segmentindex.StrategyInverted {
			tombstones, err := segment.GetTombstones()
			if err != nil {
				return nil, err
			}
			allTombstones[i] = tombstones
			hasTombstones = true
		}
	}
	if hasTombstones {

		if b.flushing != nil {
			tombstones, err := b.flushing.GetTombstones()
			if err != nil {
				return nil, err
			}
			allTombstones[len(segmentsDisk)] = tombstones
		}

		tombstones, err := b.active.GetTombstones()
		if err != nil {
			return nil, err
		}
		allTombstones[len(segmentsDisk)+1] = tombstones
	}
	return allTombstones, nil
}

// MapSet writes one [MapPair] into the map for the given row key. It is
// agnostic of whether the row key already exists, as well as agnostic of
// whether the map key already exists. In both cases it will create the entry
// if it does not exist or override if it does.
//
// Example to add a new MapPair:
//
//	pair := MapPair{Key: []byte("Jane"), Value: []byte("Backend")}
//	err := bucket.MapSet([]byte("developers"), pair)
//	if err != nil {
//		/* do something */
//	}
//
// MapSet is specific to the Map Strategy, for Replace use [Bucket.Put], and for Set use [Bucket.SetAdd] instead.
func (b *Bucket) MapSet(rowKey []byte, kv MapPair) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	return b.active.appendMapSorted(rowKey, kv)
}

// MapSetMulti is the same as [Bucket.MapSet], except that it takes in multiple
// [MapPair] objects at the same time.
func (b *Bucket) MapSetMulti(rowKey []byte, kvs []MapPair) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	for _, kv := range kvs {
		if err := b.active.appendMapSorted(rowKey, kv); err != nil {
			return err
		}
	}

	return nil
}

// MapDeleteKey removes one key-value pair from the given map row. Note that
// LSM stores are append only, thus internally this action appends a tombstone.
// The entry will not be removed until a compaction has run, and even then a
// compaction does not guarantee the removal of the data right away. This is
// because an entry could have been created in an older segment than those
// present in the compaction. This can be seen as an implementation detail,
// unless the caller expects to free disk space by calling this method. Such
// freeing is not guaranteed.
//
// MapDeleteKey is specific to the Map Strategy. For Replace, you can use
// [Bucket.Delete] to delete the entire row, for Sets use [Bucket.SetDeleteSingle] to delete a single set element.
func (b *Bucket) MapDeleteKey(rowKey, mapKey []byte) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	pair := MapPair{
		Key:       mapKey,
		Tombstone: true,
	}

	return b.active.appendMapSorted(rowKey, pair)
}

// Delete removes the given row. Note that LSM stores are append only, thus
// internally this action appends a tombstone.  The entry will not be removed
// until a compaction has run, and even then a compaction does not guarantee
// the removal of the data right away. This is because an entry could have been
// created in an older segment than those present in the compaction. This can
// be seen as an implementation detail, unless the caller expects to free disk
// space by calling this method. Such freeing is not guaranteed.
//
// Delete is specific to the Replace Strategy. For Maps, you can use
// [Bucket.MapDeleteKey] to delete a single key-value pair, for Sets use
// [Bucket.SetDeleteSingle] to delete a single set element.
func (b *Bucket) Delete(key []byte, opts ...SecondaryKeyOption) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	return b.active.setTombstone(key, opts...)
}

func (b *Bucket) DeleteWith(key []byte, deletionTime time.Time, opts ...SecondaryKeyOption) error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	if !b.keepTombstones {
		return fmt.Errorf("bucket requires option `keepTombstones` set to delete keys at a given timestamp")
	}

	return b.active.setTombstoneWith(key, deletionTime, opts...)
}

// meant to be called from situations where a lock is already held, does not
// lock on its own
func (b *Bucket) setNewActiveMemtable() error {
	path := filepath.Join(b.dir, fmt.Sprintf("segment-%d", time.Now().UnixNano()))

	cl, err := newCommitLogger(path)
	if err != nil {
		return errors.Wrap(err, "init commit logger")
	}

	mt, err := newMemtable(path, b.strategy, b.secondaryIndices, cl,
		b.metrics, b.logger, b.enableChecksumValidation)
	if err != nil {
		return err
	}

	b.active = mt
	return nil
}

func (b *Bucket) Count() int {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	if b.strategy != StrategyReplace {
		panic("Count() called on strategy other than 'replace'")
	}

	memtableCount := 0
	if b.flushing == nil {
		// only consider active
		memtableCount += b.memtableNetCount(b.active.countStats(), nil)
	} else {
		flushingCountStats := b.flushing.countStats()
		activeCountStats := b.active.countStats()
		deltaActive := b.memtableNetCount(activeCountStats, flushingCountStats)
		deltaFlushing := b.memtableNetCount(flushingCountStats, nil)

		memtableCount = deltaActive + deltaFlushing
	}

	diskCount := b.disk.count()

	if b.monitorCount {
		b.metrics.ObjectCount(memtableCount + diskCount)
	}
	return memtableCount + diskCount
}

// CountAsync ignores the current memtable, that makes it async because it only
// reflects what has been already flushed. This in turn makes it very cheap to
// call, so it can be used for observability purposes where eventual
// consistency on the count is fine, but a large cost is not.
func (b *Bucket) CountAsync() int {
	return b.disk.count()
}

func (b *Bucket) memtableNetCount(stats *countStats, previousMemtable *countStats) int {
	netCount := 0

	// TODO: this uses regular get, given that this may be called quite commonly,
	// we might consider building a pure Exists(), which skips reading the value
	// and only checks for tombstones, etc.
	for _, key := range stats.upsertKeys {
		if !b.existsOnDiskAndPreviousMemtable(previousMemtable, key) {
			netCount++
		}
	}

	for _, key := range stats.tombstonedKeys {
		if b.existsOnDiskAndPreviousMemtable(previousMemtable, key) {
			netCount--
		}
	}

	return netCount
}

func (b *Bucket) existsOnDiskAndPreviousMemtable(previous *countStats, key []byte) bool {
	v, _ := b.disk.get(key) // current implementation can't error
	if v == nil {
		// not on disk, but it could still be in the previous memtable
		return previous.hasUpsert(key)
	}

	// it exists on disk ,but it could still have been deleted in the previous memtable
	return !previous.hasTombstone(key)
}

func (b *Bucket) Shutdown(ctx context.Context) error {
	defer GlobalBucketRegistry.Remove(b.dir)

	if err := b.disk.shutdown(ctx); err != nil {
		return err
	}

	if err := b.flushCallbackCtrl.Unregister(ctx); err != nil {
		return fmt.Errorf("long-running flush in progress: %w", ctx.Err())
	}

	// Searchable buckets are flushed using the inverted index strategy,
	// if the environment variable USE_INVERTED_SEARCHABLE is set to true.
	// Memtables in memory are always created using the Map strategy.
	// See the desiredStrategy for more information on why this only matters
	// for searchable buckets and at flush time.
	if strings.Contains(b.active.path, "_searchable") && os.Getenv("USE_INVERTED_SEARCHABLE") == "true" {
		b.desiredStrategy = StrategyInverted
		b.active.flushStrategy = StrategyInverted
	}
	b.flushLock.Lock()
	if err := b.active.flush(); err != nil {
		return err
	}
	b.flushLock.Unlock()

	if b.flushing == nil {
		// active has flushing, no one else was currently flushing, it's safe to
		// exit
		return nil
	}

	// it seems we still need to wait for someone to finish flushing
	t := time.NewTicker(50 * time.Millisecond)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			if b.flushing == nil {
				return nil
			}
		}
	}
}

func (b *Bucket) flushAndSwitchIfThresholdsMet(shouldAbort cyclemanager.ShouldAbortCallback) bool {
	b.flushLock.RLock()
	commitLogSize := b.active.commitlog.Size()
	memtableTooLarge := b.active.Size() >= b.memtableThreshold
	walTooLarge := uint64(commitLogSize) >= b.walThreshold
	dirtyTooLong := b.active.DirtyDuration() >= b.flushDirtyAfter
	shouldSwitch := memtableTooLarge || walTooLarge || dirtyTooLong

	// If true, the parent shard has indicated that it has
	// entered an immutable state. During this time, the
	// bucket should refrain from flushing until its shard
	// indicates otherwise
	if shouldSwitch && b.isReadOnly() {
		if b.haltedFlushTimer.IntervalElapsed() {
			b.logger.WithField("action", "lsm_memtable_flush").
				WithField("path", b.dir).
				Warn("flush halted due to shard READONLY status")
			b.haltedFlushTimer.IncreaseInterval()
		}

		b.flushLock.RUnlock()
		return false
	}

	b.flushLock.RUnlock()
	if shouldSwitch {
		b.haltedFlushTimer.Reset()
		cycleLength := b.active.ActiveDuration()
		if err := b.FlushAndSwitch(); err != nil {
			b.logger.WithField("action", "lsm_memtable_flush").
				WithField("path", b.dir).
				WithError(err).
				Errorf("flush and switch failed")
		}

		if b.memtableResizer != nil {
			next, ok := b.memtableResizer.NextTarget(int(b.memtableThreshold), cycleLength)
			if ok {
				b.memtableThreshold = uint64(next)
			}
		}
		return true
	}
	return false
}

// UpdateStatus is used by the parent shard to communicate to the bucket
// when the shard has been set to readonly, or when it is ready for
// writes.
func (b *Bucket) UpdateStatus(status storagestate.Status) {
	b.statusLock.Lock()
	defer b.statusLock.Unlock()

	b.status = status
	b.disk.UpdateStatus(status)
}

func (b *Bucket) isReadOnly() bool {
	b.statusLock.Lock()
	defer b.statusLock.Unlock()

	return b.status == storagestate.StatusReadOnly
}

// FlushAndSwitch is the main way to flush a memtable, replace it with a new
// one, and make sure that the flushed segment gets added to the segment group.
//
// Flushing and adding a segment can take considerable time, which is why the
// whole process is designed to be non-blocking.
//
// To achieve a non-blocking flush, the process is split into four parts:
//
//  1. atomicallySwitchMemtable: A new memtable is created, the previous
//     memtable is moved from "active" to "flushing". This switch is blocking
//     (holds b.flushLock.Lock()), but extremely fast, as we essentially just
//     switch a pointer.
//
//  2. flush: The previous memtable is flushed to disk. This may take
//     considerable time as we are I/O-bound. This is done "in the
//     background"meaning that it does not block any CRUD operations for the
//     user. It only blocks the flush process itself, meaning only one flush per
//     bucket can happen simultaneously. This is by design.
//
//  3. initAndPrecomputeNewSegment: (Newly added in
//     https://github.com/weaviate/weaviate/pull/5943, early October 2024). After
//     the previous flush step the segment can now be initialized. However, to
//     make it usable for real life, we still need to compute metadata, such as
//     bloom filters (all types) and net count additions (only Replace type).
//     Bloom filters can be calculated in isolation and are therefore fairly
//     trivial. Net count additions on the other hand are more complex, as they
//     depend on all previous segments. Calculating net count additions can take
//     a considerable amount of time, especially as the buckets grow larger. As a
//     result, we need to provide two guarantees: (1) the calculation is
//     non-blocking from a user's POV and (2) for the duration of the
//     calculation, the segment group is considered stable, i.e. no other
//     segments are added, removed, or merged. We can achieve this by holding a
//     `b.disk.maintenanceLock.RLock()` which prevents modification of the
//     segments array, but does not block user operation (which are themselves
//     RLock-holders on that same Lock).
//
//  4. atomicallyAddDiskSegmentAndRemoveFlushing: The previous method returned
//     a fully initialized segment that has not yet been added to the segment
//     group. This last step is the counter part to the first step and again
//     blocking, but fast. It adds the segment to the segment group  which at
//     this point is just a simple array append. At the same time it removes the
//     "flushing" memtable. It holds the `b.flushLock.Lock()` making this
//     operation atomic, but blocking.
//
// FlushAndSwitch is typically called periodically and does not require manual
// calling, but there are some situations where this might be intended, such as
// in test scenarios or when a force flush is desired.
func (b *Bucket) FlushAndSwitch() error {
	before := time.Now()

	b.logger.WithField("action", "lsm_memtable_flush_start").
		WithField("path", b.dir).
		Trace("start flush and switch")
	if err := b.atomicallySwitchMemtable(); err != nil {
		return fmt.Errorf("switch active memtable: %w", err)
	}

	// Searchable buckets are flushed using the inverted index strategy,
	// if the environment variable USE_INVERTED_SEARCHABLE is set to true.
	// Memtables in memory are always created using the Map strategy.
	// See the desiredStrategy for more information on why this only matters
	// for searchable buckets and at flush time.
	if strings.Contains(b.flushing.path, "_searchable") && os.Getenv("USE_INVERTED_SEARCHABLE") == "true" {
		b.desiredStrategy = StrategyInverted
		b.flushing.flushStrategy = StrategyInverted
	}
	if err := b.flushing.flush(); err != nil {
		return fmt.Errorf("flush: %w", err)
	}

	segment, err := b.initAndPrecomputeNewSegment()
	if err != nil {
		return fmt.Errorf("precompute metadata: %w", err)
	}

	if err := b.atomicallyAddDiskSegmentAndRemoveFlushing(segment); err != nil {
		return fmt.Errorf("add segment and remove flushing: %w", err)
	}

	took := time.Since(before)
	b.logger.WithField("action", "lsm_memtable_flush_complete").
		WithField("path", b.dir).
		Trace("finish flush and switch")

	b.logger.WithField("action", "lsm_memtable_flush_complete").
		WithField("path", b.dir).
		WithField("took", took).
		Debugf("flush and switch took %s\n", took)

	return nil
}

func (b *Bucket) initAndPrecomputeNewSegment() (*segment, error) {
	// Note that this operation does not require the flush lock, i.e. it can
	// happen in the background and we can accept new writes will this
	// pre-compute is happening.
	path := b.flushing.path
	segment, err := b.disk.initAndPrecomputeNewSegment(path + ".db")
	if err != nil {
		return nil, err
	}

	return segment, nil
}

func (b *Bucket) atomicallyAddDiskSegmentAndRemoveFlushing(seg *segment) error {
	b.flushLock.Lock()
	defer b.flushLock.Unlock()

	if b.flushing.Size() == 0 {
		b.flushing = nil
		return nil
	}

	if err := b.disk.addInitializedSegment(seg); err != nil {
		return err
	}
	b.flushing = nil

	if b.strategy == StrategyReplace && b.monitorCount {
		// having just flushed the memtable we now have the most up2date count which
		// is a good place to update the metric
		b.metrics.ObjectCount(b.disk.count())
	}

	return nil
}

func (b *Bucket) atomicallySwitchMemtable() error {
	b.flushLock.Lock()
	defer b.flushLock.Unlock()

	b.flushing = b.active
	return b.setNewActiveMemtable()
}

func (b *Bucket) Strategy() string {
	return b.strategy
}

func (b *Bucket) DesiredStrategy() string {
	return b.desiredStrategy
}

// the WAL uses a buffer and isn't written until the buffer size is crossed or
// this function explicitly called. This allows to avoid unnecessary disk
// writes in larger operations, such as batches. It is sufficient to call write
// on the WAL just once. This does not make a batch atomic, but it guarantees
// that the WAL is written before a successful response is returned to the
// user.
func (b *Bucket) WriteWAL() error {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	return b.active.writeWAL()
}

func (b *Bucket) DocPointerWithScoreList(ctx context.Context, key []byte, propBoost float32, cfgs ...MapListOption) ([]terms.DocPointerWithScore, error) {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	c := MapListOptionConfig{}
	for _, cfg := range cfgs {
		cfg(&c)
	}

	segments := [][]terms.DocPointerWithScore{}
	disk, segmentsDisk, err := b.disk.getCollectionAndSegments(key)
	if err != nil && !errors.Is(err, lsmkv.NotFound) {
		return nil, err
	}

	allTombstones, err := b.loadAllTombstones(segmentsDisk)
	if err != nil {
		return nil, err
	}

	for i := range disk {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		propLengths := make(map[uint64]uint32)
		if segmentsDisk[i].strategy == segmentindex.StrategyInverted {
			propLengths, err = segmentsDisk[i].GetPropertyLengths()
			if err != nil {
				return nil, err
			}
		}

		segmentDecoded := make([]terms.DocPointerWithScore, len(disk[i]))
		for j, v := range disk[i] {
			if segmentsDisk[i].strategy == segmentindex.StrategyInverted {
				docId := binary.BigEndian.Uint64(v.value[:8])
				propLen := propLengths[docId]
				if err := segmentDecoded[j].FromBytesInverted(v.value, propBoost, float32(propLen)); err != nil {
					return nil, err
				}
				// check if there are any tombstones between the i and len(disk) segments
				for _, tombstones := range allTombstones[i+1:] {
					if tombstones != nil && tombstones.Contains(docId) {
						segmentDecoded[j].Frequency = 0
						break
					}
				}
			} else {
				if err := segmentDecoded[j].FromBytes(v.value, v.tombstone, propBoost); err != nil {
					return nil, err
				}
			}
		}
		segments = append(segments, segmentDecoded)
	}

	if b.flushing != nil {
		mem, err := b.flushing.getMap(key)
		if err != nil && !errors.Is(err, lsmkv.NotFound) {
			return nil, err
		}
		docPointers := make([]terms.DocPointerWithScore, len(mem))
		for i, v := range mem {
			if err := docPointers[i].FromKeyVal(v.Key, v.Value, v.Tombstone, propBoost); err != nil {
				return nil, err
			}
		}
		segments = append(segments, docPointers)
	}

	mem, err := b.active.getMap(key)
	if err != nil && !errors.Is(err, lsmkv.NotFound) {
		return nil, err
	}
	docPointers := make([]terms.DocPointerWithScore, len(mem))
	for i, v := range mem {
		if err := docPointers[i].FromKeyVal(v.Key, v.Value, v.Tombstone, propBoost); err != nil {
			return nil, err
		}
	}
	segments = append(segments, docPointers)

	if c.legacyRequireManualSorting {
		// Sort to support segments which were stored in an unsorted fashion
		for i := range segments {
			sort.Slice(segments[i], func(a, b int) bool {
				return segments[i][a].Id < segments[i][b].Id
			})
		}
	}

	return terms.NewSortedDocPointerWithScoreMerger().Do(ctx, segments)
}

func (b *Bucket) CreateDiskTerm(N float64, filterDocIds helpers.AllowList, query []string, propName string, propertyBoost float32, duplicateTextBoosts []int, averagePropLength float64, config schema.BM25Config, ctx context.Context) ([][]terms.TermInterface, *sync.RWMutex, error) {
	b.flushLock.RLock()
	defer b.flushLock.RUnlock()

	lock := &b.disk.maintenanceLock
	// The lock is necessary, as data is being read from the disk during blockmax wand search.
	// BlockMax is ran outside this function, so, the lock is returned to the caller.
	// Panics at this level are caught and the lock is released in the defer function.
	// The lock is released after the blockmax search is done, and panics are also handled.
	lock.RLock()

	defer func() {
		if r := recover(); r != nil {
			b.logger.Errorf("Recovered from panic in CreateDiskTerm: %v", r)
			lock.RUnlock()
		}
	}()

	segmentsDisk := b.disk.segments
	output := make([][]terms.TermInterface, len(segmentsDisk)+2)
	idfs := make([]float64, len(query))

	// flusing memtable
	output[len(segmentsDisk)] = make([]terms.TermInterface, 0, len(query))

	// active memtable
	output[len(segmentsDisk)+1] = make([]terms.TermInterface, 0, len(query))

	for i, queryTerm := range query {
		key := []byte(queryTerm)
		n := uint64(0)

		flushing := terms.NewTerm(queryTerm, i, propertyBoost, config)
		active := terms.NewTerm(queryTerm, i, propertyBoost, config)
		if b.flushing != nil {
			mapPairs, err := b.flushing.getMap(key)
			if err != nil && !errors.Is(err, lsmkv.NotFound) {
				return nil, lock, err
			}
			n2, err := addDataToTerm(mapPairs, filterDocIds, flushing)
			if err != nil {
				return nil, lock, err
			}
			n += n2
			if n2 > 0 {
				output[len(segmentsDisk)] = append(output[len(segmentsDisk)], flushing)
			}
		}

		if b.active != nil {
			mapPairs, err := b.active.getMap(key)
			if err != nil && !errors.Is(err, lsmkv.NotFound) {
				return nil, lock, err
			}
			n2, err := addDataToTerm(mapPairs, filterDocIds, active)
			if err != nil {
				return nil, lock, err
			}
			n += n2
			if n2 > 0 {
				output[len(segmentsDisk)+1] = append(output[len(segmentsDisk)+1], active)
			}
		}

		for _, segment := range segmentsDisk {
			if segment.strategy == segmentindex.StrategyInverted && segment.hasKey(key) {
				n += segment.getDocCount(key)
			}
		}

		idfs[i] = math.Log(float64(1)+(N-float64(n)+0.5)/(float64(n)+0.5)) * float64(duplicateTextBoosts[i])
		flushing.SetIdf(idfs[i])
		active.SetIdf(idfs[i])

	}

	allTombstones := make([]*sroar.Bitmap, len(segmentsDisk)+2)

	if b.flushing != nil {
		tombstones, err := b.flushing.GetTombstones()
		if err != nil {
			return nil, lock, err
		}
		allTombstones[len(segmentsDisk)] = tombstones

	}

	tombstones, err := b.active.GetTombstones()
	if err != nil {
		return nil, lock, err
	}
	allTombstones[len(segmentsDisk)+1] = tombstones

	for i, segment := range segmentsDisk {
		tombstones, err := segment.GetTombstones()
		if err != nil {
			return nil, lock, err
		}

		allTombstones[i] = tombstones
	}

	for j, segment := range segmentsDisk {
		output[j] = make([]terms.TermInterface, 0, len(query))
		for i, key := range query {
			term := NewSegmentBlockMax(segment, []byte(key), i, idfs[i], propertyBoost, allTombstones[j], filterDocIds, averagePropLength, config)
			if term != nil {
				output[j] = append(output[j], term)
			}
		}
	}
	return output, lock, nil
}

func addDataToTerm(mem []MapPair, filterDocIds helpers.AllowList, term *terms.Term) (uint64, error) {
	n := uint64(0)
	data := make([]terms.DocPointerWithScore, 0, len(mem))

	for _, v := range mem {
		if v.Tombstone {
			continue
		}
		n++
		if len(v.Value) < 8 {
			// b.logger.Warnf("Skipping pair in BM25: MapPair.Value should be 8 bytes long, but is %d.", len(v.Value))
			continue
		}
		d := terms.DocPointerWithScore{}
		if err := d.FromKeyVal(v.Key, v.Value, false, 1.0); err != nil {
			return 0, err
		}
		if filterDocIds != nil && !filterDocIds.Contains(d.Id) {
			continue
		}
		data = append(data, d)

	}
	if len(data) == 0 {
		return n, nil
	}
	term.Data = data
	term.SetPosPointer(0)
	term.SetIdPointer(term.Data[0].Id)
	return n, nil
}
