package treesearch

import (
	"sort"

	"github.com/lk16/dots/internal/othello"
)

const (
	minHashtableDepth = 5
)

type hashtableValue struct {
	high      int
	low       int
	bestChild othello.Board
}

// Mtdf implements the mtdf tree search algorithm
type Mtdf struct {
	board     othello.Board
	high      int
	low       int
	depth     int
	hashtable map[othello.Board]hashtableValue
	stats     Stats
	heuristic func(othello.Board) int
	sorter    Pvs
}

// NewMtdf returns a new Mtdf
func NewMtdf(heuristic func(othello.Board) int) *Mtdf {
	return &Mtdf{
		heuristic: heuristic,
		hashtable: make(map[othello.Board]hashtableValue, 100000),
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

// ClearHashTable clears the Mtdf hash table
func (mtdf *Mtdf) ClearHashTable() {
	for key := range mtdf.hashtable {
		delete(mtdf.hashtable, key)
	}
}

// Search searches for the the best move up to a certain depth
func (mtdf *Mtdf) Search(board othello.Board, alpha, beta, depth int) int {
	mtdf.low = alpha
	mtdf.high = beta

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
	return mtdf.Search(board, alpha, beta, 60) / ExactScoreFactor
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
		var bound = -mtdf.search(-(f + 1))

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

func (mtdf *Mtdf) search(alpha int) int {

	if mtdf.depth < minHashtableDepth {
		return mtdf.searchNoHashtable(alpha)
	}

	mtdf.stats.Nodes++

	key := mtdf.board.Normalize()

	entry, ok := mtdf.hashtable[key]

	if ok {
		if entry.high <= alpha {
			return alpha
		}
		if entry.low >= alpha+1 {
			return alpha + 1
		}

		mtdf.depth--
		parent := mtdf.board
		mtdf.board = entry.bestChild
		childHeur := -mtdf.search(-(alpha + 1))
		mtdf.board = parent
		mtdf.depth++
		if childHeur > alpha {
			return alpha + 1
		}

	} else {
		entry = hashtableValue{
			high: MaxHeuristic,
			low:  MinHeuristic,
		}
	}

	children := mtdf.board.GetSortableChildren()

	if len(children) == 0 {
		return mtdf.handleNoMoves(alpha)
	}

	for i := range children {
		children[i].Heur = mtdf.sorter.Search(children[i].Board, MinHeuristic, MaxHeuristic, 2)
	}
	sort.Slice(children, func(i, j int) bool {
		return children[i].Heur > children[j].Heur
	})

	heur := alpha
	mtdf.depth--
	parent := mtdf.board
	bestChild := children[0].Board
	for _, child := range children {
		mtdf.board = child.Board
		childHeur := -mtdf.search(-(alpha + 1))
		if childHeur > alpha {
			heur = alpha + 1
			bestChild = child.Board
			break
		}
	}
	mtdf.board = parent
	mtdf.depth++

	if heur == alpha {
		if alpha < entry.high {
			entry.high = alpha
		}
	} else {
		if alpha+1 > entry.low {
			entry.low = alpha + 1
		}
	}
	entry.bestChild = bestChild

	mtdf.hashtable[key] = entry

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
