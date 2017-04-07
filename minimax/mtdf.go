package minimax

import (
	"dots/board"
)

type Mtdf struct {
	heuristic Heuristic
}

func (mtdf *Mtdf) Search(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int) {
	mtdf.heuristic = heuristic

	if board.CountEmpties() == depth_left {

		upper_cap := 100
		lower_cap := -upper_cap

		capped_lower_bound := alpha
		if alpha < lower_cap {
			lower_cap = alpha
		}

		capped_upper_bound := upper_cap

		capped_result := mtdf.loop(board, depth_left, capped_lower_bound, capped_upper_bound, 0, 1)

		if (capped_result > capped_lower_bound) && (capped_result < capped_upper_bound) {
			heur = capped_result
			return
		}
	}

	heur = mtdf.loop(board, depth_left, alpha, Max_heuristic, 0, 2*Exact_score_factor)
	return
}

func (mtdf *Mtdf) loop(board board.Board, depth_left uint, lower_bound, upper_bound, guess, step int) (heur int) {
	f := guess
	if f < lower_bound {
		f = lower_bound
	}
	if f > upper_bound {
		f = upper_bound
	}
	for upper_bound-lower_bound >= step {
		bound := -mtdf.doMtdf(board, depth_left, -(f + 1))
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

func (mtdf *Mtdf) polish(heur, alpha int) (outheur int) {
	outheur = heur
	if heur > alpha {
		heur++
	}
	return
}
