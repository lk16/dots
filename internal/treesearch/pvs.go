package treesearch

import (
	"github.com/lk16/dots/internal/othello"
	"sort"
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

	children := board.GetSortableChildren()

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

	if depth > 4 {
		for i := range children {
			children[i].Heur = pvs.Search(children[i].Board, MinHeuristic, MaxHeuristic, 2)
		}
		sort.Slice(children, func(i, j int) bool {
			return children[i].Heur > children[j].Heur
		})
	}

	for i, child := range children {

		var heur int
		if i == 0 {
			heur = -pvs.search(&child.Board, -beta, -alpha, depth-1)
		} else {
			heur = -pvs.searchNullWindow(&child.Board, -(alpha + 1), depth-1)
			if (alpha < heur) && (heur < beta) {
				heur = -pvs.search(&child.Board, -beta, -heur, depth-1)
			}
		}
		if heur >= beta {
			alpha = beta
			break
		}
		if heur > alpha {
			alpha = heur
		}

	}

	return alpha
}

func (pvs *Pvs) searchNullWindow(board *othello.Board, alpha, depth int) int {

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
		heur := -pvs.searchNullWindow(board, -(alpha + 1), depth)
		board.SwitchTurn()
		return heur
	}

	for _, child := range children {

		heur := -pvs.searchNullWindow(&child, -(alpha + 1), depth-1)

		if heur > alpha {
			return alpha + 1
		}

	}

	return alpha
}
