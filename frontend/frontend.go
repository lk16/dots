package frontend

import (
	"dots/board"
)

type Frontend interface {
	OnUpdate(state GameState)
	OnGameEnd(state GameState)
	OnHumanMove(state GameState) board.Board
}

func Get(name string) (frontend Frontend) {
	if name == "gtk" {
		frontend = NewGtk()
	} else if name == "cli" {
		frontend = NewCommandLine()
	} else {
		panic("Invalid frontend name")
	}
	return
}
