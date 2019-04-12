package treesearch

import (
	"bytes"
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"math/rand"
	"testing"
)

func TestFormatBigNumber(t *testing.T) {

	type testCase struct {
		input          uint64
		expectedOutput string
	}

	testCases := []testCase{
		{0, "0"},
		{1, "1"},
		{12, "12"},
		{123, "123"},
		{1234, "1.23K"},
		{123456, "123K"},
		{1234567, "1.23M"},
		{12345678, "12.3M"},
		{123456789, "123M"},
		{1234567890, "1.23G"},
		{12345678901, "12.3G"},
		{123456789012, "123G"},
		{1234567890123, "1.23T"},
		{12345678901234, "12.3T"},
		{123456789012345, "123T"},
		{1234567890123456, "1.23P"},
		{12345678901234567, "12.3P"},
		{123456789012345678, "123P"},
		{1234567890123456789, "1.23E"},
		{12345678901234567890, "12.3E"}}

	for i := range testCases {
		output := FormatBigNumber(testCases[i].input)

		if output != testCases[i].expectedOutput {
			t.Errorf("For input %d expected \"%s\", got \"%s\"",
				testCases[i].input, testCases[i].expectedOutput, output)
		}

	}
}

func TestTreeSearch(t *testing.T) {

	internal := func(t *testing.T, depth int, board othello.Board, minimax, mtdf, pvs Interface, testedBoards int) {

		minimaxResult := minimax.Search(board, MinHeuristic, MaxHeuristic, depth)
		mtdfResult := mtdf.Search(board, MinHeuristic, MaxHeuristic, depth)
		pvsResult := pvs.Search(board, MinHeuristic, MaxHeuristic, depth)

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

	for i := 0; i < 1000; i++ {
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
