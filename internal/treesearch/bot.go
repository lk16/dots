package treesearch

import (
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"io"
	"log"
	"sort"
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

	children := board.GetSortableChildren()

	// prevent returning empty Board when bot cannot prevent losing all discs
	afterwards := children[0].Board

	if len(children) == 0 {
		return nil, fmt.Errorf("no moves possible")
	}

	if len(children) == 1 {
		bot.write("Only one move. Skipping evaluation.\n")
		return &children[0].Board, nil
	}

	alpha := MinHeuristic
	beta := MaxHeuristic

	var depth int
	if board.CountEmpties() <= bot.exactDepth {
		depth = board.CountEmpties()
	} else {
		depth = bot.searchDepth
	}

	search := NewPvs()

	stats := NewStats()
	stats.StartClock()

	if depth > 4 {
		for i := range children {
			children[i].Heur = search.Search(children[i].Board, MinHeuristic, MaxHeuristic, 4)
		}
		sort.Slice(children, func(i, j int) bool {
			return children[i].Heur > children[j].Heur
		})
	}

	for i, child := range children {

		var heur int
		if board.CountEmpties() <= bot.exactDepth {
			heur = search.ExactSearch(child.Board, alpha, beta)
		} else {
			heur = search.Search(child.Board, alpha, beta, depth)
		}

		bot.write("Child %2d/%2d: %d\n", i+1, len(children), heur)

		if heur > alpha {
			alpha = heur
			afterwards = child.Board
		}
	}

	stats.StopClock()

	stats.Add(search.GetStats())

	bot.write("%s nodes in %.3f seconds = %s nodes/second\n",
		FormatBigNumber(stats.Nodes), stats.Duration.Seconds(), FormatBigNumber(stats.NodesPerSecond()))

	bot.write("\n\n")
	return &afterwards, nil
}
