package players

import (
	"dots/board"
)

type BotRandom struct{}

// Creates a new BotRandom
func NewBotRandom() *BotRandom {
	return &BotRandom{}
}

// Does a random move
func (bot *BotRandom) DoMove(board board.Board) (afterwards board.Board) {
	afterwards = board
	afterwards.DoRandomMove()
	return
}
