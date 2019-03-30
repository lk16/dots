package treesearch

import (
	"bytes"
	"fmt"
	"github.com/lk16/dots/internal/othello"
	"math/rand"
	"testing"
)

func TestTreeSearch(t *testing.T) {

	internal := func(t *testing.T, discs, depth int) {
		algos := []Interface{
			NewMinimax(),
			NewNegaMax(),
			NewAlphaBeta(MinHeuristic, MaxHeuristic),
			NewMtdf(MinHeuristic, MaxHeuristic)}
		board, err := othello.RandomBoard(discs)
		if err != nil {
			t.Errorf("Failed to generate random board: %s", err)
		}

		results := make(map[string]int, len(algos))

		for _, algo := range algos {
			if (algo.Name() == "minimax" || algo.Name() == "negamax") && depth >= 4 {
				continue
			}

			result := algo.Search(*board, depth)

			if result <= -200 || result >= 200 {
				// skip exact search
				break
			}
			results[algo.Name()] = result
		}

		resultsSet := make(map[int]struct{}, len(algos))

		for _, result := range results {
			resultsSet[result] = struct{}{}
		}

		if len(resultsSet) > 1 {
			msg := "Found inconsistent tree search results:\n"
			for _, algo := range algos {
				if result, ok := results[algo.Name()]; ok {
					msg += fmt.Sprintf("%10s: %5d\n", algo.Name(), result)
				} else {
					msg += fmt.Sprintf("%10s: <skipped>\n", algo.Name())
				}
			}
			var buff bytes.Buffer
			board.ASCIIArt(&buff, false)
			msg += fmt.Sprintf("for this board at depth %d:\n\n%s", depth, buff.String())
			t.Error(msg)
			t.FailNow()
		}
	}

	rand.Seed(0)

	for depth := 0; depth <= 4; depth++ {
		for discs := 4; discs <= 64; discs++ {
			for i := 0; i <= 5; i++ {
				internal(t, discs, depth)
			}
		}
	}

	for i := 0; i < 20; i++ {
		internal(t, 8, 6)
	}
}

func Benchmark8Deep(b *testing.B) {

	algos := []Interface{
		NewMinimax(),
		NewNegaMax(),
		NewAlphaBeta(MinHeuristic, MaxHeuristic),
		NewMtdf(MinHeuristic, MaxHeuristic)}

	for _, algo := range algos {
		b.Run(algo.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				algo.Search(*othello.NewBoard(), 8)
			}

		})
	}
}
