package othello

import (
	"bytes"
	"fmt"
	"log"
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

	return (board.me|board.opp)&startMask == startMask
}

// Test othello generator
func genTestBoards() chan Board {
	ch := make(chan Board)
	go func() {

		// generate all boards with all flipping lines from each square

		// for each field
		for y := 0; y < 8; y++ {
			for x := 0; x < 8; x++ {
				board := Board{}
				board.me.Set(y*8 + x)

				// for each direction
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if (dy == 0) && (dx == 0) {
							continue
						}
						board.opp = 0

						// for each distance
						for d := 1; d <= 6; d++ {

							// check if me can still flip within othello boundaries
							py := y + (d+1)*dy
							px := x + (d+1)*dx

							if (py < 0) || (py > 7) || (px < 0) || (px > 7) {
								break
							}

							qy := y + d*dy
							qx := x + d*dx

							board.opp.Set(qy*8 + qx)

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
				board, err := NewRandomBoard(discs)
				if err != nil {
					log.Printf("genTestBoards() breaking: %s", err)
				}
				ch <- *board
			}
		}

		close(ch)
	}()
	return ch
}

func TestRandomBoard(t *testing.T) {
	for discs := 4; discs <= 64; discs++ {

		expected := discs

		board, err := NewRandomBoard(discs)
		if err != nil {
			t.Error(err)
		}

		got := (board.me | board.opp).Count()

		if expected != got {
			t.Fatalf("Expected disc count %d, got %d\n", expected, got)
		}

		if !boardIsValid(board) {
			t.Fatalf("Invalid othello:\n%s\n\n", board.String())
		}
	}

	board, err := NewRandomBoard(3)
	if err == nil {
		t.Fatalf("Expected error, got nil\n%s\n\n", board.String())
	}

	board, err = NewRandomBoard(65)
	if err == nil {
		t.Fatalf("Expected error, got nil\n%s\n\n", board.String())
	}
}

func TestBoardDoMove(t *testing.T) {

	doMove := func(board *Board, index uint) BitSet {
		if (board.me | board.opp).Test(int(index)) {
			return 0
		}
		flipped := BitSet(0)
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				s := 1
				for {
					curx := int(index%8) + (dx * s)
					cury := int(index/8) + (dy * s)
					cur := 8*cury + curx
					if curx < 0 || curx >= 8 || cury < 0 || cury >= 8 {
						break
					}
					if board.opp.Test(cur) {
						s++
					} else {
						if board.me.Test(cur) && (s >= 2) {
							for p := 1; p < s; p++ {
								f := int(index) + (p * (8*dy + dx))
								flipped.Set(f)
							}
						}
						break
					}
				}
			}
		}
		board.me |= flipped
		board.me.Set(int(index))
		board.opp &= ^board.me
		board.opp, board.me = board.me, board.opp
		return flipped
	}

	for board := range genTestBoards() {
		moves := board.Moves()
		for i := uint(0); i < 64; i++ {

			// othello.DoMove() should not be called for invalid moves
			if !moves.Test(int(i)) {
				continue
			}

			// don't play start disc moves
			if i == 27 || i == 28 || i == 35 || i == 36 {
				continue
			}

			clone := board
			expectedReturn := doMove(&clone, i)
			expectedBoard := clone

			clone = board
			gotReturn := clone.DoMove(BitSet(1 << i))
			gotBoard := clone

			if (gotReturn != expectedReturn) || (gotBoard != expectedBoard) {
				t.Errorf("Doing move %c%d on othello\n%s\n", 'a'+i%8, (i/8)+1,
					board.String())
				t.Errorf("Expected return value:\n%s\n\nGot:\n%s\n\n",
					BitSet(expectedReturn).String(), BitSet(gotReturn).String())
				t.Errorf("Expected othello:\n%s\n\nGot:\n%s\n\n",
					expectedBoard.String(), gotBoard.String())
				t.FailNow()
			}
		}
	}
}

func TestBoardMoves(t *testing.T) {

	boardMoves := func(board Board) BitSet {
		moves := BitSet(0)
		empties := ^(board.me | board.opp)

		for i := 0; i < 64; i++ {
			clone := board
			if empties.Test(i) && clone.DoMove(1<<uint(i)) != 0 {
				moves.Set(i)
			}
		}
		return moves
	}

	for b := range genTestBoards() {

		// create copy to silence warnings
		board := b

		if !boardIsValid(&board) {
			continue
		}

		clone := board
		expected := boardMoves(board)

		got := clone.Moves()
		if expected != got {
			t.Errorf("For othello\n%s", board.String())
			t.Fatalf("Expected:\n%s\n\nGot:\n%s\n\n",
				BitSet(expected).String(), BitSet(got).String())
		}
		if clone != board {
			t.Fatalf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.String(), clone.String())
		}
	}
}

func (board *Board) getChildren() []Board {
	var children []Board
	empties := board.me | board.opp
	for i := uint(0); i < 64; i++ {
		clone := *board
		if clone.DoMove(1<<i) != 0 && !empties.Test(int(i)) {
			children = append(children, clone)
		}
	}
	return children
}

func TestBoardGetChildren(t *testing.T) {
	for b := range genTestBoards() {

		// create copy to silence warnings
		board := b

		if !boardIsValid(&board) {
			continue
		}

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
				board.String(), clone.String())
		}

		for _, child := range got {

			childDiscs := child.me | child.opp

			if (childDiscs & discs) != discs {
				t.Fatalf("Pieces were removed from othello with othello.GetChildren()\n")
			}

			// create copy to silence warnings
			childCopy := child

			if !boardIsValid(&childCopy) {
				t.Fatalf("Valid othello:\n%s\n\nInvalid child:\n%s\n\n",
					board.String(), child.String())
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

		moves := board.Moves()

		clone := board

		expected := new(bytes.Buffer)

		expected.WriteString("+-a-b-c-d-e-f-g-h-+\n")

		for y := 0; y < 8; y++ {
			expected.WriteString(fmt.Sprintf("%d ", y+1))

			for x := 0; x < 8; x++ {
				if clone.me.Test(8*y + x) {
					expected.WriteString("○ ")
				} else if clone.opp.Test(8*y + x) {
					expected.WriteString("● ")
				} else if moves.Test(8*y + x) {
					expected.WriteString("- ")
				} else {
					expected.WriteString("  ")
				}
			}

			expected.WriteString("|\n")
		}

		expected.WriteString("+-----------------+\nTo move: ○\n")
		expected.WriteString("Raw: " + fmt.Sprintf("%#v", clone) + "\n")

		got := new(bytes.Buffer)
		clone = board
		got.WriteString(board.String())

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.String(), clone.String())
		}

		if got.String() != expected.String() {
			t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
				expected.String(), got.String())
		}

	}
}

func TestBoardDoRandomMove(t *testing.T) {
	for b := range genTestBoards() {

		// make copy to silence warnings
		board := b

		clone := board

		clone.DoRandomMove()

		if board.Moves() == 0 {
			// no moves means no change
			if clone != board {
				t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
					board.String(), clone.String())
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
				board.String(), clone.String())
		}

		if boardIsValid(&board) && !boardIsValid(&clone) {
			t.Errorf("Found othello:\n%s\n\nWith invalid child:\n%s\n\n",
				board.String(), clone.String())
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
				expected.String(), got.String())
			t.FailNow()
		}
	}
}

func TestBoardCountDiscs(t *testing.T) {
	for board := range genTestBoards() {
		expected := (board.me | board.opp).Count()

		clone := board
		got := clone.CountDiscs()

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.String(), clone.String())
		}

		if expected != got {
			t.Errorf("Expected %d discs, got %d\n", expected, got)
		}
	}
}

func TestBoardCountEmpties(t *testing.T) {
	for board := range genTestBoards() {
		expected := 64 - (board.me | board.opp).Count()

		clone := board
		got := clone.CountEmpties()

		if clone != board {
			t.Errorf("Board was changed. Before:\n%s\n\nAfter\n%s\n\n",
				board.String(), clone.String())
		}

		if expected != got {
			t.Errorf("Expected %d discs, got %d\n", expected, got)
		}
	}
}

func TestBoardExactScore(t *testing.T) {
	for board := range genTestBoards() {
		var expected int

		meCount := board.me.Count()
		oppCount := board.opp.Count()
		emptyCount := board.CountEmpties()

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
				board.String(), clone.String())
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
				board.String(), clone.String())
		}

		if expected != got {
			t.Errorf("Expected %s, got %s\n",
				BitSet(expected).String(), BitSet(got).String())
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
				board.String(), clone.String())
		}

		if expected != got {
			t.Errorf("Expected %v, got %v\n",
				BitSet(expected).String(), BitSet(got).String())
		}
	}
}

func TestBoardNewBoard(t *testing.T) {

	// center of the start othello:
	// W B
	// B W

	expected := Board{}
	expected.me = BitSet(1)<<(4*8+3) | BitSet(1)<<(3*8+4)
	expected.opp = BitSet(1)<<(3*8+3) | BitSet(1)<<(4*8+4)

	got := *NewBoard()

	if expected != got {
		t.Errorf("Expected:\n%s\n\nGot:\n%s\n\n",
			expected.String(), got.String())
	}

	if !boardIsValid(&got) {
		t.Errorf("Start othello is invalid:\n%s\n\n", got.String())
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
					expected.String(), got.String())
				t.FailNow()
			}
		}
	}
}

var dummy int

func TestBoardCornerCountDifference(t *testing.T) {
	for board := range genTestBoards() {
		expected := (board.Me() & cornerMask).Count() - (board.Opp() & cornerMask).Count()
		got := board.CornerCountDifference()

		if expected != got {
			t.Errorf("\n%s\n\nExpected: %d\nGot: %d\n", board.String(), expected, got)
			t.FailNow()
		}
	}
}

func BenchmarkCornerCountDifference(b *testing.B) {
	board := NewBoard()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dummy = board.CornerCountDifference()
	}
	b.StopTimer()
}

func TestBoardXsquareCountDifference(t *testing.T) {
	for board := range genTestBoards() {
		expected := (board.Me() & xSquareMask).Count() - (board.Opp() & xSquareMask).Count()
		got := board.XsquareCountDifference()
		if expected != got {
			t.Errorf("\n%s\n\nExpected: %d\nGot: %d\n", board.String(), expected, got)
			t.FailNow()
		}
	}
}

func BenchmarkXsquareCountDifference(b *testing.B) {
	board := NewBoard()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dummy = board.XsquareCountDifference()
	}
	b.StopTimer()
}

func TestBoardCsquareCountDifference(t *testing.T) {
	for board := range genTestBoards() {
		expected := (board.Me() & cSquareMask).Count() - (board.Opp() & cSquareMask).Count()
		got := board.CsquareCountDifference()
		if expected != got {
			t.Errorf("\n%s\n\nExpected: %d\nGot: %d\n", board.String(), expected, got)
			t.FailNow()
		}
	}
}

func BenchmarkCsquareCountDifference(b *testing.B) {
	board := NewBoard()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dummy = board.CsquareCountDifference()
	}
	b.StopTimer()
}

func BenchmarkPotentialMoveCountDifference(b *testing.B) {
	board := NewBoard()
	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		dummy = board.PotentialMoveCountDifference()
	}
	b.StopTimer()
}

var dummyBoardSlice []Board

func BenchmarkGetChildrenXot(b *testing.B) {
	rand.Seed(0)
	board := NewXotBoard()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dummyBoardSlice = board.GetChildren()
	}
	b.StopTimer()
}

func TestDoMove0Right(t *testing.T) {

	rand.Seed(0)
	i := 0
	for i < 100000 {

		board := Board{
			me:  BitSet(rand.Intn(256)),
			opp: BitSet(rand.Intn(256)),
		}

		board.me &= 0xFE
		board.opp &= 0xFE

		if board.me&board.opp != 0 {
			continue
		}

		if board.opp >= board.me {
			continue
		}

		copy := board
		lhs := board.DoMove(1 << 0)

		board = copy
		rhs := board.doMove0Right()

		board = copy

		if lhs != rhs {
			log.Printf(board.String())
			log.Printf(lhs.String())
			log.Printf(rhs.String())
			t.FailNow()
		}

		i++
	}
}

func TestDoMove0Down(t *testing.T) {

	rand.Seed(0)
	i := 0
	for i < 100000 {

		board := Board{
			me:  BitSet(rand.Uint64()),
			opp: BitSet(rand.Uint64()),
		}

		board.me &= 0x0101010101010100
		board.opp &= 0x0101010101010100

		if board.me&board.opp != 0 {
			continue
		}

		if board.opp >= board.me {
			continue
		}

		copy := board
		lhs := board.DoMove(1 << 0)

		board = copy
		rhs := board.doMove0Down()

		board = copy

		if lhs != rhs {
			log.Printf("\n%s", board.String())
			log.Printf("\n%s", lhs.String())
			log.Printf("\n%s", rhs.String())
			t.FailNow()
		}

		i++
	}
}

var dummyBitSet BitSet

func BenchmarkDoMove0Right(b *testing.B) {
	var board Board
	for i := 0; i < b.N; i++ {
		board.me = BitSet(rand.Intn(256))
		board.opp = BitSet(rand.Intn(256))
		board.me &^= board.opp

		board.me &^= (1 << 0)
		board.opp &^= (1 << 0)

		dummyBitSet = board.doMove0Right()
	}
}

func BenchmarkDoMove(b *testing.B) {
	var board Board
	for i := 0; i < b.N; i++ {
		board.me = BitSet(rand.Intn(256))
		board.opp = BitSet(rand.Intn(256))
		board.me &^= board.opp

		board.me &^= (1 << 0)
		board.opp &^= (1 << 0)

		dummyBitSet = board.DoMove(1 << 0)
	}
}
