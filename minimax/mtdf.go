package minimax

import (
	"time"

	"dots/board"
)

type Mtdf struct {
	heuristic       Heuristic
	nodes           uint64
	compute_time_ns uint64
	search_start    time.Time
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

func (mtdf *Mtdf) Search(board board.Board, depth_left uint, heuristic Heuristic,
	alpha int) (heur int) {

	mtdf.preSearch(heuristic)
	defer mtdf.postSearch()

	/*
	   Temporary hack:
	   If the true heuristic is outside the -100,100 interval
	   we return either 100 or -100, because exact scores currently take too long to compute
	*/

	upper_limit := 100
	lower_limit := -upper_limit

	capped_alpha := lower_limit
	if alpha > capped_alpha {
		capped_alpha = alpha
	}

	capped_beta := upper_limit

	mtdf.board = board
	heur = mtdf.loop(depth_left, capped_alpha, capped_beta, 0, 1, false)

	return
}

func (mtdf *Mtdf) ExactSearch(board board.Board, alpha int) (heur int) {
	mtdf.preSearch(nil)
	defer mtdf.postSearch()

	mtdf.board = board
	heur = mtdf.loop(64, alpha, Max_exact_heuristic, 0, 2, true)
	return
}

func (mtdf *Mtdf) loop(depth_left uint, lower_bound, upper_bound, guess, step int,
	exact bool) (heur int) {

	f := guess
	if f < lower_bound {
		f = lower_bound
	}
	if f > upper_bound {
		f = upper_bound
	}
	for upper_bound-lower_bound >= step {
		var bound int
		if exact {
			bound = -mtdf.doMtdfExact(-(f + 1))
		} else {
			bound = -mtdf.doMtdf(depth_left, -(f + 1))
		}
		if bound == f {
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

func (mtdf *Mtdf) doMtdf(depth_left uint, alpha int) (heur int) {

	mtdf.nodes += 1

	if depth_left == 0 {
		heur = mtdf.polish(mtdf.heuristic(mtdf.board), alpha)
		return
	}

	gen := mtdf.board.ChildGen()

	if gen.HasMoves() {
		heur = alpha
		for gen.Next() {
			child_heur := -mtdf.doMtdf(depth_left-1, -(alpha + 1))
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
		heur = -mtdf.doMtdf(depth_left, -(alpha + 1))
		mtdf.board.SwitchTurn()
		return
	}

	mtdf.board.SwitchTurn()
	heur = mtdf.polish(Exact_score_factor*mtdf.board.ExactScore(), alpha)
	return
}

func (mtdf *Mtdf) doMtdfExact(alpha int) (heur int) {

	mtdf.nodes += 1

	gen := mtdf.board.ChildGen()

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
