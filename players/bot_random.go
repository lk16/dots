package players

import (
    "dots/board"
)

type BotRandom struct{}

func NewBotRandom() *BotRandom {
    return &BotRandom{}
}

func (bot *BotRandom) DoMove(board board.Board) (afterwards board.Board) {
    afterwards = board
    afterwards.DoRandomMove()
    return
}
