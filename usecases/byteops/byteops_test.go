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

package byteops

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const MaxUint32 = ^uint32(0)

func mustRandIntn(max int64) int {
	randInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(fmt.Sprintf("mustRandIntn error: %v", err))
	}
	return int(randInt.Int64())
}

// Create a buffer with space for several values and first write into it and then test that the values can be read again
func TestReadAnWrite(t *testing.T) {
	valuesNumbers := []uint64{234, 78, 23, 66, 8, 9, 2, 346745, 1}
	valuesByteArray := make([]byte, mustRandIntn(500))
	rand.Read(valuesByteArray)

	writeBuffer := make([]byte, 2*uint64Len+2*uint32Len+2*uint16Len+len(valuesByteArray))
	byteOpsWrite := NewReadWriter(writeBuffer)

	byteOpsWrite.WriteUint64(valuesNumbers[0])
	byteOpsWrite.WriteUint32(uint32(valuesNumbers[1]))
	byteOpsWrite.WriteUint32(uint32(valuesNumbers[2]))
	assert.Equal(t, byteOpsWrite.CopyBytesToBuffer(valuesByteArray), nil)
	byteOpsWrite.WriteUint16(uint16(valuesNumbers[3]))
	byteOpsWrite.WriteUint64(valuesNumbers[4])
	byteOpsWrite.WriteUint16(uint16(valuesNumbers[5]))

	byteOpsRead := NewReadWriter(writeBuffer)

	require.Equal(t, byteOpsRead.ReadUint64(), valuesNumbers[0])
	require.Equal(t, byteOpsRead.ReadUint32(), uint32(valuesNumbers[1]))
	require.Equal(t, byteOpsRead.ReadUint32(), uint32(valuesNumbers[2]))

	// we are going to do the next op twice (once with copying, once without)
	// to be able to rewind the buffer, let's cache the current position
	posBeforeByteArray := byteOpsRead.Position

	returnBuf, err := byteOpsRead.CopyBytesFromBuffer(uint64(len(valuesByteArray)), nil)
	assert.Equal(t, returnBuf, valuesByteArray)
	assert.Equal(t, err, nil)

	// rewind the buffer to where it was before the read
	byteOpsRead.MoveBufferToAbsolutePosition(posBeforeByteArray)

	subSlice := byteOpsRead.ReadBytesFromBuffer(uint64(len(valuesByteArray)))
	assert.Equal(t, subSlice, valuesByteArray)

	// now read again using the other method

	require.Equal(t, byteOpsRead.ReadUint16(), uint16(valuesNumbers[3]))
	require.Equal(t, byteOpsRead.ReadUint64(), valuesNumbers[4])
	require.Equal(t, byteOpsRead.ReadUint16(), uint16(valuesNumbers[5]))
}

// create buffer that is larger than uint32 and write to the end and then try to reread it
func TestReadAnWriteLargeBuffer(t *testing.T) {
	writeBuffer := make([]byte, uint64(MaxUint32)+4)
	byteOpsWrite := NewReadWriter(writeBuffer)
	byteOpsWrite.MoveBufferPositionForward(uint64(MaxUint32))
	byteOpsWrite.WriteUint16(uint16(10))

	byteOpsRead := NewReadWriter(writeBuffer)
	byteOpsRead.MoveBufferPositionForward(uint64(MaxUint32))
	require.Equal(t, byteOpsRead.ReadUint16(), uint16(10))
}

func TestWritingAndReadingBufferOfDynamicLength(t *testing.T) {
	t.Run("uint64 length indicator", func(t *testing.T) {
		bufLen := uint64(mustRandIntn(1024))
		buf := make([]byte, bufLen)
		rand.Read(buf)

		// uint64 length indicator + buffer + unrelated data at end of buffer
		totalBuf := make([]byte, bufLen+16)
		bo := NewReadWriter(totalBuf)

		assert.Nil(t, bo.CopyBytesToBufferWithUint64LengthIndicator(buf))
		bo.WriteUint64(17)
		assert.Equal(t, buf, totalBuf[8:8+bufLen])

		// read
		bo = NewReadWriter(totalBuf)
		bufRead := bo.ReadBytesFromBufferWithUint64LengthIndicator()
		assert.Len(t, bufRead, int(bufLen))
		assert.Equal(t, uint64(17), bo.ReadUint64())

		// discard
		bo = NewReadWriter(totalBuf)
		discarded := bo.DiscardBytesFromBufferWithUint64LengthIndicator()
		assert.Equal(t, bufLen, discarded)
		assert.Equal(t, uint64(17), bo.ReadUint64())
	})

	t.Run("uint32 length indicator", func(t *testing.T) {
		bufLen := uint32(mustRandIntn(1024))
		buf := make([]byte, bufLen)
		rand.Read(buf)

		// uint32 length indicator + buffer + unrelated data at end of buffer
		totalBuf := make([]byte, bufLen+8)
		bo := NewReadWriter(totalBuf)

		assert.Nil(t, bo.CopyBytesToBufferWithUint32LengthIndicator(buf))
		bo.WriteUint32(17)
		assert.Equal(t, buf, totalBuf[4:4+bufLen])

		// read
		bo = NewReadWriter(totalBuf)
		bufRead := bo.ReadBytesFromBufferWithUint32LengthIndicator()
		assert.Len(t, bufRead, int(bufLen))
		assert.Equal(t, uint32(17), bo.ReadUint32())

		// discard
		bo = NewReadWriter(totalBuf)
		discarded := bo.DiscardBytesFromBufferWithUint32LengthIndicator()
		assert.Equal(t, bufLen, discarded)
		assert.Equal(t, uint32(17), bo.ReadUint32())
	})
}

func TestStringsToByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		bytes := StringsToByteVector([]string{})
		assert.Equal(t, []byte{}, bytes)
	})

	t.Run("non-empty string", func(t *testing.T) {
		bytes := StringsToByteVector([]string{"hello", "world"})
		assert.Equal(t, []byte{0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64}, bytes)
	})
}

func TestStringsFromByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		strings := StringsFromByteVector([]byte{})
		assert.Equal(t, []string{}, strings)
	})

	t.Run("non-empty string", func(t *testing.T) {
		strings := StringsFromByteVector([]byte{0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x05, 0x77, 0x6f, 0x72, 0x6c, 0x64})
		assert.Equal(t, []string{"hello", "world"}, strings)
	})
}

func TestIntsToByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		bytes := IntsToByteVector([]float64{})
		assert.Equal(t, []byte{}, bytes)
	})

	t.Run("non-empty array of u8s", func(t *testing.T) {
		bytes := IntsToByteVector([]float64{1, 2, 3})
		assert.Equal(t, []byte{
			01, 00, 00, 00, 00, 00, 00, 00,
			02, 00, 00, 00, 00, 00, 00, 00,
			03, 00, 00, 00, 00, 00, 00, 00,
		}, bytes)
	})

	t.Run("non-empty array of u64", func(t *testing.T) {
		bytes := IntsToByteVector([]float64{
			9007199254740992, // MaxFloat64
			9007199254740991,
			9007199254740990,
		})
		assert.Equal(t, []byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1f, 0x00,
			0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1f, 0x00,
		}, bytes)
	})
}

func TestIntsFromByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		ints := IntsFromByteVector([]byte{})
		assert.Equal(t, []int64{}, ints)
	})

	t.Run("non-empty array of u8s", func(t *testing.T) {
		ints := IntsFromByteVector([]byte{
			01, 00, 00, 00, 00, 00, 00, 00,
			02, 00, 00, 00, 00, 00, 00, 00,
			03, 00, 00, 00, 00, 00, 00, 00,
		})
		assert.Equal(t, []int64{1, 2, 3}, ints)
	})

	t.Run("non-empty array of u64", func(t *testing.T) {
		ints := IntsFromByteVector([]byte{
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x00,
			0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1f, 0x00,
			0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1f, 0x00,
		})
		assert.Equal(t, []int64{
			9007199254740992, // MaxFloat64
			9007199254740991,
			9007199254740990,
		}, ints)
	})
}

func TestBoolsToByteVectors(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		bytes := BoolsToByteVector([]bool{})
		assert.Equal(t, []byte{}, bytes)
	})

	t.Run("non-empty array", func(t *testing.T) {
		bytes := BoolsToByteVector([]bool{true, false, true})
		assert.Equal(t, []byte{0x01, 0x00, 0x01}, bytes)
	})
}

func TestBoolsFromByteVectors(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		bools := BoolsFromByteVector([]byte{})
		assert.Equal(t, []bool{}, bools)
	})

	t.Run("non-empty array", func(t *testing.T) {
		bools := BoolsFromByteVector([]byte{0x01, 0x00, 0x01})
		assert.Equal(t, []bool{true, false, true}, bools)
	})
}

func TestFloat32ToByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		bytes := Float32ToByteVector([]float32{})
		assert.Equal(t, []byte{}, bytes)
	})

	t.Run("non-empty array", func(t *testing.T) {
		bytes := Float32ToByteVector([]float32{1.1, 2.2, 3.3})
		assert.Equal(t, []byte{
			0xcd, 0xcc, 0x8c, 0x3f,
			0xcd, 0xcc, 0xc, 0x40,
			0x33, 0x33, 0x53, 0x40,
		}, bytes)
	})
}

func TestFloat32FromByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		floats := Float32FromByteVector([]byte{})
		assert.Equal(t, []float32{}, floats)
	})

	t.Run("non-empty array", func(t *testing.T) {
		floats := Float32FromByteVector([]byte{
			0xcd, 0xcc, 0x8c, 0x3f,
			0xcd, 0xcc, 0xc, 0x40,
			0x33, 0x33, 0x53, 0x40,
		})
		assert.Equal(t, []float32{1.1, 2.2, 3.3}, floats)
	})
}

func TestFloat64ToByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		bytes := Float32ToByteVector([]float32{})
		assert.Equal(t, []byte{}, bytes)
	})

	t.Run("non-empty array", func(t *testing.T) {
		bytes := Float64ToByteVector([]float64{1.1, 2.2, 3.3})
		assert.Equal(t, []byte{
			0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xf1, 0x3f,
			0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0x1, 0x40,
			0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xa, 0x40,
		}, bytes)
	})
}

func TestFloat64FromByteVector(t *testing.T) {
	t.Run("empty array", func(t *testing.T) {
		floats := Float64FromByteVector([]byte{})
		assert.Equal(t, []float64{}, floats)
	})

	t.Run("non-empty array", func(t *testing.T) {
		floats := Float64FromByteVector([]byte{
			0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xf1, 0x3f,
			0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0x1, 0x40,
			0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xa, 0x40,
		})
		assert.Equal(t, []float64{1.1, 2.2, 3.3}, floats)
	})
}
