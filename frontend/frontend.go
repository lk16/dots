package frontend

import (
	"dots/board"
)

// Frontend is an interface for frontends of Controller
type Frontend interface {
	OnUpdate(state GameState)
	OnGameEnd(state GameState)
	OnHumanMove(state GameState) board.Board
}

// Get gets a Frontend by name
func Get(name string) Frontend {

	frontendMap := map[string]func() Frontend{
		"gtk": NewkGtkFrontend,
		"cli": NewCommandLine}

	if newFrontend, ok := frontendMap[name]; ok {
		return newFrontend()
	}

	panic("Invalid frontend name")
}
