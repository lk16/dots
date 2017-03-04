package board

import "fmt"

type Board struct {
    me,opp uint64
}

func NewBoard() *Board {
    return &Board{
        me: (1 << 28) | (1 << 35),
        opp: (1 << 27) | (1 << 36)}
}

func (board *Board) Print() {
    moves := board.Moves()

    fmt.Printf("+-a-b-c-d-e-f-g-h-+\n")
    for f:=uint(0); f<64; f++ {
        bit := uint64(1) << f
        if f%8 == 0 {
            fmt.Printf("%d ",(f/8)+1)
        }
        if board.me & bit != 0 {
            fmt.Printf("○ ")
        } else if board.opp & bit != 0 {
            fmt.Printf("● ")
        } else if moves & bit != 0 {
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

func (board *Board) Moves() uint64 {
    // this function is a modified version of code from Edax
    mask := board.opp & 0x7E7E7E7E7E7E7E7E
    res := movesPartial(board.me,mask,1)
    res |= movesPartial(board.me,mask,7)
    res |= movesPartial(board.me,mask,9)
    res |= movesPartial(board.me,board.opp,8)
    return res & ^(board.me | board.opp)
}