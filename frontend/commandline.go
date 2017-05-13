package frontend

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"dots/board"
)

type CommandLine struct {
	writer io.Writer
}

func NewCommandLine() (cli *CommandLine) {
	cli = &CommandLine{
		writer: os.Stdout}
	return
}

func (cli *CommandLine) Initialize() {

}

func (cli *CommandLine) asciiArt(state GameState) {
	swap_disc_colors := state.turn == 1
	state.board.AsciiArt(cli.writer, swap_disc_colors)
}

func (cli *CommandLine) OnUpdate(state GameState) {
	cli.asciiArt(state)
}

func (cli *CommandLine) OnGameEnd(state GameState) {
	cli.asciiArt(state)

	board := state.board
	if state.turn == 1 {
		board.SwitchTurn()
	}

	white_count := board.Opp().Count()
	black_count := board.Me().Count()

	var str string

	if white_count > black_count {
		str = fmt.Sprintf("White wins: %d-%d\n", white_count, black_count)
	} else if white_count < black_count {
		str = fmt.Sprintf("Black wins: %d-%d\n", black_count, white_count)
	} else {
		str = fmt.Sprintf("It's a draw: %d-%d\n", white_count, white_count)
	}

	bytes := bytes.NewBufferString(str).Bytes()
	cli.writer.Write(bytes)
}

func (cli *CommandLine) OnHumanMove(state GameState) (afterwards board.Board) {
	moves := state.board.Moves()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		prompt := bytes.NewBufferString("> ").Bytes()
		cli.writer.Write(prompt)

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

		afterwards = state.board
		afterwards.DoMove(index)
		return
	}
}
