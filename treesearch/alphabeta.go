package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type AlphaBeta struct {
	alpha, beta int
}

func NewAlphaBeta(alpha, beta int) *AlphaBeta {
	return &AlphaBeta{
		alpha: alpha,
		beta:  beta}
}

func (alphabeta *AlphaBeta) Name() string {
	return "alphabeta"
}

func (alphabeta *AlphaBeta) Search(board othello.Board, depth int) int {
	return alphabeta.search(&board, alphabeta.alpha, alphabeta.beta, depth)
}

func (alphabeta *AlphaBeta) ExactSearch(board othello.Board) int {
	return alphabeta.search(&board, 60, alphabeta.alpha, alphabeta.beta)
}

func (alphabeta *AlphaBeta) search(board *othello.Board, alpha, beta, depth int) int {

	if depth == 0 {
		return heuristics.Squared(*board)
	}

	gen := othello.NewGenerator(board, 0)

	if !gen.HasMoves() {
		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -alphabeta.search(board, -beta, -alpha, depth)
		board.SwitchTurn()
		return heur
	}

	heur := alpha
	for gen.Next() {
		childHeur := -alphabeta.search(board, -beta, -alpha, depth-1)
		if childHeur >= beta {
			gen.RestoreParent()
			return beta
		}
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur
}
