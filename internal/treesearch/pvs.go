package treesearch

import (
	"github.com/lk16/dots/internal/othello"
	"log"
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

func (pvs *Pvs) printf(format string, args ...interface{}) {
	if false {
		log.Printf(format, args...)
	}
}

func (pvs *Pvs) search(board othello.Board, alpha, beta, depth int) int {

	a := alpha

	if depth == 0 {
		pvs.printf("pvs search(board, alpha=%d, beta=%d, depth=%d) = %d <depth limit>", alpha, beta, depth,
			FastHeuristic(board))
		return FastHeuristic(board)
	}

	children := board.GetChildren()

	if len(children) == 0 {
		if board.OpponentMoves() == 0 {
			heur := ExactScoreFactor * board.ExactScore()
			pvs.printf("pvs search(board, alpha=%d, beta=%d, depth=%d) = %d <game end>", alpha, beta, depth, heur)
			return heur
		} else {
			board.SwitchTurn()
			heur := -pvs.search(board, -beta, -alpha, depth)
			board.SwitchTurn()
			pvs.printf("pvs search(board, alpha=%d, beta=%d, depth=%d) = %d <skipped turn>", alpha, beta, depth, heur)
			return heur
		}
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
	pvs.printf("pvs search(board, alpha=%d, beta=%d, depth=%d) = %d <move>", a, beta, depth, alpha)

	return alpha
}
