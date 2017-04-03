package board

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"dots/bitset"
)

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
				board.me = bitset.Bitset(1 << uint(y*8+x))

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

							board.opp |= bitset.Bitset(1 << uint(qy*8+qx))

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

func TestRandomBoard(t *testing.T) {
	for i := uint(4); i <= 64; i++ {
		board := RandomBoard(i)
		expected := i
		got := (board.me | board.opp).Count()
		if expected != got {
			t.Errorf("Expected %d, got %d\n", expected, got)
		}
	}
}

func TestBoardClone(t *testing.T) {
	board := Board{
		me:  1,
		opp: 2}
	clone := board.Clone()
	clone.me = 3
	if board.me != 1 {
		t.Errorf("Board.Clone() does not make a deep copy!\n")
	}
}

func (board *Board) doMove(index uint) bitset.Bitset {
	if (board.me | board.opp).TestBit(index) {
		return 0
	}
	flipped := bitset.Bitset(0)
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
							f := uint(int(index) + (8 * dy * p) + (dx * p))
							flipped |= bitset.Bitset(1 << f)
						}
					}
					break
				}
			}
		}
	}
	board.me |= flipped | bitset.Bitset(1<<index)
	board.opp &= ^board.me
	board.opp, board.me = board.me, board.opp
	return flipped
}

func TestBoardDoMove(t *testing.T) {
	for board := range genTestBoards() {
		moves := board.Moves()
		for i := uint(0); i < 64; i++ {
			if !moves.TestBit(i) {
				// board.DoMove() should not be called for invalid moves
				continue
			}

			clone := board.Clone()
			expected_return_val := clone.doMove(i)
			expected_board_val := clone

			clone = board.Clone()
			got_return_val := clone.DoMove(i)
			got_board_val := clone

			if (got_return_val != expected_return_val) || (got_board_val != expected_board_val) {
				t.Errorf("Doing move %c%d on board\n%s\n", 'a'+i%8, (i/8)+1, board.AsciiArt())
				t.Errorf("Expected return val:\n%s\n\nGot:\n%s\n\n", expected_return_val.AsciiArt(), got_return_val.AsciiArt())
				t.Errorf("Expected board:\n%s\n\nGot:\n%s\n\n", expected_board_val.AsciiArt(), got_board_val.AsciiArt())
				t.FailNow()
			}
		}
	}
}

func (board *Board) moves() bitset.Bitset {
	moves := bitset.Bitset(0)
	for i := uint(0); i < 64; i++ {
		if (board.me | board.opp).TestBit(i) {
			continue
		}
		clone := board.Clone()
		if clone.doMove(i) != 0 {
			moves |= bitset.Bitset(1 << i)
		}
	}
	return moves
}

func TestBoardMoves(t *testing.T) {
	for board := range genTestBoards() {
		clone := board
		expected := board.moves()
		got := board.Moves()
		if expected != got {
			t.Errorf("For board\n%s", board.AsciiArt())
			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n", expected.AsciiArt(), got.AsciiArt())
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
		clone := board.Clone()
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
				t.Errorf("board:\n%s\n\nchild: \n%s\n\n", board.AsciiArt(), g.AsciiArt())
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

		moves := board.Moves()

		ascii_art := board.AsciiArt()

		lines := strings.Split(ascii_art, "\n")

		expected := "+-a-b-c-d-e-f-g-h-+"
		if lines[0] != expected {
			t.Errorf("At lines[0]: expected '%s', got '%s'\n", expected, lines[0])
		}

		for y := uint(0); y < 8; y++ {

			expected_buf := new(bytes.Buffer)
			expected_buf.WriteString(fmt.Sprintf("%d ", y+1))

			for x := uint(0); x < 8; x++ {

				if board.me.TestBit(8*y + x) {
					expected_buf.WriteString("○ ")
				} else if board.opp.TestBit(8*y + x) {
					expected_buf.WriteString("● ")
				} else if moves.TestBit(8*y + x) {
					expected_buf.WriteString("- ")
				} else {
					expected_buf.WriteString("  ")
				}
			}

			expected_buf.WriteString("|")

			got := expected_buf.String()
			if lines[y+1] != got {
				t.Errorf("At lines[%d]: expected '%s', got '%s'\n", y+1, lines[y+1], got)
			}

		}

		expected = "+-----------------+"
		if lines[9] != expected {
			t.Errorf("At lines[9]: expected '%s', got '%s'\n", expected, lines[9])
		}

	}
}

func TestBoardDoRandomMove(t *testing.T) {
	for board := range genTestBoards() {
		clone := board
		if clone.Moves().Count() == 0 {
			// no children means Board.DoRandomMove() will panic
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
			t.Errorf("Expected child of:\n%s\n\nGot:\n%s\n\n", board.AsciiArt(), clone.AsciiArt())
		}

	}
}

func TestBoardSwitchTurn(t *testing.T) {
	for board := range genTestBoards() {
		clone := board.Clone()
		clone.SwitchTurn()
		if (board.me != clone.opp) || (board.opp != clone.me) {
			t.Errorf("Failure in Board.SwitchTurn()")
		}
	}
}

func TestBoardCountDiscs(t *testing.T) {
	for board := range genTestBoards() {
		excpected := board.me.Count() + board.opp.Count()
		got := board.CountDiscs()

		if excpected != got {
			t.Errorf("Expected %d discs, got %d\n", excpected, got)
		}
	}
}
