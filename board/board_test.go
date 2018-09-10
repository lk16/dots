package board

import (
	"bytes"
	"fmt"
	"math/bits"
	"math/rand"
	"testing"
)

func boardIsValid(board *Board) bool {

	// no field can be occupied by two discs
	if (board.me & board.opp) != 0 {
		return false
	}

	// start discs are never removed
	startBoard := NewBoard()
	startMask := startBoard.me | startBoard.opp

	if (board.me|board.opp)&startMask != startMask {
		return false
	}

	// no indication board is invalid
	return true
}

// Helper for this testing file
// Returns a string written by board.AsciiArt()
func (board Board) asciiArtString(swapDiscColors bool) (output string) {
	buffer := new(bytes.Buffer)
	board.ASCIIArt(buffer, swapDiscColors)
	output = buffer.String()
	return
}

// Test board generator
func genTestBoards() (ch chan Board) {
	ch = make(chan Board)
	go func() {

		// generate all boards with all flipping lines from each square

		// for each field
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				board := Board{}
				board.me |= uint64(1) << uint(y*8+x)

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

							board.opp |= uint64(1) << uint(qy*8+qx)

							ch <- board
						}
					}
				}
			}
		}

		ch <- Board{me: 0, opp: 0}
		ch <- *NewBoard()

		// random reachable boards with 4-64 discs
		for i := 0; i < 10; i++ {
			for discs := 4; discs <= 64; discs++ {
				ch <- *RandomBoard(discs)
			}
		}

		// random boards not necessarily reachable
		for i := 0; i < 100; i++ {
			board := Board{
				me:  rand.Uint64(),
				opp: rand.Uint64()}
			board.opp &^= board.me
			ch <- board
		}

		close(ch)
	}()
	return
}

func TestBoardIsValid(t *testing.T) {

	type testCase struct {
		board    Board
		expected bool
	}

	startBoard := *NewBoard()
	emptyBoard := Board{me: 0, opp: 0}

	all := ^uint64(0)

	startMask := startBoard.me | startBoard.opp
	dupBoard := Board{me: startMask, opp: startMask}
	winBoard := Board{me: all, opp: 0}
	loseBoard := Board{me: 0, opp: all}
	drawBoard := Board{me: all << 32, opp: all >> 32}

	testCases := []testCase{
		{emptyBoard, false},
		{startBoard, true},
		{dupBoard, false},
		{winBoard, true},
		{loseBoard, true},
		{drawBoard, true}}

	for _, testCase := range testCases {
		expected := testCase.expected
		board := testCase.board

		got := boardIsValid(&board)

		if expected != got {
			t.Errorf("Expected %t, got %t for board\n%s\n\n",
				expected, got, board.asciiArtString(false))
		}
	}
}

func TestRandomBoard(t *testing.T) {
	for discs := 4; discs <= 64; discs++ {

		expected := discs

		board := RandomBoard(discs)
		got := bits.OnesCount64(board.me | board.opp)

		if expected != got {
			t.Fatalf("Expected disc count %d, got %d\n", expected, got)
		}

		if !boardIsValid(board) {
			t.Fatalf("Invalid board:\n%s\n\n", board.asciiArtString(false))
		}
	}

	board := RandomBoard(3)
	if board != nil {
		t.Fatalf("Expected nil, got:\n%s\n\n", board.asciiArtString(false))
	}

	board = RandomBoard(65)
	if board != nil {
		t.Fatalf("Expected nil, got:\n%s\n\n", board.asciiArtString(false))
	}

}

func TestBoardCustom(t *testing.T) {

	me := uint64(1)
	opp := uint64(2)

	board := CustomBoard(me, opp)

	if board.me != me || board.opp != opp {
		t.Errorf("Custom board failed\n")
	}

}

func (board *Board) doMove(index uint) (flipped uint64) {
	if (board.me|board.opp)&(uint64(1)<<index) != 0 {
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
				if board.opp&(uint64(1)<<cur) != 0 {
					s++
				} else {
					if (board.me&(uint64(1)<<cur) != 0) && (s >= 2) {
						for p := 1; p < s; p++ {
							f := uint(int(index) + (p * (8*dy + dx)))
							flipped |= uint64(1) << f
						}
					}
					break
				}
			}
		}
	}
	board.me |= flipped
	board.me |= uint64(1) << index
	board.opp &= ^board.me
	board.opp, board.me = board.me, board.opp
	return
}

func TestBoardDoMove(t *testing.T) {
	for board := range genTestBoards() {
		moves := board.Moves()
		for i := uint(0); i < 64; i++ {
			if moves&(uint64(1)<<i) == 0 {
				// board.DoMove() should not be called for invalid moves
				continue
			}

			clone := board
			expectedReturn := clone.doMove(i)
			expectedBoard := clone

			clone = board
			gotReturn := clone.DoMove(int(i))
			gotBoard := clone

			if (gotReturn != expectedReturn) || (gotBoard != expectedBoard) {
				t.Errorf("Doing move %c%d on board\n%s\n", 'a'+i%8, (i/8)+1,
					board.asciiArtString(false))
				t.Errorf("Expected return val:\n%s\n\nGot:\n%s\n\n",
					bitsetASCIIArtString(expectedReturn), bitsetASCIIArtString(gotReturn))
				t.Errorf("Expected board:\n%s\n\nGot:\n%s\n\n",
					expectedBoard.asciiArtString(false), gotBoard.asciiArtString(false))
				t.FailNow()
			}
		}
	}
}

func TestBoardDoMoveN(t *testing.T) {

	for board := range genTestBoards() {

		clone := board

		doMoveFuncs := []func() uint64{
			clone.doMove0, clone.doMove1, clone.doMove2, clone.doMove3,
			clone.doMove4, clone.doMove5, clone.doMove6, clone.doMove7,
			clone.doMove8, clone.doMove9, clone.doMove10, clone.doMove11,
			clone.doMove12, clone.doMove13, clone.doMove14, clone.doMove15,
			clone.doMove16, clone.doMove17, clone.doMove18, clone.doMove19,
			clone.doMove20, clone.doMove21, clone.doMove22, clone.doMove23,
			clone.doMove24, clone.doMove25, clone.doMove26, clone.doMove27,
			clone.doMove28, clone.doMove29, clone.doMove30, clone.doMove31,
			clone.doMove32, clone.doMove33, clone.doMove34, clone.doMove35,
			clone.doMove36, clone.doMove37, clone.doMove38, clone.doMove39,
			clone.doMove40, clone.doMove41, clone.doMove42, clone.doMove43,
			clone.doMove44, clone.doMove45, clone.doMove46, clone.doMove47,
			clone.doMove48, clone.doMove49, clone.doMove50, clone.doMove51,
			clone.doMove52, clone.doMove53, clone.doMove54, clone.doMove55,
			clone.doMove56, clone.doMove57, clone.doMove58, clone.doMove59,
			clone.doMove60, clone.doMove61, clone.doMove62, clone.doMove63}

		moves := board.Moves()
		for i := uint(0); i < 64; i++ {
			if moves&(uint64(1)<<i) == 0 {
				// board.DoMove() should not be called for invalid moves
				continue
			}

			clone = board
			expected := clone.doMove(i)

			clone = board
			got := doMoveFuncs[i]()

			if clone != board {
				t.Errorf("Board was changed. Before:\n%s\n\nAfter:\n%s\n\n",
					board.asciiArtString(false), clone.asciiArtString(false))
			}

			if expected != got {
				t.Errorf("Doing move %c%d on board\n%s\n", 'a'+i%8, (i/8)+1,
					board.asciiArtString(false))
				t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
					bitsetASCIIArtString(expected), bitsetASCIIArtString(got))
				t.FailNow()
			}
		}
	}
}

func (board Board) moves() (moves uint64) {
	empties := ^(board.me | board.opp)

	for i := uint(0); i < 64; i++ {
		clone := board
		if (empties&(uint64(1)<<i) != 0) && clone.DoMove(int(i)) != 0 {
			moves |= uint64(1) << i
		}
	}
	return
}

func TestBoardMoves(t *testing.T) {
	for board := range genTestBoards() {

		clone := board
		expected := board.moves()

		got := clone.Moves()
		if expected != got {
			t.Errorf("For board\n%s", board.asciiArtString(false))
			t.Fatalf("Expected:\n%s\n\nGot:\n%s\n\n",
				bitsetASCIIArtString(expected), bitsetASCIIArtString(got))
		}
		if clone != board {
			t.Fatalf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}
	}
}

func (board *Board) getChildren() (children []Board) {
	for i := uint(0); i < 64; i++ {
		clone := *board
		if clone.doMove(i) != 0 {
			children = append(children, clone)
		}
	}
	return
}

func TestBoardGetChildren(t *testing.T) {
	for board := range genTestBoards() {

		valid := boardIsValid(&board)

		expected := board.getChildren()
		expectedSet := make(map[Board]struct{}, 10)
		for _, e := range expected {
			expectedSet[e] = struct{}{}
		}

		discs := board.me | board.opp

		clone := board
		got := clone.GetChildren()

		if clone != board {
			t.Fatalf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		for _, child := range got {

			childDiscs := child.me | child.opp

			if (childDiscs & discs) != discs {
				t.Fatalf("Pieces were removed from board with board.GetChildren()\n")
			}

			if valid && !boardIsValid(&child) {
				t.Fatalf("Valid board:\n%s\n\nInvalid child:\n%s\n\n",
					board.asciiArtString(false), child.asciiArtString(false))
			}

		}

		if len(got) != len(expected) {
			t.Fatalf("Expected %d children, got %d.\n", len(expected), len(got))
		}

		for _, g := range got {
			if _, ok := expectedSet[g]; !ok {
				t.Fatalf("Children sets are unequal.\n")
			}
		}

	}
}

func TestBoardAsciiArt(t *testing.T) {
	for board := range genTestBoards() {

		for _, swapDiscColors := range []bool{true, false} {

			moves := board.Moves()

			clone := board
			if swapDiscColors {
				clone.SwitchTurn()
			}

			expected := new(bytes.Buffer)

			expected.WriteString("+-a-b-c-d-e-f-g-h-+\n")

			for y := uint(0); y < 8; y++ {
				expected.WriteString(fmt.Sprintf("%d ", y+1))

				for x := uint(0); x < 8; x++ {
					mask := uint64(1) << (8*y + x)
					if clone.me&mask != 0 {
						expected.WriteString("○ ")
					} else if clone.opp&mask != 0 {
						expected.WriteString("● ")
					} else if moves&mask != 0 {
						expected.WriteString("- ")
					} else {
						expected.WriteString("  ")
					}
				}

				expected.WriteString("|\n")
			}

			expected.WriteString("+-----------------+\n")

			got := new(bytes.Buffer)
			clone = board
			clone.ASCIIArt(got, swapDiscColors)

			if clone != board {
				t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
					board.asciiArtString(false), clone.asciiArtString(false))
			}

			if got.String() != expected.String() {
				t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
					expected.String(), got.String())
			}
		}
	}
}

func TestBoardDoRandomMove(t *testing.T) {
	for board := range genTestBoards() {
		clone := board

		clone.DoRandomMove()

		if board.Moves() == 0 {
			// no moves means no change
			if clone != board {
				t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
					board.asciiArtString(false), clone.asciiArtString(false))
			}
			continue
		}

		found := false
		for _, child := range board.GetChildren() {
			if clone == child {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected child of:\n%s\n\nGot:\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		if boardIsValid(&board) && !boardIsValid(&clone) {
			t.Errorf("Found board:\n%s\n\nWith invalid child:\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}
	}
}

func TestBoardSwitchTurn(t *testing.T) {
	for board := range genTestBoards() {

		expected := Board{}
		expected.me, expected.opp = board.opp, board.me

		got := board
		got.SwitchTurn()

		if expected != got {
			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
				expected.asciiArtString(false), got.asciiArtString(false))
			t.FailNow()
		}
	}
}

func TestBoardCountDiscs(t *testing.T) {
	for board := range genTestBoards() {
		expected := bits.OnesCount64(board.me | board.opp)

		clone := board
		got := clone.CountDiscs()

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		if expected != got {
			t.Errorf("Expected %d discs, got %d\n", expected, got)
		}
	}
}

func TestBoardCountEmpties(t *testing.T) {
	for board := range genTestBoards() {
		expected := 64 - bits.OnesCount64(board.me|board.opp)

		clone := board
		got := clone.CountEmpties()

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		if expected != got {
			t.Errorf("Expected %d discs, got %d\n", expected, got)
		}
	}
}

func TestBoardExactScore(t *testing.T) {
	for board := range genTestBoards() {
		var expected int

		meCount := bits.OnesCount64(board.me)
		oppCount := bits.OnesCount64(board.opp)
		emptyCount := int(board.CountEmpties())

		if meCount > oppCount {
			expected = meCount + emptyCount - oppCount
		} else if meCount < oppCount {
			expected = meCount - emptyCount - oppCount
		} else {
			expected = 0
		}

		clone := board
		got := clone.ExactScore()
		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		if expected != got {
			t.Errorf("Expected %d, got %d\n", expected, got)
		}
	}
}

func TestBoardMe(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.me

		clone := board
		got := clone.Me()

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		if expected != got {
			t.Errorf("Expected %s, got %s\n",
				bitsetASCIIArtString(expected), bitsetASCIIArtString(got))
		}
	}
}

func TestBoardOpp(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.opp

		clone := board
		got := clone.Opp()

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.asciiArtString(false), clone.asciiArtString(false))
		}

		if expected != got {
			t.Errorf("Expected %v, got %v\n",
				bitsetASCIIArtString(expected), bitsetASCIIArtString(got))
		}
	}
}

func TestBoardNewBoard(t *testing.T) {

	// center of the start board:
	// W B
	// B W

	expected := Board{}
	expected.me = uint64(1)<<(4*8+3) | uint64(1)<<(3*8+4)
	expected.opp = uint64(1)<<(3*8+3) | uint64(1)<<(4*8+4)

	got := *NewBoard()

	if expected != got {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
			expected.asciiArtString(false), got.asciiArtString(false))
	}

	if !boardIsValid(&got) {
		t.Errorf("Start board is invalid:\n%s\n\n", got.asciiArtString(false))
	}
}

func TestBoardOpponentMoves(t *testing.T) {
	for board := range genTestBoards() {

		clone := board
		clone.SwitchTurn()

		expected := clone.Moves()
		got := board.OpponentMoves()

		if expected != got {
			t.Errorf("Expected %d, got %d", expected, got)
		}
	}
}

func TestBoardNormalize(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.rotate(0).Normalize()

		for r := 1; r < 8; r++ {
			got := board.rotate(r).Normalize()

			if expected != got {
				t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
					expected.asciiArtString(false), got.asciiArtString(false))
				t.FailNow()
			}
		}
	}
}