package players

import (
	"dots/board"
)

type Player interface {
	DoMove(board.Board) board.Board
}
