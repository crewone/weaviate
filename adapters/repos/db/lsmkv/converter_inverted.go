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
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"math"

	"github.com/pkg/errors"
	"github.com/weaviate/sroar"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv/segmentindex"
	"github.com/weaviate/weaviate/adapters/repos/db/lsmkv/varenc"
)

type convertedInverted struct {
	c1 *segmentCursorMap

	// the level matching those of the cursors
	currentLevel        uint16
	secondaryIndexCount uint16
	// Tells if tombstones or keys without corresponding values
	// can be removed from merged segment.
	// (left segment is root (1st) one, keepTombstones is off for bucket)
	cleanupTombstones bool

	w    io.WriteSeeker
	bufw *bufio.Writer

	scratchSpacePath string

	offset int

	tombstonesToWrite *sroar.Bitmap
	tombstonesToClean *sroar.Bitmap

	propertyLengthsToWrite map[uint64]uint32
	propertyLengthsCount   uint64
	propertyLengthsSum     uint64

	invertedHeader *segmentindex.HeaderInverted
}

func newConvertedInverted(w io.WriteSeeker,
	c1 *segmentCursorMap, level, secondaryIndexCount uint16,
	scratchSpacePath string, cleanupTombstones bool, allTombstones *sroar.Bitmap,
) *convertedInverted {
	return &convertedInverted{
		c1:                  c1,
		w:                   w,
		bufw:                bufio.NewWriterSize(w, 256*1024),
		currentLevel:        level,
		cleanupTombstones:   cleanupTombstones,
		tombstonesToClean:   allTombstones,
		secondaryIndexCount: secondaryIndexCount,
		scratchSpacePath:    scratchSpacePath,
		offset:              0,
	}
}

func (c *convertedInverted) do() error {
	var err error

	if err := c.init(); err != nil {
		return errors.Wrap(err, "init")
	}

	c.tombstonesToWrite = sroar.NewBitmap()
	c.propertyLengthsToWrite = make(map[uint64]uint32)

	kis, err := c.writeKeys()
	if err != nil {
		return errors.Wrap(err, "write keys")
	}

	tombstones := sroar.Or(c.tombstonesToWrite, c.tombstonesToClean)
	tombstoneOffset := c.offset

	_, err = c.writeTombstones(tombstones)
	if err != nil {
		return errors.Wrap(err, "write tombstones")
	}

	propertyLengthsOffset := c.offset
	_, err = c.writePropertyLengths(c.propertyLengthsToWrite)
	if err != nil {
		return errors.Wrap(err, "write property lengths")
	}
	treeOffset := uint64(c.offset)
	if err := c.writeIndices(kis); err != nil {
		return errors.Wrap(err, "write index")
	}

	// flush buffered, so we can safely seek on underlying writer
	if err := c.bufw.Flush(); err != nil {
		return errors.Wrap(err, "flush buffered")
	}
	if err := c.writeHeader(c.currentLevel, 0, c.secondaryIndexCount,
		treeOffset); err != nil {
		return errors.Wrap(err, "write header")
	}

	if err := c.writeInvertedHeader(tombstoneOffset, propertyLengthsOffset); err != nil {
		return errors.Wrap(err, "write keys length")
	}

	return nil
}

func (c *convertedInverted) init() error {
	// write a dummy header, we don't know the contents of the actual header yet,
	// we will seek to the beginning and overwrite the actual header at the very
	// end
	if _, err := c.bufw.Write(make([]byte, segmentindex.HeaderSize)); err != nil {
		return errors.Wrap(err, "write empty header")
	}

	dataFields := []varenc.VarEncDataType{varenc.DeltaVarIntUint64, varenc.VarIntUint64}
	if _, err := c.bufw.Write(make([]byte, segmentindex.SegmentInvertedDefaultHeaderSize+len(dataFields))); err != nil {
		return errors.Wrap(err, "write empty inverted header")
	}
	c.offset = segmentindex.HeaderSize + segmentindex.SegmentInvertedDefaultHeaderSize + len(dataFields)

	c.invertedHeader = &segmentindex.HeaderInverted{
		KeysOffset:            uint64(c.offset),
		TombstoneOffset:       0,
		PropertyLengthsOffset: 0,
		Version:               0,
		BlockSize:             uint8(segmentindex.SegmentInvertedDefaultBlockSize),
		DataFieldCount:        uint8(len(dataFields)),
		DataFields:            dataFields,
	}

	return nil
}

func (c *convertedInverted) writeTombstones(tombstones *sroar.Bitmap) (int, error) {
	tombstonesBuffer := make([]byte, 0)

	if tombstones != nil && tombstones.GetCardinality() > 0 {
		tombstonesBuffer = tombstones.ToBuffer()
	}

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(len(tombstonesBuffer)))
	if _, err := c.bufw.Write(buf); err != nil {
		return 0, err
	}

	if _, err := c.bufw.Write(tombstonesBuffer); err != nil {
		return 0, err
	}
	c.offset += len(tombstonesBuffer) + 8
	return len(tombstonesBuffer) + 8, nil
}

func (c *convertedInverted) writePropertyLengths(propLengths map[uint64]uint32) (int, error) {
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(propLengths)
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 8)

	binary.LittleEndian.PutUint64(buf, c.propertyLengthsSum)
	if _, err := c.bufw.Write(buf); err != nil {
		return 0, err
	}

	binary.LittleEndian.PutUint64(buf, c.propertyLengthsCount)
	if _, err := c.bufw.Write(buf); err != nil {
		return 0, err
	}

	binary.LittleEndian.PutUint64(buf, uint64(b.Len()))
	if _, err := c.bufw.Write(buf); err != nil {
		return 0, err
	}

	if _, err := c.bufw.Write(b.Bytes()); err != nil {
		return 0, err
	}
	c.offset += b.Len() + 8 + 8 + 8
	return b.Len() + 8 + 8 + 8, nil
}

func (c *convertedInverted) writeKeys() ([]segmentindex.Key, error) {
	key1, value1, _ := c.c1.first()

	// the (dummy) header was already written, this is our initial offset

	var kis []segmentindex.Key

	for {
		if key1 == nil {
			break
		}

		if values, skip := c.cleanupValues(value1); !skip {
			ki, err := c.writeIndividualNode(c.offset, key1, values, c.propertyLengthsToWrite)
			if err != nil {
				return nil, errors.Wrap(err, "write individual node (key1 smaller)")
			}

			c.offset = ki.ValueEnd
			kis = append(kis, ki)
		}
		key1, value1, _ = c.c1.next()

	}

	return kis, nil
}

func (c *convertedInverted) writeIndividualNode(offset int, key []byte,
	values []MapPair, propertyLengths map[uint64]uint32,
) (segmentindex.Key, error) {
	// NOTE: There are no guarantees in the cursor logic that any memory is valid
	// for more than a single iteration. Every time you call next() to advance
	// the cursor, any memory might be reused.
	//
	// This includes the key buffer which was the cause of
	// https://github.com/weaviate/weaviate/issues/3517
	//
	// A previous logic created a new assignment in each iteration, but thatwas
	// not an explicit guarantee. A change in v1.21 (for pread/mmap) added a
	// reusable buffer for the key which surfaced this bug.
	keyCopy := make([]byte, len(key))
	copy(keyCopy, key)

	return segmentInvertedNode{
		values:      values,
		primaryKey:  keyCopy,
		offset:      offset,
		propLengths: propertyLengths,
	}.KeyIndexAndWriteTo(c.bufw)
}

func (c *convertedInverted) writeIndices(keys []segmentindex.Key) error {
	indices := segmentindex.Indexes{
		Keys:                keys,
		SecondaryIndexCount: c.secondaryIndexCount,
		ScratchSpacePath:    c.scratchSpacePath,
	}

	_, err := indices.WriteTo(c.bufw)
	return err
}

// writeHeader assumes that everything has been written to the underlying
// writer and it is now safe to seek to the beginning and override the initial
// header
func (c *convertedInverted) writeHeader(level, version, secondaryIndices uint16,
	startOfIndex uint64,
) error {
	if _, err := c.w.Seek(0, io.SeekStart); err != nil {
		return errors.Wrap(err, "seek to beginning to write header")
	}

	h := &segmentindex.Header{
		Level:            level,
		Version:          version,
		SecondaryIndices: secondaryIndices,
		Strategy:         segmentindex.StrategyInverted,
		IndexStart:       startOfIndex,
	}

	if _, err := h.WriteTo(c.w); err != nil {
		return err
	}

	return nil
}

// writeInvertedHeader assumes that everything has been written to the underlying
// writer and it is now safe to seek to the beginning and override the initial
// header
func (c *convertedInverted) writeInvertedHeader(tombstoneOffset, propertyLengthsOffset int) error {
	if _, err := c.w.Seek(16, io.SeekStart); err != nil {
		return errors.Wrap(err, "seek to beginning to write inverted header")
	}

	c.invertedHeader.TombstoneOffset = uint64(tombstoneOffset)
	c.invertedHeader.PropertyLengthsOffset = uint64(propertyLengthsOffset)

	if _, err := c.invertedHeader.WriteTo(c.w); err != nil {
		return err
	}

	return nil
}

// Removes values with tombstone set from input slice. Output slice may be smaller than input one.
// Returned skip of true means there are no values left (key can be omitted in segment)
// WARN: method can alter input slice by swapping its elements and reducing length (not capacity)
func (c *convertedInverted) cleanupValues(values []MapPair) (vals []MapPair, skip bool) {
	// Reuse input slice not to allocate new memory
	// Rearrange slice in a way that tombstoned values are moved to the end
	// and reduce slice's length.
	last := 0
	propLength := uint32(0)
	for i := 0; i < len(values); i++ {
		docId := binary.BigEndian.Uint64(values[i].Key)
		if values[i].Tombstone {
			c.tombstonesToWrite.Set(docId)
		} else if !(c.tombstonesToClean != nil && c.tombstonesToClean.Contains(docId)) {

			// if docId in propertyLengthsToWrite
			// then add propertyLengthsToWrite[docId] to propertyLengthsSum
			if _, ok := c.propertyLengthsToWrite[docId]; !ok {
				propLength = uint32(math.Float32frombits(binary.LittleEndian.Uint32(values[i].Value[4:])))
				c.propertyLengthsToWrite[docId] = propLength
				c.propertyLengthsCount++
				c.propertyLengthsSum += uint64(propLength)
			}
			values[last], values[i] = values[i], values[last]
			last++
		}

	}

	if last == 0 {
		return nil, true
	}
	return values[:last], false
}
