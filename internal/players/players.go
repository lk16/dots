package players

import (
	"os"

	"github.com/lk16/dots/internal/othello"
)

// Player is an interface for all structs that can play othello
type Player interface {
	DoMove(othello.Board) othello.Board
}

// Get gets a player by name and sets the level if applicable
func Get(name string, lvl int, parallel bool) Player {

	if name == "human" {
		return nil
	}

	if name == "heur" {
		perfectDepth := 2 * lvl
		if lvl > 8 {
			perfectDepth = lvl + 8
		}
		return NewBotHeuristic(os.Stdout, lvl, perfectDepth)
	}

	panic("Invalid player name")
}
