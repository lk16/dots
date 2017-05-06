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

func (bot *BotHeuristic) DoMove(board board.Board) (afterwards board.Board) {

	children := board.GetChildren()

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

	heuristic := bot.heuristic
	depth := bot.search_depth

	do_exact_search := board.CountEmpties() <= bot.exact_depth

	var alpha int
	if do_exact_search {
		alpha = minimax.Min_exact_heuristic
	} else {
		alpha = minimax.Min_heuristic
	}

	for i, child := range children {
		var heur int
		if do_exact_search {
			heur = bot.minimax.ExactSearch(child, alpha)
		} else {
			heur = bot.minimax.Search(child, depth, heuristic, alpha)
		}

		str := fmt.Sprintf("move %d/%d: ", i+1, len(children))
		buff := bytes.NewBufferString(str)

		if heur > alpha {
			buff.WriteString(fmt.Sprintf("%d\n", heur))
			alpha = heur
			afterwards = child
		} else {
			buff.WriteString("not better\n")
		}
		bot.writer.Write(buff.Bytes())
	}

	bot.writer.Write(bytes.NewBufferString("\n").Bytes())
	return
}
