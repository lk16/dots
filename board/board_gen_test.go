package board

import (
	"bytes"
	"testing"
)

func TestBoardChildGenNext(t *testing.T) {

	for depth := 0; depth < 4; depth++ {
		for board := range genTestBoards() {

			expectedSet := map[Board]struct{}{}

			for _, child := range board.GetChildren() {
				expectedSet[child] = struct{}{}
			}

			gotSet := map[Board]struct{}{}

			clone := board
			gen := NewGenerator(&clone, depth)

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
				t.Errorf("For depth %d: expected set (%d):\n",
					depth, len(expectedSet))

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
}

func TestBoardChildGenRestoreParent(t *testing.T) {
	for depth := 0; depth < 4; depth++ {
		board := NewBoard()
		gen := NewGenerator(board, depth)
		gen.Next()
		gen.RestoreParent()
		if *board != *NewBoard() {
			t.Errorf("Restore parent failed with lookAhead %d", depth)
		}
	}
}

func TestBoardChildGenHasMoves(t *testing.T) {
	for depth := 0; depth < 4; depth++ {
		board := NewBoard()
		gen := NewGenerator(board, depth)
		if !gen.HasMoves() {
			t.Errorf("Expected initial board has moves")
		}

		board = RandomBoard(64)
		gen = NewGenerator(board, depth)

		if gen.HasMoves() {
			t.Errorf("Expected full board does not have moves")
		}
	}
}
