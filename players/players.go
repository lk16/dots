package players

import (
	"os"

	"github.com/lk16/dots/board"
)

// Player is an interface for all structs that can play othello
type Player interface {
	DoMove(board.Board) board.Board
}

// Get gets a player by name and sets the level if applicable
func Get(name string, lvl int, parallel bool) Player {

	if name == "human" {
		return nil
	}

	if name == "random" {
		return NewBotRandom()
	}

	perfectDepth := 2 * lvl
	if lvl > 8 {
		perfectDepth = lvl + 8
	}

	if name == "evolve" {
		return NewBotEvolve(lvl, perfectDepth, Parameters{})
	}

	if name == "heur" {
		return NewBotHeuristic(board.Squared, lvl, perfectDepth,
			os.Stdout, parallel)
	}

	panic("Invalid player name")
}
