package players

import (
	"dots/board"
)

// BotRandom is a bot that does random moves
type BotRandom struct{}

// NewBotRandom creates a new BotRandom
func NewBotRandom() *BotRandom {
	return &BotRandom{}
}

// DoMove does a random move
func (bot *BotRandom) DoMove(board board.Board) board.Board {
	board.DoRandomMove()
	return board
}
