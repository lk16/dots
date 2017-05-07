package frontend

import (
	"dots/board"
)

type Frontend interface {
	OnUpdate(state GameState)
	OnGameEnd(state GameState)
	OnHumanMove(state GameState) board.Board
}
