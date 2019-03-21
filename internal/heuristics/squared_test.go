package heuristics

import (
	"bytes"
	"github.com/lk16/dots/internal/othello"
	"testing"
)

func TestSquared(t *testing.T) {

	type testCase struct {
		board    *othello.Board
		expected int
	}

	testCases := []testCase{
		{othello.CustomBoard(0, 0), 0},
		{othello.NewBoard(), 0},
		{othello.RandomBoard(5), 0},
		{othello.CustomBoard(1, 0), 3},
		{othello.CustomBoard(0, 1), -3},
		{othello.CustomBoard(1, 2), 4}}

	for _, test := range testCases {
		got := Squared(*test.board)

		if got != test.expected {
			buff := new(bytes.Buffer)
			test.board.ASCIIArt(buff, false)
			t.Errorf("Expected %d, got %d for othello\n%s\n\n",
				test.expected, got, buff.String())
		}
	}
}
