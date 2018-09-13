package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type NegaMax struct {
	board othello.Board
}

func NewNegaMax() *NegaMax {
	return &NegaMax{}
}

func (negamax *NegaMax) Name() string {
	return "negamax"
}

func (negamax *NegaMax) Search(board othello.Board, depth int) int {
	return negamax.search(&board, depth)
}

func (negamax *NegaMax) ExactSearch(board othello.Board) int {
	return negamax.search(&board, 60)
}

func (negamax *NegaMax) search(board *othello.Board, depth int) int {

	if depth == 0 {
		return heuristics.Squared(*board)
	}

	gen := othello.NewGenerator(board, 0)

	if !gen.HasMoves() {
		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -negamax.search(board, depth)
		board.SwitchTurn()
		return heur
	}

	heur := MinHeuristic
	for gen.Next() {
		childHeur := -negamax.search(board, depth-1)
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur
}
