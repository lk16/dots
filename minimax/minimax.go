package minimax

import (
	"dots/board"
)

const (
	Max_exact_heuristic = 64
	Min_exact_heuristic = -Max_exact_heuristic
	Exact_score_factor  = 10000
	Max_heuristic       = Exact_score_factor * Max_exact_heuristic
	Min_heuristic       = Exact_score_factor * Min_exact_heuristic
)

type Heuristic func(board board.Board) (heur int)

type MinimaxInterface interface {
	Evaluate(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int)
}

type Minimax struct {
	heuristic Heuristic
}

func (minimax *Minimax) Evaluate(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int) {
	minimax.heuristic = heuristic
	heur = minimax.doMinimax(board, depth_left, true)
	return
}

func (minimax *Minimax) doMinimax(board board.Board, depth_left uint, is_max bool) (heur int) {
	if depth_left == 0 {
		if is_max {
			heur = minimax.heuristic(board)
		} else {
			heur = -minimax.heuristic(board)
		}
		return
	}

	if moves := board.Moves(); moves != 0 {
		if is_max {
			heur = Min_heuristic
		} else {
			heur = Max_heuristic
		}

		for child := range board.GenChildren() {
			child_heur := minimax.doMinimax(child, depth_left-1, !is_max)
			if is_max && (child_heur > heur) {
				heur = child_heur
			}
			if (!is_max) && (child_heur < heur) {
				heur = child_heur
			}
		}
		return
	}

	board.SwitchTurn()
	if moves := board.Moves(); moves != 0 {
		heur = minimax.doMinimax(board, depth_left, !is_max)
		return
	}

	if is_max {
		heur = Exact_score_factor * board.ExactScore()
	} else {
		heur = -Exact_score_factor * board.ExactScore()
	}
	return
}
