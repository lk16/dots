package bitset

import (
	"bytes"
	"testing"
)

func genTestBitsets() (ch chan Bitset) {
	ch = make(chan Bitset)
	go func() {
		ch <- Bitset(0)
		ch <- ^Bitset(0)
		for i := uint(0); i < 64; i++ {
			ch <- Bitset(1) << i
		}
		for i := uint(0); i < 64; i++ {
			ch <- ^(Bitset(1) << i)
		}
		for i := 0; i < 100000; i++ {
			ch <- RandomBitset()
		}
		close(ch)
	}()
	return
}

func (bs Bitset) count() (count uint) {
	for i := uint(0); i < 64; i++ {
		if bs&(1<<i) != 0 {
			count++
		}
	}
	return
}

func TestBitsetCount(t *testing.T) {
	for bs := range genTestBitsets() {
		expected := bs.count()
		clone := bs
		got := clone.Count()

		if bs != clone {
			t.Errorf("Bitset was modified: %d -> %d", bs, clone)
		}

		if expected != got {
			t.Errorf("Expected %d but got %d\n", expected, got)
		}
	}
}

func (bs Bitset) firstBit() Bitset {
	for mask := Bitset(1); mask != 0; mask <<= 1 {
		if bs&mask != 0 {
			return mask
		}
	}
	return 0
}

func TestBitsetFirstBit(t *testing.T) {
	for bs := range genTestBitsets() {
		expected := bs.firstBit()

		clone := bs
		got := bs.FirstBit()

		if bs != clone {
			t.Errorf("Bitset was modified: %d -> %d", bs, clone)
		}

		if expected != got {
			t.Errorf("Expected %d but got %d\n", expected, got)
		}
	}
}

func (bs Bitset) lastBit() Bitset {
	for mask := Bitset(1 << 63); mask != 0; mask >>= 1 {
		if bs&mask != 0 {
			return mask
		}
	}
	return 0
}

func TestBitsetLastBit(t *testing.T) {
	for bs := range genTestBitsets() {
		expected := bs.lastBit()

		clone := bs
		got := bs.LastBit()

		if bs != clone {
			t.Errorf("Bitset was modified: %d -> %d", bs, clone)
		}

		if expected != got {
			t.Errorf("Expected %d but got %d\n", expected, got)
		}
	}
}

func (bs Bitset) testBit(index uint) bool {
	return bs&Bitset(1<<index) != 0
}

func TestBitsetTestBit(t *testing.T) {
	for bs := range genTestBitsets() {
		clone := bs
		for index := uint(0); index < 64; index++ {
			expected := bs.testBit(index)
			got := bs.TestBit(index)

			if bs != clone {
				t.Errorf("Bitset was modified: %d -> %d", bs, clone)
			}

			if expected != got {
				t.Errorf("Expected %d but got %d\n", expected, got)
			}
		}
	}
}

func (bs Bitset) firstBitIndex() uint {
	for index := uint(0); index < 64; index++ {
		if bs&Bitset(1<<index) != 0 {
			return index
		}
	}
	return 0
}

func TestBitsetFirstBitIndex(t *testing.T) {
	for bs := range genTestBitsets() {
		expected := bs.firstBitIndex()

		clone := bs
		got := bs.FirstBitIndex()

		if bs != clone {
			t.Errorf("Bitset was modified: %d -> %d", bs, clone)
		}

		if expected != got {
			t.Errorf("Expected %d but got %d\n", expected, got)
		}
	}
}

func TestBitsetAsciiArt(t *testing.T) {
	for bs := range genTestBitsets() {

		expected_buff := new(bytes.Buffer)

		expected_buff.WriteString("+-----------------+\n")
		for y := uint(0); y < 8; y++ {
			expected_buff.WriteString("| ")

			for x := uint(0); x < 8; x++ {
				if bs.testBit(y*8 + x) {
					expected_buff.WriteString("@ ")
				} else {
					expected_buff.WriteString("  ")
				}
			}
			expected_buff.WriteString("|\n")
		}
		expected_buff.WriteString("+-----------------+\n")

		expected := expected_buff.String()

		clone := bs

		got_buff := new(bytes.Buffer)
		bs.AsciiArt(got_buff)
		got := got_buff.String()

		if bs != clone {
			t.Errorf("Bitset was modified: %d -> %d", bs, clone)
		}

		if expected != got {
			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n", expected, got)
		}

	}
}

func TestBitsetSetBit(t *testing.T) {
	for bs := range genTestBitsets() {
		for i := uint(0); i < 64; i++ {
			clone := bs
			clone.SetBit(i)
			expected := bs | Bitset(1<<i)
			if clone != expected {
				t.Errorf("Expected %d, got %d", expected, clone)
			}
		}
	}

	bs := Bitset(0)
	bs.SetBit(1).SetBit(2)
	expected := Bitset(1<<1) | Bitset(1<<2)

	if bs != expected {
		t.Errorf("Expected %d, got %d", expected, bs)
	}

}

func TestBitsetResetBit(t *testing.T) {
	for bs := range genTestBitsets() {
		for i := uint(0); i < 64; i++ {
			clone := bs
			clone.ResetBit(i)
			expected := bs &^ Bitset(1<<i)
			if clone != expected {
				t.Errorf("Expected %d, got %d", expected, clone)
			}
		}
	}

	bs := (^Bitset(0))
	bs.ResetBit(1).ResetBit(2)
	expected := (^Bitset(0)) &^ (Bitset(1<<1) | Bitset(1<<2))

	if bs != expected {
		t.Errorf("Expected %d, got %d", expected, bs)
	}

}

func TestBitsetAny(t *testing.T) {

	bs := Bitset(0)
	if bs.Any() {
		t.Errorf("Bitset(0).Any() should not be true")
	}

	bs = Bitset(1)
	if !bs.Any() {
		t.Errorf("Bitset(1).Any() should not be false")
	}
}

func TestBitsetNone(t *testing.T) {

	bs := Bitset(0)
	if !bs.None() {
		t.Errorf("Bitset(0).None() should be true")
	}

	bs = Bitset(1)
	if bs.None() {
		t.Errorf("Bitset(1).None() should be false")
	}
}
