package treesearch

import (
	"github.com/lk16/dots/internal/othello"
)

// Pvs implements the principal variation search algorithm
type Pvs struct {
	stats Stats
}

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

// GetStats returns the statistics for the latest search
func (pvs Pvs) GetStats() Stats {
	return pvs.stats
}

// ResetStats resets the statistics for the latest search to zeroes
func (pvs *Pvs) ResetStats() {
	pvs.stats.Reset()
}

// Search searches for the the best move up to a certain depth
func (pvs *Pvs) Search(board othello.Board, alpha, beta, depth int) int {
	if depth >= board.CountEmpties() {
		depth = 60
	}

	pvs.stats.StartClock()
	heur := -pvs.search(&board, -beta, -alpha, depth)
	pvs.stats.StopClock()

	if heur < alpha {
		heur = alpha
	}

	if heur > beta {
		heur = beta
	}

	return heur
}

func (pvs *Pvs) search(board *othello.Board, alpha, beta, depth int) int {

	pvs.stats.Nodes++

	if depth == 0 {
		return FastHeuristic(*board)
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

	for i, child := range children {

		var score int
		if i == 0 {
			score = -pvs.search(&child, -beta, -alpha, depth-1)
		} else {
			score = -pvs.searchNullWindow(&child, -alpha-1, depth-1)
			if (alpha < score) && (score < beta) {
				score = -pvs.search(&child, -beta, -score, depth-1)
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

func (pvs *Pvs) searchNullWindow(board *othello.Board, alpha, depth int) int {

	beta := alpha + 1

	pvs.stats.Nodes++

	if depth == 0 {
		return FastHeuristic(*board)
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

	for i, child := range children {

		var score int
		if i == 0 {
			score = -pvs.search(&child, -beta, -alpha, depth-1)
		} else {
			score = -pvs.search(&child, -alpha-1, -alpha, depth-1)
			if (alpha < score) && (score < beta) {
				score = -pvs.search(&child, -beta, -score, depth-1)
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
