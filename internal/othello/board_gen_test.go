package othello

import (
	"bytes"
	"testing"
)

func TestBoardChildGenNext(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	for b := range genTestBoards() {

		// create copy to silence warnings
		board := b

		if !boardIsValid(&board) {
			continue
		}

		expectedSet := map[Board]struct{}{}

		for _, child := range board.GetChildren() {
			expectedSet[child] = struct{}{}
		}

		gotSet := map[Board]struct{}{}

		clone := board
		gen := NewUnsortedChildGenerator(&clone)

		for gen.Next() {
			gotSet[clone] = struct{}{}
		}

		if clone != board {
			t.Errorf("Parent state not restored")
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
	gen := NewUnsortedChildGenerator(board)
	gen.Next()
	gen.RestoreParent()
	if *board != *NewBoard() {
		t.Errorf("Restore parent failed")
	}
}

func TestBoardChildGenHasMoves(t *testing.T) {
	board := NewBoard()
	gen := NewUnsortedChildGenerator(board)
	if !gen.HasMoves() {
		t.Errorf("Expected initial othello has moves")
	}

	board, err := NewRandomBoard(64)
	if err != nil {
		t.Errorf("Error generating random full board: %s", err)
	}
	gen = NewUnsortedChildGenerator(board)

	if gen.HasMoves() {
		t.Errorf("Expected full othello does not have moves")
	}
}
