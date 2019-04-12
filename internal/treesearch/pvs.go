package treesearch

import (
	"github.com/lk16/dots/internal/othello"
	"sort"
)

// Pvs implements the principal variation search algorithm
type Pvs struct {
	stats   Stats
	sortPvs *Pvs
}

// NewPvs returns a new Pvs
func NewPvs() *Pvs {
	return &Pvs{
		sortPvs: &Pvs{}}
}

// Name returns the name of the tree search algorithm
func (pvs Pvs) Name() string {
	return "pvs"
}

// ExactSearch searches for the best move without a depth limitation
func (pvs *Pvs) ExactSearch(board othello.Board, alpha, beta int) int {
	return pvs.Search(board, alpha*ExactScoreFactor, beta*ExactScoreFactor, 60) / ExactScoreFactor
}

// GetStats returns the statistics for the latest search
func (pvs Pvs) GetStats() Stats {
	stats := Stats{}
	stats.Add(pvs.stats)
	if pvs.sortPvs != nil {
		stats.Add(pvs.sortPvs.GetStats())
	}
	return stats
}

// ResetStats resets the statistics for the latest search to zeroes
func (pvs *Pvs) ResetStats() {
	pvs.stats.Reset()
	if pvs.sortPvs != nil {
		pvs.sortPvs.ResetStats()
	}
}

// Search searches for the the best move up to a certain depth
func (pvs *Pvs) Search(board othello.Board, alpha, beta, depth int) int {

	pvs.stats.StartClock()

	var heur int
	if depth >= board.CountEmpties() {
		heur = -ExactScoreFactor * pvs.searchExact(&board, -beta, -alpha)
	} else {
		heur = -pvs.search(&board, -beta, -alpha, depth)
	}

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

	if depth <= 7 || pvs.sortPvs == nil || beta-alpha == 1 {
		return pvs.searchNoSort(board, alpha, beta, depth)
	}

	pvs.stats.Nodes++

	children := board.GetSortableChildren()

	if len(children) == 0 {

		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		return -pvs.search(board, -beta, -alpha, depth)
	}

	for i := range children {
		children[i].Heur = pvs.sortPvs.Search(children[i].Board, MinHeuristic, MaxHeuristic, 2)
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].Heur > children[j].Heur
	})

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
			return beta
		}
		if heur > alpha {
			alpha = heur
		}

	}

	return alpha
}

func (pvs *Pvs) searchNoSort(board *othello.Board, alpha, beta, depth int) int {

	pvs.stats.Nodes++

	if depth == 0 {
		return FastHeuristic(*board)
	}

	gen := othello.NewUnsortedChildGenerator(board)

	if !gen.HasMoves() {

		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -pvs.searchNoSort(board, -beta, -alpha, depth)
		board.SwitchTurn()
		return heur
	}

	for i := 0; gen.Next(); i++ {

		var heur int
		if i == 0 {
			heur = -pvs.searchNoSort(board, -beta, -alpha, depth-1)
		} else {
			heur = -pvs.searchNullWindowNoSort(board, -(alpha + 1), depth-1)
			if (alpha < heur) && (heur < beta) {
				heur = -pvs.searchNoSort(board, -beta, -heur, depth-1)
			}
		}
		if heur >= beta {
			gen.RestoreParent()
			return beta
		}
		if heur > alpha {
			alpha = heur
		}

	}

	return alpha
}

func (pvs *Pvs) searchNullWindow(board *othello.Board, alpha, depth int) int {

	if depth <= 7 || pvs.sortPvs == nil {
		return pvs.searchNullWindowNoSort(board, alpha, depth)
	}

	pvs.stats.Nodes++

	children := board.GetSortableChildren()

	if len(children) == 0 {

		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -pvs.searchNullWindow(board, -(alpha + 1), depth)
		board.SwitchTurn()
		return heur
	}

	for i := range children {
		children[i].Heur = pvs.sortPvs.Search(children[i].Board, MinHeuristic, MaxHeuristic, 2)
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].Heur > children[j].Heur
	})

	for _, child := range children {

		heur := -pvs.searchNullWindow(&child.Board, -(alpha + 1), depth-1)
		if heur > alpha {
			return alpha + 1
		}
	}

	return alpha
}

func (pvs *Pvs) searchNullWindowNoSort(board *othello.Board, alpha, depth int) int {

	pvs.stats.Nodes++

	if depth == 0 {
		return FastHeuristic(*board)
	}

	gen := othello.NewUnsortedChildGenerator(board)

	if !gen.HasMoves() {

		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -pvs.searchNullWindowNoSort(board, -(alpha + 1), depth)
		board.SwitchTurn()
		return heur
	}

	for i := 0; gen.Next(); i++ {

		heur := -pvs.searchNullWindowNoSort(board, -(alpha + 1), depth-1)
		if heur > alpha {
			gen.RestoreParent()
			return alpha + 1
		}
	}

	return alpha
}

func (pvs *Pvs) searchExact(board *othello.Board, alpha, beta int) int {

	pvs.stats.Nodes++

	children := board.GetSortableChildren()

	if len(children) == 0 {

		if board.OpponentMoves() == 0 {
			return board.ExactScore()
		}

		board.SwitchTurn()
		return -pvs.searchExact(board, -beta, -alpha)
	}

	/*for i := range children {
		children[i].Heur = pvs.sortPvs.Search(children[i].Board, MinHeuristic, MaxHeuristic, 2)
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].Heur > children[j].Heur
	})*/

	for i, child := range children {

		var heur int
		if i == 0 {
			heur = -pvs.searchExact(&child.Board, -beta, -alpha)
		} else {
			heur = -pvs.searchExactNullWindow(&child.Board, -(alpha + 1))
			if (alpha < heur) && (heur < beta) {
				heur = -pvs.searchExact(&child.Board, -beta, -heur)
			}
		}
		if heur >= beta {
			return beta
		}
		if heur > alpha {
			alpha = heur
		}

	}

	return alpha
}

func (pvs *Pvs) searchExactNullWindow(board *othello.Board, alpha int) int {

	pvs.stats.Nodes++

	children := board.GetSortableChildren()

	if len(children) == 0 {

		if board.OpponentMoves() == 0 {
			return board.ExactScore()
		}

		board.SwitchTurn()
		heur := -pvs.searchExactNullWindow(board, -(alpha + 1))
		board.SwitchTurn()
		return heur
	}

	/*for i := range children {
		children[i].Heur = pvs.sortPvs.Search(children[i].Board, MinHeuristic, MaxHeuristic, 2)
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].Heur > children[j].Heur
	})*/

	for _, child := range children {

		heur := -pvs.searchExactNullWindow(&child.Board, -(alpha + 1))
		if heur > alpha {
			return alpha + 1
		}
	}

	return alpha
}
