package othello

import (
	"math/rand"
	"testing"
)

func TestMostSignificantBit(t *testing.T) {

	trivialMsb := func(x uint64) uint64 {
		for i := 63; i >= 0; i-- {
			mask := uint64(1) << uint(i)
			if x&mask != 0 {
				return mask
			}
		}
		return 0
	}

	for i := 0; i < 64; i++ {
		input := uint64(1) << uint(i)
		expected := trivialMsb(input)
		got := MostSignificantBit(input)

		if expected != got {
			t.Errorf("For input 1 << %d == %d, expected: %d, got %d\n", i, input, expected, got)
		}
	}

	rand.Seed(0)
	for i := 0; i < 100000; i++ {
		input := rand.Uint64()
		expected := trivialMsb(input)
		got := MostSignificantBit(input)

		if expected != got {
			t.Errorf("For input 1 << %d == %d, expected: %d, got %d\n", i, input, expected, got)
		}
	}

}
