package board

import (
    "fmt"
    "github.com/willf/bitset"
)

type Board struct {
    me,opp bitset.BitSet
}

func NewBoard() *Board {
    board := &Board{
        me: *bitset.New(64),
        opp: *bitset.New(64)}

    board.me.Set(28).Set(35)
    board.opp.Set(27).Set(36)

    return board
}

func (board *Board) Print() {
    moves := board.Moves()

    fmt.Printf("+-a-b-c-d-e-f-g-h-+\n")
    for f:=uint(0); f<64; f++ {
        if f%8 == 0 {
            fmt.Printf("%d ",(f/8)+1)
        }
        if board.me.Test(f) {
            fmt.Printf("○ ")
        } else if board.opp.Test(f) {
            fmt.Printf("● ")
        } else if moves.Test(f) {
            fmt.Printf("- ")
        } else {
            fmt.Printf("  ")
        }

        if f%8 == 7 {
            fmt.Printf("|\n")
        }
    }
    fmt.Printf("+-----------------+\n")

}

func movesPartial(me,mask,n uint64) uint64 {
    flip_l := mask & (me << n)
    flip_l |= mask & (flip_l << n)
    mask_l := mask & (mask << n)
    flip_l |= mask_l & (flip_l << (2*n))
    flip_l |= mask_l & (flip_l << (2*n))
    flip_r := mask & (me >> n)
    flip_r |= mask & (flip_r >> n)
    mask_r := mask & (mask >> n)
    flip_r |= mask_r & (flip_r >> (2*n))
    flip_r |= mask_r & (flip_r >> (2*n))
    return (flip_l << n) | (flip_r >> n)
}

func (board *Board) Moves() bitset.BitSet {
    // this function is a modified version of code from Edax
    mask := board.opp.Bytes()[0] & 0x7E7E7E7E7E7E7E7E
    res := movesPartial(board.me.Bytes()[0],mask,1)
    res |= movesPartial(board.me.Bytes()[0],mask,7)
    res |= movesPartial(board.me.Bytes()[0],mask,9)
    res |= movesPartial(board.me.Bytes()[0],board.opp.Bytes()[0],8)
    
    return *bitset.From([]uint64{res & ^(board.me.Bytes()[0] | board.opp.Bytes()[0])})
}



func (board *Board) DoMove(index uint) uint64 {

    doMoveFuncs := []func() uint64{
         board.doMove0, board.doMove1, board.doMove2, board.doMove3, board.doMove4, board.doMove5, board.doMove6, board.doMove7,
         board.doMove8, board.doMove9,board.doMove10,board.doMove11,board.doMove12,board.doMove13,board.doMove14,board.doMove15,
        board.doMove16,board.doMove17,board.doMove18,board.doMove19,board.doMove20,board.doMove21,board.doMove22,board.doMove23,
        board.doMove24,board.doMove25,board.doMove26,board.doMove27,board.doMove28,board.doMove29,board.doMove30,board.doMove31,
        board.doMove32,board.doMove33,board.doMove34,board.doMove35,board.doMove36,board.doMove37,board.doMove38,board.doMove39,
        board.doMove40,board.doMove41,board.doMove42,board.doMove43,board.doMove44,board.doMove45,board.doMove46,board.doMove47,
        board.doMove48,board.doMove49,board.doMove50,board.doMove51,board.doMove52,board.doMove53,board.doMove54,board.doMove55,
        board.doMove56,board.doMove57,board.doMove58,board.doMove59,board.doMove60,board.doMove61,board.doMove62,board.doMove63}

    flipped := doMoveFuncs[index]()

    me := board.opp.Bytes()[0]

    opp := board.me.Bytes()[0] | flipped | (1 << index)

    me = me &^ opp

    board.me = *bitset.From([]uint64{me})
    board.opp = *bitset.From([]uint64{opp})

    return flipped
}