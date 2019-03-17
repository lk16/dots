package treesearch

import (
	"github.com/lk16/dots/heuristics"
	"github.com/lk16/dots/othello"
)

type hashtableKey struct {
	board othello.Board
	depth int
}

type bounds struct {
	high int
	low  int
}

type Mtdf struct {
	board     othello.Board
	high      int
	low       int
	depth     int
	hashtable map[hashtableKey]bounds
}

func NewMtdf(low, high int) *Mtdf {
	return &Mtdf{
		high: high,
		low:  low}
}

func (mtdf *Mtdf) Name() string {
	return "mtdf"
}

func (mtdf *Mtdf) Search(board othello.Board, depth int) int {
	mtdf.board = board
	mtdf.depth = depth
	mtdf.hashtable = make(map[hashtableKey]bounds, 100000)
	return mtdf.slideWindow()
}

func (mtdf *Mtdf) ExactSearch(board othello.Board) int {
	mtdf.high *= ExactScoreFactor
	mtdf.low *= ExactScoreFactor
	defer func() {
		mtdf.high /= ExactScoreFactor
		mtdf.low /= ExactScoreFactor
	}()
	return mtdf.Search(board, 60) / ExactScoreFactor
}

func (mtdf *Mtdf) slideWindow() int {

	var f int

	var step int
	if mtdf.board.CountEmpties() > mtdf.depth {
		f = heuristics.Squared(mtdf.board)
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

func (mtdf *Mtdf) polish(heur, alpha int) int {
	if heur > alpha {
		return alpha + 1
	}
	return alpha
}

func (mtdf *Mtdf) search(alpha int) int {

	if mtdf.depth <= 6 {
		return mtdf.searchNoHashtable(alpha)
	}

	key := hashtableKey{
		board: mtdf.board.Normalize(),
		depth: mtdf.depth}

	entry, ok := mtdf.hashtable[key]
	if ok {
		if entry.high <= alpha {
			return alpha
		}
		if entry.low >= alpha+1 {
			return alpha + 1
		}
	}

	gen := othello.NewGenerator(&mtdf.board, 0)

	if !gen.HasMoves() {

		if mtdf.board.OpponentMoves() != 0 {
			mtdf.board.SwitchTurn()
			heur := -mtdf.search(-(alpha + 1))
			mtdf.board.SwitchTurn()
			return heur
		}

		return mtdf.polish(ExactScoreFactor*mtdf.board.ExactScore(), alpha)
	}

	heur := alpha
	for gen.Next() {
		mtdf.depth--
		childHeur := -mtdf.search(-(alpha + 1))
		mtdf.depth++
		if childHeur > alpha {
			gen.RestoreParent()
			heur = alpha + 1
			break
		}
	}

	entry, ok = mtdf.hashtable[key]

	if !ok {
		entry = bounds{
			high: MaxHeuristic,
			low:  MinHeuristic}
	}

	if heur == alpha {
		if alpha < entry.high {
			entry.high = alpha
		}
	} else {
		if alpha+1 > entry.low {
			entry.low = alpha + 1
		}
	}

	mtdf.hashtable[key] = entry
	return heur
}

func (mtdf *Mtdf) searchNoHashtable(alpha int) (heur int) {

	if mtdf.depth == 0 {
		return mtdf.polish(heuristics.Squared(mtdf.board), alpha)
	}

	gen := othello.NewGenerator(&mtdf.board, 0)

	if !gen.HasMoves() {

		if mtdf.board.OpponentMoves() != 0 {
			mtdf.board.SwitchTurn()
			heur = -mtdf.searchNoHashtable(-(alpha + 1))
			mtdf.board.SwitchTurn()
			return
		}

		return mtdf.polish(ExactScoreFactor*mtdf.board.ExactScore(), alpha)
	}

	for gen.Next() {
		mtdf.depth--
		childHeur := -mtdf.searchNoHashtable(-(alpha + 1))
		mtdf.depth++
		if childHeur > alpha {
			gen.RestoreParent()
			return alpha + 1
		}
	}
	return alpha
}
