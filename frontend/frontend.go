package frontend

import (
	"dots/othello"
)

// Frontend is an interface for frontends of Controller
type Frontend interface {
	OnUpdate(state GameState)
	OnGameEnd(state GameState)
	OnHumanMove(state GameState) othello.Board
}

// Get gets a Frontend by name
func Get(name string) Frontend {

	frontendMap := map[string]func() Frontend{
		"cli": NewCommandLine}

	if newFrontend, ok := frontendMap[name]; ok {
		return newFrontend()
	}

	return nil
}
