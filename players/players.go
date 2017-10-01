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
func Get(name string, lvl int, parallel bool) Player {

	if name == "human" {
		return nil
	}

	if name == "random" {
		return NewBotRandom()
	}

	if name == "beam" {
		return &BotBeam{}
	}

	if name == "heur" {
		perfectDepth := 2 * lvl
		if lvl > 8 {
			perfectDepth = lvl + 8
		}
		return NewBotHeuristic(board.Squared, lvl, perfectDepth,
			os.Stdout, parallel)
	}

	panic("Invalid player name")
}
