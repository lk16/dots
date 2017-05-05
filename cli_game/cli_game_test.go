package cli_game

import (
	"bytes"
	"fmt"
	"testing"

	"dots/board"
	"dots/players"
)

// Fails if panic() is not called
func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("panic() was not called")
	}
}

func TestCliGameNew(t *testing.T) {
	buff := new(bytes.Buffer)

	test := func(black, white players.Player) {

		if black == nil || white == nil {
			defer assertPanic(t)
		}

		cli_game := NewCliGame(black, white, buff)

		if black != nil && white != nil {
			if cli_game.players[0] != black {
				t.Errorf("Black player.Player not assigned correctly")
			}
			if cli_game.players[1] != white {
				t.Errorf("White player.Player not assigned correctly")
			}
			if cli_game.writer != buff {
				t.Errorf("Writer is not assigned correctly")
			}
			if buff.String() != "" {
				t.Errorf("Expected no output, got \"%s\"", buff.String())
			}
		}
	}

	for _, white := range []players.Player{&players.BotRandom{}, nil} {
		for _, black := range []players.Player{&players.BotRandom{}, nil} {
			test(black, white)
		}
	}
}

func TestCliGameRun(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}
	buff := new(bytes.Buffer)

	cli_game := NewCliGame(black, white, buff)
	cli_game.Run()

	if !cli_game.board.IsLeaf() {
		board_ascii_buff := new(bytes.Buffer)
		cli_game.board.AsciiArt(board_ascii_buff, false)
		t.Errorf("Game state at end of game is not a leaf:\n%s\n\n", board_ascii_buff.String())
	}
}

func TestCliGameSkipTurn(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}
	got_output := new(bytes.Buffer)
	cli_game := NewCliGame(black, white, got_output)
	got_output.Reset()

	cli_game.onNewGame()

	for i := 0; i < 2; i++ {
		expected_board := cli_game.board
		expected_board.SwitchTurn()
		expected_turn := 1 - cli_game.turn
		cli_game.skipTurn()

		if cli_game.board != expected_board {
			expected_board_ascii := new(bytes.Buffer)
			expected_board.AsciiArt(expected_board_ascii, false)

			got_board_ascii := new(bytes.Buffer)
			cli_game.board.AsciiArt(got_board_ascii, false)

			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
				expected_board_ascii.String(), got_board_ascii.String())
		}

		if cli_game.turn != expected_turn {
			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n", expected_turn, cli_game.turn)
		}

		if got_output.String() != "" {
			t.Errorf("Expected no output, got \"%s\"", got_output.String())
		}

		got_output.Reset()
	}
}

func TestCliGameDoMove(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}

	got_output := new(bytes.Buffer)
	expected_output := new(bytes.Buffer)

	cli_game := NewCliGame(black, white, got_output)
	cli_game.onNewGame()

	for i := 0; i < 2; i++ {
		cli_game.board.AsciiArt(expected_output, cli_game.turn == 1)
		expected_output.WriteString("\n")

		expected_turn := 1 - cli_game.turn
		expected_discs := cli_game.board.CountDiscs() + 1

		cli_game.doMove()

		if got_output.String() != expected_output.String() {
			t.Errorf("Expected output:\n%s\n\nGot output:\n%s\n\n",
				expected_output.String(), got_output.String())
		}

		if cli_game.turn != expected_turn {
			t.Errorf("Expected turn %d, got %d\n", expected_turn, cli_game.turn)
		}

		if cli_game.board.CountDiscs() != expected_discs {
			t.Errorf("Expected %d discs, got %d", expected_discs, cli_game.board.CountDiscs())
		}

		got_output.Reset()
		expected_output.Reset()
	}
}

func TestCliGameCanMove(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}

	got_output := new(bytes.Buffer)

	cli_game := NewCliGame(black, white, got_output)
	cli_game.onNewGame()
	got_output.Reset()

	test := func(board board.Board) {
		cli_game.board = board
		clone := *cli_game
		expected := cli_game.board.Moves().Count() != 0
		got := cli_game.canMove()

		if expected != got {
			t.Errorf("Expected %t, got %t", expected, got)
		}

		if clone != *cli_game {
			t.Errorf("CliGame was modified")
		}

		if got_output.String() != "" {
			t.Errorf("Expected no output, got \"%s\"", got_output)
		}
	}

	// board with moves
	test(*board.NewBoard())

	// board without moves
	test(*board.RandomBoard(64))

}

func TestCliGameOnNewGame(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}

	got_output := new(bytes.Buffer)

	cli_game := NewCliGame(black, white, got_output)
	got_output.Reset()

	cli_game.onNewGame()

	if cli_game.board != *board.NewBoard() {
		got_board_buff := new(bytes.Buffer)
		cli_game.board.AsciiArt(got_board_buff, false)
		t.Errorf("Expected initial state board\nGot:\n%s\n\n", got_board_buff.String())
	}

	if cli_game.turn != 0 {
		t.Errorf("Expected turn 0,got %d", cli_game.turn)
	}

}

func TestCliGameOnGameEnd(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}

	got_output := new(bytes.Buffer)
	expected_output := new(bytes.Buffer)
	cli_game := NewCliGame(black, white, got_output)

	test := func(board board.Board, turn int) {
		cli_game.turn = turn
		cli_game.board = board

		clone := *cli_game
		cli_game.onGameEnd()

		if clone != *cli_game {
			t.Errorf("CliGame was modified")
		}

		board.AsciiArt(expected_output, turn == 1)

		black_count := board.Me().Count()
		white_count := board.Opp().Count()

		if turn == 1 {
			black_count, white_count = white_count, black_count
		}

		var result_line string

		if white_count > black_count {
			result_line = fmt.Sprintf("White wins: %d-%d\n", white_count, black_count)
		} else if black_count > white_count {
			result_line = fmt.Sprintf("Black wins: %d-%d\n", black_count, white_count)
		} else {
			result_line = fmt.Sprintf("It's a draw: %d-%d\n", white_count, black_count)
		}
		expected_output.WriteString(result_line)

		if expected_output.String() != got_output.String() {
			t.Errorf("Expected output:\n%s\n\nGot:\n%s\n\n",
				expected_output.String(), got_output.String())
		}

		got_output.Reset()
		expected_output.Reset()
	}

	// draw
	test(*board.NewBoard(), 0)

	// win for black
	test(*board.RandomBoard(5), 1)

	// win for white
	test(*board.RandomBoard(5), 0)

}

func TestCliGameRunning(t *testing.T) {

	black := &players.BotRandom{}
	white := &players.BotRandom{}

	got_output := new(bytes.Buffer)
	cli_game := NewCliGame(black, white, got_output)

	test := func(board board.Board) {
		expected := true

		if board.Moves().Count() == 0 {
			board.SwitchTurn()
			if board.Moves().Count() == 0 {
				expected = false
			}
			board.SwitchTurn()
		}

		cli_game.board = board
		clone := *cli_game
		got := cli_game.gameRunning()

		if clone != *cli_game {
			t.Errorf("CliGame was modified.\n")
		}

		if expected != got {
			board_buff := new(bytes.Buffer)
			board.AsciiArt(board_buff, false)
			t.Errorf("Expected %t, got %t for board\n%s\n\n",
				expected, got, board_buff.String())
		}

		if got_output.String() != "" {
			t.Errorf("Expected no output, got \"%s\"", got_output.String())
			got_output.Reset()
		}
	}

	// unfinished game
	test(*board.NewBoard())

	// finished game
	test(*board.RandomBoard(64))
}

func TestCliGameAsciiArt(t *testing.T) {

	black := &players.BotRandom{}
	white := &players.BotRandom{}
	got_output := new(bytes.Buffer)

	cli_game := NewCliGame(black, white, got_output)

	test := func(board board.Board, turn int) {
		cli_game.board = board
		cli_game.turn = turn

		expected_output := new(bytes.Buffer)
		board.AsciiArt(expected_output, turn == 1)

		cli_game.writer = got_output
		cli_game.asciiArt()

		if got_output.String() != expected_output.String() {
			t.Errorf("Expected output:\n%s\n\nGot output:\n%s\n\n",
				expected_output.String(), got_output.String())
		}

		got_output.Reset()
	}

	// black to move
	test(*board.NewBoard(), 0)

	// white to move
	test(*board.NewBoard(), 1)

}
