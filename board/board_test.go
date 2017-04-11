package board

import (
	"bytes"
	"fmt"
	"testing"

	"dots/bitset"
)

/**
TODO:
- test all outputs for validity with Board.IsValid()
- test functions for constness
**/

// Helper for this testing file
// Returns a string written by board.AsciiArt()
func (board Board) asciiArtString(swap_disc_colors bool) (output string) {
	buffer := new(bytes.Buffer)
	board.AsciiArt(buffer, swap_disc_colors)
	output = buffer.String()
	return
}

func bitsetAsciiArtString(bs bitset.Bitset) (output string) {
	buffer := new(bytes.Buffer)
	bs.AsciiArt(buffer)
	output = buffer.String()
	return
}

// Test board generator
func genTestBoards() (ch chan Board) {
	ch = make(chan Board)
	go func() {
		ch <- *NewBoard()

		// random reachable boards with 4-64 discs
		for i := 0; i < 10; i++ {
			for discs := uint(4); discs <= 64; discs++ {
				ch <- *RandomBoard(discs)
			}
		}

		// random boards not necessarily reachable
		for i := 0; i < 1000; i++ {
			board := Board{
				me:  bitset.RandomBitset(),
				opp: bitset.RandomBitset()}
			board.opp &^= board.me
			ch <- board
		}

		// generate all boards with all drawing lines from each square

		// for each field
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				board := Board{}
				board.me.SetBit(uint(y*8 + x))

				// for each direction
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if (dy == 0) && (dx == 0) {
							continue
						}
						board.opp = 0

						// for each distance
						for d := 1; d <= 6; d++ {

							// check if me can still flip within board boundaries
							py := y + (d+1)*dy
							px := x + (d+1)*dx

							if (py < 0) || (py > 7) || (px < 0) || (px > 7) {
								break
							}

							qy := y + d*dy
							qx := x + d*dx

							board.opp.SetBit(uint(qy*8 + qx))

							ch <- board
						}
					}
				}
			}
		}

		close(ch)
	}()
	return
}

// Fails if panic() is not called
func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("panic() was not called")
	}
}

func TestBoardIsValid(t *testing.T) {

	type testCase struct {
		board    Board
		expected bool
	}

	start_board := *NewBoard()
	empty_board := Board{me: 0, opp: 0}

	all := ^bitset.Bitset(0)

	start_mask := start_board.me | start_board.opp
	duplicated_board := Board{me: start_mask, opp: start_mask}
	me_win_board := Board{me: all, opp: 0}
	opp_win_board := Board{me: 0, opp: all}
	draw_board := Board{me: all << 32, opp: all >> 32}

	test_cases := []testCase{
		{empty_board, false},
		{start_board, true},
		{duplicated_board, false},
		{me_win_board, true},
		{opp_win_board, true},
		{draw_board, true}}

	for _, test_case := range test_cases {
		expected := test_case.expected
		board := test_case.board

		got := board.IsValid()

		if expected != got {
			t.Errorf("Expected %t, got %t for board\n%s\n\n", expected, got, board.asciiArtString(false))
		}
	}
}

func TestRandomBoard(t *testing.T) {
	for i := uint(4); i <= 64; i++ {
		board := RandomBoard(i)
		expected := i
		got := (board.me | board.opp).Count()
		if expected != got {
			t.Errorf("Expected disc count %d, got %d\n", expected, got)
		}
	}

	func() {
		defer assertPanic(t)
		RandomBoard(3)
	}()

	func() {
		defer assertPanic(t)
		RandomBoard(65)
	}()
}

// Board.Clone() was removed, but this test is kept to test for assignment copy works
func TestBoardClone(t *testing.T) {
	board := Board{
		me:  1,
		opp: 2}
	clone := board
	clone.me = 3
	if board.me != 1 {
		t.Errorf("'clone := board' does not make a deep copy!\n")
	}
}

func (board *Board) doMove(index uint) (flipped bitset.Bitset) {
	if (board.me | board.opp).TestBit(index) {
		return
	}
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			s := 1
			for {
				curx := int(index%8) + (dx * s)
				cury := int(index/8) + (dy * s)
				cur := uint(8*cury + curx)
				if curx < 0 || curx >= 8 || cury < 0 || cury >= 8 {
					break
				}
				if board.opp.TestBit(cur) {
					s++
				} else {
					if board.me.TestBit(cur) && (s >= 2) {
						for p := 1; p < s; p++ {
							f := uint(int(index) + (p * (8*dy + dx)))
							flipped.SetBit(f)
						}
					}
					break
				}
			}
		}
	}
	board.me |= flipped
	board.me.SetBit(index)
	board.opp &= ^board.me
	board.opp, board.me = board.me, board.opp
	return
}

func TestBoardDoMove(t *testing.T) {
	for board := range genTestBoards() {
		moves := board.Moves()
		for i := uint(0); i < 64; i++ {
			if !moves.TestBit(i) {
				// board.DoMove() should not be called for invalid moves
				continue
			}

			clone := board
			expected_return_val := clone.doMove(i)
			expected_board_val := clone

			clone = board
			got_return_val := clone.DoMove(i)
			got_board_val := clone

			if (got_return_val != expected_return_val) || (got_board_val != expected_board_val) {
				t.Errorf("Doing move %c%d on board\n%s\n", 'a'+i%8, (i/8)+1,
					board.asciiArtString(false))
				t.Errorf("Expected return val:\n%s\n\nGot:\n%s\n\n",
					bitsetAsciiArtString(expected_return_val), bitsetAsciiArtString(got_return_val))
				t.Errorf("Expected board:\n%s\n\nGot:\n%s\n\n",
					expected_board_val.asciiArtString(false), got_board_val.asciiArtString(false))
				t.FailNow()
			}
		}
	}
}

func (board Board) moves() (moves bitset.Bitset) {
	empties := ^(board.me | board.opp)

	for i := uint(0); i < 64; i++ {
		clone := board
		if empties.TestBit(i) && clone.DoMove(i) != 0 {
			moves.SetBit(i)
		}
	}
	return
}

func TestBoardMoves(t *testing.T) {
	for board := range genTestBoards() {
		clone := board
		expected := board.moves()
		got := board.Moves()
		if expected != got {
			t.Errorf("For board\n%s", board.asciiArtString(false))
			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n", bitsetAsciiArtString(expected), bitsetAsciiArtString(got))
			t.FailNow()
		}
		if clone != board {
			t.Errorf("board.Moves() changed the board!\n")
			t.FailNow()
		}
	}
}

func (board *Board) getChildren() (children []Board) {
	for i := uint(0); i < 64; i++ {
		clone := *board
		if clone.doMove(i) != bitset.Bitset(0) {
			children = append(children, clone)
		}
	}
	return
}

func TestBoardGetChildren(t *testing.T) {
	for board := range genTestBoards() {

		expected := board.getChildren()
		expected_set := make(map[Board]struct{}, 10)
		for _, e := range expected {
			expected_set[e] = struct{}{}
		}

		board_pieces := board.me | board.opp

		got := board.GetChildren()
		got_set := make(map[Board]struct{}, 10)
		for _, g := range got {
			got_set[g] = struct{}{}

			child_pieces := board.me | board.opp

			if child_pieces&board_pieces != board_pieces {
				t.Errorf("Pieces where removed from board with board.GetChildren()\n")
				t.Errorf("board:\n%s\n\nchild: \n%s\n\n",
					board.asciiArtString(false), g.asciiArtString(false))
				t.FailNow()
			}
		}

		if len(got_set) != len(expected_set) {
			t.Errorf("Expected %d children, got %d.\n", len(expected_set), len(got_set))
			t.FailNow()
		}

		equal_sets := true

		for e, _ := range expected_set {
			if _, ok := got_set[e]; !ok {
				equal_sets = false
			}
		}

		if !equal_sets {
			t.Errorf("Children sets are unequal.\n")
			t.FailNow()
		}
	}
}

func TestBoardAsciiArt(t *testing.T) {
	for board := range genTestBoards() {

		for _, swap_disc_colors := range []bool{true, false} {

			moves := board.Moves()

			clone := board
			if swap_disc_colors {
				clone.SwitchTurn()
			}

			expected_buff := new(bytes.Buffer)

			expected_buff.WriteString("+-a-b-c-d-e-f-g-h-+\n")

			for y := uint(0); y < 8; y++ {
				expected_buff.WriteString(fmt.Sprintf("%d ", y+1))

				for x := uint(0); x < 8; x++ {
					if clone.me.TestBit(8*y + x) {
						expected_buff.WriteString("○ ")
					} else if clone.opp.TestBit(8*y + x) {
						expected_buff.WriteString("● ")
					} else if moves.TestBit(8*y + x) {
						expected_buff.WriteString("- ")
					} else {
						expected_buff.WriteString("  ")
					}
				}

				expected_buff.WriteString("|\n")
			}

			expected_buff.WriteString("+-----------------+\n")

			got_buff := new(bytes.Buffer)
			board.AsciiArt(got_buff, swap_disc_colors)

			got := got_buff.String()
			expected := expected_buff.String()

			if got != expected {
				t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n", expected, got)
			}
		}
	}
}

func TestBoardDoRandomMove(t *testing.T) {
	for board := range genTestBoards() {
		clone := board
		if clone.Moves().Count() == 0 {
			// No moves -> panic() should be called
			func() {
				defer assertPanic(t)
				clone.DoRandomMove()
			}()
			continue
		}

		clone.DoRandomMove()

		found := false
		for child := range board.GenChildren() {
			if clone == child {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected child of:\n%s\n\nGot:\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

	}

	func() {
		defer assertPanic(t)
		board := Board{me: 0, opp: 0}
		board.DoRandomMove()
	}()
}

func TestBoardSwitchTurn(t *testing.T) {
	for board := range genTestBoards() {
		clone := board
		clone.SwitchTurn()
		if (board.me != clone.opp) || (board.opp != clone.me) {
			t.Errorf("Failure in Board.SwitchTurn()")
		}
	}
}

func TestBoardCountDiscs(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.me.Count() + board.opp.Count()
		got := board.CountDiscs()

		if expected != got {
			t.Errorf("Expected %d discs, got %d\n", expected, got)
		}
	}
}

func TestBoardCountEmpties(t *testing.T) {
	for board := range genTestBoards() {
		expected := 64 - (board.me.Count() + board.opp.Count())
		got := board.CountEmpties()

		if expected != got {
			t.Errorf("Expected %d discs, got %d\n", expected, got)
		}
	}
}

func TestBoardExactScore(t *testing.T) {
	for board := range genTestBoards() {
		var expected int

		me_count := int(board.me.Count())
		opp_count := int(board.opp.Count())
		empty_count := int(board.CountEmpties())

		if me_count > opp_count {
			expected = me_count + empty_count - opp_count
		} else if me_count < opp_count {
			expected = me_count - empty_count - opp_count
		} else {
			expected = 0
		}

		got := board.ExactScore()

		if expected != got {
			t.Errorf("Expected %d, got %d\n", expected, got)
		}
	}
}

func TestBoardMe(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.me
		got := board.Me()

		if expected != got {
			t.Errorf("Expected %d, got %d\n", expected, got)
		}
	}
}

func TestBoardOpp(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.opp
		got := board.Opp()

		if expected != got {
			t.Errorf("Expected %d, got %d\n", expected, got)
		}
	}
}
