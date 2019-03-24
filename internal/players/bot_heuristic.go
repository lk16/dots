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
		bot.write("No moves found. Something went wrong.\n")
		return board
	}

	// prevent returning empty Board when bot cannot prevent losing all discs
	afterwards := children[0]

	if len(children) == 1 {
		bot.write("Only one move. Skipping evaluation.\n")
		return children[0]
	}

	alpha := treesearch.MinHeuristic
	beta := treesearch.MaxHeuristic

	var depth int
	if board.CountEmpties() <= bot.exactDepth {
		depth = board.CountEmpties()
	} else {
		depth = bot.searchDepth
	}

	search := treesearch.NewMtdf(alpha, beta)

	for i, child := range children {

		search.SetAlphaBeta(alpha, beta)

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

	bot.write("%d nodes in %.3f seconds = %dK nodes/second\n",
		search.Stats.Nodes, search.Stats.Duration.Seconds(), int(search.Stats.NodesPerSecond())/1000)

	bot.write("\n\n")
	return afterwards
}
