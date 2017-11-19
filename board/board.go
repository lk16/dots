package board

import (
	"bytes"
	"fmt"
	"io"
	"math/bits"
	"math/rand"
)

func bitsetASCIIArtString(bs uint64) (output string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	for y := uint(0); y < 8; y++ {
		buffer.WriteString(fmt.Sprintf("%d ", y+1))

		for x := uint(0); x < 8; x++ {
			f := y*8 + x
			if bs&uint64(1<<f) != 0 {
				buffer.WriteString("@ ")
			} else {
				buffer.WriteString("- ")
			}

		}
		buffer.WriteString("|\n")
	}
	buffer.WriteString("+-----------------+\n")

	output = buffer.String()
	return
}

// Board represents the state of an othello board game.
// It does not keep track which discs are white or black.
// Instead it keeps track which discs are owned by the player to move.
type Board struct {
	me, opp uint64
}

// NewBoard returns a Board representing the initial state
func NewBoard() *Board {
	return &Board{
		me:  1<<28 | 1<<35,
		opp: 1<<27 | 1<<36}
}

// CustomBoard returns a Board with a custom state
func CustomBoard(me, opp uint64) (board *Board) {
	return &Board{
		me:  me,
		opp: opp}
}

// RandomBoard returns a random Board with a given number of discs
func RandomBoard(discs int) (board *Board) {

	if discs < 4 || discs > 64 {
		return nil
	}

	board = NewBoard()
	skips := 0

	for board.CountDiscs() != discs {

		if skips == 2 {
			// Stuck. Try again.
			board = NewBoard()
			skips = 0
			continue
		}

		if bits.OnesCount64(board.Moves()) == 0 {
			skips++
			board.SwitchTurn()
			continue
		}

		skips = 0
		board.DoRandomMove()
	}

	return
}

func (board Board) rotate(rotation int) (rotated Board) {

	rotate := func(bitset uint64, rotation int) (result uint64) {
		result = bitset
		if rotation&1 != 0 {
			result = (result&0x00000000FFFFFFFF)<<32 | (result&0xFFFFFFFF00000000)>>32
			result = (result&0x0000FFFF0000FFFF)<<16 | (result&0xFFFF0000FFFF0000)>>16
			result = (result&0x00FF00FF00FF00FF)<<8 | (result&0xFF00FF00FF00FF00)>>8
		}
		if rotation&2 != 0 {
			result = (result&0x0F0F0F0F0F0F0F0F)<<4 | (result&0xF0F0F0F0F0F0F0F0)>>4
			result = (result&0x3333333333333333)<<2 | (result&0xCCCCCCCCCCCCCCCC)>>2
			result = (result&0x5555555555555555)<<1 | (result&0xAAAAAAAAAAAAAAAA)>>1
		}
		if rotation&4 != 0 {
			var tmp uint64
			k1 := uint64(0xaa00aa00aa00aa00)
			k2 := uint64(0xcccc0000cccc0000)
			k4 := uint64(0xf0f0f0f00f0f0f0f)
			tmp = result ^ (result << 36)
			result ^= k4 & (tmp ^ (result >> 36))
			tmp = k2 & (result ^ (result << 18))
			result ^= tmp ^ (tmp >> 18)
			tmp = k1 & (result ^ (result << 9))
			result ^= tmp ^ (tmp >> 9)
			return result
		}

		return
	}

	rotated = board
	rotated.me = rotate(rotated.me, rotation)
	rotated.opp = rotate(rotated.opp, rotation)
	return

}

// Normalize normalizes a board with regards to symmetry
func (board *Board) Normalize() (rotation int) {
	min := board.rotate(0)
	rotation = 0

	for r := 1; r < 8; r++ {
		cur := board.rotate(r)
		if cur.me < min.me {
			min = cur
			rotation = r
		}
		if cur.me == min.me {
			if cur.opp < min.opp {
				min = cur
				rotation = r
			}
		}
	}
	*board = min
	return
}

// Unnormalize undoes normalization using the return value of Normalize()
func (board *Board) Unnormalize(rotation int) {
	*board = board.rotate(rotation)
}

// ASCIIArt writes ascii-art of a Board to a writer
func (board Board) ASCIIArt(writer io.Writer, SwapDiscColors bool) {

	buffer := new(bytes.Buffer)
	buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	moves := board.Moves()

	if SwapDiscColors {
		board.SwitchTurn()
	}

	for y := uint(0); y < 8; y++ {
		buffer.WriteString(fmt.Sprintf("%d ", y+1))

		for x := uint(0); x < 8; x++ {
			mask := uint64(1) << (y*8 + x)
			if board.me&mask != 0 {
				buffer.WriteString("○ ")
			} else if board.opp&mask != 0 {
				buffer.WriteString("● ")
			} else if moves&mask != 0 {
				buffer.WriteString("- ")
			} else {
				buffer.WriteString("  ")
			}
		}

		buffer.WriteString("|\n")
	}
	buffer.WriteString("+-----------------+\n")

	writer.Write(buffer.Bytes())
}

// Moves returns a bitset of valid moves for a Board
func (board Board) Moves() uint64 {
	return moves(board.me, board.opp)
}

// OpponentMoves returns a bitset with all valid moves for the opponent
func (board Board) OpponentMoves() uint64 {
	return moves(board.opp, board.me)
}

func moves(me, opp uint64) (movesSet uint64) {
	// Returns a subset of the moves for a Board
	movesPartial := func(me, mask, n uint64) (moves uint64) {
		flipL := mask & (me << n)
		flipL |= mask & (flipL << n)
		maskL := mask & (mask << n)
		flipL |= maskL & (flipL << (2 * n))
		flipL |= maskL & (flipL << (2 * n))
		flipR := mask & (me >> n)
		flipR |= mask & (flipR >> n)
		maskR := mask & (mask >> n)
		flipR |= maskR & (flipR >> (2 * n))
		flipR |= maskR & (flipR >> (2 * n))
		moves = (flipL << n) | (flipR >> n)
		return
	}

	// this function is a modified version of code from Edax
	mask := opp & 0x7E7E7E7E7E7E7E7E

	movesSet = movesPartial(me, mask, 1)
	movesSet |= movesPartial(me, mask, 7)
	movesSet |= movesPartial(me, mask, 9)
	movesSet |= movesPartial(me, opp, 8)

	movesSet &^= (me | opp)
	return
}

// DoMove does a move and returns the flipped discs
func (board *Board) DoMove(index int) (flipped uint64) {

	switch index {
	case 0:
		flipped = board.doMove0()
	case 1:
		flipped = board.doMove1()
	case 2:
		flipped = board.doMove2()
	case 3:
		flipped = board.doMove3()
	case 4:
		flipped = board.doMove4()
	case 5:
		flipped = board.doMove5()
	case 6:
		flipped = board.doMove6()
	case 7:
		flipped = board.doMove7()
	case 8:
		flipped = board.doMove8()
	case 9:
		flipped = board.doMove9()
	case 10:
		flipped = board.doMove10()
	case 11:
		flipped = board.doMove11()
	case 12:
		flipped = board.doMove12()
	case 13:
		flipped = board.doMove13()
	case 14:
		flipped = board.doMove14()
	case 15:
		flipped = board.doMove15()
	case 16:
		flipped = board.doMove16()
	case 17:
		flipped = board.doMove17()
	case 18:
		flipped = board.doMove18()
	case 19:
		flipped = board.doMove19()
	case 20:
		flipped = board.doMove20()
	case 21:
		flipped = board.doMove21()
	case 22:
		flipped = board.doMove22()
	case 23:
		flipped = board.doMove23()
	case 24:
		flipped = board.doMove24()
	case 25:
		flipped = board.doMove25()
	case 26:
		flipped = board.doMove26()
	case 27:
		flipped = board.doMove27()
	case 28:
		flipped = board.doMove28()
	case 29:
		flipped = board.doMove29()
	case 30:
		flipped = board.doMove30()
	case 31:
		flipped = board.doMove31()
	case 32:
		flipped = board.doMove32()
	case 33:
		flipped = board.doMove33()
	case 34:
		flipped = board.doMove34()
	case 35:
		flipped = board.doMove35()
	case 36:
		flipped = board.doMove36()
	case 37:
		flipped = board.doMove37()
	case 38:
		flipped = board.doMove38()
	case 39:
		flipped = board.doMove39()
	case 40:
		flipped = board.doMove40()
	case 41:
		flipped = board.doMove41()
	case 42:
		flipped = board.doMove42()
	case 43:
		flipped = board.doMove43()
	case 44:
		flipped = board.doMove44()
	case 45:
		flipped = board.doMove45()
	case 46:
		flipped = board.doMove46()
	case 47:
		flipped = board.doMove47()
	case 48:
		flipped = board.doMove48()
	case 49:
		flipped = board.doMove49()
	case 50:
		flipped = board.doMove50()
	case 51:
		flipped = board.doMove51()
	case 52:
		flipped = board.doMove52()
	case 53:
		flipped = board.doMove53()
	case 54:
		flipped = board.doMove54()
	case 55:
		flipped = board.doMove55()
	case 56:
		flipped = board.doMove56()
	case 57:
		flipped = board.doMove57()
	case 58:
		flipped = board.doMove58()
	case 59:
		flipped = board.doMove59()
	case 60:
		flipped = board.doMove60()
	case 61:
		flipped = board.doMove61()
	case 62:
		flipped = board.doMove62()
	case 63:
		flipped = board.doMove63()
	}

	moveBit := uint64(1 << uint(index))

	tmp := board.me | flipped | moveBit

	board.me = board.opp &^ tmp
	board.opp = tmp

	return
}

// GetChildren returns a slice with all children of a Board
func (board Board) GetChildren() (children []Board) {

	moves := board.Moves()
	children = make([]Board, bits.OnesCount64(moves))

	for i := range children {
		moveIndex := bits.TrailingZeros64(moves)

		// reset lowest bit
		moves &= moves - 1

		children[i] = board
		children[i].DoMove(moveIndex)
	}
	return
}

// UndoMove undoes a move
func (board *Board) UndoMove(moveBit, flipped uint64) {
	tmp := board.me
	board.me = board.opp &^ (flipped | moveBit)
	board.opp = tmp | flipped
}

// DoRandomMove does a random move on a Board
// If no moves are possible, DoRandomMove does nothing
func (board *Board) DoRandomMove() {
	children := board.GetChildren()
	if len(children) == 0 {
		return
	}
	*board = children[rand.Intn(len(children))]
}

// SwitchTurn effectively passes a turn
func (board *Board) SwitchTurn() {
	board.me, board.opp = board.opp, board.me
}

// CountDiscs counts the number of discs on a Board
func (board Board) CountDiscs() int {
	return bits.OnesCount64(board.me | board.opp)
}

// CountEmpties returns the number of empty fields on a Board
func (board Board) CountEmpties() int {
	return 64 - board.CountDiscs()
}

// ExactScore returns the final score of a Board
func (board Board) ExactScore() int {
	meCount := bits.OnesCount64(board.me)
	oppCount := bits.OnesCount64(board.opp)

	if meCount > oppCount {
		return 64 - (2 * oppCount)
	}
	if meCount < oppCount {
		return -64 + (2 * meCount)
	}
	return 0
}

// Me returns a bitset with the discs of the player to move
func (board Board) Me() uint64 {
	return board.me
}

// Opp returns a bitset with the discs of the opponent of the player to move
func (board Board) Opp() uint64 {
	return board.opp
}

// Flips discs on a Board, given a flipping line.
// This only affects the directions right, left down, down and right down
// Returns the flipped discs.
func (board *Board) doMoveToHigherBits(line uint64) uint64 {
	b := (^line | board.opp) + 1
	lineMask := (b & -b & board.me) - 1
	x := (lineMask >> 63) - 1
	return x & lineMask & board.opp & line
}

// Flips discs on a Board, given a flipping line.
// This only affects the directions left up, up, right up and left
// Returns the flipped discs.
func (board *Board) doMoveToLowerBits(line uint64) uint64 {
	lineMask := line & board.me
	if lineMask == 0 {
		return 0
	}
	bit := uint64(1) << uint(bits.Len64(lineMask)-1)
	line &^= uint64((bit << 1) - 1)

	if line&board.opp == line {
		return line
	}
	return 0
}

//Does the move at field 0.
//Returns the flipped discs.
func (board *Board) doMove0() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000000000FE)
	flipped |= board.doMoveToHigherBits(0x0101010101010100)
	flipped |= board.doMoveToHigherBits(0x8040201008040200)
	return
}

//Does the move at field 1.
//Returns the flipped discs.
func (board *Board) doMove1() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000000000FC)
	flipped |= board.doMoveToHigherBits(0x0202020202020200)
	flipped |= board.doMoveToHigherBits(0x0080402010080400)
	return
}

//Does the move at field 2.
//Returns the flipped discs.
func (board *Board) doMove2() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000000000F8)
	flipped |= board.doMoveToHigherBits(0x0000000000010200)
	flipped |= board.doMoveToHigherBits(0x0404040404040400)
	flipped |= board.doMoveToHigherBits(0x0000804020100800)
	flipped |= board.doMoveToLowerBits(0x0000000000000003)
	return
}

//Does the move at field 3.
//Returns the flipped discs.
func (board *Board) doMove3() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000000000F0)
	flipped |= board.doMoveToHigherBits(0x0000000001020400)
	flipped |= board.doMoveToHigherBits(0x0808080808080800)
	flipped |= board.doMoveToHigherBits(0x0000008040201000)
	flipped |= board.doMoveToLowerBits(0x0000000000000007)
	return
}

//Does the move at field 4.
//Returns the flipped discs.
func (board *Board) doMove4() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000000000E0)
	flipped |= board.doMoveToHigherBits(0x0000000102040800)
	flipped |= board.doMoveToHigherBits(0x1010101010101000)
	flipped |= board.doMoveToHigherBits(0x0000000080402000)
	flipped |= board.doMoveToLowerBits(0x000000000000000F)
	return
}

//Does the move at field 5.
//Returns the flipped discs.
func (board *Board) doMove5() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000000000C0)
	flipped |= board.doMoveToHigherBits(0x0000010204081000)
	flipped |= board.doMoveToHigherBits(0x2020202020202000)
	flipped |= board.doMoveToHigherBits(0x0000000000804000)
	flipped |= board.doMoveToLowerBits(0x000000000000001F)
	return
}

//Does the move at field 6.
//Returns the flipped discs.
func (board *Board) doMove6() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0001020408102000)
	flipped |= board.doMoveToHigherBits(0x4040404040404000)
	flipped |= board.doMoveToLowerBits(0x000000000000003F)
	return
}

//Does the move at field 7.
//Returns the flipped discs.
func (board *Board) doMove7() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0102040810204000)
	flipped |= board.doMoveToHigherBits(0x8080808080808000)
	flipped |= board.doMoveToLowerBits(0x000000000000007F)
	return
}

//Does the move at field 8.
//Returns the flipped discs.
func (board *Board) doMove8() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000000000FE00)
	flipped |= board.doMoveToHigherBits(0x0101010101010000)
	flipped |= board.doMoveToHigherBits(0x4020100804020000)
	return
}

//Does the move at field 9.
//Returns the flipped discs.
func (board *Board) doMove9() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000000000FC00)
	flipped |= board.doMoveToHigherBits(0x0202020202020000)
	flipped |= board.doMoveToHigherBits(0x8040201008040000)
	return
}

//Does the move at field 10.
//Returns the flipped discs.
func (board *Board) doMove10() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000000000F800)
	flipped |= board.doMoveToHigherBits(0x0000000001020000)
	flipped |= board.doMoveToHigherBits(0x0404040404040000)
	flipped |= board.doMoveToHigherBits(0x0080402010080000)
	flipped |= board.doMoveToLowerBits(0x0000000000000300)
	return
}

//Does the move at field 11.
//Returns the flipped discs.
func (board *Board) doMove11() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000000000F000)
	flipped |= board.doMoveToHigherBits(0x0000000102040000)
	flipped |= board.doMoveToHigherBits(0x0808080808080000)
	flipped |= board.doMoveToHigherBits(0x0000804020100000)
	flipped |= board.doMoveToLowerBits(0x0000000000000700)
	return
}

//Does the move at field 12.
//Returns the flipped discs.
func (board *Board) doMove12() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000000000E000)
	flipped |= board.doMoveToHigherBits(0x0000010204080000)
	flipped |= board.doMoveToHigherBits(0x1010101010100000)
	flipped |= board.doMoveToHigherBits(0x0000008040200000)
	flipped |= board.doMoveToLowerBits(0x0000000000000F00)
	return
}

//Does the move at field 13.
//Returns the flipped discs.
func (board *Board) doMove13() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000000000C000)
	flipped |= board.doMoveToHigherBits(0x0001020408100000)
	flipped |= board.doMoveToHigherBits(0x2020202020200000)
	flipped |= board.doMoveToHigherBits(0x0000000080400000)
	flipped |= board.doMoveToLowerBits(0x0000000000001F00)
	return
}

//Does the move at field 14.
//Returns the flipped discs.
func (board *Board) doMove14() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0102040810200000)
	flipped |= board.doMoveToHigherBits(0x4040404040400000)
	flipped |= board.doMoveToLowerBits(0x0000000000003F00)
	return
}

//Does the move at field 15.
//Returns the flipped discs.
func (board *Board) doMove15() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0204081020400000)
	flipped |= board.doMoveToHigherBits(0x8080808080800000)
	flipped |= board.doMoveToLowerBits(0x0000000000007F00)
	return
}

//Does the move at field 16.
//Returns the flipped discs.
func (board *Board) doMove16() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000000000FE0000)
	flipped |= board.doMoveToHigherBits(0x0101010101000000)
	flipped |= board.doMoveToHigherBits(0x2010080402000000)
	flipped |= board.doMoveToLowerBits(0x0000000000000204)
	flipped |= board.doMoveToLowerBits(0x0000000000000101)
	return
}

//Does the move at field 17.
//Returns the flipped discs.
func (board *Board) doMove17() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000000000FC0000)
	flipped |= board.doMoveToHigherBits(0x0202020202000000)
	flipped |= board.doMoveToHigherBits(0x4020100804000000)
	flipped |= board.doMoveToLowerBits(0x0000000000000408)
	flipped |= board.doMoveToLowerBits(0x0000000000000202)
	return
}

//Does the move at field 18.
//Returns the flipped discs.
func (board *Board) doMove18() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000000000F80000)
	flipped |= board.doMoveToHigherBits(0x0000000102000000)
	flipped |= board.doMoveToHigherBits(0x0404040404000000)
	flipped |= board.doMoveToHigherBits(0x8040201008000000)
	flipped |= board.doMoveToLowerBits(0x0000000000030000)
	flipped |= board.doMoveToLowerBits(0x0000000000000810)
	flipped |= board.doMoveToLowerBits(0x0000000000000404)
	flipped |= board.doMoveToLowerBits(0x0000000000000201)
	return
}

//Does the move at field 19.
//Returns the flipped discs.
func (board *Board) doMove19() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000000000F00000)
	flipped |= board.doMoveToHigherBits(0x0000010204000000)
	flipped |= board.doMoveToHigherBits(0x0808080808000000)
	flipped |= board.doMoveToHigherBits(0x0080402010000000)
	flipped |= board.doMoveToLowerBits(0x0000000000070000)
	flipped |= board.doMoveToLowerBits(0x0000000000001020)
	flipped |= board.doMoveToLowerBits(0x0000000000000808)
	flipped |= board.doMoveToLowerBits(0x0000000000000402)
	return
}

//Does the move at field 20.
//Returns the flipped discs.
func (board *Board) doMove20() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000000000E00000)
	flipped |= board.doMoveToHigherBits(0x0001020408000000)
	flipped |= board.doMoveToHigherBits(0x1010101010000000)
	flipped |= board.doMoveToHigherBits(0x0000804020000000)
	flipped |= board.doMoveToLowerBits(0x00000000000F0000)
	flipped |= board.doMoveToLowerBits(0x0000000000002040)
	flipped |= board.doMoveToLowerBits(0x0000000000001010)
	flipped |= board.doMoveToLowerBits(0x0000000000000804)
	return
}

//Does the move at field 21.
//Returns the flipped discs.
func (board *Board) doMove21() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000000000C00000)
	flipped |= board.doMoveToHigherBits(0x0102040810000000)
	flipped |= board.doMoveToHigherBits(0x2020202020000000)
	flipped |= board.doMoveToHigherBits(0x0000008040000000)
	flipped |= board.doMoveToLowerBits(0x00000000001F0000)
	flipped |= board.doMoveToLowerBits(0x0000000000004080)
	flipped |= board.doMoveToLowerBits(0x0000000000002020)
	flipped |= board.doMoveToLowerBits(0x0000000000001008)
	return
}

//Does the move at field 22.
//Returns the flipped discs.
func (board *Board) doMove22() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0204081020000000)
	flipped |= board.doMoveToHigherBits(0x4040404040000000)
	flipped |= board.doMoveToLowerBits(0x00000000003F0000)
	flipped |= board.doMoveToLowerBits(0x0000000000004040)
	flipped |= board.doMoveToLowerBits(0x0000000000002010)
	return
}

//Does the move at field 23.
//Returns the flipped discs.
func (board *Board) doMove23() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0408102040000000)
	flipped |= board.doMoveToHigherBits(0x8080808080000000)
	flipped |= board.doMoveToLowerBits(0x00000000007F0000)
	flipped |= board.doMoveToLowerBits(0x0000000000008080)
	flipped |= board.doMoveToLowerBits(0x0000000000004020)
	return
}

//Does the move at field 24.
//Returns the flipped discs.
func (board *Board) doMove24() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000FE000000)
	flipped |= board.doMoveToHigherBits(0x0101010100000000)
	flipped |= board.doMoveToHigherBits(0x1008040200000000)
	flipped |= board.doMoveToLowerBits(0x0000000000020408)
	flipped |= board.doMoveToLowerBits(0x0000000000010101)
	return
}

//Does the move at field 25.
//Returns the flipped discs.
func (board *Board) doMove25() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000FC000000)
	flipped |= board.doMoveToHigherBits(0x0202020200000000)
	flipped |= board.doMoveToHigherBits(0x2010080400000000)
	flipped |= board.doMoveToLowerBits(0x0000000000040810)
	flipped |= board.doMoveToLowerBits(0x0000000000020202)
	return
}

//Does the move at field 26.
//Returns the flipped discs.
func (board *Board) doMove26() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000F8000000)
	flipped |= board.doMoveToHigherBits(0x0000010200000000)
	flipped |= board.doMoveToHigherBits(0x0404040400000000)
	flipped |= board.doMoveToHigherBits(0x4020100800000000)
	flipped |= board.doMoveToLowerBits(0x0000000003000000)
	flipped |= board.doMoveToLowerBits(0x0000000000081020)
	flipped |= board.doMoveToLowerBits(0x0000000000040404)
	flipped |= board.doMoveToLowerBits(0x0000000000020100)
	return
}

//Does the move at field 27.
//Returns the flipped discs.
func (board *Board) doMove27() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000F0000000)
	flipped |= board.doMoveToHigherBits(0x0001020400000000)
	flipped |= board.doMoveToHigherBits(0x0808080800000000)
	flipped |= board.doMoveToHigherBits(0x8040201000000000)
	flipped |= board.doMoveToLowerBits(0x0000000007000000)
	flipped |= board.doMoveToLowerBits(0x0000000000102040)
	flipped |= board.doMoveToLowerBits(0x0000000000080808)
	flipped |= board.doMoveToLowerBits(0x0000000000040201)
	return
}

//Does the move at field 28.
//Returns the flipped discs.
func (board *Board) doMove28() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000E0000000)
	flipped |= board.doMoveToHigherBits(0x0102040800000000)
	flipped |= board.doMoveToHigherBits(0x1010101000000000)
	flipped |= board.doMoveToHigherBits(0x0080402000000000)
	flipped |= board.doMoveToLowerBits(0x000000000F000000)
	flipped |= board.doMoveToLowerBits(0x0000000000204080)
	flipped |= board.doMoveToLowerBits(0x0000000000101010)
	flipped |= board.doMoveToLowerBits(0x0000000000080402)
	return
}

//Does the move at field 29.
//Returns the flipped discs.
func (board *Board) doMove29() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00000000C0000000)
	flipped |= board.doMoveToHigherBits(0x0204081000000000)
	flipped |= board.doMoveToHigherBits(0x2020202000000000)
	flipped |= board.doMoveToHigherBits(0x0000804000000000)
	flipped |= board.doMoveToLowerBits(0x000000001F000000)
	flipped |= board.doMoveToLowerBits(0x0000000000408000)
	flipped |= board.doMoveToLowerBits(0x0000000000202020)
	flipped |= board.doMoveToLowerBits(0x0000000000100804)
	return
}

//Does the move at field 30.
//Returns the flipped discs.
func (board *Board) doMove30() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0408102000000000)
	flipped |= board.doMoveToHigherBits(0x4040404000000000)
	flipped |= board.doMoveToLowerBits(0x000000003F000000)
	flipped |= board.doMoveToLowerBits(0x0000000000404040)
	flipped |= board.doMoveToLowerBits(0x0000000000201008)
	return
}

//Does the move at field 31.
//Returns the flipped discs.
func (board *Board) doMove31() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0810204000000000)
	flipped |= board.doMoveToHigherBits(0x8080808000000000)
	flipped |= board.doMoveToLowerBits(0x000000007F000000)
	flipped |= board.doMoveToLowerBits(0x0000000000808080)
	flipped |= board.doMoveToLowerBits(0x0000000000402010)
	return
}

//Does the move at field 32.
//Returns the flipped discs.
func (board *Board) doMove32() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000FE00000000)
	flipped |= board.doMoveToHigherBits(0x0101010000000000)
	flipped |= board.doMoveToHigherBits(0x0804020000000000)
	flipped |= board.doMoveToLowerBits(0x0000000002040810)
	flipped |= board.doMoveToLowerBits(0x0000000001010101)
	return
}

//Does the move at field 33.
//Returns the flipped discs.
func (board *Board) doMove33() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000FC00000000)
	flipped |= board.doMoveToHigherBits(0x0202020000000000)
	flipped |= board.doMoveToHigherBits(0x1008040000000000)
	flipped |= board.doMoveToLowerBits(0x0000000004081020)
	flipped |= board.doMoveToLowerBits(0x0000000002020202)
	return
}

//Does the move at field 34.
//Returns the flipped discs.
func (board *Board) doMove34() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000F800000000)
	flipped |= board.doMoveToHigherBits(0x0001020000000000)
	flipped |= board.doMoveToHigherBits(0x0404040000000000)
	flipped |= board.doMoveToHigherBits(0x2010080000000000)
	flipped |= board.doMoveToLowerBits(0x0000000300000000)
	flipped |= board.doMoveToLowerBits(0x0000000008102040)
	flipped |= board.doMoveToLowerBits(0x0000000004040404)
	flipped |= board.doMoveToLowerBits(0x0000000002010000)
	return
}

//Does the move at field 35.
//Returns the flipped discs.
func (board *Board) doMove35() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000F000000000)
	flipped |= board.doMoveToHigherBits(0x0102040000000000)
	flipped |= board.doMoveToHigherBits(0x0808080000000000)
	flipped |= board.doMoveToHigherBits(0x4020100000000000)
	flipped |= board.doMoveToLowerBits(0x0000000700000000)
	flipped |= board.doMoveToLowerBits(0x0000000010204080)
	flipped |= board.doMoveToLowerBits(0x0000000008080808)
	flipped |= board.doMoveToLowerBits(0x0000000004020100)
	return
}

//Does the move at field 36.
//Returns the flipped discs.
func (board *Board) doMove36() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000E000000000)
	flipped |= board.doMoveToHigherBits(0x0204080000000000)
	flipped |= board.doMoveToHigherBits(0x1010100000000000)
	flipped |= board.doMoveToHigherBits(0x8040200000000000)
	flipped |= board.doMoveToLowerBits(0x0000000F00000000)
	flipped |= board.doMoveToLowerBits(0x0000000020408000)
	flipped |= board.doMoveToLowerBits(0x0000000010101010)
	flipped |= board.doMoveToLowerBits(0x0000000008040201)
	return
}

//Does the move at field 37.
//Returns the flipped discs.
func (board *Board) doMove37() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x000000C000000000)
	flipped |= board.doMoveToHigherBits(0x0408100000000000)
	flipped |= board.doMoveToHigherBits(0x2020200000000000)
	flipped |= board.doMoveToHigherBits(0x0080400000000000)
	flipped |= board.doMoveToLowerBits(0x0000001F00000000)
	flipped |= board.doMoveToLowerBits(0x0000000040800000)
	flipped |= board.doMoveToLowerBits(0x0000000020202020)
	flipped |= board.doMoveToLowerBits(0x0000000010080402)
	return
}

//Does the move at field 38.
//Returns the flipped discs.
func (board *Board) doMove38() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0810200000000000)
	flipped |= board.doMoveToHigherBits(0x4040400000000000)
	flipped |= board.doMoveToLowerBits(0x0000003F00000000)
	flipped |= board.doMoveToLowerBits(0x0000000040404040)
	flipped |= board.doMoveToLowerBits(0x0000000020100804)
	return
}

//Does the move at field 39.
//Returns the flipped discs.
func (board *Board) doMove39() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x1020400000000000)
	flipped |= board.doMoveToHigherBits(0x8080800000000000)
	flipped |= board.doMoveToLowerBits(0x0000007F00000000)
	flipped |= board.doMoveToLowerBits(0x0000000080808080)
	flipped |= board.doMoveToLowerBits(0x0000000040201008)
	return
}

//Does the move at field 40.
//Returns the flipped discs.
func (board *Board) doMove40() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000FE0000000000)
	flipped |= board.doMoveToHigherBits(0x0101000000000000)
	flipped |= board.doMoveToHigherBits(0x0402000000000000)
	flipped |= board.doMoveToLowerBits(0x0000000204081020)
	flipped |= board.doMoveToLowerBits(0x0000000101010101)
	return
}

//Does the move at field 41.
//Returns the flipped discs.
func (board *Board) doMove41() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000FC0000000000)
	flipped |= board.doMoveToHigherBits(0x0202000000000000)
	flipped |= board.doMoveToHigherBits(0x0804000000000000)
	flipped |= board.doMoveToLowerBits(0x0000000408102040)
	flipped |= board.doMoveToLowerBits(0x0000000202020202)
	return
}

//Does the move at field 42.
//Returns the flipped discs.
func (board *Board) doMove42() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000F80000000000)
	flipped |= board.doMoveToHigherBits(0x0102000000000000)
	flipped |= board.doMoveToHigherBits(0x0404000000000000)
	flipped |= board.doMoveToHigherBits(0x1008000000000000)
	flipped |= board.doMoveToLowerBits(0x0000030000000000)
	flipped |= board.doMoveToLowerBits(0x0000000810204080)
	flipped |= board.doMoveToLowerBits(0x0000000404040404)
	flipped |= board.doMoveToLowerBits(0x0000000201000000)
	return
}

//Does the move at field 43.
//Returns the flipped discs.
func (board *Board) doMove43() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000F00000000000)
	flipped |= board.doMoveToHigherBits(0x0204000000000000)
	flipped |= board.doMoveToHigherBits(0x0808000000000000)
	flipped |= board.doMoveToHigherBits(0x2010000000000000)
	flipped |= board.doMoveToLowerBits(0x0000070000000000)
	flipped |= board.doMoveToLowerBits(0x0000001020408000)
	flipped |= board.doMoveToLowerBits(0x0000000808080808)
	flipped |= board.doMoveToLowerBits(0x0000000402010000)
	return
}

//Does the move at field 44.
//Returns the flipped discs.
func (board *Board) doMove44() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000E00000000000)
	flipped |= board.doMoveToHigherBits(0x0408000000000000)
	flipped |= board.doMoveToHigherBits(0x1010000000000000)
	flipped |= board.doMoveToHigherBits(0x4020000000000000)
	flipped |= board.doMoveToLowerBits(0x00000F0000000000)
	flipped |= board.doMoveToLowerBits(0x0000002040800000)
	flipped |= board.doMoveToLowerBits(0x0000001010101010)
	flipped |= board.doMoveToLowerBits(0x0000000804020100)
	return
}

//Does the move at field 45.
//Returns the flipped discs.
func (board *Board) doMove45() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x0000C00000000000)
	flipped |= board.doMoveToHigherBits(0x0810000000000000)
	flipped |= board.doMoveToHigherBits(0x2020000000000000)
	flipped |= board.doMoveToHigherBits(0x8040000000000000)
	flipped |= board.doMoveToLowerBits(0x00001F0000000000)
	flipped |= board.doMoveToLowerBits(0x0000004080000000)
	flipped |= board.doMoveToLowerBits(0x0000002020202020)
	flipped |= board.doMoveToLowerBits(0x0000001008040201)
	return
}

//Does the move at field 46.
//Returns the flipped discs.
func (board *Board) doMove46() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x1020000000000000)
	flipped |= board.doMoveToHigherBits(0x4040000000000000)
	flipped |= board.doMoveToLowerBits(0x00003F0000000000)
	flipped |= board.doMoveToLowerBits(0x0000004040404040)
	flipped |= board.doMoveToLowerBits(0x0000002010080402)
	return
}

//Does the move at field 47.
//Returns the flipped discs.
func (board *Board) doMove47() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x2040000000000000)
	flipped |= board.doMoveToHigherBits(0x8080000000000000)
	flipped |= board.doMoveToLowerBits(0x00007F0000000000)
	flipped |= board.doMoveToLowerBits(0x0000008080808080)
	flipped |= board.doMoveToLowerBits(0x0000004020100804)
	return
}

//Does the move at field 48.
//Returns the flipped discs.
func (board *Board) doMove48() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00FE000000000000)
	flipped |= board.doMoveToLowerBits(0x0000020408102040)
	flipped |= board.doMoveToLowerBits(0x0000010101010101)
	return
}

//Does the move at field 49.
//Returns the flipped discs.
func (board *Board) doMove49() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00FC000000000000)
	flipped |= board.doMoveToLowerBits(0x0000040810204080)
	flipped |= board.doMoveToLowerBits(0x0000020202020202)
	return
}

//Does the move at field 50.
//Returns the flipped discs.
func (board *Board) doMove50() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00F8000000000000)
	flipped |= board.doMoveToLowerBits(0x0003000000000000)
	flipped |= board.doMoveToLowerBits(0x0000081020408000)
	flipped |= board.doMoveToLowerBits(0x0000040404040404)
	flipped |= board.doMoveToLowerBits(0x0000020100000000)
	return
}

//Does the move at field 51.
//Returns the flipped discs.
func (board *Board) doMove51() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00F0000000000000)
	flipped |= board.doMoveToLowerBits(0x0007000000000000)
	flipped |= board.doMoveToLowerBits(0x0000102040800000)
	flipped |= board.doMoveToLowerBits(0x0000080808080808)
	flipped |= board.doMoveToLowerBits(0x0000040201000000)
	return
}

//Does the move at field 52.
//Returns the flipped discs.
func (board *Board) doMove52() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00E0000000000000)
	flipped |= board.doMoveToLowerBits(0x000F000000000000)
	flipped |= board.doMoveToLowerBits(0x0000204080000000)
	flipped |= board.doMoveToLowerBits(0x0000101010101010)
	flipped |= board.doMoveToLowerBits(0x0000080402010000)
	return
}

//Does the move at field 53.
//Returns the flipped discs.
func (board *Board) doMove53() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0x00C0000000000000)
	flipped |= board.doMoveToLowerBits(0x001F000000000000)
	flipped |= board.doMoveToLowerBits(0x0000408000000000)
	flipped |= board.doMoveToLowerBits(0x0000202020202020)
	flipped |= board.doMoveToLowerBits(0x0000100804020100)
	return
}

//Does the move at field 54.
//Returns the flipped discs.
func (board *Board) doMove54() (flipped uint64) {
	flipped = board.doMoveToLowerBits(0x003F000000000000)
	flipped |= board.doMoveToLowerBits(0x0000404040404040)
	flipped |= board.doMoveToLowerBits(0x0000201008040201)
	return
}

//Does the move at field 55.
//Returns the flipped discs.
func (board *Board) doMove55() (flipped uint64) {
	flipped = board.doMoveToLowerBits(0x007F000000000000)
	flipped |= board.doMoveToLowerBits(0x0000808080808080)
	flipped |= board.doMoveToLowerBits(0x0000402010080402)
	return
}

//Does the move at field 56.
//Returns the flipped discs.
func (board *Board) doMove56() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0xFE00000000000000)
	flipped |= board.doMoveToLowerBits(0x0002040810204080)
	flipped |= board.doMoveToLowerBits(0x0001010101010101)
	return
}

//Does the move at field 57.
//Returns the flipped discs.
func (board *Board) doMove57() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0xFC00000000000000)
	flipped |= board.doMoveToLowerBits(0x0004081020408000)
	flipped |= board.doMoveToLowerBits(0x0002020202020202)
	return
}

//Does the move at field 58.
//Returns the flipped discs.
func (board *Board) doMove58() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0xF800000000000000)
	flipped |= board.doMoveToLowerBits(0x0300000000000000)
	flipped |= board.doMoveToLowerBits(0x0008102040800000)
	flipped |= board.doMoveToLowerBits(0x0004040404040404)
	flipped |= board.doMoveToLowerBits(0x0002010000000000)
	return
}

//Does the move at field 59.
//Returns the flipped discs.
func (board *Board) doMove59() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0xF000000000000000)
	flipped |= board.doMoveToLowerBits(0x0700000000000000)
	flipped |= board.doMoveToLowerBits(0x0010204080000000)
	flipped |= board.doMoveToLowerBits(0x0008080808080808)
	flipped |= board.doMoveToLowerBits(0x0004020100000000)
	return
}

//Does the move at field 60.
//Returns the flipped discs.
func (board *Board) doMove60() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0xE000000000000000)
	flipped |= board.doMoveToLowerBits(0x0F00000000000000)
	flipped |= board.doMoveToLowerBits(0x0020408000000000)
	flipped |= board.doMoveToLowerBits(0x0010101010101010)
	flipped |= board.doMoveToLowerBits(0x0008040201000000)
	return
}

//Does the move at field 61.
//Returns the flipped discs.
func (board *Board) doMove61() (flipped uint64) {
	flipped = board.doMoveToHigherBits(0xC000000000000000)
	flipped |= board.doMoveToLowerBits(0x1F00000000000000)
	flipped |= board.doMoveToLowerBits(0x0040800000000000)
	flipped |= board.doMoveToLowerBits(0x0020202020202020)
	flipped |= board.doMoveToLowerBits(0x0010080402010000)
	return
}

//Does the move at field 62.
//Returns the flipped discs.
func (board *Board) doMove62() (flipped uint64) {
	flipped = board.doMoveToLowerBits(0x3F00000000000000)
	flipped |= board.doMoveToLowerBits(0x0040404040404040)
	flipped |= board.doMoveToLowerBits(0x0020100804020100)
	return
}

//Does the move at field 63.
//Returns the flipped discs.
func (board *Board) doMove63() (flipped uint64) {
	flipped = board.doMoveToLowerBits(0x7F00000000000000)
	flipped |= board.doMoveToLowerBits(0x0080808080808080)
	flipped |= board.doMoveToLowerBits(0x0040201008040201)
	return
}
