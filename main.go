package main

import (
    "dots/board"
)

func main() {
    board := board.NewBoard()
    board.Print()
    board.DoMove(44)
    board.Print()
    board.DoMove(29)
    board.Print()
    board.DoMove(22)
    board.Print()
}
