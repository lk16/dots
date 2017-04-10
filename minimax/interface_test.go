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

	search_depth := uint(5)
	exact_depth := uint(5)
	alpha := Min_exact_heuristic
	heuristic := heuristic.Squared

	for board := range genTestBoards() {
		if board.CountEmpties() > exact_depth {
			expected := base_case.Search(board, search_depth, heuristic, alpha)
			for _, test_case := range test_cases {
				got := test_case.Search(board, search_depth, heuristic, alpha)
				if got != expected {
					t.Errorf("Expected %d, got %d for board\n%s\n\n", expected, got, board.AsciiArt())
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

	exact_depth := uint(5)
	alpha := Min_exact_heuristic

	for board := range genTestBoards() {
		if board.CountEmpties() <= exact_depth {

			expected := base_case.ExactSearch(board, alpha)
			for _, test_case := range test_cases {
				got := test_case.ExactSearch(board, alpha)
				if got != expected {
					t.Errorf("Expected %d, got %d for board\n%s\n\n", expected, got, board.AsciiArt())
				}
			}

		}
	}

}
