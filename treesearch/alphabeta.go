package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type AlphaBeta struct {
	board othello.Board
	depth int
	alpha int
	beta  int
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
	alphabeta.board = board
	alphabeta.depth = depth
	return -alphabeta.search(alphabeta.beta, alphabeta.alpha)
}

func (alphabeta *AlphaBeta) ExactSearch(board othello.Board) int {
	return alphabeta.Search(board, 60)
}

func (alphabeta *AlphaBeta) search(alpha, beta int) int {

	if alphabeta.depth == 0 {
		return heuristics.Squared(alphabeta.board)
	}

	gen := othello.NewGenerator(&alphabeta.board, 0)

	if !gen.HasMoves() {
		if alphabeta.board.OpponentMoves() == 0 {
			return ExactScoreFactor * alphabeta.board.ExactScore()
		}

		alphabeta.board.SwitchTurn()
		heur := -alphabeta.search(-beta, -alpha)
		alphabeta.board.SwitchTurn()
		return heur
	}

	heur := alpha
	for gen.Next() {
		alphabeta.depth--
		childHeur := -alphabeta.search(-beta, -alpha)
		alphabeta.depth++
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
