package treesearch

import (
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

// Mtdf implements the mtdf tree search algorithm
type Mtdf struct {
	board     othello.Board
	high      int
	low       int
	hashtable map[hashtableKey]bounds
	Stats     stats
}

// NewMtdf returns a new Mtdf
func NewMtdf() *Mtdf {
	return &Mtdf{
		hashtable: make(map[hashtableKey]bounds, 100000)}
}

// Name returns the tree search algorithm name
func (mtdf *Mtdf) Name() string {
	return "mtdf"
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

	if board.Moves() == 0 {

		if board.OpponentMoves() == 0 {
			return ExactScoreFactor * board.ExactScore()
		}

		board.SwitchTurn()
		heur := -mtdf.Search(board, -beta, -alpha, depth)
		board.SwitchTurn()
		return heur
	}

	mtdf.board = board
	mtdf.Stats.StartClock()
	heuristic := mtdf.slideWindow(depth)
	mtdf.Stats.StopClock()
	return heuristic
}

// ExactSearch searches for the best move without a depth limitation
func (mtdf *Mtdf) ExactSearch(board othello.Board, alpha, beta int) int {
	return mtdf.Search(board, alpha, beta, 60) / ExactScoreFactor
}

func (mtdf *Mtdf) slideWindow(depth int) int {

	var f int

	var step int
	if depth < mtdf.board.CountEmpties() {
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

func (mtdf *Mtdf) handleNoMoves(alpha, depth int) int {

	if mtdf.board.OpponentMoves() == 0 {
		return mtdf.polish(-ExactScoreFactor*mtdf.board.ExactScore(), alpha)
	}

	mtdf.board.SwitchTurn()
	heur := -mtdf.search(-(alpha + 1), depth)
	mtdf.board.SwitchTurn()
	return heur
}

func (mtdf *Mtdf) checkHashTable(alpha, depth int) (int, bool) {
	key := hashtableKey{
		board: mtdf.board.Normalize(),
		depth: depth}

	entry, ok := mtdf.hashtable[key]
	if ok {
		if entry.high <= alpha {
			return alpha, true
		}
		if entry.low >= alpha+1 {
			return alpha + 1, true
		}
	}
	return 0, false
}

func (mtdf *Mtdf) updateHashTable(alpha, depth, heur int) {

	key := hashtableKey{
		board: mtdf.board.Normalize(),
		depth: depth}

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

func (mtdf *Mtdf) search(alpha, depth int) int {

	mtdf.Stats.Nodes++

	if depth == 0 {
		return mtdf.polish(FastHeuristic(mtdf.board), alpha)
	}

	if depth > 4 {
		if heur, ok := mtdf.checkHashTable(alpha, depth); ok {
			return heur
		}
	}

	gen := othello.NewUnsortedChildGenerator(&mtdf.board)
	if !gen.HasMoves() {
		return mtdf.handleNoMoves(alpha, depth)
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
		mtdf.updateHashTable(alpha, depth, heur)
	}

	return heur
}
