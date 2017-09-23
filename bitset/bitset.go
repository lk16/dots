package bitset

import (
	"bytes"
	"io"
	"math/bits"
	"math/rand"
)

type Bitset uint64

// Returns a pseudo-random Bitset
func RandomBitset() (random Bitset) {
	random = Bitset(rand.Uint64())
	return
}

// Returns the number of set bits in a Bitset
func (bs Bitset) Count() (count uint) {
	return uint(bits.OnesCount64(uint64(bs)))
}

// Tests if the bit in a Bitset at index is set
func (bs Bitset) TestBit(index uint) bool {
	return (bs & (1 << index)) != 0
}

// Returns new bitset with only the first (least significant) bit set.
// Returns the 0 Bitset if the input Bitset is 0
func (bs Bitset) FirstBit() (first_bit Bitset) {
	first_bit = bs & -bs
	return
}

// Returns index of first (least significant) set bit in a Bitset
// Returns the 0 Bitset if the input Bitset is 0
func (bs Bitset) FirstBitIndex() (first_index uint) {
	first_index = uint(bits.TrailingZeros64(uint64(bs)))
	if first_index == 64 {
		return 0
	}
	return
}

// Returns last (most significant) set bit in a Bitset
// Returns 0 Bitset if the input Bitset is 0
func (bs Bitset) LastBit() (last_bit Bitset) {
	length := bits.Len64(uint64(bs))
	last_bit = Bitset(1) << uint(length-1)
	return
}

// Returns an Ascii-Art string representing a Bitset
func (bs Bitset) AsciiArt(writer io.Writer) {

	buffer := new(bytes.Buffer)
	buffer.WriteString("+-----------------+\n")

	for y := uint(0); y < 8; y++ {

		buffer.WriteString("| ")

		for x := uint(0); x < 8; x++ {

			f := y*8 + x

			if bs.TestBit(f) {
				buffer.WriteString("@ ")
			} else {
				buffer.WriteString("  ")
			}

		}

		buffer.WriteString("|\n")
	}
	buffer.WriteString("+-----------------+\n")

	writer.Write(buffer.Bytes())
}

// Sets a bit of a bitset
// Returns the modified bitset, to allow chaining
func (bs *Bitset) SetBit(index uint) (out *Bitset) {
	*bs |= Bitset(1 << index)
	out = bs
	return
}

// Resets a bit of a bitset
// Returns the modified bitset, to allow chaining
func (bs *Bitset) ResetBit(index uint) (out *Bitset) {
	*bs &^= Bitset(1 << index)
	out = bs
	return
}

func (bs *Bitset) Any() (any bool) {
	any = !bs.None()
	return
}

func (bs *Bitset) None() (none bool) {
	none = (*bs == 0)
	return
}
