package players

import (
	"dots/board"
	"fmt"
)

// BotBeam uses a Beamsearch-like approach for finding the best move
type BotBeam struct {
}

func doBeamSearch(b board.Board, depth, perfectDepth int) (heur int) {

	empties := b.CountEmpties()

	//b.ASCIIArt(os.Stdout, false)

	if empties <= perfectDepth {
		return board.AlphaBeta(&b, board.MinHeuristic, board.MaxHeuristic, 60)
	}

	children := b.GetChildren()

	if len(children) == 0 {
		return board.ExactScoreFactor * b.ExactScore()
	}

	bestHeur := board.MinHeuristic
	bestChild := children[0]
	for _, child := range children {
		childHeur := board.AlphaBeta(&child, bestHeur, board.MaxHeuristic, depth)
		if childHeur > bestHeur {
			bestHeur = childHeur
			bestChild = child
		}
	}
	fmt.Printf("%d ", bestHeur)
	return doBeamSearch(bestChild, depth, perfectDepth)
}

// DoMove computes the best move for a a board
func (bot *BotBeam) DoMove(before board.Board) board.Board {

	children := before.GetChildren()

	if len(children) == 1 {
		return children[0]
	}

	bestHeur := board.MinHeuristic
	bestChild := children[0]
	for i, child := range children {
		childHeur := doBeamSearch(child, 4, 8)
		fmt.Printf("\nChild %d/%d:\t%d\n", i+1, len(children), childHeur)
		if childHeur > bestHeur {
			bestHeur = childHeur
			bestChild = child
		}
	}
	fmt.Printf("\n\n")
	return bestChild
}
