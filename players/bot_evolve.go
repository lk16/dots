package players

import (
	"dots/board"
	"fmt"
	"math/bits"
)

// BotEvolve is a parametrised bot designed to work with Evolution
type BotEvolve struct {
	searchDepth int
	exactDepth  int
}

// NewBotEvolve creates a new BotEvolve
func NewBotEvolve(searchDepth, exactDepth int) (bot *BotEvolve) {
	return &BotEvolve{
		searchDepth: searchDepth,
		exactDepth:  exactDepth}
}

// DoMove does a move
func (bot *BotEvolve) DoMove(b board.Board) (afterwards board.Board) {

	children := b.GetChildren()
	afterwards = children[0]

	if len(children) == 1 {
		return
	}

	alpha := -100
	beta := 100
	depth := bot.searchDepth

	if b.CountEmpties() <= bot.exactDepth {
		alpha = board.MinScore
		beta = board.MaxScore
		depth = b.CountEmpties()
	}

	var heur int

	for i, child := range children {

		if b.CountEmpties() > bot.exactDepth {
			heur = -bot.alphaBeta(&child, -beta, -alpha, depth)
		} else {
			heur = -bot.alphaBetaExact(&child, -beta, -alpha)
		}

		if heur > alpha {
			alpha = heur
			afterwards = child
		}
		fmt.Printf("Child %d/%d\talpha = %d\n", i+1, len(children), alpha)
	}

	fmt.Printf("---\n")
	return
}

func (bot *BotEvolve) heuristic(b board.Board) int {
	return bits.OnesCount64(b.Opp()) - bits.OnesCount64(b.Me())
}

func (bot *BotEvolve) alphaBeta(b *board.Board, alpha, beta, depth int) int {

	if depth == 0 {
		return bot.heuristic(*b)
	}

	gen := board.NewGenerator(b, 0)

	if !gen.HasMoves() {
		if b.OpponentMoves() == 0 {
			return board.ExactScoreFactor * b.ExactScore()
		}

		b.SwitchTurn()
		heur := -bot.alphaBeta(b, -beta, -alpha, depth)
		b.SwitchTurn()
		return heur
	}

	heur := alpha
	for gen.Next() {
		childHeur := -bot.alphaBeta(b, -beta, -alpha, depth-1)
		if childHeur >= beta {
			gen.RestoreParent()
			return beta
		}
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur
}

func (bot *BotEvolve) alphaBetaExact(b *board.Board, alpha, beta int) int {

	gen := board.NewGenerator(b, 0)

	if !gen.HasMoves() {
		if b.OpponentMoves() == 0 {
			return b.ExactScore()
		}

		b.SwitchTurn()
		heur := -bot.alphaBetaExact(b, -beta, -alpha)
		b.SwitchTurn()
		return heur
	}

	heur := alpha
	for gen.Next() {
		childHeur := -bot.alphaBetaExact(b, -beta, -alpha)
		if childHeur >= beta {
			gen.RestoreParent()
			return beta
		}
		if childHeur > heur {
			heur = childHeur
		}
	}
	return heur
}
