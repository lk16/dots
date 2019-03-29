package treesearch

import (
	"github.com/lk16/dots/internal/heuristics"
	"github.com/lk16/dots/internal/othello"
	"time"
)

type stats struct {
	Nodes     uint64
	StartTime time.Time
	Duration  time.Duration
}

func (s *stats) StartClock() {
	s.StartTime = time.Now()
}

func (s *stats) StopClock() {
	s.Duration += time.Now().Sub(s.StartTime)
}

func (s *stats) NodesPerSecond() float64 {

	duration := s.Duration.Seconds()

	if duration == 0.0 {
		return 0.0
	}

	return float64(s.Nodes) / duration
}

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
	hashtable map[hashtableKey]bounds
	Stats     stats
}

func NewMtdf(low, high int) *Mtdf {
	mtdf := &Mtdf{
		hashtable: make(map[hashtableKey]bounds, 100000)}
	mtdf.SetAlphaBeta(low, high)
	return mtdf
}

func (mtdf *Mtdf) Name() string {
	return "mtdf"
}

func (mtdf *Mtdf) SetAlphaBeta(alpha, beta int) {
	mtdf.low = alpha
	mtdf.high = beta
}

func (mtdf *Mtdf) ClearHashTable() {
	for key := range mtdf.hashtable {
		delete(mtdf.hashtable, key)
	}
}

func (mtdf *Mtdf) Search(board othello.Board, depth int) int {
	mtdf.board = board
	mtdf.Stats.StartClock()
	heuristic := mtdf.slideWindow(depth)
	mtdf.Stats.StopClock()
	return heuristic
}

func (mtdf *Mtdf) ExactSearch(board othello.Board) int {
	mtdf.high = MaxHeuristic
	mtdf.low = MinHeuristic
	return mtdf.Search(board, 60) / ExactScoreFactor
}

func (mtdf *Mtdf) slideWindow(depth int) int {

	var f int

	var step int
	if mtdf.board.CountEmpties() > depth {
		f = heuristics.FastHeuristic(mtdf.board)
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
		var bound = -mtdf.search(-(f + 1), depth)

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

func (mtdf *Mtdf) search(alpha, depth int) int {

	mtdf.Stats.Nodes++

	if depth == 0 {
		return mtdf.polish(heuristics.FastHeuristic(mtdf.board), alpha)
	}

	var key hashtableKey

	if depth > 4 {
		key = hashtableKey{
			board: mtdf.board.Normalize(),
			depth: depth}

		entry, ok := mtdf.hashtable[key]
		if ok {
			if entry.high <= alpha {
				return alpha
			}
			if entry.low >= alpha+1 {
				return alpha + 1
			}
		}
	}

	gen := othello.NewUnsortedChildGenerator(&mtdf.board)

	if !gen.HasMoves() {

		if mtdf.board.OpponentMoves() != 0 {
			mtdf.board.SwitchTurn()
			heur := -mtdf.search(-(alpha + 1), depth)
			mtdf.board.SwitchTurn()
			return heur
		}

		return mtdf.polish(ExactScoreFactor*mtdf.board.ExactScore(), alpha)
	}

	heur := alpha
	for gen.Next() {
		childHeur := -mtdf.search(-(alpha + 1), depth-1)
		if childHeur > alpha {
			gen.RestoreParent()
			heur = alpha + 1
			break
		}
	}

	if depth > 4 {
		entry, ok := mtdf.hashtable[key]

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
	}
	return heur
}

/*
func (mtdf *Mtdf) searchNoHashtable(alpha int) (heur int) {

	mtdf.Stats.Nodes++

	if mtdf.depth == 0 {
		return mtdf.polish(heuristics.FastHeuristic(mtdf.board), alpha)
	}

	gen := othello.NewUnsortedChildGenerator(&mtdf.board)

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
*/
