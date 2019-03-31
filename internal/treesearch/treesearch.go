// Package treesearch contains algorithms for Board evaluation
package treesearch

import "github.com/lk16/dots/internal/othello"

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

// Interface is the interface for tree search algorithms
type Interface interface {
	Name() string
	Search(board othello.Board, depth int) int
	ExactSearch(board othello.Board) int
}
