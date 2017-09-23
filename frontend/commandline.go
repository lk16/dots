package frontend

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/bits"
	"os"

	"dots/board"
)

type CommandLine struct {
	writer io.Writer
}

// Create new CommandLine
func NewCommandLine() (cli *CommandLine) {
	cli = &CommandLine{
		writer: os.Stdout}
	return
}

// Initialize CommandLine (does nothing)
func (cli *CommandLine) Initialize() {

}

// Print GameState to cli.writer
func (cli *CommandLine) asciiArt(state GameState) {
	swap_disc_colors := state.turn == 1
	state.board.ASCIIArt(cli.writer, swap_disc_colors)
}

// Print GameState to cli.writer on update
func (cli *CommandLine) OnUpdate(state GameState) {
	cli.asciiArt(state)
}

// Show game end details when game ends
func (cli *CommandLine) OnGameEnd(state GameState) {
	cli.asciiArt(state)

	board := state.board
	if state.turn == 1 {
		board.SwitchTurn()
	}

	white_count := bits.OnesCount64(board.Opp())
	black_count := bits.OnesCount64(board.Me())

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

// Read human move from os.Stdin
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

		col := line[0] - 'a'
		row := line[1] - '1'

		if col < 0 || row < 0 || col > 7 || row > 7 {
			continue
		}

		index := int(8*row + col)

		mask := uint64(1) << uint(index)

		if moves&mask == 0 {
			continue
		}

		afterwards = state.board
		afterwards.DoMove(index)
		return
	}
}
