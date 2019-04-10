package othello

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/bits"
	"math/rand"
)

const (
	cornerMask  = uint64(1<<0 | 1<<7 | 1<<56 | 1<<63)
	xSquareMask = uint64(1<<9 | 1<<14 | 1<<49 | 1<<54)
	cSquareMask = uint64(1<<1 | 1<<6 | 1<<8 | 1<<15 | 1<<48 | 1<<55 | 1<<57 | 1<<62)
)

func bitsetASCIIArtString(bs uint64) string {
	buffer := new(bytes.Buffer)
	_, _ = buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	for y := uint(0); y < 8; y++ {
		_, _ = buffer.WriteString(fmt.Sprintf("%d ", y+1))

		for x := uint(0); x < 8; x++ {
			f := y*8 + x
			if bs&uint64(1<<f) != 0 {
				_, _ = buffer.WriteString("@ ")
			} else {
				_, _ = buffer.WriteString("- ")
			}

		}
		_, _ = buffer.WriteString("|\n")
	}
	_, _ = buffer.WriteString("+-----------------+\n")

	return buffer.String()
}

// Board represents the state of an othello othello game.
// It does not keep track which discs are white or black.
// Instead it keeps track which discs are owned by the player to move.
type Board struct {
	me, opp uint64
}

// SortableBoard is a board with associated heuristic estimation suitable for sorting
type SortableBoard struct {
	Board Board
	Heur  int
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
func RandomBoard(discs int) (*Board, error) {

	if discs < 4 || discs > 64 {
		err := fmt.Errorf("cannot create random board with %d discs", discs)
		return nil, err
	}

	board := NewBoard()
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

	return board, nil
}

func (board Board) rotate(rotation int) Board {

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

	rotated := board
	rotated.me = rotate(rotated.me, rotation)
	rotated.opp = rotate(rotated.opp, rotation)
	return rotated

}

// Normalize returns a normalized othello with regards to symmetry
func (board Board) Normalize() Board {

	less := func(lhs, rhs Board) bool {
		if lhs.me < rhs.me {
			return true
		}
		if lhs.me == rhs.me && lhs.opp < rhs.opp {
			return true
		}
		return false
	}

	normalized := board.rotate(0)

	for r := 1; r < 8; r++ {
		cur := board.rotate(r)
		if less(cur, normalized) {
			normalized = cur
		}
	}

	return normalized
}

// ASCIIArt writes ascii-art of a Board to a writer
func (board Board) ASCIIArt(writer io.Writer, swapDiscColors bool) {

	buffer := new(bytes.Buffer)
	_, _ = buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	moves := board.Moves()

	if swapDiscColors {
		board.SwitchTurn()
	}

	for y := uint(0); y < 8; y++ {
		_, _ = buffer.WriteString(fmt.Sprintf("%d ", y+1))

		for x := uint(0); x < 8; x++ {
			mask := uint64(1) << (y*8 + x)
			if board.me&mask != 0 {
				_, _ = buffer.WriteString("○ ")
			} else if board.opp&mask != 0 {
				_, _ = buffer.WriteString("● ")
			} else if moves&mask != 0 {
				_, _ = buffer.WriteString("- ")
			} else {
				_, _ = buffer.WriteString("  ")
			}
		}

		_, _ = buffer.WriteString("|\n")
	}

	var moveDisc string
	if swapDiscColors {
		moveDisc = "●"
	} else {
		moveDisc = "○"
	}

	_, _ = buffer.WriteString("+-----------------+\n")
	_, _ = buffer.WriteString("To move: " + moveDisc + "\n")
	_, _ = buffer.WriteString("Raw: " + fmt.Sprintf("%#v", board) + "\n")

	_, err := writer.Write(buffer.Bytes())
	if err != nil {
		log.Printf("board.ASCIIArt() error: %s", err)
	}
}

// Moves returns a bitset of valid moves for a Board
func (board Board) Moves() uint64 {
	return moves(board.me, board.opp)
}

// OpponentMoves returns a bitset with all valid moves for the opponent
func (board Board) OpponentMoves() uint64 {
	return moves(board.opp, board.me)
}

func moves(me, opp uint64) uint64 {

	// this function is a modified version of code from Edax
	mask := opp & 0x7E7E7E7E7E7E7E7E

	flipL := mask & (me << 1)
	flipL |= mask & (flipL << 1)
	maskL := mask & (mask << 1)
	flipL |= maskL & (flipL << (2 * 1))
	flipL |= maskL & (flipL << (2 * 1))
	flipR := mask & (me >> 1)
	flipR |= mask & (flipR >> 1)
	maskR := mask & (mask >> 1)
	flipR |= maskR & (flipR >> (2 * 1))
	flipR |= maskR & (flipR >> (2 * 1))
	movesSet := (flipL << 1) | (flipR >> 1)

	flipL = mask & (me << 7)
	flipL |= mask & (flipL << 7)
	maskL = mask & (mask << 7)
	flipL |= maskL & (flipL << (2 * 7))
	flipL |= maskL & (flipL << (2 * 7))
	flipR = mask & (me >> 7)
	flipR |= mask & (flipR >> 7)
	maskR = mask & (mask >> 7)
	flipR |= maskR & (flipR >> (2 * 7))
	flipR |= maskR & (flipR >> (2 * 7))
	movesSet |= (flipL << 7) | (flipR >> 7)

	flipL = mask & (me << 9)
	flipL |= mask & (flipL << 9)
	maskL = mask & (mask << 9)
	flipL |= maskL & (flipL << (2 * 9))
	flipL |= maskL & (flipL << (2 * 9))
	flipR = mask & (me >> 9)
	flipR |= mask & (flipR >> 9)
	maskR = mask & (mask >> 9)
	flipR |= maskR & (flipR >> (2 * 9))
	flipR |= maskR & (flipR >> (2 * 9))
	movesSet |= (flipL << 9) | (flipR >> 9)

	flipL = opp & (me << 8)
	flipL |= opp & (flipL << 8)
	maskL = opp & (opp << 8)
	flipL |= maskL & (flipL << (2 * 8))
	flipL |= maskL & (flipL << (2 * 8))
	flipR = opp & (me >> 8)
	flipR |= opp & (flipR >> 8)
	maskR = opp & (opp >> 8)
	flipR |= maskR & (flipR >> (2 * 8))
	flipR |= maskR & (flipR >> (2 * 8))
	movesSet |= (flipL << 8) | (flipR >> 8)

	movesSet &^= me | opp
	return movesSet
}

// DoMove does a move and returns the flipped discs
func (board *Board) DoMove(moveBit uint64) uint64 {

	var flipped uint64

	switch moveBit {
	case uint64(1 << 0):
		flipped = board.doMoveToHigherBits(0x00000000000000FE)
		flipped |= board.doMoveToHigherBits(0x0101010101010100)
		flipped |= board.doMoveToHigherBits(0x8040201008040200)
	case uint64(1 << 1):
		flipped = board.doMoveToHigherBits(0x00000000000000FC)
		flipped |= board.doMoveToHigherBits(0x0202020202020200)
		flipped |= board.doMoveToHigherBits(0x0080402010080400)
	case uint64(1 << 2):
		flipped = board.doMoveToHigherBits(0x00000000000000F8)
		flipped |= board.doMoveToHigherBits(0x0000000000010200)
		flipped |= board.doMoveToHigherBits(0x0404040404040400)
		flipped |= board.doMoveToHigherBits(0x0000804020100800)
		flipped |= board.doMoveToLowerBits(0x0000000000000003)
	case uint64(1 << 3):
		flipped = board.doMoveToHigherBits(0x00000000000000F0)
		flipped |= board.doMoveToHigherBits(0x0000000001020400)
		flipped |= board.doMoveToHigherBits(0x0808080808080800)
		flipped |= board.doMoveToHigherBits(0x0000008040201000)
		flipped |= board.doMoveToLowerBits(0x0000000000000007)
	case uint64(1 << 4):
		flipped = board.doMoveToHigherBits(0x00000000000000E0)
		flipped |= board.doMoveToHigherBits(0x0000000102040800)
		flipped |= board.doMoveToHigherBits(0x1010101010101000)
		flipped |= board.doMoveToHigherBits(0x0000000080402000)
		flipped |= board.doMoveToLowerBits(0x000000000000000F)
	case uint64(1 << 5):
		flipped = board.doMoveToHigherBits(0x00000000000000C0)
		flipped |= board.doMoveToHigherBits(0x0000010204081000)
		flipped |= board.doMoveToHigherBits(0x2020202020202000)
		flipped |= board.doMoveToHigherBits(0x0000000000804000)
		flipped |= board.doMoveToLowerBits(0x000000000000001F)
	case uint64(1 << 6):
		flipped = board.doMoveToHigherBits(0x0001020408102000)
		flipped |= board.doMoveToHigherBits(0x4040404040404000)
		flipped |= board.doMoveToLowerBits(0x000000000000003F)
	case uint64(1 << 7):
		flipped = board.doMoveToHigherBits(0x0102040810204000)
		flipped |= board.doMoveToHigherBits(0x8080808080808000)
		flipped |= board.doMoveToLowerBits(0x000000000000007F)
	case uint64(1 << 8):
		flipped = board.doMoveToHigherBits(0x000000000000FE00)
		flipped |= board.doMoveToHigherBits(0x0101010101010000)
		flipped |= board.doMoveToHigherBits(0x4020100804020000)
	case uint64(1 << 9):
		flipped = board.doMoveToHigherBits(0x000000000000FC00)
		flipped |= board.doMoveToHigherBits(0x0202020202020000)
		flipped |= board.doMoveToHigherBits(0x8040201008040000)
	case uint64(1 << 10):
		flipped = board.doMoveToHigherBits(0x000000000000F800)
		flipped |= board.doMoveToHigherBits(0x0000000001020000)
		flipped |= board.doMoveToHigherBits(0x0404040404040000)
		flipped |= board.doMoveToHigherBits(0x0080402010080000)
		flipped |= board.doMoveToLowerBits(0x0000000000000300)
	case uint64(1 << 11):
		flipped = board.doMoveToHigherBits(0x000000000000F000)
		flipped |= board.doMoveToHigherBits(0x0000000102040000)
		flipped |= board.doMoveToHigherBits(0x0808080808080000)
		flipped |= board.doMoveToHigherBits(0x0000804020100000)
		flipped |= board.doMoveToLowerBits(0x0000000000000700)
	case uint64(1 << 12):
		flipped = board.doMoveToHigherBits(0x000000000000E000)
		flipped |= board.doMoveToHigherBits(0x0000010204080000)
		flipped |= board.doMoveToHigherBits(0x1010101010100000)
		flipped |= board.doMoveToHigherBits(0x0000008040200000)
		flipped |= board.doMoveToLowerBits(0x0000000000000F00)
	case uint64(1 << 13):
		flipped = board.doMoveToHigherBits(0x000000000000C000)
		flipped |= board.doMoveToHigherBits(0x0001020408100000)
		flipped |= board.doMoveToHigherBits(0x2020202020200000)
		flipped |= board.doMoveToHigherBits(0x0000000080400000)
		flipped |= board.doMoveToLowerBits(0x0000000000001F00)
	case uint64(1 << 14):
		flipped = board.doMoveToHigherBits(0x0102040810200000)
		flipped |= board.doMoveToHigherBits(0x4040404040400000)
		flipped |= board.doMoveToLowerBits(0x0000000000003F00)
	case uint64(1 << 15):
		flipped = board.doMoveToHigherBits(0x0204081020400000)
		flipped |= board.doMoveToHigherBits(0x8080808080800000)
		flipped |= board.doMoveToLowerBits(0x0000000000007F00)
	case uint64(1 << 16):
		flipped = board.doMoveToHigherBits(0x0000000000FE0000)
		flipped |= board.doMoveToHigherBits(0x0101010101000000)
		flipped |= board.doMoveToHigherBits(0x2010080402000000)
		flipped |= board.doMoveToLowerBits(0x0000000000000204)
		flipped |= board.doMoveToLowerBits(0x0000000000000101)
	case uint64(1 << 17):
		flipped = board.doMoveToHigherBits(0x0000000000FC0000)
		flipped |= board.doMoveToHigherBits(0x0202020202000000)
		flipped |= board.doMoveToHigherBits(0x4020100804000000)
		flipped |= board.doMoveToLowerBits(0x0000000000000408)
		flipped |= board.doMoveToLowerBits(0x0000000000000202)
	case uint64(1 << 18):
		flipped = board.doMoveToHigherBits(0x0000000000F80000)
		flipped |= board.doMoveToHigherBits(0x0000000102000000)
		flipped |= board.doMoveToHigherBits(0x0404040404000000)
		flipped |= board.doMoveToHigherBits(0x8040201008000000)
		flipped |= board.doMoveToLowerBits(0x0000000000030000)
		flipped |= board.doMoveToLowerBits(0x0000000000000810)
		flipped |= board.doMoveToLowerBits(0x0000000000000404)
		flipped |= board.doMoveToLowerBits(0x0000000000000201)
	case uint64(1 << 19):
		flipped = board.doMoveToHigherBits(0x0000000000F00000)
		flipped |= board.doMoveToHigherBits(0x0000010204000000)
		flipped |= board.doMoveToHigherBits(0x0808080808000000)
		flipped |= board.doMoveToHigherBits(0x0080402010000000)
		flipped |= board.doMoveToLowerBits(0x0000000000070000)
		flipped |= board.doMoveToLowerBits(0x0000000000001020)
		flipped |= board.doMoveToLowerBits(0x0000000000000808)
		flipped |= board.doMoveToLowerBits(0x0000000000000402)
	case uint64(1 << 20):
		flipped = board.doMoveToHigherBits(0x0000000000E00000)
		flipped |= board.doMoveToHigherBits(0x0001020408000000)
		flipped |= board.doMoveToHigherBits(0x1010101010000000)
		flipped |= board.doMoveToHigherBits(0x0000804020000000)
		flipped |= board.doMoveToLowerBits(0x00000000000F0000)
		flipped |= board.doMoveToLowerBits(0x0000000000002040)
		flipped |= board.doMoveToLowerBits(0x0000000000001010)
		flipped |= board.doMoveToLowerBits(0x0000000000000804)
	case uint64(1 << 21):
		flipped = board.doMoveToHigherBits(0x0000000000C00000)
		flipped |= board.doMoveToHigherBits(0x0102040810000000)
		flipped |= board.doMoveToHigherBits(0x2020202020000000)
		flipped |= board.doMoveToHigherBits(0x0000008040000000)
		flipped |= board.doMoveToLowerBits(0x00000000001F0000)
		flipped |= board.doMoveToLowerBits(0x0000000000004080)
		flipped |= board.doMoveToLowerBits(0x0000000000002020)
		flipped |= board.doMoveToLowerBits(0x0000000000001008)
	case uint64(1 << 22):
		flipped = board.doMoveToHigherBits(0x0204081020000000)
		flipped |= board.doMoveToHigherBits(0x4040404040000000)
		flipped |= board.doMoveToLowerBits(0x00000000003F0000)
		flipped |= board.doMoveToLowerBits(0x0000000000004040)
		flipped |= board.doMoveToLowerBits(0x0000000000002010)
	case uint64(1 << 23):
		flipped = board.doMoveToHigherBits(0x0408102040000000)
		flipped |= board.doMoveToHigherBits(0x8080808080000000)
		flipped |= board.doMoveToLowerBits(0x00000000007F0000)
		flipped |= board.doMoveToLowerBits(0x0000000000008080)
		flipped |= board.doMoveToLowerBits(0x0000000000004020)
	case uint64(1 << 24):
		flipped = board.doMoveToHigherBits(0x00000000FE000000)
		flipped |= board.doMoveToHigherBits(0x0101010100000000)
		flipped |= board.doMoveToHigherBits(0x1008040200000000)
		flipped |= board.doMoveToLowerBits(0x0000000000020408)
		flipped |= board.doMoveToLowerBits(0x0000000000010101)
	case uint64(1 << 25):
		flipped = board.doMoveToHigherBits(0x00000000FC000000)
		flipped |= board.doMoveToHigherBits(0x0202020200000000)
		flipped |= board.doMoveToHigherBits(0x2010080400000000)
		flipped |= board.doMoveToLowerBits(0x0000000000040810)
		flipped |= board.doMoveToLowerBits(0x0000000000020202)
	case uint64(1 << 26):
		flipped = board.doMoveToHigherBits(0x00000000F8000000)
		flipped |= board.doMoveToHigherBits(0x0000010200000000)
		flipped |= board.doMoveToHigherBits(0x0404040400000000)
		flipped |= board.doMoveToHigherBits(0x4020100800000000)
		flipped |= board.doMoveToLowerBits(0x0000000003000000)
		flipped |= board.doMoveToLowerBits(0x0000000000081020)
		flipped |= board.doMoveToLowerBits(0x0000000000040404)
		flipped |= board.doMoveToLowerBits(0x0000000000020100)
	// skip 27 and 28 as they are start discs and would never be played
	case uint64(1 << 29):
		flipped = board.doMoveToHigherBits(0x00000000C0000000)
		flipped |= board.doMoveToHigherBits(0x0204081000000000)
		flipped |= board.doMoveToHigherBits(0x2020202000000000)
		flipped |= board.doMoveToHigherBits(0x0000804000000000)
		flipped |= board.doMoveToLowerBits(0x000000001F000000)
		flipped |= board.doMoveToLowerBits(0x0000000000408000)
		flipped |= board.doMoveToLowerBits(0x0000000000202020)
		flipped |= board.doMoveToLowerBits(0x0000000000100804)
	case uint64(1 << 30):
		flipped = board.doMoveToHigherBits(0x0408102000000000)
		flipped |= board.doMoveToHigherBits(0x4040404000000000)
		flipped |= board.doMoveToLowerBits(0x000000003F000000)
		flipped |= board.doMoveToLowerBits(0x0000000000404040)
		flipped |= board.doMoveToLowerBits(0x0000000000201008)
	case uint64(1 << 31):
		flipped = board.doMoveToHigherBits(0x0810204000000000)
		flipped |= board.doMoveToHigherBits(0x8080808000000000)
		flipped |= board.doMoveToLowerBits(0x000000007F000000)
		flipped |= board.doMoveToLowerBits(0x0000000000808080)
		flipped |= board.doMoveToLowerBits(0x0000000000402010)
	case uint64(1 << 32):
		flipped = board.doMoveToHigherBits(0x000000FE00000000)
		flipped |= board.doMoveToHigherBits(0x0101010000000000)
		flipped |= board.doMoveToHigherBits(0x0804020000000000)
		flipped |= board.doMoveToLowerBits(0x0000000002040810)
		flipped |= board.doMoveToLowerBits(0x0000000001010101)
	case uint64(1 << 33):
		flipped = board.doMoveToHigherBits(0x000000FC00000000)
		flipped |= board.doMoveToHigherBits(0x0202020000000000)
		flipped |= board.doMoveToHigherBits(0x1008040000000000)
		flipped |= board.doMoveToLowerBits(0x0000000004081020)
		flipped |= board.doMoveToLowerBits(0x0000000002020202)
	case uint64(1 << 34):
		flipped = board.doMoveToHigherBits(0x000000F800000000)
		flipped |= board.doMoveToHigherBits(0x0001020000000000)
		flipped |= board.doMoveToHigherBits(0x0404040000000000)
		flipped |= board.doMoveToHigherBits(0x2010080000000000)
		flipped |= board.doMoveToLowerBits(0x0000000300000000)
		flipped |= board.doMoveToLowerBits(0x0000000008102040)
		flipped |= board.doMoveToLowerBits(0x0000000004040404)
		flipped |= board.doMoveToLowerBits(0x0000000002010000)
	// skip 35 and 36 as they are start discs and would never be played
	case uint64(1 << 37):
		flipped = board.doMoveToHigherBits(0x000000C000000000)
		flipped |= board.doMoveToHigherBits(0x0408100000000000)
		flipped |= board.doMoveToHigherBits(0x2020200000000000)
		flipped |= board.doMoveToHigherBits(0x0080400000000000)
		flipped |= board.doMoveToLowerBits(0x0000001F00000000)
		flipped |= board.doMoveToLowerBits(0x0000000040800000)
		flipped |= board.doMoveToLowerBits(0x0000000020202020)
		flipped |= board.doMoveToLowerBits(0x0000000010080402)
	case uint64(1 << 38):
		flipped = board.doMoveToHigherBits(0x0810200000000000)
		flipped |= board.doMoveToHigherBits(0x4040400000000000)
		flipped |= board.doMoveToLowerBits(0x0000003F00000000)
		flipped |= board.doMoveToLowerBits(0x0000000040404040)
		flipped |= board.doMoveToLowerBits(0x0000000020100804)
	case uint64(1 << 39):
		flipped = board.doMoveToHigherBits(0x1020400000000000)
		flipped |= board.doMoveToHigherBits(0x8080800000000000)
		flipped |= board.doMoveToLowerBits(0x0000007F00000000)
		flipped |= board.doMoveToLowerBits(0x0000000080808080)
		flipped |= board.doMoveToLowerBits(0x0000000040201008)
	case uint64(1 << 40):
		flipped = board.doMoveToHigherBits(0x0000FE0000000000)
		flipped |= board.doMoveToHigherBits(0x0101000000000000)
		flipped |= board.doMoveToHigherBits(0x0402000000000000)
		flipped |= board.doMoveToLowerBits(0x0000000204081020)
		flipped |= board.doMoveToLowerBits(0x0000000101010101)
	case uint64(1 << 41):
		flipped = board.doMoveToHigherBits(0x0000FC0000000000)
		flipped |= board.doMoveToHigherBits(0x0202000000000000)
		flipped |= board.doMoveToHigherBits(0x0804000000000000)
		flipped |= board.doMoveToLowerBits(0x0000000408102040)
		flipped |= board.doMoveToLowerBits(0x0000000202020202)
	case uint64(1 << 42):
		flipped = board.doMoveToHigherBits(0x0000F80000000000)
		flipped |= board.doMoveToHigherBits(0x0102000000000000)
		flipped |= board.doMoveToHigherBits(0x0404000000000000)
		flipped |= board.doMoveToHigherBits(0x1008000000000000)
		flipped |= board.doMoveToLowerBits(0x0000030000000000)
		flipped |= board.doMoveToLowerBits(0x0000000810204080)
		flipped |= board.doMoveToLowerBits(0x0000000404040404)
		flipped |= board.doMoveToLowerBits(0x0000000201000000)
	case uint64(1 << 43):
		flipped = board.doMoveToHigherBits(0x0000F00000000000)
		flipped |= board.doMoveToHigherBits(0x0204000000000000)
		flipped |= board.doMoveToHigherBits(0x0808000000000000)
		flipped |= board.doMoveToHigherBits(0x2010000000000000)
		flipped |= board.doMoveToLowerBits(0x0000070000000000)
		flipped |= board.doMoveToLowerBits(0x0000001020408000)
		flipped |= board.doMoveToLowerBits(0x0000000808080808)
		flipped |= board.doMoveToLowerBits(0x0000000402010000)
	case uint64(1 << 44):
		flipped = board.doMoveToHigherBits(0x0000E00000000000)
		flipped |= board.doMoveToHigherBits(0x0408000000000000)
		flipped |= board.doMoveToHigherBits(0x1010000000000000)
		flipped |= board.doMoveToHigherBits(0x4020000000000000)
		flipped |= board.doMoveToLowerBits(0x00000F0000000000)
		flipped |= board.doMoveToLowerBits(0x0000002040800000)
		flipped |= board.doMoveToLowerBits(0x0000001010101010)
		flipped |= board.doMoveToLowerBits(0x0000000804020100)
	case uint64(1 << 45):
		flipped = board.doMoveToHigherBits(0x0000C00000000000)
		flipped |= board.doMoveToHigherBits(0x0810000000000000)
		flipped |= board.doMoveToHigherBits(0x2020000000000000)
		flipped |= board.doMoveToHigherBits(0x8040000000000000)
		flipped |= board.doMoveToLowerBits(0x00001F0000000000)
		flipped |= board.doMoveToLowerBits(0x0000004080000000)
		flipped |= board.doMoveToLowerBits(0x0000002020202020)
		flipped |= board.doMoveToLowerBits(0x0000001008040201)
	case uint64(1 << 46):
		flipped = board.doMoveToHigherBits(0x1020000000000000)
		flipped |= board.doMoveToHigherBits(0x4040000000000000)
		flipped |= board.doMoveToLowerBits(0x00003F0000000000)
		flipped |= board.doMoveToLowerBits(0x0000004040404040)
		flipped |= board.doMoveToLowerBits(0x0000002010080402)
	case uint64(1 << 47):
		flipped = board.doMoveToHigherBits(0x2040000000000000)
		flipped |= board.doMoveToHigherBits(0x8080000000000000)
		flipped |= board.doMoveToLowerBits(0x00007F0000000000)
		flipped |= board.doMoveToLowerBits(0x0000008080808080)
		flipped |= board.doMoveToLowerBits(0x0000004020100804)
	case uint64(1 << 48):
		flipped = board.doMoveToHigherBits(0x00FE000000000000)
		flipped |= board.doMoveToLowerBits(0x0000020408102040)
		flipped |= board.doMoveToLowerBits(0x0000010101010101)
	case uint64(1 << 49):
		flipped = board.doMoveToHigherBits(0x00FC000000000000)
		flipped |= board.doMoveToLowerBits(0x0000040810204080)
		flipped |= board.doMoveToLowerBits(0x0000020202020202)
	case uint64(1 << 50):
		flipped = board.doMoveToHigherBits(0x00F8000000000000)
		flipped |= board.doMoveToLowerBits(0x0003000000000000)
		flipped |= board.doMoveToLowerBits(0x0000081020408000)
		flipped |= board.doMoveToLowerBits(0x0000040404040404)
		flipped |= board.doMoveToLowerBits(0x0000020100000000)
	case uint64(1 << 51):
		flipped = board.doMoveToHigherBits(0x00F0000000000000)
		flipped |= board.doMoveToLowerBits(0x0007000000000000)
		flipped |= board.doMoveToLowerBits(0x0000102040800000)
		flipped |= board.doMoveToLowerBits(0x0000080808080808)
		flipped |= board.doMoveToLowerBits(0x0000040201000000)
	case uint64(1 << 52):
		flipped = board.doMoveToHigherBits(0x00E0000000000000)
		flipped |= board.doMoveToLowerBits(0x000F000000000000)
		flipped |= board.doMoveToLowerBits(0x0000204080000000)
		flipped |= board.doMoveToLowerBits(0x0000101010101010)
		flipped |= board.doMoveToLowerBits(0x0000080402010000)
	case uint64(1 << 53):
		flipped = board.doMoveToHigherBits(0x00C0000000000000)
		flipped |= board.doMoveToLowerBits(0x001F000000000000)
		flipped |= board.doMoveToLowerBits(0x0000408000000000)
		flipped |= board.doMoveToLowerBits(0x0000202020202020)
		flipped |= board.doMoveToLowerBits(0x0000100804020100)
	case uint64(1 << 54):
		flipped = board.doMoveToLowerBits(0x003F000000000000)
		flipped |= board.doMoveToLowerBits(0x0000404040404040)
		flipped |= board.doMoveToLowerBits(0x0000201008040201)
	case uint64(1 << 55):
		flipped = board.doMoveToLowerBits(0x007F000000000000)
		flipped |= board.doMoveToLowerBits(0x0000808080808080)
		flipped |= board.doMoveToLowerBits(0x0000402010080402)
	case uint64(1 << 56):
		flipped = board.doMoveToHigherBits(0xFE00000000000000)
		flipped |= board.doMoveToLowerBits(0x0002040810204080)
		flipped |= board.doMoveToLowerBits(0x0001010101010101)
	case uint64(1 << 57):
		flipped = board.doMoveToHigherBits(0xFC00000000000000)
		flipped |= board.doMoveToLowerBits(0x0004081020408000)
		flipped |= board.doMoveToLowerBits(0x0002020202020202)
	case uint64(1 << 58):
		flipped = board.doMoveToHigherBits(0xF800000000000000)
		flipped |= board.doMoveToLowerBits(0x0300000000000000)
		flipped |= board.doMoveToLowerBits(0x0008102040800000)
		flipped |= board.doMoveToLowerBits(0x0004040404040404)
		flipped |= board.doMoveToLowerBits(0x0002010000000000)
	case uint64(1 << 59):
		flipped = board.doMoveToHigherBits(0xF000000000000000)
		flipped |= board.doMoveToLowerBits(0x0700000000000000)
		flipped |= board.doMoveToLowerBits(0x0010204080000000)
		flipped |= board.doMoveToLowerBits(0x0008080808080808)
		flipped |= board.doMoveToLowerBits(0x0004020100000000)
	case uint64(1 << 60):
		flipped = board.doMoveToHigherBits(0xE000000000000000)
		flipped |= board.doMoveToLowerBits(0x0F00000000000000)
		flipped |= board.doMoveToLowerBits(0x0020408000000000)
		flipped |= board.doMoveToLowerBits(0x0010101010101010)
		flipped |= board.doMoveToLowerBits(0x0008040201000000)
	case uint64(1 << 61):
		flipped = board.doMoveToHigherBits(0xC000000000000000)
		flipped |= board.doMoveToLowerBits(0x1F00000000000000)
		flipped |= board.doMoveToLowerBits(0x0040800000000000)
		flipped |= board.doMoveToLowerBits(0x0020202020202020)
		flipped |= board.doMoveToLowerBits(0x0010080402010000)
	case uint64(1 << 62):
		flipped = board.doMoveToLowerBits(0x3F00000000000000)
		flipped |= board.doMoveToLowerBits(0x0040404040404040)
		flipped |= board.doMoveToLowerBits(0x0020100804020100)
	case uint64(1 << 63):
		flipped = board.doMoveToLowerBits(0x7F00000000000000)
		flipped |= board.doMoveToLowerBits(0x0080808080808080)
		flipped |= board.doMoveToLowerBits(0x0040201008040201)
	}

	tmp := board.me | flipped | moveBit

	board.me = board.opp &^ tmp
	board.opp = tmp

	return flipped
}

// GetChildren returns a slice with all children of a Board
func (board Board) GetChildren() []Board {

	moves := board.Moves()
	children := make([]Board, bits.OnesCount64(moves))

	for i := range children {
		moveBit := moves & (-moves)
		moves &^= moveBit

		children[i] = board
		children[i].DoMove(moveBit)
	}
	return children
}

// GetSortableChildren returns a slice with all children of a Board
// such that they can easily be sorted
func (board Board) GetSortableChildren() []SortableBoard {

	moves := board.Moves()
	children := make([]SortableBoard, bits.OnesCount64(moves))

	for i := range children {
		moveBit := moves & (-moves)
		moves &^= moveBit

		children[i].Board = board
		children[i].Board.DoMove(moveBit)
		children[i].Heur = 0
	}
	return children
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

// CornerCountDifference returns the corner count difference.
// Positive result means player to move has more corners.
func (board Board) CornerCountDifference() int {
	return bits.OnesCount64(board.me&cornerMask) - bits.OnesCount64(board.opp&cornerMask)
}

// XsquareCountDifference returns the x-square count difference
// X-squares are fields fields diagonal to a corner.
// Positive result means player to move has more x-squares.
func (board Board) XsquareCountDifference() int {
	return bits.OnesCount64(board.me&xSquareMask) - bits.OnesCount64(board.opp&xSquareMask)
}

// CsquareCountDifference returns the c-square count difference
// C-squares are fields on side of the board next to a corner.
// Positive result means player to move has more c-squares.
func (board Board) CsquareCountDifference() int {
	return bits.OnesCount64(board.me&cSquareMask) - bits.OnesCount64(board.opp&cSquareMask)
}

func potentialMoves(me, opp uint64) uint64 {

	const (
		leftMask  = 0x7F7F7F7F7F7F7F7F
		rightMask = 0xFEFEFEFEFEFEFEFE
	)

	oppSurrounded := uint64(0)
	oppSurrounded |= (opp & leftMask) << 1
	oppSurrounded |= (opp & rightMask) >> 1
	oppSurrounded |= (opp & leftMask) << 9
	oppSurrounded |= (opp & rightMask) >> 9
	oppSurrounded |= (opp & rightMask) << 7
	oppSurrounded |= (opp & leftMask) >> 7

	oppSurrounded |= opp << 8
	oppSurrounded |= opp >> 8

	oppSurrounded &^= (me | opp)
	return oppSurrounded
}

// PotentialMoveCountDifference returns the difference in a rough estimation of the amount of moves
func (board Board) PotentialMoveCountDifference() int {
	mePotentialMoveCount := bits.OnesCount64(potentialMoves(board.me, board.opp))
	oppPotentialMoveCount := bits.OnesCount64(potentialMoves(board.opp, board.me))
	return mePotentialMoveCount - oppPotentialMoveCount
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
	line &^= bitsBefore[bits.Len64(lineMask)]

	if line&board.opp == line {
		return line
	}
	return 0
}

var bitsBefore = [64]uint64{
	(uint64(1) << 0) - 1,
	(uint64(1) << 1) - 1,
	(uint64(1) << 2) - 1,
	(uint64(1) << 3) - 1,
	(uint64(1) << 4) - 1,
	(uint64(1) << 5) - 1,
	(uint64(1) << 6) - 1,
	(uint64(1) << 7) - 1,
	(uint64(1) << 8) - 1,
	(uint64(1) << 9) - 1,
	(uint64(1) << 10) - 1,
	(uint64(1) << 11) - 1,
	(uint64(1) << 12) - 1,
	(uint64(1) << 13) - 1,
	(uint64(1) << 14) - 1,
	(uint64(1) << 15) - 1,
	(uint64(1) << 16) - 1,
	(uint64(1) << 17) - 1,
	(uint64(1) << 18) - 1,
	(uint64(1) << 19) - 1,
	(uint64(1) << 20) - 1,
	(uint64(1) << 21) - 1,
	(uint64(1) << 22) - 1,
	(uint64(1) << 23) - 1,
	(uint64(1) << 24) - 1,
	(uint64(1) << 25) - 1,
	(uint64(1) << 26) - 1,
	(uint64(1) << 27) - 1,
	(uint64(1) << 28) - 1,
	(uint64(1) << 29) - 1,
	(uint64(1) << 30) - 1,
	(uint64(1) << 31) - 1,
	(uint64(1) << 32) - 1,
	(uint64(1) << 33) - 1,
	(uint64(1) << 34) - 1,
	(uint64(1) << 35) - 1,
	(uint64(1) << 36) - 1,
	(uint64(1) << 37) - 1,
	(uint64(1) << 38) - 1,
	(uint64(1) << 39) - 1,
	(uint64(1) << 40) - 1,
	(uint64(1) << 41) - 1,
	(uint64(1) << 42) - 1,
	(uint64(1) << 43) - 1,
	(uint64(1) << 44) - 1,
	(uint64(1) << 45) - 1,
	(uint64(1) << 46) - 1,
	(uint64(1) << 47) - 1,
	(uint64(1) << 48) - 1,
	(uint64(1) << 49) - 1,
	(uint64(1) << 50) - 1,
	(uint64(1) << 51) - 1,
	(uint64(1) << 52) - 1,
	(uint64(1) << 53) - 1,
	(uint64(1) << 54) - 1,
	(uint64(1) << 55) - 1,
	(uint64(1) << 56) - 1,
	(uint64(1) << 57) - 1,
	(uint64(1) << 58) - 1,
	(uint64(1) << 59) - 1,
	(uint64(1) << 60) - 1,
	(uint64(1) << 61) - 1,
	(uint64(1) << 62) - 1,
	(uint64(1) << 63) - 1}
