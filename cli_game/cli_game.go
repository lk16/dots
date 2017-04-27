package cli_game

import (
	"bytes"
	"fmt"
	"io"

	"dots/board"
	"dots/players"
)

type CliGame struct {
	players [2]players.Player
	board   board.Board
	turn    int
	writer  io.Writer
}

// Returns a new CliGame with two players
func NewCliGame(black, white players.Player, writer io.Writer) (cli *CliGame) {
	cli = &CliGame{}

	if white == nil || black == nil {
		panic("A player cannot be nil!")
	}

	cli.players = [2]players.Player{black, white}
	cli.writer = writer
	return
}

// Skips a turn (for when a player has no moves)
func (cli *CliGame) skipTurn() {
	cli.board.SwitchTurn()
	cli.turn = 1 - cli.turn
}

// Lets the player to move do a move
func (cli *CliGame) doMove() {
	cli.asciiArt()
	cli.writer.Write(bytes.NewBufferString("\n").Bytes())

	cli.board = cli.players[cli.turn].DoMove(cli.board)
	cli.turn = 1 - cli.turn
}

// Returns whether the player to move can do a move
func (cli *CliGame) canMove() (can_move bool) {
	moves_count := cli.board.Moves().Count()
	can_move = (moves_count != 0)
	return
}

func (cli *CliGame) onNewGame() {
	cli.board = *board.NewBoard()
	cli.turn = 0
}

func (cli *CliGame) onGameEnd() {
	cli.asciiArt()

	board := cli.board
	if cli.turn == 1 {
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

func (cli *CliGame) gameRunning() (running bool) {
	return !cli.board.IsLeaf()
}

func (cli CliGame) asciiArt() {
	swap_disc_colors := cli.turn == 1
	cli.board.AsciiArt(cli.writer, swap_disc_colors)
}

// Runs the game
func (cli *CliGame) Run() {

	cli.onNewGame()
	for cli.gameRunning() {
		if !cli.canMove() {
			cli.skipTurn()
			continue
		}
		cli.doMove()
	}
	cli.onGameEnd()
}
