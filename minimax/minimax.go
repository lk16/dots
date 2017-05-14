package minimax

import (
	"dots/board"
)

type Minimax struct {
	heuristic Heuristic
	nodes     uint64
}

func (minimax *Minimax) Search(board board.Board, depth_left uint,
	heuristic Heuristic, alpha int) (heur int) {

	minimax.nodes = 0
	minimax.heuristic = heuristic
	heur = -minimax.doMinimax(board, depth_left, true)
	return
}

func (minimax *Minimax) doMinimax(board board.Board, depth_left uint,
	is_max bool) (heur int) {

	minimax.nodes++

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
			for _, child := range board.GetChildren() {
				child_heur := minimax.doMinimax(child, depth_left-1, false)
				if child_heur > heur {
					heur = child_heur
				}
			}
		} else {
			heur = Max_heuristic
			for _, child := range board.GetChildren() {
				child_heur := minimax.doMinimax(child, depth_left-1, true)
				if child_heur < heur {
					heur = child_heur
				}
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

func (minimax *Minimax) ExactSearch(board board.Board, alpha int) (heur int) {
	minimax.nodes = 0
	heur = -minimax.doMinimaxExact(board, true)
	return
}

func (minimax *Minimax) doMinimaxExact(board board.Board, is_max bool) (heur int) {

	minimax.nodes++

	if moves := board.Moves(); moves != 0 {
		if is_max {
			heur = Min_exact_heuristic
			for _, child := range board.GetChildren() {
				child_heur := minimax.doMinimaxExact(child, false)
				if child_heur > heur {
					heur = child_heur
				}
			}
		} else {
			heur = Max_exact_heuristic
			for _, child := range board.GetChildren() {
				child_heur := minimax.doMinimaxExact(child, true)
				if child_heur < heur {
					heur = child_heur
				}
			}
		}
		return
	}

	clone := board
	clone.SwitchTurn()
	if moves := clone.Moves(); moves != 0 {
		heur = minimax.doMinimaxExact(clone, !is_max)
		return
	}

	if is_max {
		heur = board.ExactScore()
	} else {
		heur = -board.ExactScore()
	}
	return
}

func (minimax Minimax) Name() (name string) {
	name = "minimax"
	return
}

func (minimax *Minimax) NodesVisited() (nodes uint64) {
	nodes = minimax.nodes
	return
}

func (minimax *Minimax) ComputeTimeNs() (ns uint64) {
	// TODO
	ns = 0
	return
}
