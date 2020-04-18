package othello

import (
	"bytes"
	"fmt"
	"log"
	"math/bits"
)

// BitSet contains 64 bits for efficient Board computations
type BitSet uint64

// String returns an ASCII-art string representation of a BitSet
func (bs BitSet) String() string {
	var buffer bytes.Buffer
	_, _ = buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	for y := uint(0); y < 8; y++ {
		_, _ = buffer.WriteString(fmt.Sprintf("%d ", y+1))

		for x := uint(0); x < 8; x++ {
			f := y*8 + x
			if uint64(bs)&uint64(1<<f) != 0 {
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
func (bs *BitSet) Set(offset int) {
	if offset < 0 || offset >= 64 {
		log.Fatalf("BitSet.Set() called with offset %d.\n", offset)
	}
	*bs |= BitSet(1) << uint(offset)
}

// Test returns whether a bit at a given offset is set
func (bs BitSet) Test(offset int) bool {
	if offset < 0 || offset >= 64 {
		log.Fatalf("BitSet.Test() called with offset %d.\n", offset)
	}
	mask := BitSet(1) << uint(offset)
	return bs&mask != 0
}

// Lowest returns the offset of the least significant bit
// Result is 64 for bs == 0
func (bs BitSet) Lowest() int {
	return bits.TrailingZeros64(uint64(bs))
}

var doMoveToLowerLookup = [64]BitSet{
	// no bits are flipped if player to move has no discs on the inspected line
	^(BitSet(0)),
	(BitSet(1) << 1) - 1,
	(BitSet(1) << 2) - 1,
	(BitSet(1) << 3) - 1,
	(BitSet(1) << 4) - 1,
	(BitSet(1) << 5) - 1,
	(BitSet(1) << 6) - 1,
	(BitSet(1) << 7) - 1,
	(BitSet(1) << 8) - 1,
	(BitSet(1) << 9) - 1,
	(BitSet(1) << 10) - 1,
	(BitSet(1) << 11) - 1,
	(BitSet(1) << 12) - 1,
	(BitSet(1) << 13) - 1,
	(BitSet(1) << 14) - 1,
	(BitSet(1) << 15) - 1,
	(BitSet(1) << 16) - 1,
	(BitSet(1) << 17) - 1,
	(BitSet(1) << 18) - 1,
	(BitSet(1) << 19) - 1,
	(BitSet(1) << 20) - 1,
	(BitSet(1) << 21) - 1,
	(BitSet(1) << 22) - 1,
	(BitSet(1) << 23) - 1,
	(BitSet(1) << 24) - 1,
	(BitSet(1) << 25) - 1,
	(BitSet(1) << 26) - 1,
	(BitSet(1) << 27) - 1,
	(BitSet(1) << 28) - 1,
	(BitSet(1) << 29) - 1,
	(BitSet(1) << 30) - 1,
	(BitSet(1) << 31) - 1,
	(BitSet(1) << 32) - 1,
	(BitSet(1) << 33) - 1,
	(BitSet(1) << 34) - 1,
	(BitSet(1) << 35) - 1,
	(BitSet(1) << 36) - 1,
	(BitSet(1) << 37) - 1,
	(BitSet(1) << 38) - 1,
	(BitSet(1) << 39) - 1,
	(BitSet(1) << 40) - 1,
	(BitSet(1) << 41) - 1,
	(BitSet(1) << 42) - 1,
	(BitSet(1) << 43) - 1,
	(BitSet(1) << 44) - 1,
	(BitSet(1) << 45) - 1,
	(BitSet(1) << 46) - 1,
	(BitSet(1) << 47) - 1,
	(BitSet(1) << 48) - 1,
	(BitSet(1) << 49) - 1,
	(BitSet(1) << 50) - 1,
	(BitSet(1) << 51) - 1,
	(BitSet(1) << 52) - 1,
	(BitSet(1) << 53) - 1,
	(BitSet(1) << 54) - 1,
	(BitSet(1) << 55) - 1,
	(BitSet(1) << 56) - 1,
	(BitSet(1) << 57) - 1,
	(BitSet(1) << 58) - 1,
	(BitSet(1) << 59) - 1,
	(BitSet(1) << 60) - 1,
	(BitSet(1) << 61) - 1,
	(BitSet(1) << 62) - 1,
	(BitSet(1) << 63) - 1}
