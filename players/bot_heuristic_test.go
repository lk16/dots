package players

import (
	"bytes"
	"testing"

	"dots/bitset"
	"dots/board"
	"dots/minimax"
)

var fake_heuristic_calls uint

func fake_heuristic(board board.Board) (heur int) {
	fake_heuristic_calls++
	return
}

type fake_minimax struct{}

var fake_search_calls uint

func (fake *fake_minimax) Search(board board.Board, depth_left uint,
	heuristic minimax.Heuristic, alpha int) (heur int) {
	fake_search_calls++
	return
}

var fake_exact_search_calls uint

func (fake *fake_minimax) ExactSearch(board board.Board, alpha int) (heur int) {
	fake_exact_search_calls++
	return
}

func (fake *fake_minimax) Name() (name string) {
	return "fake"
}

func TestNewBotHeuristic(t *testing.T) {

	minimax := &minimax.Mtdf{}
	search_depth := uint(3)
	exact_depth := uint(6)
	writer := new(bytes.Buffer)

	fake_heuristic_calls = 0

	bot := NewBotHeuristic(fake_heuristic, minimax, search_depth, exact_depth, writer)

	bot.heuristic(*board.NewBoard())
	if fake_heuristic_calls != 1 {
		t.Errorf("Setting heuristics failed")
	}

	if bot.minimax != minimax {
		t.Errorf("Setting minimax failed")
	}

	if bot.search_depth != search_depth || bot.exact_depth != exact_depth {
		t.Errorf("Setting depths failed")
	}

	if bot.writer != writer {
		t.Errorf("Setting writer failed")
	}
}

// Fails if panic() is not called
func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("panic() was not called")
	}
}

func TestBotHeuristicDoMove(t *testing.T) {

	fake_search_calls = 0
	fake_exact_search_calls = 0

	bot := NewBotHeuristic(fake_heuristic, &fake_minimax{}, uint(3), uint(6), new(bytes.Buffer))

	test := func(b board.Board) {

		fake_search_calls = 0
		fake_exact_search_calls = 0

		moves := b.Moves().Count()
		discs := b.CountDiscs()
		afterwards := bot.DoMove(b)

		var expected_search_calls, expected_exact_calls uint

		if b.Moves().Count() == 1 {
			expected_search_calls = 0
			expected_exact_calls = 0
		} else if b.CountEmpties() <= bot.exact_depth {
			expected_search_calls = 0
			expected_exact_calls = moves
		} else {
			expected_search_calls = moves
			expected_exact_calls = 0
		}

		if fake_search_calls != expected_search_calls ||
			fake_exact_search_calls != expected_exact_calls {
			t.Errorf("Expected (%d,%d), got (%d,%d)\n", expected_search_calls,
				expected_exact_calls, fake_search_calls, fake_exact_search_calls)
		}

		discs_afterwards := afterwards.CountDiscs()
		if discs_afterwards != discs+1 {
			t.Errorf("Expected %d discs, got %d\n", discs+1, discs_afterwards)
		}
	}

	// normal search, 0 moves
	func() {
		defer assertPanic(t)
		test(*board.CustomBoard(0, 0))
	}()

	// normal search, 1 move
	test(*board.CustomBoard(1, 2))

	// normal search, >1 move
	test(*board.NewBoard())

	// exact search, 0 moves
	func() {
		defer assertPanic(t)
		test(*board.CustomBoard(^bitset.Bitset(0), 0))
	}()

	// exact search, 1 move
	test(*board.CustomBoard(^bitset.Bitset(3), 2))

	// exact search, >1 move
	me := ^bitset.Bitset(0)
	me.ResetBit(0).ResetBit(1).ResetBit(3).ResetBit(4)
	opp := bitset.Bitset(0)
	opp.SetBit(1).SetBit(3)
	test(*board.CustomBoard(me, opp))
}
