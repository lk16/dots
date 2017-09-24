package players

import (
	"bytes"
	"testing"

	"dots/board"
)

func TestHeuristicSquared(t *testing.T) {

	type testCase struct {
		board    board.Board
		expected int
	}

	testCases := []testCase{
		{*board.CustomBoard(0, 0), 0},
		{*board.NewBoard(), 0},
		{*board.RandomBoard(5), 0},
		{*board.CustomBoard(1, 0), 3},
		{*board.CustomBoard(0, 1), -3},
		{*board.CustomBoard(1, 2), 4}}

	for _, test := range testCases {

		got := Squared(test.board)

		if got != test.expected {
			buff := new(bytes.Buffer)
			test.board.ASCIIArt(buff, false)
			t.Errorf("Expected %d, got %d for board\n%s\n\n",
				test.expected, got, buff.String())
		}

	}

}
