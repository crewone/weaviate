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

package inverted

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"

	"github.com/weaviate/weaviate/adapters/repos/db/vector/hnsw/priorityqueue"

	"github.com/weaviate/weaviate/entities/models"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/adapters/repos/db/helpers"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv"
	"github.com/weaviate/weaviate/adapters/repos/db/propertyspecific"
	"github.com/weaviate/weaviate/entities/additional"
	"github.com/weaviate/weaviate/entities/filters"
	"github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/entities/searchparams"
	"github.com/weaviate/weaviate/entities/storobj"
)

type BM25Searcher struct {
	config        schema.BM25Config
	store         *lsmkv.Store
	schema        schema.Schema
	rowCache      cacher
	classSearcher ClassSearcher // to allow recursive searches on ref-props
	propIndices   propertyspecific.Indices
	deletedDocIDs DeletedDocIDChecker
	propLengths   propLengthRetriever
	logger        logrus.FieldLogger
	shardVersion  uint16
}

type propLengthRetriever interface {
	PropertyMean(prop string) (float32, error)
}

func NewBM25Searcher(config schema.BM25Config, store *lsmkv.Store, schema schema.Schema,
	rowCache cacher, propIndices propertyspecific.Indices,
	classSearcher ClassSearcher, deletedDocIDs DeletedDocIDChecker,
	propLengths propLengthRetriever, logger logrus.FieldLogger,
	shardVersion uint16,
) *BM25Searcher {
	return &BM25Searcher{
		config:        config,
		store:         store,
		schema:        schema,
		rowCache:      rowCache,
		propIndices:   propIndices,
		classSearcher: classSearcher,
		deletedDocIDs: deletedDocIDs,
		propLengths:   propLengths,
		logger:        logger.WithField("action", "bm25_search"),
		shardVersion:  shardVersion,
	}
}

// Objects returns a list of full objects
func (b *BM25Searcher) Objects(ctx context.Context, limit int,
	keywordRanking *searchparams.KeywordRanking,
	filter *filters.LocalFilter, sort []filters.Sort, additional additional.Properties,
	className schema.ClassName,
) ([]*storobj.Object, []float32, error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()

	class, err := schema.GetClassByName(b.schema.Objects, string(className))
	if err != nil {
		return nil, []float32{}, errors.Wrap(err,
			"get class by name")
	}
	property := keywordRanking.Properties[0]
	p, err := schema.GetPropertyByName(class, property)
	if err != nil {
		return nil, []float32{}, errors.Wrap(err, "read property from class")
	}
	indexed := p.IndexInverted

	if indexed == nil || *indexed {
		return b.wand(ctx, class, keywordRanking.Query, keywordRanking.Properties[:1], limit)
	} else {
		return []*storobj.Object{}, []float32{}, nil
	}
}

type term struct {
	// doubles as max impact (with tf=1, the max impact would be 1*idf), if there
	// is a boost for a term, simply apply it here once
	idf float64

	idPointer         uint64
	posPointer        uint64
	data              []docPointerWithScore
	exhausted         bool
	k1                float64
	b                 float64
	averagePropLength float64
}

func (t terms) pivot(minScore float64) {
	minID, pivotPoint := t.findMinID(minScore)
	if pivotPoint == 0 {
		return
	}

	t.advanceAllAtLeast(minID)
	t.sort()
}

func (t terms) advanceAllAtLeast(minID uint64) {
	for i := range t {
		t[i].advanceAtLeast(minID)
	}
}

func (t terms) findMinID(minScore float64) (uint64, int) {
	cumScore := float64(0)

	for i, term := range t {
		cumScore += term.idf
		if cumScore >= minScore {
			return term.idPointer, i
		}
	}

	panic(fmt.Sprintf("score of %f is unreachable", minScore))
}

func (t terms) sort() {
	sort.Slice(t, func(a, b int) bool { return t[a].idPointer < t[b].idPointer })
}

func (t terms) findFirstNonExhausted() (int, bool) {
	for i := range t {
		if !t[i].exhausted {
			return i, true
		}
	}

	return -1, false
}

func (t terms) scoreNext() (uint64, float64, bool) {
	pos, ok := t.findFirstNonExhausted()
	if !ok {
		// done, nothing left to score
		return 0, 0, false
	}

	id := t[pos].idPointer
	var cumScore float64
	for i := pos; i < len(t); i++ {
		if t[i].idPointer != id || t[i].exhausted {
			continue
		}
		_, score := t[i].scoreAndAdvance()
		cumScore += score
	}

	return id, cumScore, true
}

func (t *term) scoreAndAdvance() (uint64, float64) {
	id := t.idPointer
	pair := t.data[t.posPointer]
	tf := pair.frequency / (pair.frequency + t.k1*(1-t.b+t.b*pair.propLength/t.averagePropLength))

	// advance
	t.posPointer++
	if t.posPointer >= uint64(len(t.data)) {
		t.exhausted = true
	} else {
		t.idPointer = t.data[t.posPointer].id
	}

	return id, tf * t.idf
}

func (t *term) advanceAtLeast(minID uint64) {
	for t.idPointer < minID {
		t.posPointer++
		if t.posPointer >= uint64(len(t.data)) {
			t.exhausted = true
			return
		}
		t.idPointer = t.data[t.posPointer].id
	}

	return
}

type terms []term

func (b *BM25Searcher) wand(
	ctx context.Context, class *models.Class, fullQuery string, properties []string, limit int,
) ([]*storobj.Object, []float32, error) {
	objectsBucket := b.store.Bucket(helpers.ObjectsBucketLSM)
	if objectsBucket == nil {
		return nil, nil, errors.Errorf("objects bucket not found")
	}
	N := float64(b.store.Bucket(helpers.ObjectsBucketLSM).Count())

	// There are currently cases, for different tokenization:
	// Text, string and field.
	// For the first two the query is tokenized accordingly and for the last one the full query is used. The respective
	// properties are then searched for the search terms and the results at the end are combined using WAND

	queryTextTerms, duplicateTextBoost := helpers.TokenizeTextAndCountDuplicates(fullQuery)
	queryStringTerms, duplicateStringBoost := helpers.TokenizeStringAndCountDuplicates(fullQuery)

	propertyNamesFullQuery := make([]string, 0)
	propertyNamesText := make([]string, 0)
	propertyNamesString := make([]string, 0)
	propertyBoosts := make(map[string]float32, len(properties))

	for _, propertyWithBoost := range properties {
		property := propertyWithBoost
		propBoost := 1
		if strings.Contains(propertyWithBoost, "^") {
			property = strings.Split(propertyWithBoost, "^")[0]
			boostStr := strings.Split(propertyWithBoost, "^")[1]
			propBoost, _ = strconv.Atoi(boostStr)
		}
		propertyBoosts[property] = float32(propBoost)

		prop, err := schema.GetPropertyByName(class, property)
		if err != nil {
			return nil, nil, err
		}

		if prop.Tokenization == "word" {
			if prop.DataType[0] == "text" {
				propertyNamesText = append(propertyNamesText, property)
			} else if prop.DataType[0] == "string" {
				propertyNamesString = append(propertyNamesString, property)
			} else {
				return nil, nil, fmt.Errorf("cannot handle datatype %v", prop.DataType[0])
			}
		} else {
			propertyNamesFullQuery = append(propertyNamesFullQuery, property)
		}
	}
	fullQueryLen := 0
	if len(propertyNamesFullQuery) > 0 {
		fullQueryLen = 1
	}
	results := make(terms, len(queryTextTerms)+len(queryStringTerms)+fullQueryLen)

	for i, queryTerm := range queryTextTerms {
		err := b.createTerm(&results[i], N, queryTerm, propertyNamesText, propertyBoosts, duplicateTextBoost[i])
		if err != nil {
			return nil, nil, err
		}
	}

	for i, queryTerm := range queryStringTerms {
		err := b.createTerm(&results[i+len(queryTextTerms)], N, queryTerm, propertyNamesString, propertyBoosts, duplicateStringBoost[i])
		if err != nil {
			return nil, nil, err
		}
	}
	if len(propertyNamesFullQuery) > 0 {
		err := b.createTerm(&results[len(queryTextTerms)+len(queryStringTerms)], N, fullQuery, propertyNamesFullQuery, propertyBoosts, 1)
		if err != nil {
			return nil, nil, err
		}
	}

	topKHeap := priorityqueue.NewMin(limit)
	worstDist := float64(0)
	for {

		results.pivot(worstDist)

		id, score, ok := results.scoreNext()
		if !ok {
			// nothing left to score
			break
		}

		if topKHeap.Len() < limit || topKHeap.Top().Dist < float32(score) {
			topKHeap.Insert(id, float32(score))
			for topKHeap.Len() > limit {
				topKHeap.Pop()
			}
		}
		worstDist = float64(topKHeap.Top().Dist)
	}

	objects := make([]*storobj.Object, 0, topKHeap.Len())
	scores := make([]float32, 0, topKHeap.Len())

	buf := make([]byte, 8)
	for topKHeap.Len() > 0 {
		res := topKHeap.Pop()
		scores = append(scores, res.Dist)
		binary.LittleEndian.PutUint64(buf, res.ID)
		objectByte, err := objectsBucket.GetBySecondary(0, buf)
		if err != nil {
			return nil, nil, err
		}

		obj, err := storobj.FromBinary(objectByte)
		if err != nil {
			return nil, nil, err
		}
		objects = append(objects, obj)
	}
	return objects, scores, nil
}

func (b *BM25Searcher) createTerm(termResult *term, N float64, query string, propertyNames []string, propertyBoosts map[string]float32, duplicateTextBoost int) error {
	var docMapPairs []docPointerWithScore = nil
	docMapPairsIndices := make(map[uint64]int, 0)
	termResult.k1 = b.config.K1
	termResult.b = b.config.B
	averagePropLength := float32(0.)
	for _, propName := range propertyNames {
		propMean, err := b.propLengths.PropertyMean(propName)
		averagePropLength += propMean

		bucket := b.store.Bucket(helpers.BucketFromPropNameLSM(propName))
		if bucket == nil {
			return fmt.Errorf("could not find bucket for property %v", propName)
		}
		m, err := bucket.MapList([]byte(query))
		if err != nil {
			return err
		}
		if len(m) == 0 {
			continue
		}

		if docMapPairs == nil {
			docMapPairs = make([]docPointerWithScore, 0, len(m))
			for k, val := range m {
				freqBits := binary.LittleEndian.Uint32(val.Value[0:4])
				propLenBits := binary.LittleEndian.Uint32(val.Value[4:8])
				docMapPairs = append(docMapPairs, docPointerWithScore{id: binary.BigEndian.Uint64(val.Key), frequency: float64(math.Float32frombits(freqBits) * propertyBoosts[propName]), propLength: float64(math.Float32frombits(propLenBits))})
				docMapPairsIndices[binary.BigEndian.Uint64(val.Key)] = k
			}
		} else {
			for k, val := range m {
				key := binary.BigEndian.Uint64(val.Key)
				ind, ok := docMapPairsIndices[key]
				freqBits := binary.LittleEndian.Uint32(val.Value[0:4])
				propLenBits := binary.LittleEndian.Uint32(val.Value[4:8])
				if ok {
					docMapPairs[ind].propLength += float64(math.Float32frombits(propLenBits))
					docMapPairs[ind].frequency += float64(math.Float32frombits(freqBits) * propertyBoosts[propName])
				} else {
					docMapPairs = append(docMapPairs, docPointerWithScore{id: binary.BigEndian.Uint64(val.Key), frequency: float64(math.Float32frombits(freqBits) * propertyBoosts[propName]), propLength: float64(math.Float32frombits(propLenBits))})
					docMapPairsIndices[binary.BigEndian.Uint64(val.Key)] = k
				}
			}
		}
	}
	if docMapPairs == nil {
		termResult.exhausted = true
		return nil
	}
	termResult.averagePropLength = float64(averagePropLength / float32(len(propertyNames)))
	termResult.data = docMapPairs

	n := float64(len(docMapPairs))
	termResult.idf = math.Log(float64(1)+(N-n+0.5)/(n+0.5)) * float64(duplicateTextBoost)

	termResult.posPointer = 0
	termResult.idPointer = termResult.data[0].id
	return nil
}

func (b *BM25Searcher) BM25F(ctx context.Context, className schema.ClassName, limit int,
	keywordRanking *searchparams.KeywordRanking,
	filter *filters.LocalFilter, sort []filters.Sort, additional additional.Properties,
	objectByIndexID func(index uint64) *storobj.Object,
) ([]*storobj.Object, []float32, error) {
	// WEAVIATE-471 - If a property is not searchable, return an error
	for _, property := range keywordRanking.Properties {
		if !schema.PropertyIsIndexed(b.schema.Objects, string(className), property) {
			return nil, nil, errors.New("Property " + property + " is not indexed.  Please choose another property or add an index to this property")
		}
	}
	class, err := schema.GetClassByName(b.schema.Objects, string(className))

	objs, scores, err := b.wand(ctx, class, keywordRanking.Query, keywordRanking.Properties, limit)
	if err != nil {
		return nil, nil, errors.Wrap(err, "wand")
	}

	return objs, scores, nil
}
