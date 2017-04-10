package minimax

import (
	"dots/board"
)

const (
	Max_exact_heuristic = 64
	Min_exact_heuristic = -Max_exact_heuristic
	Exact_score_factor  = 1000
	Max_heuristic       = Exact_score_factor * Max_exact_heuristic
	Min_heuristic       = Exact_score_factor * Min_exact_heuristic
)

type Heuristic func(board board.Board) (heur int)

type Interface interface {
	Search(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int)
	ExactSearch(board board.Board, alpha int) (heur int)
	Name() (name string)
}
