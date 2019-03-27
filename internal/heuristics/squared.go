package heuristics

import (
	"github.com/lk16/dots/internal/othello"
	"math/bits"
)

const (
	cSquareMaskVertical   = uint64(1<<1 | 1<<6 | 1<<55 | 1<<62)
	cSquareMaskHorizontal = uint64(1<<8 | 1<<15 | 1<<48 | 1<<53)
)

func countCsquares(bitset uint64) int {
	// TODO more efficient solution
	return bits.OnesCount64(bitset & (cSquareMaskHorizontal | cSquareMaskVertical))
}

// Squared is a heuristic taken from a similar project with that name
// see http://github.com/lk16/squared
func Squared(board othello.Board) int {

	cornerDiff := board.CornerCountDifference()

	meMoves := bits.OnesCount64(board.Moves())
	oppMoves := bits.OnesCount64(board.OpponentMoves())
	moveDiff := meMoves - oppMoves

	return (3 * cornerDiff) + moveDiff
}

func FastHeuristic(board othello.Board) int {

	me := board.Me()
	opp := board.Opp()

	heur := 0
	heur += 49 * board.CornerCountDifference()
	heur += -27 * board.XsquareCountDifference()
	heur += -17 * (countCsquares(me) - countCsquares(opp))
	heur += -11 * (bits.OnesCount64(me) - bits.OnesCount64(opp))
	return heur

}
