package minimax

import (
	"time"

	"dots/board"
)

type HashtableKey struct {
	board board.Board
	depth uint
}

type HashTableValue struct {
	upper_bound int
	lower_bound int
}

type Mtdf struct {
	heuristic       Heuristic
	nodes           uint64
	compute_time_ns uint64
	search_start    time.Time
	search_depth    uint
	board           board.Board
}

func (mtdf *Mtdf) preSearch(heuristic Heuristic) {
	mtdf.heuristic = heuristic
	mtdf.nodes = 0
	mtdf.compute_time_ns = 0
	mtdf.search_start = time.Now()
}

func (mtdf *Mtdf) postSearch() {
	mtdf.compute_time_ns = uint64(time.Since(mtdf.search_start).Nanoseconds())
}

func clamp(x, min, max int) int {
	if x > max {
		return max
	}
	if x < min {
		return min
	}
	return x
}

func (mtdf *Mtdf) Search(board board.Board, search_depth uint, heuristic Heuristic,
	alpha int) (heur int) {

	mtdf.preSearch(heuristic)
	defer mtdf.postSearch()

	mtdf.board = board
	mtdf.search_depth = search_depth

	upper_limit := 100
	lower_limit := -100

	capped_alpha := clamp(alpha, lower_limit, upper_limit)
	heur = mtdf.loop(capped_alpha, upper_limit, 0, 1, false)

	if heur <= lower_limit || heur >= 100 {
		heur = mtdf.loop(Min_heuristic, Max_heuristic, 0, 2*Exact_score_factor, false)
	}

	return
}

func (mtdf *Mtdf) ExactSearch(board board.Board, alpha int) (heur int) {
	mtdf.preSearch(nil)
	defer mtdf.postSearch()

	mtdf.board = board
	heur = mtdf.loop(alpha, Max_exact_heuristic, 0, 2, true)
	return
}

func (mtdf *Mtdf) loop(lower_bound, upper_bound, guess, step int, exact bool) (heur int) {

	var call func(int) int

	if exact {
		call = func(alpha int) int {
			return -mtdf.doMtdfExact(-(alpha + 1))
		}
	} else {
		call = func(alpha int) int {
			return -mtdf.doMtdf(mtdf.search_depth, -(alpha + 1))
		}
	}

	f := clamp(guess, lower_bound, upper_bound)

	for upper_bound-lower_bound >= step {
		bound := call(f)
		if f == bound {
			f -= step
			upper_bound = bound
		} else {
			f += step
			lower_bound = bound
		}
	}
	heur = upper_bound
	return
}

func (mtdf *Mtdf) doMtdf(search_depth uint, alpha int) (heur int) {

	mtdf.nodes += 1

	if search_depth == 0 {
		heur = mtdf.polish(mtdf.heuristic(mtdf.board), alpha)
		return
	}

	gen := board.NewChildGen(&mtdf.board)

	if gen.HasMoves() {
		heur = alpha
		for gen.Next() {
			child_heur := -mtdf.doMtdf(search_depth-1, -(alpha + 1))
			if child_heur > alpha {
				heur = alpha + 1
				gen.RestoreParent()
				break
			}
		}
		return
	}

	mtdf.board.SwitchTurn()
	if moves := mtdf.board.Moves(); moves != 0 {
		heur = -mtdf.doMtdf(search_depth, -(alpha + 1))
		mtdf.board.SwitchTurn()
		return
	}

	mtdf.board.SwitchTurn()
	heur = mtdf.polish(Exact_score_factor*mtdf.board.ExactScore(), alpha)
	return
}

func (mtdf *Mtdf) doMtdfExact(alpha int) (heur int) {

	mtdf.nodes += 1

	gen := board.NewChildGen(&mtdf.board)

	if gen.HasMoves() {
		heur = alpha
		for gen.Next() {
			child_heur := -mtdf.doMtdfExact(-(alpha + 1))
			if child_heur > alpha {
				heur = alpha + 1
				gen.RestoreParent()
				break
			}
		}
		return
	}

	mtdf.board.SwitchTurn()
	if moves := mtdf.board.Moves(); moves != 0 {
		heur = -mtdf.doMtdfExact(-(alpha + 1))
		mtdf.board.SwitchTurn()
		return
	}

	mtdf.board.SwitchTurn()
	heur = mtdf.polish(mtdf.board.ExactScore(), alpha)
	return
}

func (mtdf *Mtdf) polish(heur, alpha int) (outheur int) {
	if heur > alpha {
		outheur = alpha + 1
		return
	}
	outheur = alpha
	return
}

func (mtdf Mtdf) Name() (name string) {
	name = "mtdf"
	return
}

func (mtdf *Mtdf) Nodes() (nodes uint64) {
	nodes = mtdf.nodes
	return
}

func (mtdf *Mtdf) ComputeTimeNs() (ns uint64) {
	ns = mtdf.compute_time_ns
	return
}
