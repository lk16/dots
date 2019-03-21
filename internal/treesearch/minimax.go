package treesearch

import (
	"github.com/lk16/dots/internal/heuristics"
	"github.com/lk16/dots/internal/othello"
)

type MiniMax struct{}

func NewMinimax() *MiniMax {
	return &MiniMax{}
}

func (minimax *MiniMax) Name() string {
	return "minimax"
}

func (minimax *MiniMax) Search(board othello.Board, depth int) int {
	return -minimax.search(board, depth, true)
}

func (minimax *MiniMax) ExactSearch(board othello.Board) int {
	return minimax.Search(board, 60)
}

func (minimax *MiniMax) search(board othello.Board, depth int, maxPlayer bool) int {

	if depth == 0 {
		heur := heuristics.Squared(board)
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
