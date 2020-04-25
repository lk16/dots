package othello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dummyBitSet BitSet

func BenchmarkDoMove(b *testing.B) {
	var boards []Board

	for i := 0; i < 6100; i++ {
		board, err := NewRandomBoard(4 + (i % 61))
		assert.Nil(b, err)

		boards = append(boards, *board)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 64; j++ {
			dummyBitSet = boards[i%len(boards)].DoMove(BitSet(1 << j))
		}
	}
}
