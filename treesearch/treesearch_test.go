package treesearch

import (
	"bytes"
	"dots/othello"
	"fmt"
	"testing"
)

func TestTreeSearch(t *testing.T) {

	for depth := 0; depth <= 4; depth++ {

		algos := []Interface{
			NewMinimax(),
			NewNegaMax(),
			NewAlphaBeta(MinHeuristic, MaxHeuristic)}

		for discs := 4; discs <= 64; discs++ {
			for i := 0; i <= 10; i++ {

				board := othello.RandomBoard(discs)

				results := make(map[string]int, len(algos))

				for _, algo := range algos {
					boardCopy := *board
					results[algo.Name()] = algo.Search(boardCopy, depth)

					if *board != boardCopy {
						t.Errorf("Algotithm '%s' modified the input board.", algo.Name())
						t.FailNow()
					}
				}

				for _, algo := range algos {
					if results[algo.Name()] != results[algos[0].Name()] {
						msg := "Found inconsistent tree search results:\n"
						for _, algo := range algos {
							msg += fmt.Sprintf("%s: %d\n", algo.Name(), results[algo.Name()])
						}
						var buff bytes.Buffer
						board.ASCIIArt(&buff, false)
						msg += fmt.Sprintf("for this board at depth %d:\n\n%s", depth, buff.String())
						t.Error(msg)
						t.FailNow()
					}
				}

			}
		}
	}

}
