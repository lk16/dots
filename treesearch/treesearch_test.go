package treesearch

import (
"bytes"
"testing"
)


func TestNegaMax(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	for board := range genTestBoards() {
		for depth := 1; depth <= 3; depth++ {
			boardCopy := board
			negamaxHeur := Negamax(&board, 3)

			if boardCopy != board {
				t.Error("negamax modified the othello")
			}

			minimaxHeur := minimax(board, 3, 1)

			if negamaxHeur != minimaxHeur {
				var buff bytes.Buffer
				board.ASCIIArt(&buff, false)

				t.Errorf("minimax=%d, negamx=%d for othello\n\n%s\n",
					minimaxHeur, negamaxHeur, buff.String())
			}
		}
	}
}

func TestAlphaBeta(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	for board := range genTestBoards() {
		for depth := 1; depth <= 3; depth++ {
			boardCopy := board
			alphaBetaHeur := AlphaBeta(&board, MinHeuristic, MaxHeuristic, 3)

			if boardCopy != board {
				t.Error("alphabeta modified the othello")
			}

			minimaxHeur := minimax(board, 3, 1)

			if alphaBetaHeur != minimaxHeur {
				var buff bytes.Buffer
				board.ASCIIArt(&buff, false)

				t.Errorf("minimax=%d, alphabeta=%d for othello\n\n%s\n",
					minimaxHeur, alphaBetaHeur, buff.String())
			}
		}
	}
}
