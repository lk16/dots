package othello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardGetSortableChildren(t *testing.T) {

	for b := range genTestBoards() {

		// create copy to silence warnings
		board := b

		var expected []SortableBoard

		for _, child := range board.getChildren() {
			expected = append(expected, SortableBoard{Board: child})
		}

		discs := board.me | board.opp

		clone := board
		got := clone.GetSortableChildren()

		// board shouldn't change
		assert.Equal(t, board, clone)

		// children set should be matching expected children set
		assert.ElementsMatch(t, expected, got)

		for _, child := range got {

			childDiscs := child.Board.me | child.Board.opp

			// pieces shouldn't be removed
			assert.Equal(t, discs, childDiscs&discs)
		}
	}

}
