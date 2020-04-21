package treesearch

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/lk16/dots/internal/othello"
	"github.com/stretchr/testify/mock"
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

// MiniMax implements the minimax tree search algorithm
type MiniMax struct {
	stats     Stats
	heuristic func(othello.Board) int
}

// NewMinimax returns a new MiniMax
func NewMinimax(heuristic func(othello.Board) int) *MiniMax {
	return &MiniMax{
		heuristic: heuristic,
	}
}

// Name returns the tree search algorithm name
func (minimax *MiniMax) Name() string {
	return "minimax"
}

// GetStats returns the statistics for the latest search
func (minimax MiniMax) GetStats() Stats {
	return minimax.stats
}

// ResetStats resets the statistics for the latest search to zeroes
func (minimax MiniMax) ResetStats() {
	minimax.stats.Reset()
}

// Search searches for the the best move up to a certain depth
func (minimax *MiniMax) Search(board othello.Board, alpha, beta, depth int) int {

	if depth >= board.CountEmpties() {
		depth = 60
	}

	minimax.stats.StartClock()
	heur := -minimax.search(board, depth, true)
	minimax.stats.StopClock()

	if heur < alpha {
		return alpha
	}

	if heur > beta {
		return beta
	}

	return heur
}

// ExactSearch searches for the best move without a depth limitation
func (minimax *MiniMax) ExactSearch(board othello.Board, alpha, beta int) int {
	return minimax.Search(board, alpha*ExactScoreFactor, beta*ExactScoreFactor, 60) / ExactScoreFactor
}

func (minimax *MiniMax) search(board othello.Board, depth int, maxPlayer bool) int {

	minimax.stats.Nodes++

	if depth == 0 {
		heur := minimax.heuristic(board)
		if !maxPlayer {
			heur = -heur
		}
		return heur
	}

	child := board
	gen := othello.NewUnsortedChildGenerator(&child)

	if !gen.HasMoves() {

		if board.OpponentMoves() == 0 {
			heur := ExactScoreFactor * board.ExactScore()
			if !maxPlayer {
				heur = -heur
			}
			return heur
		}

		board.SwitchTurn()
		return minimax.search(board, depth, !maxPlayer)
	}

	if maxPlayer {
		heur := MinHeuristic
		for gen.Next() {
			childHeur := minimax.search(child, depth-1, !maxPlayer)
			if childHeur > heur {
				heur = childHeur
			}
		}
		return heur
	}

	heur := MaxHeuristic
	for gen.Next() {
		childHeur := minimax.search(child, depth-1, !maxPlayer)
		if childHeur < heur {
			heur = childHeur
		}
	}
	return heur
}

func TestTreeSearch(t *testing.T) {

	internal := func(t *testing.T, depth int, board othello.Board, minimax, mtdf, pvs Searcher, testedBoards int) {

		minimaxResult := minimax.Search(board, MinHeuristic, MaxHeuristic, depth)
		mtdfResult := mtdf.Search(board, MinHeuristic, MaxHeuristic, depth)
		pvsResult := pvs.Search(board, MinHeuristic, MaxHeuristic, depth)

		if minimaxResult != mtdfResult || minimaxResult != pvsResult {
			fmt.Printf("\nFailed at board %d\n", testedBoards)
			msg := "Found inconsistent tree search results:\n"
			msg += fmt.Sprintf("%10s: %5d\n", minimax.Name(), minimaxResult)
			msg += fmt.Sprintf("%10s: %5d\n", pvs.Name(), pvsResult)
			msg += fmt.Sprintf("%10s: %5d\n", mtdf.Name(), mtdfResult)
			msg += fmt.Sprintf("for this board at depth %d:\n\n%s\n", depth, board.String())
			t.Error(msg)
			t.FailNow()
		}
	}

	rand.Seed(0)
	testedBoards := make(map[othello.Board]struct{})

	minimax := NewMinimax(Squared)
	mtdf := NewMtdf(Squared)
	pvs := NewPvs(Squared)

	for i := 0; i < 10; i++ {
		for discs := 4; discs < 64; discs++ {

			board, err := othello.NewRandomBoard(discs)
			if err != nil {
				t.Errorf("Failed to generate random board: %s", err)
			}

			normalized := board.Normalize()

			if _, ok := testedBoards[normalized]; ok {
				continue
			}

			testedBoards[normalized] = struct{}{}

			for depth := 0; depth < 4; depth++ {
				internal(t, depth, *board, minimax, mtdf, pvs, len(testedBoards))
			}
		}
	}
}

func TestTreeSearchExact(t *testing.T) {

	internal := func(t *testing.T, board othello.Board, minimax, mtdf, pvs Searcher, testedBoards int) {

		minimaxResult := minimax.ExactSearch(board, MinHeuristic, MaxHeuristic)
		mtdfResult := mtdf.ExactSearch(board, MinHeuristic, MaxHeuristic)
		pvsResult := pvs.ExactSearch(board, MinHeuristic, MaxHeuristic)

		if minimaxResult != mtdfResult || minimaxResult != pvsResult {
			fmt.Printf("\nFailed at board %d\n", testedBoards)
			msg := "Found inconsistent exact tree search results:\n"
			msg += fmt.Sprintf("%10s: %5d\n", minimax.Name(), minimaxResult)
			msg += fmt.Sprintf("%10s: %5d\n", pvs.Name(), pvsResult)
			msg += fmt.Sprintf("%10s: %5d\n", mtdf.Name(), mtdfResult)
			msg += fmt.Sprintf("for this board at perfect depth\n\n%s\n", board.String())
			t.Error(msg)
			t.FailNow()
		}
	}

	rand.Seed(0)
	testedBoards := make(map[othello.Board]struct{})

	minimax := NewMinimax(Squared)
	mtdf := NewMtdf(Squared)
	pvs := NewPvs(Squared)

	for i := 0; i < 20; i++ {
		for discs := 56; discs < 64; discs++ {

			board, err := othello.NewRandomBoard(discs)
			if err != nil {
				t.Errorf("Failed to generate random board: %s", err)
			}

			normalized := board.Normalize()

			if _, ok := testedBoards[normalized]; ok {
				continue
			}

			testedBoards[normalized] = struct{}{}

			internal(t, *board, minimax, mtdf, pvs, len(testedBoards))
		}
	}
}

// mockSearcher is an autogenerated mock type for the Searcher type
type mockSearcher struct {
	mock.Mock
}

// ExactSearch provides a mock function with given fields: board, alpha, beta
func (_m *mockSearcher) ExactSearch(board othello.Board, alpha int, beta int) int {
	ret := _m.Called(board, alpha, beta)

	var r0 int
	if rf, ok := ret.Get(0).(func(othello.Board, int, int) int); ok {
		r0 = rf(board, alpha, beta)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetStats provides a mock function with given fields:
func (_m *mockSearcher) GetStats() Stats {
	ret := _m.Called()

	var r0 Stats
	if rf, ok := ret.Get(0).(func() Stats); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(Stats)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *mockSearcher) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResetStats provides a mock function with given fields:
func (_m *mockSearcher) ResetStats() {
	_m.Called()
}

// Search provides a mock function with given fields: board, alpha, beta, depth
func (_m *mockSearcher) Search(board othello.Board, alpha int, beta int, depth int) int {
	ret := _m.Called(board, alpha, beta, depth)

	var r0 int
	if rf, ok := ret.Get(0).(func(othello.Board, int, int, int) int); ok {
		r0 = rf(board, alpha, beta, depth)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
