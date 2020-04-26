package treesearch

import (
	"log"
	"sort"

	"github.com/lk16/dots/internal/othello"
)

const (
	minHashtableDepth = 5
)

// Mtdf implements the mtdf tree search algorithm
type Mtdf struct {
	board     othello.Board
	high      int
	low       int
	depth     int
	cache     Cacher
	stats     Stats
	heuristic func(othello.Board) int
	sorter    Pvs
}

// NewMtdf returns a new Mtdf
func NewMtdf(cache Cacher, heuristic func(othello.Board) int) *Mtdf {
	if cache == nil {
		cache = &NoOpCache{}
	}

	return &Mtdf{
		heuristic: heuristic,
		cache:     cache,
		sorter:    *NewPvs(nil, heuristic),
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
	mtdf.low = alpha
	mtdf.high = beta

	if cache, ok := mtdf.cache.(*MemoryCache); ok {
		cache.Clear()
	}

	if depth >= board.CountEmpties() {
		depth = 60
	}

	mtdf.board = board
	mtdf.stats.StartClock()
	mtdf.depth = depth
	heuristic := mtdf.slideWindow()
	mtdf.stats.StopClock()
	return heuristic
}

// ExactSearch searches for the best move without a depth limitation
func (mtdf *Mtdf) ExactSearch(board othello.Board, alpha, beta int) int {
	return mtdf.Search(board, alpha*ExactScoreFactor, beta*ExactScoreFactor, 60) / ExactScoreFactor
}

func (mtdf *Mtdf) slideWindow() int {
	var f int

	var step int
	if mtdf.depth < mtdf.board.CountEmpties() {
		f = FastHeuristic(mtdf.board)
		step = 1
	} else {
		f = 0
		step = 2 * ExactScoreFactor
	}

	// prevent odd results for exact search
	f -= f % step

	if f < mtdf.low {
		f = mtdf.low
	}

	if f > mtdf.high {
		f = mtdf.high
	}

	for mtdf.high-mtdf.low >= step {
		var bound int

		if mtdf.depth < mtdf.board.CountEmpties() {
			bound = -mtdf.search(-(f + 1))
		} else {
			bound = -mtdf.searchNoHashtable(-(f + 1))
		}

		if f == bound {
			f -= step
			mtdf.high = bound
		} else {
			f += step
			mtdf.low = bound
		}
	}

	return mtdf.high
}

func (mtdf *Mtdf) handleNoMoves(alpha int) int {
	if mtdf.board.OpponentMoves() == 0 {
		heur := ExactScoreFactor * mtdf.board.ExactScore()
		if heur > alpha {
			return alpha + 1
		}
		return alpha
	}

	mtdf.board.SwitchTurn()
	heur := -mtdf.search(-(alpha + 1))
	mtdf.board.SwitchTurn()
	return heur
}

func searchForSort(board *othello.Board, alpha, beta, depth int) int {

	if depth == 0 {
		return FastHeuristic(*board)
	}

	gen := othello.NewChildGenerator(board)

	if !gen.HasMoves() {
		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -searchForSort(board, -beta, -alpha, depth)
		board.SwitchTurn()
		return heur
	}

	for i := 0; gen.Next(); i++ {
		heur := -searchForSort(board, -beta, -alpha, depth-1)
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

func (mtdf *Mtdf) search(alpha int) int {
	if mtdf.depth < minHashtableDepth {
		return mtdf.searchNoHashtable(alpha)
	}

	cacheKey := CacheKey{board: mtdf.board.Normalize(), depth: mtdf.depth}
	var cacheValue CacheValue

	var ok bool
	if cacheValue, ok = mtdf.cache.Lookup(cacheKey); ok {
		if cacheValue.alpha > alpha {
			return alpha + 1
		}

		if cacheValue.beta < alpha+1 {
			return alpha
		}
	} else {
		cacheValue = CacheValue{
			alpha: MinHeuristic,
			beta:  MaxHeuristic,
		}
	}

	mtdf.stats.Nodes++

	children := mtdf.board.GetSortableChildren()

	if len(children) == 0 {
		return mtdf.handleNoMoves(alpha)
	}

	for i := range children {
		copy := children[i].Board
		children[i].Heur = -searchForSort(&copy, MinHeuristic, MaxHeuristic, mtdf.depth/4)
	}

	sort.Slice(children, func(i, j int) bool {
		return children[i].Heur > children[j].Heur
	})

	heur := alpha
	mtdf.depth--
	parent := mtdf.board
	for _, child := range children {
		mtdf.board = child.Board
		childHeur := -mtdf.search(-(alpha + 1))
		if childHeur > alpha {
			heur = alpha + 1
			break
		}
	}
	mtdf.board = parent
	mtdf.depth++

	if heur == alpha {
		if heur < cacheValue.beta {
			cacheValue.beta = heur
		}
	} else {
		if heur > cacheValue.alpha {
			cacheValue.alpha = heur
		}
	}

	if err := mtdf.cache.Save(cacheKey, cacheValue); err != nil {
		log.Printf("warning: saving cache value failed: %s", err.Error())
	}

	return heur
}

func (mtdf *Mtdf) searchNoHashtable(alpha int) int {
	mtdf.stats.Nodes++

	if mtdf.depth == 0 {
		heur := mtdf.heuristic(mtdf.board)
		if heur > alpha {
			return alpha + 1
		}
		return alpha
	}

	gen := othello.NewChildGenerator(&mtdf.board)

	if !gen.HasMoves() {
		if mtdf.board.OpponentMoves() == 0 {
			heur := ExactScoreFactor * mtdf.board.ExactScore()
			if heur > alpha {
				return alpha + 1
			}
			return alpha
		}

		mtdf.board.SwitchTurn()
		heur := -mtdf.searchNoHashtable(-(alpha + 1))
		mtdf.board.SwitchTurn()
		return heur
	}

	heur := alpha
	mtdf.depth--
	for gen.Next() {
		childHeur := -mtdf.searchNoHashtable(-(alpha + 1))
		if childHeur > alpha {
			gen.RestoreParent()
			heur = alpha + 1
			break
		}
	}
	mtdf.depth++

	return heur
}
