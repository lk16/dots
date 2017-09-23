package board

import (
	"bytes"
	"testing"

	"math/bits"
)

func TestBoardChildGenNext(t *testing.T) {

	for board := range genTestBoards() {

		expectedSet := map[Board]struct{}{}

		for _, child := range board.GetChildren() {
			expectedSet[child] = struct{}{}
		}

		gotSet := map[Board]struct{}{}

		clone := board

		gen := NewChildGen(&clone)
		for gen.Next() {
			gotSet[clone] = struct{}{}
		}

		if clone != board {
			t.Errorf("Parent state not restored after looping over all children")
		}

		for g := range gotSet {
			if _, ok := expectedSet[g]; !ok {
				t.Errorf("Children sets are unequal.\n")
				break
			}
		}

		if t.Failed() {
			buff := new(bytes.Buffer)
			t.Errorf("Expected set (%d):\n", len(expectedSet))
			for child := range expectedSet {
				child.ASCIIArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			buff.Reset()
			t.Errorf("Got set (%d):\n", len(gotSet))
			for child := range gotSet {
				child.ASCIIArt(buff, false)
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

func lameHeuristic(board Board) int {
	return bits.OnesCount64(board.Me()) - bits.OnesCount64(board.Opp())
}

func TestBoardChildGenSortedNext(t *testing.T) {

	for board := range genTestBoards() {

		expectedSet := map[Board]struct{}{}

		for _, child := range board.GetChildren() {
			expectedSet[child] = struct{}{}
		}

		gotSet := map[Board]struct{}{}

		clone := board

		gen := NewChildGenSorted(&clone, lameHeuristic)
		for gen.Next() {
			gotSet[clone] = struct{}{}
		}

		if clone != board {
			t.Errorf("Parent state not restored after looping over all children")
		}

		for g := range gotSet {
			if _, ok := expectedSet[g]; !ok {
				t.Errorf("Children sets are unequal.\n")
				break
			}
		}

		if t.Failed() {
			buff := new(bytes.Buffer)
			t.Errorf("Expected set (%d):\n", len(expectedSet))
			for child := range expectedSet {
				child.ASCIIArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			buff.Reset()
			t.Errorf("Got set (%d):\n", len(gotSet))
			for child := range gotSet {
				child.ASCIIArt(buff, false)
				buff.WriteString("\n\n")
			}
			t.Errorf(buff.String())
			break
		}

	}
}

func TestBoardChildGenSortedRestoreParent(t *testing.T) {
	board := NewBoard()
	gen := NewChildGenSorted(board, lameHeuristic)
	gen.Next()
	gen.RestoreParent()
	if *board != *NewBoard() {
		t.Errorf("Restore parent failed")
	}
}

func TestBoardChildGenSortedHasMoves(t *testing.T) {
	if !NewChildGenSorted(NewBoard(), lameHeuristic).HasMoves() {
		t.Errorf("Expected initial board has moves!")
	}

	if NewChildGenSorted(RandomBoard(64), lameHeuristic).HasMoves() {
		t.Errorf("Expected full board does not have moves")
	}

}
