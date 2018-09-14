package treesearch

import "dots/othello"

const (
	// MaxScore is the highest game result score possible
	MaxScore = 64

	// MinScore is the lowest game result score possible
	MinScore = -MaxScore

	// ExactScoreFactor is the multiplication.
	// This is used when a non exact search runs into an exact result
	ExactScoreFactor = 1000

	// MaxHeuristic is the highest heuristic value possible
	MaxHeuristic = ExactScoreFactor * MaxScore

	// MinHeuristic is the lowest heuristic value possible
	MinHeuristic = ExactScoreFactor * MinScore
)

type Interface interface {
	Name() string
	Search(board othello.Board, depth int) int
	ExactSearch(board othello.Board) int
}
