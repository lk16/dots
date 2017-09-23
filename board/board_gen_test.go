package board

import (
	"bytes"
	"testing"

	"math/bits"
)

func TestBoardChildGenNext(t *testing.T) {

	for board := range genTestBoards() {

		expected_set := map[Board]struct{}{}

		for _, child := range board.GetChildren() {
			expected_set[child] = struct{}{}
		}

		got_set := map[Board]struct{}{}

		clone := board

		gen := NewChildGen(&clone)
		for gen.Next() {
			got_set[clone] = struct{}{}
		}

		if clone != board {
			t.Errorf("Parent state not restored after looping over all children")
		}

		for g, _ := range got_set {
			if _, ok := expected_set[g]; !ok {
				t.Errorf("Children sets are unequal.\n")
				break
			}
		}

		if t.Failed() {
			buff := new(bytes.Buffer)
			t.Errorf("Expected set (%d):\n", len(expected_set))
			for child, _ := range expected_set {
				child.AsciiArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			buff.Reset()
			t.Errorf("Got set (%d):\n", len(got_set))
			for child, _ := range got_set {
				child.AsciiArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			break
		}

	}
}

func TestBoardChildGenRestoreParent(t *testing.T) {
	board := NewBoard()
	gen := NewChildGen(board)
	gen.Next()
	gen.RestoreParent()
	if *board != *NewBoard() {
		t.Errorf("Restore parent failed")
	}
}

func TestBoardChildGenHasMoves(t *testing.T) {
	if !NewChildGen(NewBoard()).HasMoves() {
		t.Errorf("Expected initial board has moves!")
	}

	if NewChildGen(RandomBoard(64)).HasMoves() {
		t.Errorf("Expected full board does not have moves")
	}

}

func lame_heuristic(board Board) int {
	return bits.OnesCount64(board.Me()) - bits.OnesCount64(board.Opp())
}

func TestBoardChildGenSortedNext(t *testing.T) {

	for board := range genTestBoards() {

		expected_set := map[Board]struct{}{}

		for _, child := range board.GetChildren() {
			expected_set[child] = struct{}{}
		}

		got_set := map[Board]struct{}{}

		clone := board

		gen := NewChildGenSorted(&clone, lame_heuristic)
		for gen.Next() {
			got_set[clone] = struct{}{}
		}

		if clone != board {
			t.Errorf("Parent state not restored after looping over all children")
		}

		for g, _ := range got_set {
			if _, ok := expected_set[g]; !ok {
				t.Errorf("Children sets are unequal.\n")
				break
			}
		}

		if t.Failed() {
			buff := new(bytes.Buffer)
			t.Errorf("Expected set (%d):\n", len(expected_set))
			for child, _ := range expected_set {
				child.AsciiArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			buff.Reset()
			t.Errorf("Got set (%d):\n", len(got_set))
			for child, _ := range got_set {
				child.AsciiArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			break
		}

	}
}

func TestBoardChildGenSortedRestoreParent(t *testing.T) {
	board := NewBoard()
	gen := NewChildGenSorted(board, lame_heuristic)
	gen.Next()
	gen.RestoreParent()
	if *board != *NewBoard() {
		t.Errorf("Restore parent failed")
	}
}

func TestBoardChildGenSortedHasMoves(t *testing.T) {
	if !NewChildGenSorted(NewBoard(), lame_heuristic).HasMoves() {
		t.Errorf("Expected initial board has moves!")
	}

	if NewChildGenSorted(RandomBoard(64), lame_heuristic).HasMoves() {
		t.Errorf("Expected full board does not have moves")
	}

}
