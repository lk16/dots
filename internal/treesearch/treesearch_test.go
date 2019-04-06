package treesearch

import (
	"bytes"
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"math/rand"
	"testing"
)

func TestTreeSearch(t *testing.T) {

	internal := func(t *testing.T, depth int, board othello.Board, minimax, mtdf, pvs Interface, testedBoards int) {

		bound := 2 * ExactScoreFactor

		minimaxResult := minimax.Search(board, -bound, bound, depth)
		mtdfResult := mtdf.Search(board, -bound, bound, depth)
		pvsResult := pvs.Search(board, -bound, bound, depth)

		if minimaxResult != mtdfResult || minimaxResult != pvsResult {
			fmt.Printf("\nFailed at board %d\n", testedBoards)
			msg := "Found inconsistent tree search results:\n"
			msg += fmt.Sprintf("%10s: %5d\n", minimax.Name(), minimaxResult)
			msg += fmt.Sprintf("%10s: %5d\n", pvs.Name(), pvsResult)
			msg += fmt.Sprintf("%10s: %5d\n", mtdf.Name(), mtdfResult)

			var buff bytes.Buffer
			board.ASCIIArt(&buff, false)
			msg += fmt.Sprintf("for this board at depth %d:\n\n%s\n", depth, buff.String())
			t.Error(msg)
			t.FailNow()
		}
	}

	rand.Seed(0)
	testedBoards := make(map[othello.Board]struct{})

	minimax := NewMinimax()
	mtdf := NewMtdf()
	pvs := NewPvs()

	for i := 0; i < 10000; i++ {
		for discs := 4; discs < 64; discs++ {

			board, err := othello.RandomBoard(discs)
			if err != nil {
				t.Errorf("Failed to generate random board: %s", err)
			}

			normalized := board.Normalize()

			if _, ok := testedBoards[normalized]; ok {
				continue
			}

			testedBoards[normalized] = struct{}{}

			if len(testedBoards)%1000 == 0 {
				fmt.Printf("\rTesting board %10d", len(testedBoards))
			}

			for depth := 0; depth < 4; depth++ {
				internal(t, depth, *board, minimax, mtdf, pvs, len(testedBoards))
			}
		}
	}
	fmt.Printf("\rTesting board %10d\n", len(testedBoards))
}
