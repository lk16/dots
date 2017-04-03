package cli_game

import (
	"fmt"

	"dots/board"
	"dots/players"
)

type CliGame struct {
	players [2]players.Player
	board   board.Board
	turn    int
	skips   int
}

// Returns a new CliGame with two players
func NewCliGame(black, white players.Player) *CliGame {
	return &CliGame{
		players: [2]players.Player{
			black,
			white},
		board: *board.NewBoard(),
		turn:  0}
}

// Skips a turn (for when a player has no moves)
func (cli *CliGame) SkipTurn() {
	cli.board.SwitchTurn()
	cli.turn = 1 - cli.turn
	cli.skips++
}

// Runs the game
func (cli *CliGame) Run() {

	cli.skips = 0

	for cli.skips < 2 {

		if cli.skips == 0 {
			fmt.Printf("%s\n", cli.AsciiArt())
		}

		if cli.board.Moves().Count() == 0 {
			cli.SkipTurn()
			continue
		}

		cli.board = cli.players[cli.turn].DoMove(cli.board)
		cli.turn = 1 - cli.turn

	}

	fmt.Printf("%s\n", cli.AsciiArt())
	fmt.Printf("Game over!\n")
}

// Returns a string with ascii-art representing the current board state
func (cli CliGame) AsciiArt() (output string) {
	if cli.turn == 1 {
		cli.board.SwitchTurn()
	}
	return cli.board.AsciiArt()
}
