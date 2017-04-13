package players

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"dots/board"
)

type Human struct {
	reader io.Reader
	writer io.Writer
}

func NewHuman(reader io.Reader) (human *Human) {
	human = &Human{}
	human.reader = reader
	human.writer = os.Stdout
	return
}

func (human *Human) DoMove(board board.Board) (afterwards board.Board) {

	afterwards = board
	moves := board.Moves()

	fmt.Printf("\n")

	scanner := bufio.NewScanner(human.reader)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) != 2 {
			continue
		}

		col := uint(line[0] - 'a')
		row := uint(line[1] - '1')

		if col > 7 || row > 7 {
			continue
		}

		index := 8*row + col

		if !moves.TestBit(index) {
			continue
		}

		afterwards.DoMove(index)
		return
	}

	panic("Error processing input.")

}
