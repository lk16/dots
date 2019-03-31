package treesearch

import (
	"github.com/lk16/dots/internal/othello"
)

// MiniMax implements the minimax tree search algorithm
type MiniMax struct{}

// NewMinimax returns a new MiniMax
func NewMinimax() *MiniMax {
	return &MiniMax{}
}

// Name returns the tree search algorithm name
func (minimax *MiniMax) Name() string {
	return "minimax"
}

// Search searches for the the best move up to a certain depth
func (minimax *MiniMax) Search(board othello.Board, depth int) int {

	if depth > board.CountEmpties() {
		depth = board.CountEmpties()
	}

	if board.Moves() == 0 && board.OpponentMoves() != 0 {
		board.SwitchTurn()
		heur := minimax.Search(board, depth)
		board.SwitchTurn()
		return heur
	}

	return -minimax.search(board, depth, true)
}

// ExactSearch searches for the best move without a depth limitation
func (minimax *MiniMax) ExactSearch(board othello.Board) int {
	return minimax.Search(board, 60)
}

func (minimax *MiniMax) search(board othello.Board, depth int, maxPlayer bool) int {

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
			return ExactScoreFactor * board.ExactScore()
		}

		child.SwitchTurn()
		return minimax.search(child, depth, !maxPlayer)
	}

	if maxPlayer {
		heur := MinHeuristic
		for gen.Next() {
			childHeur := minimax.search(child, depth-1, !maxPlayer)
			if childHeur > heur {
				heur = childHeur
			}
		}
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
