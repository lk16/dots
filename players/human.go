package players

import (
	"bufio"
	"fmt"
	"os"

	"dots/board"
)

type Human struct{}

func (human *Human) DoMove(board board.Board) (afterwards board.Board) {
	afterwards = board
	reader := bufio.NewReader(os.Stdin)
	moves := board.Moves()
	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')

		if len(text) != 3 {
			continue
		}
		if text[0] < 'a' || text[0] > 'h' {
			continue
		}
		if text[1] < '1' || text[1] > '8' {
			continue
		}
		index := uint(8*(text[1]-'1') + (text[0] - 'a'))

		if !moves.TestBit(index) {
			continue
		}

		afterwards.DoMove(index)
		return
	}

}
