package minimax

import (
	"testing"

	"dots/board"
	"dots/heuristic"
)

func genTestBoards() (ch chan board.Board) {
	ch = make(chan board.Board)
	go func() {
		for i := uint(4); i <= 64; i++ {
			ch <- *board.RandomBoard(i)
		}
		close(ch)
	}()
	return
}

func TestMinimax(t *testing.T) {

	interfaces := []Interface{
		&Minimax{},
		&AlphaBeta{},
		&Mtdf{}}

	search_depth := uint(5)
	exact_depth := uint(5)
	alpha := Min_exact_heuristic

	for board := range genTestBoards() {
		results := make([]int, 0)

		for _, minimax := range interfaces {
			var result int
			if board.CountEmpties() > exact_depth {
				result = minimax.Search(board, search_depth, heuristic.Squared, alpha)
			} else {
				result = minimax.ExactSearch(board, alpha)
			}
			results = append(results, result)
		}

		results_equal := true
		for i := 1; i < len(results); i++ {
			if results[0] != results[i] {
				results_equal = false
			}
		}

		if !results_equal {
			t.Errorf("Results are unequal %v for board\n%s\n\n", results, board.AsciiArt())
			t.FailNow()
		}

	}

}
