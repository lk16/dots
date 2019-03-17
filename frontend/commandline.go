package frontend

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/bits"
	"os"

	"github.com/lk16/dots/othello"
)

// CommandLine is used for command line interaction
type CommandLine struct {
	writer io.Writer
}

// NewCommandLine returns a new CommandLine
func NewCommandLine() Frontend {
	return &CommandLine{
		writer: os.Stdout}
}

// Initialize initializes CommandLine. It does nothing.
func (cli *CommandLine) Initialize() {

}

func (cli *CommandLine) asciiArt(state GameState) {
	swapDiscColors := state.turn == 1
	state.board.ASCIIArt(cli.writer, swapDiscColors)
}

// OnUpdate shows the updated Board with asciiArt
func (cli *CommandLine) OnUpdate(state GameState) {
	cli.asciiArt(state)
}

// OnGameEnd shows game end details
func (cli *CommandLine) OnGameEnd(state GameState) {
	cli.asciiArt(state)

	var whiteCount, blackCount int

	if state.turn == 1 {
		whiteCount = bits.OnesCount64(state.board.Me())
		blackCount = bits.OnesCount64(state.board.Opp())
	} else {
		whiteCount = bits.OnesCount64(state.board.Opp())
		blackCount = bits.OnesCount64(state.board.Me())
	}

	var message string

	if whiteCount > blackCount {
		message = fmt.Sprintf("White wins: %d-%d\n", whiteCount, blackCount)
	} else if whiteCount < blackCount {
		message = fmt.Sprintf("Black wins: %d-%d\n", blackCount, whiteCount)
	} else {
		message = fmt.Sprintf("It's a draw: %d-%d\n", whiteCount, whiteCount)
	}

	messageBytes := bytes.NewBufferString(message).Bytes()
	cli.writer.Write(messageBytes)
}

// OnHumanMove reads a human move from os.Stdin
func (cli *CommandLine) OnHumanMove(state GameState) (afterwards othello.Board) {
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
