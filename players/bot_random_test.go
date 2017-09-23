package players

import (
	"dots/board"
	"testing"
)

func TestBotRandomDoMove(t *testing.T) {
	bot := NewBotRandom()

	for discs := 4; discs <= 64; discs++ {
		for i := 0; i < 10; i++ {

			board := board.RandomBoard(discs)

			if board.Moves() == 0 {
				continue
			}

			afterwards := bot.DoMove(*board)
			expected := discs + 1
			got := afterwards.CountDiscs()

			if expected != got {
				t.Errorf("Expected %d discs, got %d\n", expected, got)
			}

		}

	}
}
