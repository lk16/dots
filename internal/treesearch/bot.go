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
	searchDepth   int
	exactDepth    int
	writer        io.Writer
	search        Interface
	LifetimeStats Stats
}

// NewBot creates a new Bot
func NewBot(writer io.Writer, searchDepth, exactDepth int) *Bot {

	return &Bot{
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer,
		search:      NewPvs()}
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

	isExact := bot.exactDepth >= board.CountEmpties()

	var depth int
	if isExact {
		depth = board.CountEmpties()
	} else {
		depth = bot.searchDepth
	}

	if isExact {
		bot.write("Searching for exact solution at depth %d\n", depth)
	} else {
		bot.write("Searching with heuristic at depth %d\n", depth)
	}

	children := board.GetSortableChildren()

	if len(children) == 0 {
		return nil, fmt.Errorf("no moves possible")
	}

	// prevent returning empty Board when bot cannot prevent losing all discs
	afterwards := children[0].Board

	if len(children) == 1 {
		bot.write("Only one move. Skipping evaluation.\n")
		return &children[0].Board, nil
	}

	if (!isExact) && (depth > 6) {
		for i := range children {
			children[i].Heur = bot.search.Search(children[i].Board, MinHeuristic, MaxHeuristic, 6)
		}
		sort.Slice(children, func(i, j int) bool {
			return children[i].Heur > children[j].Heur
		})
	}

	sortStats := bot.search.GetStats()
	bot.write("\n\n%12s %63s\n\n", "Sorting:", sortStats.String())
	bot.search.ResetStats()

	totalStats := sortStats

	var alpha, beta int

	if isExact {
		alpha = MinScore
		beta = MaxScore
	} else {
		alpha = MinHeuristic
		beta = MaxHeuristic
	}

	for i, child := range children {

		var heur int
		if isExact {
			heur = bot.search.ExactSearch(child.Board, alpha, beta)
		} else {
			heur = bot.search.Search(child.Board, alpha, beta, bot.searchDepth)
		}

		childStats := bot.search.GetStats()
		bot.search.ResetStats()
		totalStats.Add(childStats)
		if heur > alpha {
			alpha = heur
			afterwards = child.Board
			bot.write("Child %2d/%2d: %8d%55s\n", i+1, len(children), heur, childStats.String())
		} else {
			bot.write("Child %2d/%2d: %8s%55s\n", i+1, len(children),
				fmt.Sprintf("≤ %d", heur), childStats.String())
		}
	}

	bot.write("\n%12s %63s\n\n\n", "Total:", totalStats.String())
	bot.LifetimeStats.Add(totalStats)

	return &afterwards, nil
}
