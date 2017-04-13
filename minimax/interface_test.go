package minimax

import (
	"bytes"
	"testing"

	"dots/board"
	"dots/heuristic"
)

type interfaceTestConfig struct {
	algorithm Interface
	max_depth uint
}

func genInterfaceTestConfig() (ch chan interfaceTestConfig) {
	ch = make(chan interfaceTestConfig)
	go func() {
		ch <- interfaceTestConfig{algorithm: &Minimax{}, max_depth: 4}
		ch <- interfaceTestConfig{algorithm: &AlphaBeta{}, max_depth: 60}
		ch <- interfaceTestConfig{algorithm: &Mtdf{}, max_depth: 60}
		close(ch)
	}()
	return
}

// Helper for this testing file
// Returns a string written by board.AsciiArt()
func boardAsciiArtString(board board.Board, swap_disc_colors bool) (output string) {
	buffer := new(bytes.Buffer)
	board.AsciiArt(buffer, swap_disc_colors)
	output = buffer.String()
	return
}

func genTestBoards(n uint) (ch chan board.Board) {
	ch = make(chan board.Board)
	go func() {
		for discs := uint(4); discs <= 64; discs++ {
			for i := uint(0); i < n/64; i++ {
				ch <- *board.RandomBoard(discs)
			}
		}
		close(ch)
	}()
	return
}

func TestInterfaceSearch(t *testing.T) {

	alpha := Min_heuristic
	heuristic := heuristic.Squared
	max_test_depth := uint(4)

	for board := range genTestBoards(2 * 64) {
		for depth := uint(1); depth <= max_test_depth; depth++ {

			// skip exact serches
			if board.CountEmpties() <= depth {
				continue
			}

			var base_case_algorithm Interface
			var expected int

			for test_case := range genInterfaceTestConfig() {
				if depth > test_case.max_depth {
					continue
				}
				if base_case_algorithm == nil {
					base_case_algorithm = test_case.algorithm
					expected = base_case_algorithm.Search(board, depth, heuristic, alpha)
					continue
				}

				got := test_case.algorithm.Search(board, depth, heuristic, alpha)

				if got != expected {
					t.Errorf("When searching at depth %d:\n expected: %d (from %s)\n",
						depth, expected, base_case_algorithm.Name())
					t.Errorf("got %d (from %s)\n for board\n%s\n\n",
						got, test_case.algorithm.Name(), boardAsciiArtString(board, false))
				}
			}
		}
	}

}

func TestInterfaceExactSearch(t *testing.T) {

	alpha := Min_exact_heuristic
	max_test_depth := uint(8)

	for board := range genTestBoards(1 * 64) {

		if board.CountEmpties() > max_test_depth {
			continue
		}

		var base_case_algorithm Interface
		var expected int

		for test_case := range genInterfaceTestConfig() {
			if board.CountEmpties() > test_case.max_depth {
				continue
			}

			if base_case_algorithm == nil {
				base_case_algorithm = test_case.algorithm
				expected = base_case_algorithm.ExactSearch(board, alpha)
				continue
			}

			got := test_case.algorithm.ExactSearch(board, alpha)
			if got != expected {
				t.Errorf("When searching exact expected: %d (from %s)\n",
					expected, base_case_algorithm.Name())
				t.Errorf("got %d (from %s)\n for board\n%s\n\n",
					got, test_case.algorithm.Name(), boardAsciiArtString(board, false))
				t.FailNow()
			}
		}
	}

}

func TestInterfaceName(t *testing.T) {

	for test_case := range genInterfaceTestConfig() {
		if test_case.algorithm.Name() == "" {
			t.Errorf("Expected non-empty string\n")
		}
	}
}
