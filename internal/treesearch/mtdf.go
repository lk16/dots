package treesearch

import (
	"sync"

	"github.com/lk16/dots/internal/othello"
)

var childrenPool = sync.Pool{
	New: func() interface{} {
		return new([32]othello.SortableBoard)
	},
}

// Mtdf implements the mtdf tree search algorithm
type Mtdf struct {
	cache     Cacher
	stats     Stats
	heuristic func(othello.Board) int
}

// NewMtdf returns a new Mtdf
func NewMtdf(cache Cacher, heuristic func(othello.Board) int) *Mtdf {
	if cache == nil {
		cache = &NoOpCache{}
	}

	return &Mtdf{
		heuristic: heuristic,
		cache:     cache,
	}
}

// Name returns the tree search algorithm name
func (mtdf *Mtdf) Name() string {
	return "mtdf"
}

// GetStats returns the statistics for the latest search
func (mtdf Mtdf) GetStats() Stats {
	return mtdf.stats
}

// ResetStats resets the statistics for the latest search to zeroes
func (mtdf *Mtdf) ResetStats() {
	mtdf.stats.Reset()
}

// Search searches for the the best move up to a certain depth
func (mtdf *Mtdf) Search(board othello.Board, alpha, beta, depth int) int {
	if cache, ok := mtdf.cache.(*MemoryCache); ok {
		cache.Clear()
	}

	if depth >= board.CountEmpties() {
		depth = 60
	}

	mtdf.stats.StartClock()
	heuristic := slideWindow(&board, alpha, beta, depth)
	mtdf.stats.StopClock()
	return heuristic
}

// ExactSearch searches for the best move without a depth limitation
func (mtdf *Mtdf) ExactSearch(board othello.Board, alpha, beta int) int {
	return mtdf.Search(board, alpha*ExactScoreFactor, beta*ExactScoreFactor, 60) / ExactScoreFactor
}

func slideWindowExact(board *othello.Board, alpha, beta, depth int) int {
	// TODO
	var f int

	var step int
	if depth < board.CountEmpties() {
		f = FastHeuristic(*board)
		step = 1
	} else {
		f = 0
		step = 2 * ExactScoreFactor
	}

	// prevent odd results for exact search
	f -= f % step

	if f < alpha {
		f = alpha
	}

	if f > beta {
		f = beta
	}

	for beta-alpha >= step {
		var bound int

		if depth < board.CountEmpties() {
			bound = -nullWindow(board, -(f + 1), depth)
		} else {
			bound = -nullWindow(board, -(f + 1), depth)
		}

		if f == bound {
			f -= step
			beta = bound
		} else {
			f += step
			alpha = bound
		}
	}

	return beta
}

func slideWindow(board *othello.Board, alpha, beta, depth int) int {

	if board.CountEmpties() <= depth {
		return slideWindowExact(board, alpha, beta, depth)
	}

	f := FastHeuristic(*board)
	step := 1

	if f < alpha {
		f = alpha
	}

	if f > beta {
		f = beta
	}

	for beta-alpha >= step {
		bound := -nullWindow(board, -(f + 1), depth)

		if f == bound {
			f -= step
			beta = bound
		} else {
			f += step
			alpha = bound
		}
	}

	return beta
}

func nullWindow(board *othello.Board, alpha, depth int) int {
	if depth == 0 {
		heur := FastHeuristic(*board)
		if heur > alpha {
			return alpha + 1
		}
		return alpha
	}

	movesLeft := board.Moves()
	moveBit := othello.BitSet(0)
	flipped := othello.BitSet(0)

	if movesLeft == 0 {
		if board.OpponentMoves() == 0 {
			heur := ExactScoreFactor * board.ExactScore()
			if heur > alpha {
				return alpha + 1
			}
			return alpha
		}

		board.SwitchTurn()
		heur := -nullWindow(board, -(alpha + 1), depth)
		board.SwitchTurn()
		return heur
	}

	if depth > 5 {
		moves := movesLeft

		bestChildHeur := MinHeuristic
		var bestMoveBit othello.BitSet

		for movesLeft != 0 {
			moveBit = movesLeft & (-movesLeft)
			flipped = board.DoMove(moveBit)
			movesLeft &^= moveBit

			childHeur := slideWindow(board, bestChildHeur, MaxHeuristic, depth/4)
			if childHeur > bestChildHeur {
				bestChildHeur = childHeur
				bestMoveBit = moveBit
			}

			board.UndoMove(moveBit, flipped)
		}

		if bestChildHeur != MinHeuristic {
			flipped = board.DoMove(bestMoveBit)
			childHeur := -nullWindow(board, -(alpha + 1), depth-1)
			board.UndoMove(bestMoveBit, flipped)
			if childHeur > alpha {
				return alpha + 1
			}
		}

		movesLeft = moves
	}

	for movesLeft != 0 {
		moveBit = movesLeft & (-movesLeft)
		flipped = board.DoMove(moveBit)
		movesLeft &^= moveBit

		childHeur := -nullWindow(board, -(alpha + 1), depth-1)

		if childHeur > alpha {
			board.UndoMove(moveBit, flipped)
			return alpha + 1
		}

		board.UndoMove(moveBit, flipped)
	}

	return alpha
}
