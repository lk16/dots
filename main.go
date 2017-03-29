package main

import (
    "fmt"
    "dots/board"
)

func main() {
    board := board.NewBoard()
    fmt.Printf("%s\n",board.AsciiArt())
    board.DoMove(44)
    fmt.Printf("%s\n",board.AsciiArt())
    board.DoMove(29)
    fmt.Printf("%s\n",board.AsciiArt())
    board.DoMove(22)
    fmt.Printf("%s\n",board.AsciiArt())
}
