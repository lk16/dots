package players

import (
	"os"

	"dots/board"
	"dots/heuristic"
	"dots/minimax"
)

type Player interface {
	DoMove(board.Board) board.Board
}

func Get(name string, lvl uint) (player Player) {

	if name == "human" {
		player = nil
	} else if name == "random" {
		player = NewBotRandom()
	} else if name == "heur" {
		search_depth := lvl
		perfect_depth := 2 * lvl
		if perfect_depth > 6 {
			perfect_depth -= 2
		}
		player = NewBotHeuristic(heuristic.Squared, &minimax.Mtdf{},
			search_depth, perfect_depth, os.Stdout)
	} else {
		panic("Invalid player name")
	}
	return
}
