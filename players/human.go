package players

import (
	"bufio"
	"bytes"
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

	scanner := bufio.NewScanner(human.reader)

	for {
		prompt := bytes.NewBufferString("> ").Bytes()
		human.writer.Write(prompt)

		if !scanner.Scan() {
			panic("Error processing input.")
		}
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

}
