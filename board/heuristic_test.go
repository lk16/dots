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

func minimax(board Board, depth int, sign int) int {

	if depth == 0 {
		return sign * Squared(board)
	}

	children := board.getChildren()

	if len(children) == 0 {
		if board.OpponentMoves() == 0 {
			return sign * ExactScoreFactor * board.ExactScore()
		}

		child := board
		child.SwitchTurn()
		return minimax(child, depth, -sign)
	}

	if sign == 1 {
		best := MinHeuristic
		for _, child := range children {
			heur := minimax(child, depth-1, -sign)
			if heur > best {
				best = heur
			}
		}
		return best
	}

	best := MaxHeuristic
	for _, child := range children {
		heur := minimax(child, depth-1, -sign)
		if heur < best {
			best = heur
		}
	}
	return best

}

func TestNegaMax(t *testing.T) {
	for board := range genTestBoards() {
		negamaxHeur := Negamax(board, 3)
		minimaxHeur := minimax(board, 3, 1)

		if negamaxHeur != minimaxHeur {
			var buff bytes.Buffer
			board.ASCIIArt(&buff, false)

			t.Errorf("minimax=%d, negamx=%d for board\n\n%s\n",
				minimaxHeur, negamaxHeur, buff.String())
		}

	}

}
