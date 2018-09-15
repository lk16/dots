package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type Mtdf struct {
	board othello.Board
	high  int
	low   int
	depth int
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
	return mtdf.slideWindow()
}

func (mtdf *Mtdf) ExactSearch(board othello.Board) int {
	return mtdf.Search(board, 60)
}

func (mtdf *Mtdf) slideWindow() int {

	f := heuristics.Squared(mtdf.board)

	var step int
	if mtdf.board.CountEmpties() > mtdf.depth {
		step = 1
	} else {
		step = 2
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
		var bound = mtdf.search(f)

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

	if mtdf.depth == 0 {
		return mtdf.polish(heuristics.Squared(mtdf.board), alpha)
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

	for gen.Next() {
		mtdf.depth--
		childHeur := -mtdf.search(-(alpha + 1))
		mtdf.depth++
		if childHeur > alpha {
			gen.RestoreParent()
			return alpha + 1
		}
	}
	return alpha
}

// -----

type bounds struct {
	upper int
	lower int
}

type MtdfHashTable struct {
	board  othello.Board
	high   int
	low    int
	depth  int
	lookup map[othello.Board]bounds
}

func NewMtdfHashTable(low, high int) *MtdfHashTable {
	return &MtdfHashTable{
		high: high,
		low:  low}
}

func (mtdf *MtdfHashTable) Name() string {
	return "mtdfhashtable"
}

func (mtdf *MtdfHashTable) Search(board othello.Board, depth int) int {
	mtdf.board = board
	mtdf.depth = depth
	mtdf.lookup = make(map[othello.Board]bounds, 10240)
	return mtdf.slideWindow()
}

func (mtdf *MtdfHashTable) ExactSearch(board othello.Board) int {
	return mtdf.Search(board, 60)
}

func (mtdf *MtdfHashTable) slideWindow() int {

	f := heuristics.Squared(mtdf.board)

	var step int
	if mtdf.board.CountEmpties() > mtdf.depth {
		step = 1
	} else {
		step = 2
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
		var bound = mtdf.search(f)

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

func (mtdf *MtdfHashTable) polish(heur, alpha int) int {
	if heur > alpha {
		return alpha + 1
	}
	return alpha
}

func (mtdf *MtdfHashTable) search(alpha int) (heur int) {

	defer func() {

		if mtdf.depth >= 3 {
			entry, ok := mtdf.lookup[mtdf.board.Normalize()]

			if !ok {
				entry = bounds{
					lower: MinHeuristic,
					upper: MaxHeuristic}
			}

			if heur == alpha && heur < entry.upper {
				entry.upper = alpha
			} else if heur > entry.lower {
				entry.lower = alpha + 1
			}
			mtdf.lookup[mtdf.board] = entry
		}
	}()

	if mtdf.depth >= 3 {
		if entry, ok := mtdf.lookup[mtdf.board.Normalize()]; ok {
			if entry.upper <= alpha {
				return alpha
			}
			if entry.lower >= alpha+1 {
				return alpha + 1
			}
		}
	}

	if mtdf.depth == 0 {
		return mtdf.polish(heuristics.Squared(mtdf.board), alpha)
	}

	gen := othello.NewGenerator(&mtdf.board, 0)

	if !gen.HasMoves() {

		if mtdf.board.OpponentMoves() != 0 {
			mtdf.board.SwitchTurn()
			heur = -mtdf.search(-(alpha + 1))
			mtdf.board.SwitchTurn()
			return
		}

		return mtdf.polish(ExactScoreFactor*mtdf.board.ExactScore(), alpha)
	}

	for gen.Next() {
		mtdf.depth--
		childHeur := -mtdf.search(-(alpha + 1))
		mtdf.depth++
		if childHeur > alpha {
			gen.RestoreParent()
			return alpha + 1
		}
	}
	return alpha
}