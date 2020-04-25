package treesearch

import (
	"math/bits"

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

const (
	leftMask   = 0x7F7F7F7F7F7F7F7F
	rightMask  = 0xFEFEFEFEFEFEFEFE
	cornerMask = othello.BitSet(1<<0 | 1<<7 | 1<<56 | 1<<63)
)

// FastHeuristic is way faster by not computing all possible moves
func FastHeuristic(board othello.Board) int {
	me := board.Me()
	opp := board.Opp()

	heur := 5 * (bits.OnesCount64(uint64(me&cornerMask)) - bits.OnesCount64(uint64(opp&cornerMask)))

	mePotentialMoves := (opp & leftMask) << 1
	mePotentialMoves |= (opp & rightMask) >> 1
	mePotentialMoves |= (opp & leftMask) << 9
	mePotentialMoves |= (opp & rightMask) >> 9
	mePotentialMoves |= (opp & rightMask) << 7
	mePotentialMoves |= (opp & leftMask) >> 7
	mePotentialMoves |= opp << 8
	mePotentialMoves |= opp >> 8

	mePotentialMoves &^= (me | opp)

	oppPotentialMoves := (me & leftMask) << 1
	oppPotentialMoves |= (me & rightMask) >> 1
	oppPotentialMoves |= (me & leftMask) << 9
	oppPotentialMoves |= (me & rightMask) >> 9
	oppPotentialMoves |= (me & rightMask) << 7
	oppPotentialMoves |= (me & leftMask) >> 7
	oppPotentialMoves |= me << 8
	oppPotentialMoves |= me >> 8

	oppPotentialMoves &^= (me | opp)

	mePotentialMoveCount := bits.OnesCount64(uint64(mePotentialMoves))
	oppPotentialMoveCount := bits.OnesCount64(uint64(oppPotentialMoves))

	heur += 1 * (mePotentialMoveCount - oppPotentialMoveCount)
	return heur
}
