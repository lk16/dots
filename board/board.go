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
            fmt.Printf("\033[31;1m●\033[0m ")
        } else if board.opp & bit != 0 {
            fmt.Printf("\033[34;1m●\033[0m ")
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

func (board *Board) Moves() uint64 {
    return 0
}