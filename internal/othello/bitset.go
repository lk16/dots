package othello

import (
	"bytes"
	"fmt"
	"math/bits"
)

// BitSet contains 64 bits for efficient Board computations
type BitSet uint64

// String returns an ASCII-art string representation of a BitSet
func (bs BitSet) String() string {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	for y := 0; y < 8; y++ {
		_, _ = buffer.WriteString(fmt.Sprintf("%d ", y+1))
		for x := 0; x < 8; x++ {
			f := uint(y*8 + x)

			if bs.Test(f) {
				_, _ = buffer.WriteString("@ ")
			} else {
				_, _ = buffer.WriteString("- ")
			}
		}
		_, _ = buffer.WriteString("|\n")
	}
	_, _ = buffer.WriteString("+-----------------+\n")

	return buffer.String()
}

// Count returns the number of set bits
func (bs BitSet) Count() int {
	return bits.OnesCount64(uint64(bs))
}

// Len returns the minimum number of bits necessary to represent the value of a Bitset
// Result is 0 for bs == 0
func (bs BitSet) Len() int {
	return bits.Len64(uint64(bs))
}

// Set sets a bit on a given offset.
func (bs *BitSet) Set(offset uint) {
	if offset >= 64 {
		panic("bit offset out of range")
	}
	*bs |= BitSet(1) << offset
}

// Test returns whether a bit at a given offset is set
func (bs BitSet) Test(offset uint) bool {
	if offset >= 64 {
		panic("bit offset out of range")
	}
	mask := BitSet(1) << offset
	return bs&mask != 0
}

// Lowest returns the offset of the least significant bit
// Result is 64 for bs == 0
func (bs BitSet) Lowest() int {
	return bits.TrailingZeros64(uint64(bs))
}
