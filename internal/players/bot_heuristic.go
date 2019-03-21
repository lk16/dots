package players

import (
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"github.com/lk16/dots/internal/treesearch"
	"io"
	"log"
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

func (bot *BotHeuristic) write(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	_, err := bot.writer.Write([]byte(formatted))
	if err != nil {
		log.Printf("BotHeuristic write() error: %s", err)
	}
}

// DoMove does a move
func (bot *BotHeuristic) DoMove(board othello.Board) othello.Board {

	children := board.GetChildren()

	if len(children) == 0 {
		return board
	}

	// prevent returning empty Board when bot cannot prevent losing all discs
	afterwards := children[0]

	if len(children) == 1 {
		bot.write("Only one move. Skipping evaluation.\n")
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

		bot.write("Child %2d/%2d: %d\n", i+1, len(children), heur)

		if heur > alpha {
			alpha = heur
			afterwards = child
		}

	}

	bot.write("\n\n")
	return afterwards
}
