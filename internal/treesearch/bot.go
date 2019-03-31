package treesearch

import (
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"io"
	"log"
)

// Bot is a bot that uses a Heuristic for choosing its moves
type Bot struct {
	searchDepth int
	exactDepth  int
	writer      io.Writer
}

// NewBot creates a new Bot
func NewBot(writer io.Writer, searchDepth, exactDepth int) *Bot {

	return &Bot{
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer}
}

func (bot *Bot) write(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	_, err := bot.writer.Write([]byte(formatted))
	if err != nil {
		log.Printf("Bot write() error: %s", err)
	}
}

// DoMove computes the best child of a Board
func (bot *Bot) DoMove(board othello.Board) (*othello.Board, error) {

	children := board.GetChildren()

	// prevent returning empty Board when bot cannot prevent losing all discs
	afterwards := children[0]

	if len(children) == 1 {
		bot.write("Only one move. Skipping evaluation.\n")
		return &children[0], nil
	}

	alpha := MinHeuristic
	beta := MaxHeuristic

	var depth int
	if board.CountEmpties() <= bot.exactDepth {
		depth = board.CountEmpties()
	} else {
		depth = bot.searchDepth
	}

	search := NewMtdf(alpha, beta)

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
	return &afterwards, nil
}
