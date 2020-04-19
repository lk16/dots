package treesearch

import (
	"github.com/lk16/dots/internal/othello"
)

// Squared is a heuristic taken from a similar project with that name
// see http://github.com/lk16/squared
func Squared(board othello.Board) int {

	cornerDiff := board.CornerCountDifference()

	meMoves := board.Moves().Count()
	oppMoves := board.OpponentMoves().Count()
	moveDiff := meMoves - oppMoves

	return (3 * cornerDiff) + moveDiff
}

// FastHeuristic is way faster by not computing all possible moves
func FastHeuristic(board othello.Board) int {
	heur := 5 * board.CornerCountDifference()
	heur += 1 * board.PotentialMoveCountDifference()
	return heur
}
