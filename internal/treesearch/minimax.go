package treesearch

import (
	"github.com/lk16/dots/internal/othello"
)

// MiniMax implements the minimax tree search algorithm
type MiniMax struct {
	stats Stats
}

// NewMinimax returns a new MiniMax
func NewMinimax() *MiniMax {
	return &MiniMax{}
}

// Name returns the tree search algorithm name
func (minimax *MiniMax) Name() string {
	return "minimax"
}

// GetStats returns the statistics for the latest search
func (minimax MiniMax) GetStats() Stats {
	return minimax.stats
}

// ResetStats resets the statistics for the latest search to zeroes
func (minimax MiniMax) ResetStats() {
	minimax.stats.Reset()
}

// Search searches for the the best move up to a certain depth
func (minimax *MiniMax) Search(board othello.Board, alpha, beta, depth int) int {

	if depth >= board.CountEmpties() {
		depth = 60
	}

	minimax.stats.StartClock()
	heur := -minimax.search(board, depth, true)
	minimax.stats.StopClock()

	if heur < alpha {
		return alpha
	}

	if heur > beta {
		return beta
	}

	return heur
}

// ExactSearch searches for the best move without a depth limitation
func (minimax *MiniMax) ExactSearch(board othello.Board, alpha, beta int) int {
	return minimax.Search(board, alpha, beta, 60)
}

func (minimax *MiniMax) search(board othello.Board, depth int, maxPlayer bool) int {

	minimax.stats.Nodes++

	if depth == 0 {
		heur := FastHeuristic(board)
		if !maxPlayer {
			heur = -heur
		}
		return heur
	}

	child := board
	gen := othello.NewUnsortedChildGenerator(&child)

	if !gen.HasMoves() {

		if board.OpponentMoves() == 0 {
			heur := ExactScoreFactor * board.ExactScore()
			if !maxPlayer {
				heur = -heur
			}
			return heur
		}

		board.SwitchTurn()
		return minimax.search(board, depth, !maxPlayer)
	}

	if maxPlayer {
		heur := MinHeuristic
		for gen.Next() {
			childHeur := minimax.search(child, depth-1, !maxPlayer)
			if childHeur > heur {
				heur = childHeur
			}
		}
		return heur
	}

	heur := MaxHeuristic
	for gen.Next() {
		childHeur := minimax.search(child, depth-1, !maxPlayer)
		if childHeur < heur {
			heur = childHeur
		}
	}
	return heur
}
