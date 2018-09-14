package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type MiniMax struct {
}

func NewMinimax() *MiniMax {
	return &MiniMax{}
}

func (minimax *MiniMax) Name() string {
	return "minimax"
}

func (minimax *MiniMax) Search(board othello.Board, depth int) int {
	return minimax.search(board, depth, 1)
}

func (minimax *MiniMax) ExactSearch(board othello.Board) int {
	return minimax.search(board, 60, 1)
}

func (minimax *MiniMax) search(board othello.Board, depth int, sign int) int {

	if depth == 0 {
		return sign * heuristics.Squared(board)
	}

	children := board.GetChildren()

	if len(children) == 0 {
		if board.OpponentMoves() == 0 {
			return sign * ExactScoreFactor * board.ExactScore()
		}

		child := board
		child.SwitchTurn()
		return minimax.search(child, depth, -sign)
	}

	if sign == 1 {
		best := MinHeuristic
		for _, child := range children {
			heur := minimax.search(child, depth-1, -sign)
			if heur > best {
				best = heur
			}
		}
		return best
	}

	best := MaxHeuristic
	for _, child := range children {
		heur := minimax.search(child, depth-1, -sign)
		if heur < best {
			best = heur
		}
	}

	return best
}
