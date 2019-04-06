package treesearch

import (
	"github.com/lk16/dots/internal/othello"
)

// Pvs implements the principal variation search algorithm
type Pvs struct{}

// NewPvs returns a new Pvs
func NewPvs() *Pvs {
	return &Pvs{}
}

// Name returns the name of the tree search algorithm
func (pvs Pvs) Name() string {
	return "pvs"
}

// ExactSearch searches for the best move without a depth limitation
func (pvs *Pvs) ExactSearch(board othello.Board, alpha, beta int) int {
	return pvs.Search(board, alpha, beta, 60) / ExactScoreFactor
}

// Search searches for the the best move up to a certain depth
func (pvs *Pvs) Search(board othello.Board, alpha, beta, depth int) int {
	if depth >= board.CountEmpties() {
		depth = 60
	}

	heur := -pvs.search(board, -beta, -alpha, depth)

	if heur < alpha {
		heur = alpha
	}

	if heur > beta {
		heur = beta
	}

	return heur
}

func (pvs *Pvs) search(board othello.Board, alpha, beta, depth int) int {

	if depth == 0 {
		return FastHeuristic(board)
	}

	children := board.GetChildren()

	if len(children) == 0 {

		if board.OpponentMoves() == 0 {
			heur := ExactScoreFactor * board.ExactScore()
			return heur
		}

		board.SwitchTurn()
		heur := -pvs.search(board, -beta, -alpha, depth)
		board.SwitchTurn()
		return heur
	}

	for i, it := range children {
		var score int
		if i == 0 {
			score = -pvs.search(it, -beta, -alpha, depth-1)
		} else {
			score = -pvs.search(it, -alpha-1, -alpha, depth-1)
			if (alpha < score) && (score < beta) {
				score = -pvs.search(it, -beta, -score, depth-1)
			}
		}
		if score >= beta {
			alpha = beta
			break
		}
		if score > alpha {
			alpha = score
		}

	}

	return alpha
}
