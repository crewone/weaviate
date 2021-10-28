//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2021 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package inverted

import (
	"sync"
	"sync/atomic"

	"github.com/semi-technologies/weaviate/adapters/repos/db/helpers"
)

type RowCacher struct {
	maxSize     uint64
	rowStore    *sync.Map
	currentSize uint64
}

func NewRowCacher(maxSize uint64) *RowCacher {
	c := &RowCacher{
		maxSize:  maxSize,
		rowStore: &sync.Map{},
	}

	return c
}

type CacheEntry struct {
	Type      CacheEntryType
	Hash      []byte
	Partial   *docPointers
	AllowList helpers.AllowList
}

type CacheEntryType uint8

func (t CacheEntryType) String() string {
	switch t {
	case CacheTypePartial:
		return "partial"
	case CacheTypeAllowList:
		return "allow list"
	default:
		return "unknown"
	}
}

const (
	CacheTypePartial CacheEntryType = iota
	CacheTypeAllowList
)

func (rc *RowCacher) Store(id []byte, row *CacheEntry) {
	size := uint64(100) // TODO: calculate size
	if size > rc.maxSize {
		return
	}

	if atomic.LoadUint64(&rc.currentSize)+size > rc.maxSize {
		rc.deleteExistingEntries(size)
	}
	rc.rowStore.Store(string(id), row)
	atomic.AddUint64(&rc.currentSize, size)
}

func (rc *RowCacher) deleteExistingEntries(sizeToDelete uint64) {
	var deleted uint64
	rc.rowStore.Range(func(key, value interface{}) bool {
		parsed := value.(*docPointers)
		size := uint64(parsed.count * 4)
		rc.rowStore.Delete(key)
		deleted += size
		atomic.AddUint64(&rc.currentSize, -size)
		return deleted <= sizeToDelete
	})
}

func (rc *RowCacher) Load(id []byte) (*CacheEntry, bool) {
	retrieved, ok := rc.rowStore.Load(string(id))
	if !ok {
		return nil, false
	}

	parsed := retrieved.(*CacheEntry)
	// if !bytes.Equal(parsed.checksum, expectedChecksum) {
	// 	rc.rowStore.Delete(string(id))
	// 	return nil, false
	// }

	return parsed, true
}
