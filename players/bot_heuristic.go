package players

import (
	"bytes"
	"fmt"
	"io"

	"dots/board"
	"dots/minimax"
)

type BotHeuristic struct {
	heuristic    minimax.Heuristic
	minimax      minimax.Interface
	search_depth uint
	exact_depth  uint
	writer       io.Writer
}

func NewBotHeuristic(heuristic minimax.Heuristic, minimax minimax.Interface,
	search_depth, exact_depth uint, writer io.Writer) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:    heuristic,
		minimax:      minimax,
		search_depth: search_depth,
		exact_depth:  exact_depth,
		writer:       writer}
	return
}

func (bot *BotHeuristic) DoMove(b board.Board) (afterwards board.Board) {

	children := b.GetChildren()

	if len(children) == 0 {
		panic("Cannot do move, because there are no moves.")
	}

	// prevent returning empty board when bot cannot prevent losing all discs
	afterwards = children[0]

	if len(children) == 1 {
		buff := bytes.NewBufferString("Only one move. Skipping evaluation.\n")
		bot.writer.Write(buff.Bytes())
		return
	}

	log_child_eval := func(child_id, child_count, heur, alpha int) {
		str := fmt.Sprintf("move %d/%d: ", child_id+1, child_count)
		buff := bytes.NewBufferString(str)
		if heur > alpha {
			buff.WriteString(fmt.Sprintf("%d\n", heur))
		} else {
			buff.WriteString("not better\n")
		}
		bot.writer.Write(buff.Bytes())
	}

	var child_eval func(board.Board, uint, minimax.Heuristic, int) int
	var alpha int

	if b.CountEmpties() <= bot.exact_depth {
		alpha = minimax.Min_exact_heuristic
		// wrapper function to achieve same prototype for ExactSearch() as Search()
		// it just drops arguments depth and heuristic
		child_eval = func(child board.Board, depth uint,
			heuristic minimax.Heuristic, alpha int) int {
			return bot.minimax.ExactSearch(child, alpha)
		}
	} else {
		alpha = minimax.Min_heuristic
		child_eval = bot.minimax.Search
	}

	for i, child := range children {
		heur := child_eval(child, bot.search_depth, bot.heuristic, alpha)
		log_child_eval(i+1, len(children), heur, alpha)
		if heur > alpha {
			alpha = heur
			afterwards = child
		}
	}

	bot.writer.Write(bytes.NewBufferString("\n").Bytes())
	return
}
