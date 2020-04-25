package othello

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoMoveFaster(t *testing.T) {
	for board := range genTestBoards() {
		for i := uint(0); i < 64; i++ {

			if i == 27 || i == 28 || i == 35 || i == 36 {
				continue
			}

			moveBit := BitSet(1 << i)

			if (board.Me()|board.Opp())&moveBit != 0 {
				continue
			}

			child := board
			flipped := child.DoMove(moveBit)

			fasterChild := board
			fasterFlipped := fasterChild.DoMoveFaster(moveBit)

			if child != fasterChild || flipped != fasterFlipped {
				log.Printf("board:\n%s", board)

				log.Printf("child:\n%s", child)
				log.Printf("fasterChild:\n%s", fasterChild)

				log.Printf("flipped:\n%s", flipped)
				log.Printf("fasterFlipped:\n%s", fasterFlipped)

				t.FailNow()
			}
		}

	}
}

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

func BenchmarkDoMoveFaster(b *testing.B) {
	var boards []Board

	for i := 0; i < 6100; i++ {
		board, err := NewRandomBoard(4 + (i % 61))
		assert.Nil(b, err)

		boards = append(boards, *board)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 64; j++ {
			dummyBitSet = boards[i%len(boards)].DoMoveFaster(BitSet(1 << j))
		}
	}
}
