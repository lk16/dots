package players

import (
	"github.com/lk16/dots/internal/othello"
)

// BotRandom is a bot that does random moves
type BotRandom struct{}

// NewBotRandom creates a new BotRandom
func NewBotRandom() *BotRandom {
	return &BotRandom{}
}

// DoMove does a random move
func (bot *BotRandom) DoMove(board othello.Board) othello.Board {
	board.DoRandomMove()
	return board
}
