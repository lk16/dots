package othello

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
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
		for y := uint(0); y < 8; y++ {
			for x := uint(0); x < 8; x++ {
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
							py := int(y) + (d+1)*dy
							px := int(x) + (d+1)*dx

							if (py < 0) || (py > 7) || (px < 0) || (px > 7) {
								break
							}

							qy := y + uint(d*dy)
							qx := x + uint(d*dx)

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

		// TODO be sure we don't send any invalid boards

		close(ch)
	}()
	return ch
}

func TestRandomBoard(t *testing.T) {
	for discs := 4; discs <= 64; discs++ {

		expected := discs

		board, err := NewRandomBoard(discs)
		assert.Nil(t, err)

		got := (board.me | board.opp).Count()

		assert.Equal(t, expected, got)
		assert.True(t, boardIsValid(board))
	}

	board, err := NewRandomBoard(3)
	assert.Nil(t, board)
	assert.Equal(t, ErrInvalidDiscAmount, err)

	board, err = NewRandomBoard(65)
	assert.Nil(t, board)
	assert.Equal(t, ErrInvalidDiscAmount, err)
}

func TestBoardDoMove(t *testing.T) {

	doMove := func(board *Board, index uint) BitSet {
		if (board.me | board.opp).Test(index) {
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
					if curx < 0 || curx >= 8 || cury < 0 || cury >= 8 {
						break
					}

					cur := uint(8*cury + curx)

					if board.opp.Test(cur) {
						s++
					} else {
						if board.me.Test(cur) && (s >= 2) {
							for p := 1; p < s; p++ {
								f := index + uint(p*(8*dy+dx))
								flipped.Set(f)
							}
						}
						break
					}
				}
			}
		}
		board.me |= flipped
		board.me.Set(index)
		board.opp &= ^board.me
		board.opp, board.me = board.me, board.opp
		return flipped
	}

	for board := range genTestBoards() {
		moves := board.Moves()
		for i := uint(0); i < 64; i++ {

			// othello.DoMove() should not be called for invalid moves
			if !moves.Test(i) {
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

			assert.Equal(t, expectedBoard, gotBoard)
			assert.Equal(t, expectedReturn, gotReturn)
		}
	}
}

func TestBoardMoves(t *testing.T) {

	boardMoves := func(board Board) BitSet {
		moves := BitSet(0)
		empties := ^(board.me | board.opp)

		for i := uint(0); i < 64; i++ {
			clone := board
			if empties.Test(i) && clone.DoMove(1<<i) != 0 {
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
		assert.Equal(t, expected, got)

		// board shouldn't change
		assert.Equal(t, board, clone)
	}
}

func (board *Board) getChildren() []Board {
	var children []Board
	empties := board.me | board.opp
	for i := uint(0); i < 64; i++ {
		clone := *board
		if clone.DoMove(1<<i) != 0 && !empties.Test(i) {
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

		discs := board.me | board.opp

		clone := board
		got := clone.GetChildren()

		// board shouldn't change
		assert.Equal(t, board, clone)

		// children set should be matching expected children set
		assert.ElementsMatch(t, expected, got)

		for _, child := range got {

			childDiscs := child.me | child.opp

			// pieces shouldn't be removed
			assert.Equal(t, discs, childDiscs&discs)

			// create copy to silence warnings
			childCopy := child

			assert.True(t, boardIsValid(&childCopy))
		}

	}
}

func TestBoardAsciiArt(t *testing.T) {
	for board := range genTestBoards() {

		moves := board.Moves()

		clone := board

		expected := new(bytes.Buffer)

		expected.WriteString("+-a-b-c-d-e-f-g-h-+\n")

		for y := uint(0); y < 8; y++ {
			expected.WriteString(fmt.Sprintf("%d ", y+1))

			for x := uint(0); x < 8; x++ {

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

		// board should not change
		assert.Equal(t, board, clone)

		assert.Equal(t, expected, got)
	}
}

func TestBoardDoRandomMove(t *testing.T) {
	for b := range genTestBoards() {

		// make copy to silence warnings
		board := b

		child := board

		child.DoRandomMove()

		if board.Moves() == 0 {
			// no moves means no change
			assert.Equal(t, child, board)
			continue
		}

		assert.Contains(t, board.GetChildren(), child)

		if boardIsValid(&board) {
			assert.True(t, boardIsValid(&child))
		}
	}
}

func TestBoardSwitchTurn(t *testing.T) {
	for board := range genTestBoards() {

		expected := Board{}
		expected.me, expected.opp = board.opp, board.me

		got := board
		got.SwitchTurn()

		assert.Equal(t, expected, got)
	}
}

func TestBoardCountDiscs(t *testing.T) {
	for board := range genTestBoards() {
		expected := (board.me | board.opp).Count()

		clone := board
		got := clone.CountDiscs()

		// board should not change
		assert.Equal(t, board, clone)

		assert.Equal(t, expected, got)
	}
}

func TestBoardCountEmpties(t *testing.T) {
	for board := range genTestBoards() {
		expected := 64 - (board.me | board.opp).Count()

		clone := board
		got := clone.CountEmpties()

		// board should not change
		assert.Equal(t, board, clone)

		assert.Equal(t, expected, got)
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

		// board shouldn't change
		assert.Equal(t, board, clone)

		assert.Equal(t, expected, got)
	}
}

func TestBoardMe(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.me

		clone := board
		got := clone.Me()

		// board shouldn't change
		assert.Equal(t, board, clone)

		assert.Equal(t, expected, got)
	}
}

func TestBoardOpp(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.opp

		clone := board
		got := clone.Opp()

		// board shouldn't change
		assert.Equal(t, board, clone)

		assert.Equal(t, expected, got)
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

	assert.Equal(t, expected, got)
	assert.True(t, boardIsValid(&got))
}

func TestBoardOpponentMoves(t *testing.T) {
	for board := range genTestBoards() {

		clone := board
		clone.SwitchTurn()

		expected := clone.Moves()
		got := board.OpponentMoves()

		assert.Equal(t, expected, got)
	}
}

func TestBoardNormalize(t *testing.T) {
	for board := range genTestBoards() {
		expected := board.rotate(0).Normalize()

		for r := 1; r < 8; r++ {
			got := board.rotate(r).Normalize()

			assert.Equal(t, expected, got)
		}
	}
}

var dummy int

func TestBoardCornerCountDifference(t *testing.T) {
	for board := range genTestBoards() {
		expected := (board.Me() & cornerMask).Count() - (board.Opp() & cornerMask).Count()
		got := board.CornerCountDifference()

		assert.Equal(t, expected, got)
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
		assert.Equal(t, expected, got)
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

		assert.Equal(t, expected, got)
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
