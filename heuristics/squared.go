package heuristics

import (
	"dots/othello"
	"math/bits"
)

// Squared is a heuristic taken from a similar project with that name
// see http://github.com/lk16/squared
func Squared(board othello.Board) int {
	cornerMask := uint64(1<<0 | 1<<7 | 1<<56 | 1<<63)

	meCorners := bits.OnesCount64(cornerMask & board.Me())
	oppCorners := bits.OnesCount64(cornerMask & board.Opp())
	cornerDiff := meCorners - oppCorners

	meMoves := bits.OnesCount64(board.Moves())
	oppMoves := bits.OnesCount64(board.OpponentMoves())
	moveDiff := meMoves - oppMoves

	return (3 * cornerDiff) + moveDiff
}
