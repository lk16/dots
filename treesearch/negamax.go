package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type NegaMax struct {
	board othello.Board
	depth int
}

func NewNegaMax() *NegaMax {
	return &NegaMax{}
}

func (negamax *NegaMax) Name() string {
	return "negamax"
}

func (negamax *NegaMax) Search(board othello.Board, depth int) int {
	negamax.board = board
	negamax.depth = depth
	return negamax.search()
}

func (negamax *NegaMax) ExactSearch(board othello.Board) int {
	return negamax.Search(board, 60)
}

func (negamax *NegaMax) search() int {

	if negamax.depth == 0 {
		return heuristics.Squared(negamax.board)
	}

	gen := othello.NewGenerator(&negamax.board, 0)

	if !gen.HasMoves() {
		if negamax.board.OpponentMoves() == 0 {
			return ExactScoreFactor * negamax.board.ExactScore()
		}

		negamax.board.SwitchTurn()
		heur := -negamax.search()
		negamax.board.SwitchTurn()
		return heur
	}

	heur := MinHeuristic
	for gen.Next() {
		negamax.depth--
		childHeur := -negamax.search()
		negamax.depth++
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur
}
