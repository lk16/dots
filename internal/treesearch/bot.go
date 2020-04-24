package treesearch

import (
	"fmt"
	"io"
	"log"
	"sort"

	"github.com/lk16/dots/internal/othello"
	"github.com/pkg/errors"
)

// Bot is a bot that uses a Heuristic for choosing its moves
type Bot struct {
	searchDepth   int
	exactDepth    int
	writer        io.Writer
	searcher      Searcher
	LifetimeStats Stats
}

// ErrNoMoves means there were no moves for provided board
var ErrNoMoves = errors.New("no moves possible")

// NewBot creates a new Bot
func NewBot(writer io.Writer, searchDepth, exactDepth int, searcher Searcher) *Bot {
	return &Bot{
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer,
		searcher:    searcher}
}

func (bot *Bot) writef(format string, args ...interface{}) {
	formatted := fmt.Sprintf(format, args...)
	_, err := bot.writer.Write([]byte(formatted))
	if err != nil {
		log.Printf("Bot write() error: %s", err)
	}
}

// DoMove computes the best child of a Board
func (bot *Bot) DoMove(board othello.Board) (*othello.Board, error) {
	children := board.GetSortableChildren()

	if len(children) == 0 {
		return nil, ErrNoMoves
	}

	// prevent returning empty Board when bot cannot prevent losing all discs
	afterwards := children[0].Board

	if len(children) == 1 {
		bot.writef("Only one move. Skipping evaluation.\n")
		return &afterwards, nil
	}

	emptiesCount := board.CountEmpties()

	var bestChild othello.Board

	if emptiesCount > bot.exactDepth {
		bestChild = bot.findBestChild(board)
	} else {
		bestChild = bot.findBestChildExact(board)
	}

	return &bestChild, nil
}

func (bot *Bot) findBestChild(board othello.Board) othello.Board {
	var (
		depth     = bot.searchDepth
		alpha     = MinHeuristic
		beta      = MaxHeuristic
		bestChild othello.Board
	)

	children := board.GetSortableChildren()
	bot.writef("Searching with heuristic at depth %d\n", depth)

	if depth > 6 {
		sortAlpha := MinHeuristic
		for i := range children {
			children[i].Heur = bot.searcher.Search(children[i].Board, sortAlpha, MaxHeuristic, 6)
			if children[i].Heur > sortAlpha {
				sortAlpha = children[i].Heur
			}
		}
		sort.Slice(children, func(i, j int) bool {
			return children[i].Heur > children[j].Heur
		})
	}

	sortStats := bot.searcher.GetStats()
	bot.writef("\n\n%12s %63s\n\n", "Sorting:", sortStats.String())
	bot.searcher.ResetStats()

	totalStats := sortStats

	bestChild = children[0].Board

	for i, child := range children {
		heur := bot.searcher.Search(child.Board, alpha, beta, bot.searchDepth)

		childStats := bot.searcher.GetStats()
		bot.searcher.ResetStats()
		totalStats.Add(childStats)
		if heur > alpha {
			alpha = heur
			bestChild = child.Board
			bot.writef("Child %2d/%2d: %8d%55s\n", i+1, len(children), heur, childStats.String())
		} else {
			bot.writef("Child %2d/%2d: %8s%55s\n", i+1, len(children),
				fmt.Sprintf("≤ %d", heur), childStats.String())
		}
	}

	bot.writef("\n%12s %63s\n\n\n", "Total:", totalStats.String())
	bot.LifetimeStats.Add(totalStats)

	return bestChild
}

func (bot *Bot) findBestChildExact(board othello.Board) othello.Board {
	var (
		alpha      = MinScore
		beta       = MaxScore
		depth      = board.CountEmpties()
		totalStats = NewStats()
	)

	bot.writef("Searching for exact solution at depth %d\n", depth)
	bot.searcher.ResetStats()

	children := board.GetChildren()

	bestChild := children[0]

	for i, child := range children {
		heur := bot.searcher.ExactSearch(child, alpha, beta)

		childStats := bot.searcher.GetStats()
		bot.searcher.ResetStats()
		totalStats.Add(childStats)
		if heur > alpha {
			alpha = heur
			bestChild = child
			bot.writef("Child %2d/%2d: %8d%55s\n", i+1, len(children), heur, childStats.String())
		} else {
			bot.writef("Child %2d/%2d: %8s%55s\n", i+1, len(children),
				fmt.Sprintf("≤ %d", heur), childStats.String())
		}
	}

	bot.writef("\n%12s %63s\n\n\n", "Total:", totalStats.String())
	bot.LifetimeStats.Add(totalStats)

	return bestChild
}
