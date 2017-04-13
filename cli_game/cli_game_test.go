package cli_game

import (
	"bytes"
	"testing"

	"dots/players"
)

func TestCliGameNew(t *testing.T) {
	black := &players.BotRandom{}
	white := &players.BotRandom{}
	buff := new(bytes.Buffer)
	cli_game := NewCliGame(black, white, buff)

	if cli_game.players[0] != black {
		t.Errorf("Black player.Player not assigned correctly")
	}
	if cli_game.players[1] != white {
		t.Errorf("White player.Player not assigned correctly")
	}

	if cli_game.writer != buff {
		t.Errorf("Writer is not assigned correctly")
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
