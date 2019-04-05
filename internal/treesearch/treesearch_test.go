package treesearch

import (
	"bytes"
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"math/rand"
	"testing"
)

func TestTreeSearch(t *testing.T) {

	internal := func(t *testing.T, depth int, board othello.Board) {

		minimax := (Interface)(NewMinimax())
		mtdf := (Interface)(NewMtdf())

		bound := 2 * ExactScoreFactor

		minimaxResult := minimax.Search(board, -bound, bound, depth)

		mtdfResult := mtdf.Search(board, -bound, bound, depth)

		if minimaxResult != mtdfResult {
			fmt.Printf("\n")
			msg := "Found inconsistent tree search results:\n"
			msg += fmt.Sprintf("%10s: %5d\n", minimax.Name(), minimaxResult)
			msg += fmt.Sprintf("%10s: %5d\n", mtdf.Name(), mtdfResult)

			var buff bytes.Buffer
			board.ASCIIArt(&buff, false)
			msg += fmt.Sprintf("for this board at depth %d:\n\n%s", depth, buff.String())
			t.Error(msg)
			t.FailNow()
		}
	}

	rand.Seed(0)
	testedBoards := make(map[othello.Board]struct{})

	for depth := 0; depth <= 4; depth++ {
		for discs := 4; discs <= 64; discs++ {
			for i := 0; i < 100; i++ {

				board, err := othello.RandomBoard(discs)
				if err != nil {
					t.Errorf("Failed to generate random board: %s", err)
				}

				normalized := board.Normalize()

				if _, ok := testedBoards[normalized]; ok {
					continue
				}

				testedBoards[normalized] = struct{}{}

				fmt.Printf("\rTesting board %10d", len(testedBoards))
				internal(t, depth, *board)
			}
		}
	}
	fmt.Printf("\rTesting board %10d\n", len(testedBoards))
}
