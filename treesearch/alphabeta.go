package treesearch

import (
	"dots/othello"
	"dots/players"
)

type AlphaBeta struct {
	board       othello.Board
	alpha, beta int
}

func NewAlphaBeta(board othello.Board, alpha, beta int) *AlphaBeta {
	return &AlphaBeta{
		board: board,
		alpha: alpha,
		beta: beta}
}

func (alphabeta *AlphaBeta) Search(depth int) int {
	return alphabeta.search(&alphabeta.board, depth, alphabeta.alpha, alphabeta.beta)
}

func (alphabeta *AlphaBeta) ExactSearch(board othello.Board) int {
	return alphabeta.search(&alphabeta.board, 60, alphabeta.alpha, alphabeta.beta)
}

func (alphabeta *AlphaBeta) search(board *othello.Board, alpha, beta, depth int) int {

	if depth == 0 {
		return players.Squared(*board)
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
