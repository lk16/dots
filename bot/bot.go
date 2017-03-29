package bot

import (
    "dots/board"
)

type BotInterface interface {
    DoMove(*board.Board)
}

type BotRandom struct {}

func NewBotRandom() *BotRandom {
    return &BotRandom{}
}

func (bot *BotRandom) DoMove(board *board.Board) {
    board.DoRandomMoves(1)
}