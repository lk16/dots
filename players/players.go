package players

import (
	"os"

	"dots/board"
)

// Player is an interface for all structs that can play othello
type Player interface {
	DoMove(board.Board) board.Board
}

// Get gets a player by name and sets the level if applicable
func Get(name string, lvl int) Player {

	if name == "human" {
		return nil
	}

	if name == "random" {
		return NewBotRandom()
	}

	if name == "heur" {
		searchDepth := lvl
		perfectDepth := 2 * lvl
		if perfectDepth > 6 {
			perfectDepth -= 2
		}
		return NewBotHeuristic(Squared, searchDepth, perfectDepth, os.Stdout)
	}

	panic("Invalid player name")
}
