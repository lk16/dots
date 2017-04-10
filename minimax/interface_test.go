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

func TestInterfaceSearch(t *testing.T) {

	base_case := &Minimax{}

	test_cases := []Interface{
		&AlphaBeta{},
		&Mtdf{}}

	alpha := Min_exact_heuristic
	heuristic := heuristic.Squared

	for board := range genTestBoards() {
		for depth := uint(1); depth <= 4; depth++ {

			// skip exact serches because it will make the unit test slow
			if board.CountEmpties() <= depth {
				continue
			}

			expected := base_case.Search(board, depth, heuristic, alpha)
			for _, test_case := range test_cases {
				got := test_case.Search(board, depth, heuristic, alpha)
				if got != expected {
					t.Errorf("At depth %d: expected %d, got %d from %s for board\n%s\n\n", depth, expected, got, test_case.Name(), board.AsciiArt())
				}
			}
		}
	}

}

func TestInterfaceExactSearch(t *testing.T) {

	base_case := &Minimax{}

	test_cases := []Interface{
		&AlphaBeta{},
		&Mtdf{}}

	exact_depth := uint(4)
	alpha := Min_exact_heuristic

	for board := range genTestBoards() {

		// skip exact searches for large depths becuase it will make unit tests slow
		if board.CountEmpties() <= exact_depth {
			continue
		}

		expected := base_case.ExactSearch(board, alpha)
		for _, test_case := range test_cases {
			got := test_case.ExactSearch(board, alpha)
			if got != expected {
				t.Errorf("Expected %d, got %d from %s for board\n%s\n\n", expected, got, test_case.Name(), board.AsciiArt())
			}
		}
	}

}
