package treesearch

import (
	"github.com/lk16/dots/internal/heuristics"
	"github.com/lk16/dots/internal/othello"
)

type NegaMax struct {
}

func NewNegaMax() *NegaMax {
	return &NegaMax{}
}

func (negamax *NegaMax) Name() string {
	return "negamax"
}

func (negamax *NegaMax) Search(board othello.Board, depth int) int {
	return -negamax.search(board, depth)
}

func (negamax *NegaMax) ExactSearch(board othello.Board) int {
	return negamax.Search(board, 60)
}

func (negamax *NegaMax) search(board othello.Board, depth int) int {

	if depth == 0 {
		return heuristics.Squared(board)
	}

	child := board
	gen := othello.NewUnsortedChildGenerator(&child)

	if !gen.HasMoves() {
		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		child.SwitchTurn()
		return -negamax.search(child, depth)
	}

	heur := MinHeuristic
	for gen.Next() {
		childHeur := -negamax.search(child, depth-1)
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur
}
