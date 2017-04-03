package board

import "dots/bitset"

// Flips discs on a Board, given a flipping line.
// This only affects the directions right, left down, down and right down
// Returns the flipped discs.
func (board *Board) doMoveToHigherBits(line bitset.Bitset) (flipped bitset.Bitset) {
    line_mask := line & board.me
    bit := line_mask.FirstBit()
    line &= bitset.Bitset(bit - 1)
    if (bit != 0) && (line&board.opp == line) {
        flipped = line
    }    
    return
}

// Flips discs on a Board, given a flipping line.
// This only affects the directions left up, up, right up and left
// Returns the flipped discs.
func (board *Board) doMoveToLowerBits(line bitset.Bitset) (flipped bitset.Bitset) {
    line_mask := line & board.me
    bit := line_mask.LastBit()
    if line_mask == 0 {
        bit = 0
    }
    line &^= bitset.Bitset((bit << 1) - 1)
    if (bit != 0) && (line&board.opp == line) {
        flipped = line
    }
    return
}

//Does the move at field 0.
//Returns the flipped discs.
func (board *Board) doMove0() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000000000FE)
    flipped |= board.doMoveToHigherBits(0x0101010101010100)
    flipped |= board.doMoveToHigherBits(0x8040201008040200)
    return
}

//Does the move at field 1.
//Returns the flipped discs.
func (board *Board) doMove1() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000000000FC)
    flipped |= board.doMoveToHigherBits(0x0202020202020200)
    flipped |= board.doMoveToHigherBits(0x0080402010080400)
    return
}

//Does the move at field 2.
//Returns the flipped discs.
func (board *Board) doMove2() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000000000F8)
    flipped |= board.doMoveToHigherBits(0x0000000000010200)
    flipped |= board.doMoveToHigherBits(0x0404040404040400)
    flipped |= board.doMoveToHigherBits(0x0000804020100800)
    flipped |= board.doMoveToLowerBits(0x0000000000000003)
    return
}

//Does the move at field 3.
//Returns the flipped discs.
func (board *Board) doMove3() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000000000F0)
    flipped |= board.doMoveToHigherBits(0x0000000001020400)
    flipped |= board.doMoveToHigherBits(0x0808080808080800)
    flipped |= board.doMoveToHigherBits(0x0000008040201000)
    flipped |= board.doMoveToLowerBits(0x0000000000000007)
    return
}

//Does the move at field 4.
//Returns the flipped discs.
func (board *Board) doMove4() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000000000E0)
    flipped |= board.doMoveToHigherBits(0x0000000102040800)
    flipped |= board.doMoveToHigherBits(0x1010101010101000)
    flipped |= board.doMoveToHigherBits(0x0000000080402000)
    flipped |= board.doMoveToLowerBits(0x000000000000000F)
    return
}

//Does the move at field 5.
//Returns the flipped discs.
func (board *Board) doMove5() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000000000C0)
    flipped |= board.doMoveToHigherBits(0x0000010204081000)
    flipped |= board.doMoveToHigherBits(0x2020202020202000)
    flipped |= board.doMoveToHigherBits(0x0000000000804000)
    flipped |= board.doMoveToLowerBits(0x000000000000001F)
    return
}

//Does the move at field 6.
//Returns the flipped discs.
func (board *Board) doMove6() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0001020408102000)
    flipped |= board.doMoveToHigherBits(0x4040404040404000)
    flipped |= board.doMoveToLowerBits(0x000000000000003F)
    return
}

//Does the move at field 7.
//Returns the flipped discs.
func (board *Board) doMove7() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0102040810204000)
    flipped |= board.doMoveToHigherBits(0x8080808080808000)
    flipped |= board.doMoveToLowerBits(0x000000000000007F)
    return
}

//Does the move at field 8.
//Returns the flipped discs.
func (board *Board) doMove8() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000000000FE00)
    flipped |= board.doMoveToHigherBits(0x0101010101010000)
    flipped |= board.doMoveToHigherBits(0x4020100804020000)
    return
}

//Does the move at field 9.
//Returns the flipped discs.
func (board *Board) doMove9() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000000000FC00)
    flipped |= board.doMoveToHigherBits(0x0202020202020000)
    flipped |= board.doMoveToHigherBits(0x8040201008040000)
    return
}

//Does the move at field 10.
//Returns the flipped discs.
func (board *Board) doMove10() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000000000F800)
    flipped |= board.doMoveToHigherBits(0x0000000001020000)
    flipped |= board.doMoveToHigherBits(0x0404040404040000)
    flipped |= board.doMoveToHigherBits(0x0080402010080000)
    flipped |= board.doMoveToLowerBits(0x0000000000000300)
    return
}

//Does the move at field 11.
//Returns the flipped discs.
func (board *Board) doMove11() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000000000F000)
    flipped |= board.doMoveToHigherBits(0x0000000102040000)
    flipped |= board.doMoveToHigherBits(0x0808080808080000)
    flipped |= board.doMoveToHigherBits(0x0000804020100000)
    flipped |= board.doMoveToLowerBits(0x0000000000000700)
    return
}

//Does the move at field 12.
//Returns the flipped discs.
func (board *Board) doMove12() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000000000E000)
    flipped |= board.doMoveToHigherBits(0x0000010204080000)
    flipped |= board.doMoveToHigherBits(0x1010101010100000)
    flipped |= board.doMoveToHigherBits(0x0000008040200000)
    flipped |= board.doMoveToLowerBits(0x0000000000000F00)
    return
}

//Does the move at field 13.
//Returns the flipped discs.
func (board *Board) doMove13() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000000000C000)
    flipped |= board.doMoveToHigherBits(0x0001020408100000)
    flipped |= board.doMoveToHigherBits(0x2020202020200000)
    flipped |= board.doMoveToHigherBits(0x0000000080400000)
    flipped |= board.doMoveToLowerBits(0x0000000000001F00)
    return
}

//Does the move at field 14.
//Returns the flipped discs.
func (board *Board) doMove14() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0102040810200000)
    flipped |= board.doMoveToHigherBits(0x4040404040400000)
    flipped |= board.doMoveToLowerBits(0x0000000000003F00)
    return
}

//Does the move at field 15.
//Returns the flipped discs.
func (board *Board) doMove15() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0204081020400000)
    flipped |= board.doMoveToHigherBits(0x8080808080800000)
    flipped |= board.doMoveToLowerBits(0x0000000000007F00)
    return
}

//Does the move at field 16.
//Returns the flipped discs.
func (board *Board) doMove16() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0000000000FE0000)
    flipped |= board.doMoveToHigherBits(0x0101010101000000)
    flipped |= board.doMoveToHigherBits(0x2010080402000000)
    flipped |= board.doMoveToLowerBits(0x0000000000000204)
    flipped |= board.doMoveToLowerBits(0x0000000000000101)
    return
}

//Does the move at field 17.
//Returns the flipped discs.
func (board *Board) doMove17() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0000000000FC0000)
    flipped |= board.doMoveToHigherBits(0x0202020202000000)
    flipped |= board.doMoveToHigherBits(0x4020100804000000)
    flipped |= board.doMoveToLowerBits(0x0000000000000408)
    flipped |= board.doMoveToLowerBits(0x0000000000000202)
    return
}

//Does the move at field 18.
//Returns the flipped discs.
func (board *Board) doMove18() (flipped bitset.Bitset) {
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
func (board *Board) doMove19() (flipped bitset.Bitset) {
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
func (board *Board) doMove20() (flipped bitset.Bitset) {
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
func (board *Board) doMove21() (flipped bitset.Bitset) {
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
func (board *Board) doMove22() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0204081020000000)
    flipped |= board.doMoveToHigherBits(0x4040404040000000)
    flipped |= board.doMoveToLowerBits(0x00000000003F0000)
    flipped |= board.doMoveToLowerBits(0x0000000000004040)
    flipped |= board.doMoveToLowerBits(0x0000000000002010)
    return
}

//Does the move at field 23.
//Returns the flipped discs.
func (board *Board) doMove23() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0408102040000000)
    flipped |= board.doMoveToHigherBits(0x8080808080000000)
    flipped |= board.doMoveToLowerBits(0x00000000007F0000)
    flipped |= board.doMoveToLowerBits(0x0000000000008080)
    flipped |= board.doMoveToLowerBits(0x0000000000004020)
    return
}

//Does the move at field 24.
//Returns the flipped discs.
func (board *Board) doMove24() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000FE000000)
    flipped |= board.doMoveToHigherBits(0x0101010100000000)
    flipped |= board.doMoveToHigherBits(0x1008040200000000)
    flipped |= board.doMoveToLowerBits(0x0000000000020408)
    flipped |= board.doMoveToLowerBits(0x0000000000010101)
    return
}

//Does the move at field 25.
//Returns the flipped discs.
func (board *Board) doMove25() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00000000FC000000)
    flipped |= board.doMoveToHigherBits(0x0202020200000000)
    flipped |= board.doMoveToHigherBits(0x2010080400000000)
    flipped |= board.doMoveToLowerBits(0x0000000000040810)
    flipped |= board.doMoveToLowerBits(0x0000000000020202)
    return
}

//Does the move at field 26.
//Returns the flipped discs.
func (board *Board) doMove26() (flipped bitset.Bitset) {
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
func (board *Board) doMove27() (flipped bitset.Bitset) {
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
func (board *Board) doMove28() (flipped bitset.Bitset) {
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
func (board *Board) doMove29() (flipped bitset.Bitset) {
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
func (board *Board) doMove30() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0408102000000000)
    flipped |= board.doMoveToHigherBits(0x4040404000000000)
    flipped |= board.doMoveToLowerBits(0x000000003F000000)
    flipped |= board.doMoveToLowerBits(0x0000000000404040)
    flipped |= board.doMoveToLowerBits(0x0000000000201008)
    return
}

//Does the move at field 31.
//Returns the flipped discs.
func (board *Board) doMove31() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0810204000000000)
    flipped |= board.doMoveToHigherBits(0x8080808000000000)
    flipped |= board.doMoveToLowerBits(0x000000007F000000)
    flipped |= board.doMoveToLowerBits(0x0000000000808080)
    flipped |= board.doMoveToLowerBits(0x0000000000402010)
    return
}

//Does the move at field 32.
//Returns the flipped discs.
func (board *Board) doMove32() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000FE00000000)
    flipped |= board.doMoveToHigherBits(0x0101010000000000)
    flipped |= board.doMoveToHigherBits(0x0804020000000000)
    flipped |= board.doMoveToLowerBits(0x0000000002040810)
    flipped |= board.doMoveToLowerBits(0x0000000001010101)
    return
}

//Does the move at field 33.
//Returns the flipped discs.
func (board *Board) doMove33() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x000000FC00000000)
    flipped |= board.doMoveToHigherBits(0x0202020000000000)
    flipped |= board.doMoveToHigherBits(0x1008040000000000)
    flipped |= board.doMoveToLowerBits(0x0000000004081020)
    flipped |= board.doMoveToLowerBits(0x0000000002020202)
    return
}

//Does the move at field 34.
//Returns the flipped discs.
func (board *Board) doMove34() (flipped bitset.Bitset) {
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
func (board *Board) doMove35() (flipped bitset.Bitset) {
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
func (board *Board) doMove36() (flipped bitset.Bitset) {
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
func (board *Board) doMove37() (flipped bitset.Bitset) {
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
func (board *Board) doMove38() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0810200000000000)
    flipped |= board.doMoveToHigherBits(0x4040400000000000)
    flipped |= board.doMoveToLowerBits(0x0000003F00000000)
    flipped |= board.doMoveToLowerBits(0x0000000040404040)
    flipped |= board.doMoveToLowerBits(0x0000000020100804)
    return
}

//Does the move at field 39.
//Returns the flipped discs.
func (board *Board) doMove39() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x1020400000000000)
    flipped |= board.doMoveToHigherBits(0x8080800000000000)
    flipped |= board.doMoveToLowerBits(0x0000007F00000000)
    flipped |= board.doMoveToLowerBits(0x0000000080808080)
    flipped |= board.doMoveToLowerBits(0x0000000040201008)
    return
}

//Does the move at field 40.
//Returns the flipped discs.
func (board *Board) doMove40() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0000FE0000000000)
    flipped |= board.doMoveToHigherBits(0x0101000000000000)
    flipped |= board.doMoveToHigherBits(0x0402000000000000)
    flipped |= board.doMoveToLowerBits(0x0000000204081020)
    flipped |= board.doMoveToLowerBits(0x0000000101010101)
    return
}

//Does the move at field 41.
//Returns the flipped discs.
func (board *Board) doMove41() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x0000FC0000000000)
    flipped |= board.doMoveToHigherBits(0x0202000000000000)
    flipped |= board.doMoveToHigherBits(0x0804000000000000)
    flipped |= board.doMoveToLowerBits(0x0000000408102040)
    flipped |= board.doMoveToLowerBits(0x0000000202020202)
    return
}

//Does the move at field 42.
//Returns the flipped discs.
func (board *Board) doMove42() (flipped bitset.Bitset) {
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
func (board *Board) doMove43() (flipped bitset.Bitset) {
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
func (board *Board) doMove44() (flipped bitset.Bitset) {
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
func (board *Board) doMove45() (flipped bitset.Bitset) {
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
func (board *Board) doMove46() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x1020000000000000)
    flipped |= board.doMoveToHigherBits(0x4040000000000000)
    flipped |= board.doMoveToLowerBits(0x00003F0000000000)
    flipped |= board.doMoveToLowerBits(0x0000004040404040)
    flipped |= board.doMoveToLowerBits(0x0000002010080402)
    return
}

//Does the move at field 47.
//Returns the flipped discs.
func (board *Board) doMove47() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x2040000000000000)
    flipped |= board.doMoveToHigherBits(0x8080000000000000)
    flipped |= board.doMoveToLowerBits(0x00007F0000000000)
    flipped |= board.doMoveToLowerBits(0x0000008080808080)
    flipped |= board.doMoveToLowerBits(0x0000004020100804)
    return
}

//Does the move at field 48.
//Returns the flipped discs.
func (board *Board) doMove48() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00FE000000000000)
    flipped |= board.doMoveToLowerBits(0x0000020408102040)
    flipped |= board.doMoveToLowerBits(0x0000010101010101)
    return
}

//Does the move at field 49.
//Returns the flipped discs.
func (board *Board) doMove49() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00FC000000000000)
    flipped |= board.doMoveToLowerBits(0x0000040810204080)
    flipped |= board.doMoveToLowerBits(0x0000020202020202)
    return
}

//Does the move at field 50.
//Returns the flipped discs.
func (board *Board) doMove50() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00F8000000000000)
    flipped |= board.doMoveToLowerBits(0x0003000000000000)
    flipped |= board.doMoveToLowerBits(0x0000081020408000)
    flipped |= board.doMoveToLowerBits(0x0000040404040404)
    flipped |= board.doMoveToLowerBits(0x0000020100000000)
    return
}

//Does the move at field 51.
//Returns the flipped discs.
func (board *Board) doMove51() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00F0000000000000)
    flipped |= board.doMoveToLowerBits(0x0007000000000000)
    flipped |= board.doMoveToLowerBits(0x0000102040800000)
    flipped |= board.doMoveToLowerBits(0x0000080808080808)
    flipped |= board.doMoveToLowerBits(0x0000040201000000)
    return
}

//Does the move at field 52.
//Returns the flipped discs.
func (board *Board) doMove52() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00E0000000000000)
    flipped |= board.doMoveToLowerBits(0x000F000000000000)
    flipped |= board.doMoveToLowerBits(0x0000204080000000)
    flipped |= board.doMoveToLowerBits(0x0000101010101010)
    flipped |= board.doMoveToLowerBits(0x0000080402010000)
    return
}

//Does the move at field 53.
//Returns the flipped discs.
func (board *Board) doMove53() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0x00C0000000000000)
    flipped |= board.doMoveToLowerBits(0x001F000000000000)
    flipped |= board.doMoveToLowerBits(0x0000408000000000)
    flipped |= board.doMoveToLowerBits(0x0000202020202020)
    flipped |= board.doMoveToLowerBits(0x0000100804020100)
    return
}

//Does the move at field 54.
//Returns the flipped discs.
func (board *Board) doMove54() (flipped bitset.Bitset) {
    flipped = board.doMoveToLowerBits(0x003F000000000000)
    flipped |= board.doMoveToLowerBits(0x0000404040404040)
    flipped |= board.doMoveToLowerBits(0x0000201008040201)
    return
}

//Does the move at field 55.
//Returns the flipped discs.
func (board *Board) doMove55() (flipped bitset.Bitset) {
    flipped = board.doMoveToLowerBits(0x007F000000000000)
    flipped |= board.doMoveToLowerBits(0x0000808080808080)
    flipped |= board.doMoveToLowerBits(0x0000402010080402)
    return
}

//Does the move at field 56.
//Returns the flipped discs.
func (board *Board) doMove56() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0xFE00000000000000)
    flipped |= board.doMoveToLowerBits(0x0002040810204080)
    flipped |= board.doMoveToLowerBits(0x0001010101010101)
    return
}

//Does the move at field 57.
//Returns the flipped discs.
func (board *Board) doMove57() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0xFC00000000000000)
    flipped |= board.doMoveToLowerBits(0x0004081020408000)
    flipped |= board.doMoveToLowerBits(0x0002020202020202)
    return
}

//Does the move at field 58.
//Returns the flipped discs.
func (board *Board) doMove58() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0xF800000000000000)
    flipped |= board.doMoveToLowerBits(0x0300000000000000)
    flipped |= board.doMoveToLowerBits(0x0008102040800000)
    flipped |= board.doMoveToLowerBits(0x0004040404040404)
    flipped |= board.doMoveToLowerBits(0x0002010000000000)
    return
}

//Does the move at field 59.
//Returns the flipped discs.
func (board *Board) doMove59() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0xF000000000000000)
    flipped |= board.doMoveToLowerBits(0x0700000000000000)
    flipped |= board.doMoveToLowerBits(0x0010204080000000)
    flipped |= board.doMoveToLowerBits(0x0008080808080808)
    flipped |= board.doMoveToLowerBits(0x0004020100000000)
    return
}

//Does the move at field 60.
//Returns the flipped discs.
func (board *Board) doMove60() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0xE000000000000000)
    flipped |= board.doMoveToLowerBits(0x0F00000000000000)
    flipped |= board.doMoveToLowerBits(0x0020408000000000)
    flipped |= board.doMoveToLowerBits(0x0010101010101010)
    flipped |= board.doMoveToLowerBits(0x0008040201000000)
    return
}

//Does the move at field 61.
//Returns the flipped discs.
func (board *Board) doMove61() (flipped bitset.Bitset) {
    flipped = board.doMoveToHigherBits(0xC000000000000000)
    flipped |= board.doMoveToLowerBits(0x1F00000000000000)
    flipped |= board.doMoveToLowerBits(0x0040800000000000)
    flipped |= board.doMoveToLowerBits(0x0020202020202020)
    flipped |= board.doMoveToLowerBits(0x0010080402010000)
    return
}

//Does the move at field 62.
//Returns the flipped discs.
func (board *Board) doMove62() (flipped bitset.Bitset) {
    flipped = board.doMoveToLowerBits(0x3F00000000000000)
    flipped |= board.doMoveToLowerBits(0x0040404040404040)
    flipped |= board.doMoveToLowerBits(0x0020100804020100)
    return
}

//Does the move at field 63.
//Returns the flipped discs.
func (board *Board) doMove63() (flipped bitset.Bitset) {
    flipped = board.doMoveToLowerBits(0x7F00000000000000)
    flipped |= board.doMoveToLowerBits(0x0080808080808080)
    flipped |= board.doMoveToLowerBits(0x0040201008040201)
    return
}
