package othello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardChildGenNext(t *testing.T) {

	for b := range genTestBoards() {

		// create copy to silence warnings
		board := b

		if !boardIsValid(&board) {
			continue
		}

		children := board.GetChildren()

		clone := board
		gen := NewUnsortedChildGenerator(&clone)

		var generatedChildren []Board

		for gen.Next() {
			generatedChildren = append(generatedChildren, clone)
		}

		// parent state should be restored
		assert.Equal(t, board, clone)

		assert.ElementsMatch(t, children, generatedChildren)
	}
}

func TestBoardChildGenRestoreParent(t *testing.T) {
	board := NewBoard()
	gen := NewUnsortedChildGenerator(board)
	gen.Next()
	gen.RestoreParent()

	assert.Equal(t, *NewBoard(), *board)
}

func TestBoardChildGenHasMoves(t *testing.T) {
	board := NewBoard()
	gen := NewUnsortedChildGenerator(board)

	// the start board should have moves
	assert.True(t, gen.HasMoves())

	board, err := NewRandomBoard(64)
	assert.Nil(t, err)

	gen = NewUnsortedChildGenerator(board)

	// full board should not have moves
	assert.False(t, gen.HasMoves())
}
