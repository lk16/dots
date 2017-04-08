package players

import (
	"fmt"

	"dots/board"
	"dots/minimax"
)

type BotHeuristic struct {
	heuristic    minimax.Heuristic
	minimax      minimax.Interface
	search_depth uint
	exact_depth  uint
}

func NewBotHeuristic(heuristic minimax.Heuristic, minimax minimax.Interface,
	search_depth, exact_depth uint) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:    heuristic,
		minimax:      minimax,
		search_depth: search_depth,
		exact_depth:  exact_depth}
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
		fmt.Printf("Only one move. Skipping evaluation.\n")
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

		fmt.Printf("move %d/%d: ", i+1, len(children))

		if heur > alpha {
			fmt.Printf("%d\n", heur)
			alpha = heur
			afterwards = child
		} else {
			fmt.Printf("not better\n")
		}
	}
	fmt.Printf("\n")

	return
}
