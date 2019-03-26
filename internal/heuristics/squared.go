package heuristics

import (
	"github.com/lk16/dots/internal/othello"
	"math/bits"
)

const (
	cornerMask = uint64(1<<0 | 1<<7 | 1<<56 | 1<<63)
)

func countCorners(bitset uint64) int {
	masked := bitset & cornerMask
	masked += (masked >> 56)
	masked += (masked >> 7)
	return int(masked)
}

// Squared is a heuristic taken from a similar project with that name
// see http://github.com/lk16/squared
func Squared(board othello.Board) int {

	meCorners := countCorners(board.Me())
	oppCorners := countCorners(board.Opp())
	cornerDiff := meCorners - oppCorners

	meMoves := bits.OnesCount64(board.Moves())
	oppMoves := bits.OnesCount64(board.OpponentMoves())
	moveDiff := meMoves - oppMoves

	return (3 * cornerDiff) + moveDiff
}
