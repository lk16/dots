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

	if name == "beam" {
		return &BotBeam{}
	}

	if name == "heur" {
		return NewBotHeuristic(board.Squared, lvl, 2*lvl, os.Stdout)
	}

	panic("Invalid player name")
}
