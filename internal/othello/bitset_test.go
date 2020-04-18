package othello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitSetString(t *testing.T) {

	expected := `+-a-b-c-d-e-f-g-h-+
1 - - - - - - - - |
2 @ @ @ @ @ @ @ @ |
3 - - - - - - - - |
4 @ @ @ @ @ @ @ @ |
5 - - - - - - - - |
6 @ @ @ @ @ @ @ @ |
7 - - - - - - - - |
8 @ @ @ @ @ @ @ @ |
+-----------------+
`

	got := BitSet(0xFF00FF00FF00FF00).String()

	assert.Equal(t, expected, got)

	expected = `+-a-b-c-d-e-f-g-h-+
1 @ @ @ @ @ @ @ @ |
2 @ - - - - - - @ |
3 @ - - - - - - @ |
4 @ - - - - - - @ |
5 @ - - - - - - @ |
6 @ - - - - - - @ |
7 @ - - - - - - @ |
8 @ @ @ @ @ @ @ @ |
+-----------------+
`
	got = BitSet(0xFF818181818181FF).String()

	assert.Equal(t, expected, got)

}

func TestBitSetCount(t *testing.T) {
	assert.Equal(t, 0, BitSet(0).Count())
	assert.Equal(t, 4, BitSet(0xF).Count())
	assert.Equal(t, 64, BitSet(0xFFFFFFFFFFFFFFFF).Count())
}

func TestBitSetLen(t *testing.T) {
	assert.Equal(t, 0, BitSet(0).Len())
	assert.Equal(t, 4, BitSet(0x8).Len())
	assert.Equal(t, 64, BitSet(0xFFFFFFFFFFFFFFFF).Len())
}

func TestBitSetSet(t *testing.T) {

	bs := BitSet(0)
	bs.Set(3)

	assert.Equal(t, BitSet(0x8), bs)

	bs = BitSet(0x8)
	bs.Set(3)
	assert.Equal(t, BitSet(0x8), bs)

	assert.Panics(t, func() {
		bs := BitSet(0)
		bs.Set(64)
	})
}

func TestBitSetTest(t *testing.T) {
	bs := BitSet(0x8)
	assert.True(t, bs.Test(3))
	assert.False(t, bs.Test(2))

	assert.Panics(t, func() {
		bs := BitSet(0)
		bs.Test(64)
	})
}

func TestBitSetLowest(t *testing.T) {
	assert.Equal(t, 3, BitSet(0x8).Lowest())
	assert.Equal(t, 0, BitSet(0x7).Lowest())
}
