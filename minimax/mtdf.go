package minimax

import (
	"dots/board"
)

type Mtdf struct {
	heuristic Heuristic
}

func (mtdf *Mtdf) Search(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int) {
	mtdf.heuristic = heuristic
	heur = mtdf.loop(board, depth_left, alpha, Max_heuristic, 0, 1, false)
	return
}

func (mtdf *Mtdf) ExactSearch(board board.Board, alpha int) (heur int) {
	heur = mtdf.loop(board, 64, alpha, Max_exact_heuristic, 0, 2, true)
	return
}

func (mtdf *Mtdf) loop(board board.Board, depth_left uint,
	lower_bound, upper_bound, guess, step int, exact bool) (heur int) {
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
			bound = -mtdf.doMtdfExact(board, -(f + 1))
		} else {
			bound = -mtdf.doMtdf(board, depth_left, -(f + 1))
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

func (mtdf *Mtdf) doMtdf(board board.Board, depth_left uint, alpha int) (heur int) {

	if depth_left == 0 {
		heur = mtdf.polish(mtdf.heuristic(board), alpha)
		return
	}

	if moves := board.Moves(); moves != 0 {
		heur = alpha
		for child := range board.GenChildren() {
			child_heur := -mtdf.doMtdf(child, depth_left-1, -(alpha + 1))
			if child_heur > alpha {
				heur = alpha + 1
				break
			}
		}
		return
	}

	board.SwitchTurn()
	if moves := board.Moves(); moves != 0 {
		heur = -mtdf.doMtdf(board, depth_left, -(alpha + 1))
		return
	}

	heur = mtdf.polish(Exact_score_factor*board.ExactScore(), alpha)
	return
}

func (mtdf *Mtdf) doMtdfExact(board board.Board, alpha int) (heur int) {

	if moves := board.Moves(); moves != 0 {
		for child := range board.GenChildren() {
			child_heur := -mtdf.doMtdfExact(child, -(alpha + 1))
			if child_heur > alpha {
				heur = alpha + 1
				break
			}
		}
		return
	}

	board.SwitchTurn()
	if moves := board.Moves(); moves != 0 {
		heur = -mtdf.doMtdfExact(board, -(alpha + 1))
		return
	}

	heur = mtdf.polish(board.ExactScore(), alpha)
	return
}

func (mtdf *Mtdf) polish(heur, alpha int) (outheur int) {
	outheur = heur
	if heur > alpha {
		outheur++
	}
	return
}
