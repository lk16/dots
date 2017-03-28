package board

import (
    "fmt"
    "bytes"
    "dots/bitset" 
)

type Board struct {
    me,opp bitset.Bitset
}

// Returns a Board in start state
func NewBoard() *Board {
    return &Board{
        me: bitset.Bitset(1 << 28) | bitset.Bitset(1 << 35),
        opp: bitset.Bitset(1 << 27) | bitset.Bitset(1 << 36)}
}

// Returns a deep copy of a Board
func (board *Board) Clone() Board {
    clone := *board
    return clone
}

// Returns a String of ASCII-art of a Board
func (board *Board) AsciiArt() string {

    buffer := new(bytes.Buffer)

    moves := board.Moves()

    buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")
    
    for y:=uint(0); y<8; y++ {

        buffer.WriteString(fmt.Sprintf("%d ",y+1))

        for x:=uint(0); x<8; x++ {
        
            f := y*8 + x

            if board.me.TestBit(f) {
                buffer.WriteString("○ ")
            } else if board.opp.TestBit(f) {
                buffer.WriteString("● ")
            } else if moves.TestBit(f) {
                buffer.WriteString("- ")
            } else {
                buffer.WriteString("  ")
            }

        }
     
        buffer.WriteString("|\n")
    }
    buffer.WriteString("+-----------------+\n")


    return buffer.String()
}

// Returns a subset of the moves for a Board
func movesPartial(me,mask,n bitset.Bitset) bitset.Bitset {
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

// Returns a Bitset with all valid moves for a Board
func (board *Board) Moves() bitset.Bitset {
    // this function is a modified version of code from Edax
    mask := board.opp & 0x7E7E7E7E7E7E7E7E

    res := movesPartial(board.me,mask,1)
    res |= movesPartial(board.me,mask,7)
    res |= movesPartial(board.me,mask,9)
    res |= movesPartial(board.me,board.opp,8)
    
    empties := ^(board.me | board.opp)

    return res & empties
}

// Does the move at field index on a Board
// Returns the flipped discs
func (board *Board) DoMove(index uint) bitset.Bitset {

    doMoveFuncs := []func() bitset.Bitset{
         board.doMove0, board.doMove1, board.doMove2, board.doMove3,
         board.doMove4, board.doMove5, board.doMove6, board.doMove7,
         board.doMove8, board.doMove9,board.doMove10,board.doMove11,
        board.doMove12,board.doMove13,board.doMove14,board.doMove15,
        board.doMove16,board.doMove17,board.doMove18,board.doMove19,
        board.doMove20,board.doMove21,board.doMove22,board.doMove23,
        board.doMove24,board.doMove25,board.doMove26,board.doMove27,
        board.doMove28,board.doMove29,board.doMove30,board.doMove31,
        board.doMove32,board.doMove33,board.doMove34,board.doMove35,
        board.doMove36,board.doMove37,board.doMove38,board.doMove39,
        board.doMove40,board.doMove41,board.doMove42,board.doMove43,
        board.doMove44,board.doMove45,board.doMove46,board.doMove47,
        board.doMove48,board.doMove49,board.doMove50,board.doMove51,
        board.doMove52,board.doMove53,board.doMove54,board.doMove55,
        board.doMove56,board.doMove57,board.doMove58,board.doMove59,
        board.doMove60,board.doMove61,board.doMove62,board.doMove63}

    flipped := doMoveFuncs[index]()

    tmp := board.me | flipped | bitset.Bitset(1 << index)

    board.me = board.opp &^ tmp
    board.opp = tmp

    return flipped
}

// Returns a slice with all children of a Board
func (board *Board) GetChildren() (children []Board) {
    children = make([]Board,10)
    moves := board.Moves()

    for {
        index := moves.FirstBitIndex();
        if index == uint(0) {
            break
        }
        moves &^= bitset.Bitset(1<<index)

        clone := board.Clone()
        clone.DoMove(index)
        children = append(children,clone)
    }
    return
}