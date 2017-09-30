package board

import (
	"math/bits"
)

const (
	// MaxScore is the highest game result score possible
	MaxScore = 64

	// MinScore is the lowest game result score possible
	MinScore = -MaxScore

	// ExactScoreFactor is the multiplication.
	// This is used when a non exact search runs into an exact result
	ExactScoreFactor = 1000

	// MaxHeuristic is the highest heuristic value possible
	MaxHeuristic = ExactScoreFactor * MaxScore

	// MinHeuristic is the lowest heuristic value possible
	MinHeuristic = ExactScoreFactor * MinScore
)

// Negamax does a heuristic tree search using the Squared heuristic
func Negamax(board Board, depth int) int {

	if depth == 0 {
		return Squared(board)
	}

	children := board.GetChildren()

	if len(children) == 0 {
		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -Negamax(board, depth)
		board.SwitchTurn()
		return heur
	}

	heur := MinHeuristic
	for _, child := range children {
		childHeur := -Negamax(child, depth-1)
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur

}

// Squared is a heuristic taken from a similar project with that name
// see http://github.com/lk16/squared
func Squared(board Board) int {
	cornerMask := uint64(1<<0 | 1<<7 | 1<<56 | 1<<63)

	meCorners := bits.OnesCount64(cornerMask & board.Me())
	oppCorners := bits.OnesCount64(cornerMask & board.Opp())
	cornerDiff := meCorners - oppCorners

	meMoves := bits.OnesCount64(board.Moves())
	oppMoves := bits.OnesCount64(board.OpponentMoves())
	moveDiff := meMoves - oppMoves

	return (3 * cornerDiff) + moveDiff
}