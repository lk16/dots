package cli_game

import (
	"testing"

	"dots/board"
)

type BotFirstMove struct{}

func (bot *BotFirstMove) DoMove(board board.Board) (afterwards board.Board) {
	afterwards = board.GetChildren()[0]
	return
}

type BotLastMove struct{}

func (bot *BotLastMove) DoMove(board board.Board) (afterwards board.Board) {
	children := board.GetChildren()
	afterwards = children[len(children)-1]
	return
}

func TestCliGameNew(t *testing.T) {
	bot_first := &BotFirstMove{}
	bot_last := &BotLastMove{}
	cli_game := NewCliGame(bot_first, bot_last)

	if cli_game.players[0] != bot_first {
		t.Errorf("Black player.Player not assigned correctly")
	}
	if cli_game.players[1] != bot_last {
		t.Errorf("White player.Player not assigned correctly")
	}
	if cli_game.board != *board.NewBoard() {
		t.Errorf("Board not initialised correctly")
	}
	if cli_game.turn != 0 {
		t.Errorf("Initial turn expected to be 0, got %d", cli_game.turn)
	}
}

func TestCliGameRun(t *testing.T) {
	bot_first := &BotFirstMove{}
	bot_last := &BotLastMove{}
	cli_game := NewCliGame(bot_first, bot_last)

	cli_game.Run()

	if cli_game.skips != 2 {
		t.Errorf("Expected skips to be 2,got %d", cli_game.skips)
	}

	if move_count := cli_game.board.Moves().Count(); move_count != 0 {
		t.Errorf("Board has %d valid move(s) after Run() returned", move_count)
	}

	board := cli_game.board
	board.SwitchTurn()

	if move_count := board.Moves().Count(); move_count != 0 {
		t.Errorf("After passing once, board has %d valid move(s) after Run() returned", move_count)
	}

}

func TestCliGameSkipTurn(t *testing.T) {
	bot_first := &BotFirstMove{}
	bot_last := &BotLastMove{}
	cli_game := NewCliGame(bot_first, bot_last)

	cli_game.SkipTurn()

	if cli_game.turn != 1 {
		t.Errorf("turn expected to be 1, got %d", cli_game.turn)
	}

	if cli_game.skips != 1 {
		t.Errorf("skips expected to be 1, got %d", cli_game.skips)
	}

	expected := *board.NewBoard()
	expected.SwitchTurn()

	if cli_game.board != expected {
		t.Errorf("excpected:\n%s\n\ngot:\n%s\n\n", cli_game.board.AsciiArt(), expected.AsciiArt())
	}
}
