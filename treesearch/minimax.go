package treesearch

import (
	"dots/heuristics"
	"dots/othello"
)

type MiniMax struct {
	board othello.Board
	depth int
	sign  int
}

func NewMinimax() *MiniMax {
	return &MiniMax{}
}

func (minimax *MiniMax) Name() string {
	return "minimax"
}

func (minimax *MiniMax) Search(board othello.Board, depth int) int {
	minimax.board = board
	minimax.depth = depth
	minimax.sign = 1
	return minimax.search()
}

func (minimax *MiniMax) ExactSearch(board othello.Board) int {
	return minimax.Search(board, 60)
}

func (minimax *MiniMax) search() int {

	if minimax.depth == 0 {
		return minimax.sign * heuristics.Squared(minimax.board)
	}

	gen := othello.NewGenerator(&minimax.board, 0)

	if !gen.HasMoves() {
		if minimax.board.OpponentMoves() == 0 {
			return minimax.sign * ExactScoreFactor * minimax.board.ExactScore()
		}

		minimax.board.SwitchTurn()
		minimax.sign = -minimax.sign
		heur := minimax.search()
		minimax.sign = -minimax.sign
		minimax.board.SwitchTurn()
		return heur
	}

	if minimax.sign == 1 {
		heur := MinHeuristic
		for gen.Next() {
			minimax.depth--
			minimax.sign = -minimax.sign
			childHeur := minimax.search()
			minimax.sign = -minimax.sign
			minimax.depth++
			if childHeur > heur {
				heur = childHeur
			}
		}
		return heur
	}

	heur := MaxHeuristic
	for gen.Next() {
		minimax.depth--
		minimax.sign = -minimax.sign
		childHeur := minimax.search()
		minimax.sign = -minimax.sign
		minimax.depth++
		if childHeur < heur {
			heur = childHeur
		}
	}
	return heur
}
