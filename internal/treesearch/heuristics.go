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

const cornerMask = othello.BitSet(1<<0 | 1<<7 | 1<<56 | 1<<63)

func potentialMoves(me, opp othello.BitSet) othello.BitSet {
	const (
		leftMask  = 0x7F7F7F7F7F7F7F7F
		rightMask = 0xFEFEFEFEFEFEFEFE
	)

	oppSurrounded := othello.BitSet(0)
	oppSurrounded |= (opp & leftMask) << 1
	oppSurrounded |= (opp & rightMask) >> 1
	oppSurrounded |= (opp & leftMask) << 9
	oppSurrounded |= (opp & rightMask) >> 9
	oppSurrounded |= (opp & rightMask) << 7
	oppSurrounded |= (opp & leftMask) >> 7

	oppSurrounded |= opp << 8
	oppSurrounded |= opp >> 8

	oppSurrounded &^= (me | opp)
	return oppSurrounded
}

// FastHeuristic is way faster by not computing all possible moves
func FastHeuristic(board othello.Board) int {
	heur := 5 * ((board.Me() & cornerMask).Count() - (board.Opp() & cornerMask).Count())
	mePotentialMoveCount := potentialMoves(board.Me(), board.Opp()).Count()
	oppPotentialMoveCount := potentialMoves(board.Opp(), board.Me()).Count()
	heur += 1 * (mePotentialMoveCount - oppPotentialMoveCount)
	return heur
}
