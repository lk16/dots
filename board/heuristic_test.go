package board

import (
	"bytes"
	"testing"
)

func TestHeuristicSquared(t *testing.T) {

	type testCase struct {
		board    *Board
		expected int
	}

	testCases := []testCase{
		{CustomBoard(0, 0), 0},
		{NewBoard(), 0},
		{RandomBoard(5), 0},
		{CustomBoard(1, 0), 3},
		{CustomBoard(0, 1), -3},
		{CustomBoard(1, 2), 4}}

	for _, test := range testCases {

		got := Squared(*test.board)

		if got != test.expected {
			buff := new(bytes.Buffer)
			test.board.ASCIIArt(buff, false)
			t.Errorf("Expected %d, got %d for board\n%s\n\n",
				test.expected, got, buff.String())
		}

	}

}
