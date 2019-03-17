package players

import (
	"bytes"
	"fmt"
	"github.com/lk16/dots/othello"
	"github.com/lk16/dots/treesearch"
	"io"
)

// Heuristic is a function that estimates how promising a Board is.
type Heuristic func(othello.Board) int

// BotHeuristic is a bot that uses a Heuristic for choosing its moves
type BotHeuristic struct {
	searchDepth int
	exactDepth  int
	writer      io.Writer
}

// NewBotHeuristic creates a new BotHeuristic
func NewBotHeuristic(writer io.Writer, searchDepth, exactDepth int) *BotHeuristic {

	return &BotHeuristic{
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer}
}

// DoMove does a move
func (bot *BotHeuristic) DoMove(board othello.Board) othello.Board {

	children := board.GetChildren()

	if len(children) == 0 {
		return board
	}

	// prevent returning empty othello when bot cannot prevent losing all discs
	afterwards := children[0]

	if len(children) == 1 {
		buff := bytes.NewBufferString("Only one move. Skipping evaluation.\n")
		_, _ = bot.writer.Write(buff.Bytes())
		return afterwards
	}

	var alpha, beta int
	var depth int

	if board.CountEmpties() <= bot.exactDepth {
		alpha = treesearch.MinScore
		beta = treesearch.MaxScore
		depth = board.CountEmpties()
	} else {
		depth = bot.searchDepth

		// HACK: stumbling upon an exact solution
		// takes forever to compute. we set limits to solve that for now.
		alpha = -100
		beta = 100
	}

	for i, child := range children {

		search := (treesearch.Interface)(treesearch.NewMtdf(alpha, beta))

		var heur int
		if board.CountEmpties() <= bot.exactDepth {
			heur = search.ExactSearch(child)
		} else {
			heur = search.Search(child, depth)
		}

		buff := bytes.NewBufferString(fmt.Sprintf("Child %2d/%2d: %d\n", i+1, len(children), heur))

		_, _ = bot.writer.Write(buff.Bytes())

		if heur > alpha {
			alpha = heur
			afterwards = child
		}

	}

	_, _ = bot.writer.Write(bytes.NewBufferString("\n\n").Bytes())
	return afterwards
}
