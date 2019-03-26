package heuristics

import (
	"github.com/lk16/dots/internal/othello"
	"math/bits"
)

const (
	cornerMask            = uint64(1<<0 | 1<<7 | 1<<56 | 1<<63)
	xSquareMask           = uint64(1<<9 | 1<<14 | 1<<49 | 1<<54)
	cSquareMaskVertical   = uint64(1<<1 | 1<<6 | 1<<55 | 1<<62)
	cSquareMaskHorizontal = uint64(1<<8 | 1<<15 | 1<<48 | 1<<53)
)

func countCorners(bitset uint64) int {
	masked := bitset & cornerMask
	masked += (masked >> 56)
	masked += (masked >> 7)
	masked &= 7
	return int(masked)
}

func countXsquares(bitset uint64) int {
	masked := bitset & xSquareMask
	masked += (masked >> 40)
	masked += (masked >> 5)
	masked = (masked >> 9) & 7
	return int(masked)
}

func countCsquares(bitset uint64) int {
	// TODO more efficient solution
	return bits.OnesCount64(bitset & (cSquareMaskHorizontal | cSquareMaskVertical))
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

func FastHeuristic(board othello.Board) int {

	me := board.Me()
	opp := board.Opp()

	heur := 0
	heur += 49 * (countCorners(me) - countCorners(opp))
	heur += -23 * (countXsquares(me) - countXsquares(opp))
	heur += -13 * (countCsquares(me) - countCsquares(opp))
	heur += -5 * (bits.OnesCount64(me) - bits.OnesCount64(opp))
	return heur

}
