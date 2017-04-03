package board

import (
    "bytes"
    "fmt"
    "math/rand"

    "dots/bitset"
)

type Board struct {
    me, opp bitset.Bitset
}

// Returns a Board in start state
func NewBoard() *Board {
    return &Board{
        me:  bitset.Bitset(1<<28) | bitset.Bitset(1<<35),
        opp: bitset.Bitset(1<<27) | bitset.Bitset(1<<36)}
}

// Returns a deep copy of a Board
func (board Board) Clone() (clone Board) {
    clone = board
    return
}

// Returns a String of ASCII-art of a Board
func (board Board) AsciiArt() (output string) {

    buffer := new(bytes.Buffer)
    buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

    moves := board.Moves()


    for y := uint(0); y < 8; y++ {
        buffer.WriteString(fmt.Sprintf("%d ", y+1))

        for x := uint(0); x < 8; x++ {
            if board.me.TestBit(y*8 + x) {
                buffer.WriteString("○ ")
            } else if board.opp.TestBit(y*8 + x) {
                buffer.WriteString("● ")
            } else if moves.TestBit(y*8 + x) {
                buffer.WriteString("- ")
            } else {
                buffer.WriteString("  ")
            }
        }

        buffer.WriteString("|\n")
    }
    buffer.WriteString("+-----------------+\n")

    output = buffer.String()
    return
}

// Returns a subset of the moves for a Board
func movesPartial(me, mask, n bitset.Bitset) (moves bitset.Bitset) {
    flip_l := mask & (me << n)
    flip_l |= mask & (flip_l << n)
    mask_l := mask & (mask << n)
    flip_l |= mask_l & (flip_l << (2 * n))
    flip_l |= mask_l & (flip_l << (2 * n))
    flip_r := mask & (me >> n)
    flip_r |= mask & (flip_r >> n)
    mask_r := mask & (mask >> n)
    flip_r |= mask_r & (flip_r >> (2 * n))
    flip_r |= mask_r & (flip_r >> (2 * n))
    moves = (flip_l << n) | (flip_r >> n)
    return
}

// Returns a Bitset with all valid moves for a Board
func (board Board) Moves() (moves bitset.Bitset) {
    // this function is a modified version of code from Edax
    mask := board.opp & 0x7E7E7E7E7E7E7E7E

    moves = movesPartial(board.me, mask, 1)
    moves |= movesPartial(board.me, mask, 7)
    moves |= movesPartial(board.me, mask, 9)
    moves |= movesPartial(board.me, board.opp, 8)

    moves &^= (board.me | board.opp)
    return
}

// Does the move at field index on a Board
// Returns the flipped discs
func (board *Board) DoMove(index uint) (flipped bitset.Bitset) {

    doMoveFuncs := []func() bitset.Bitset{
        board.doMove0, board.doMove1, board.doMove2, board.doMove3,
        board.doMove4, board.doMove5, board.doMove6, board.doMove7,
        board.doMove8, board.doMove9, board.doMove10, board.doMove11,
        board.doMove12, board.doMove13, board.doMove14, board.doMove15,
        board.doMove16, board.doMove17, board.doMove18, board.doMove19,
        board.doMove20, board.doMove21, board.doMove22, board.doMove23,
        board.doMove24, board.doMove25, board.doMove26, board.doMove27,
        board.doMove28, board.doMove29, board.doMove30, board.doMove31,
        board.doMove32, board.doMove33, board.doMove34, board.doMove35,
        board.doMove36, board.doMove37, board.doMove38, board.doMove39,
        board.doMove40, board.doMove41, board.doMove42, board.doMove43,
        board.doMove44, board.doMove45, board.doMove46, board.doMove47,
        board.doMove48, board.doMove49, board.doMove50, board.doMove51,
        board.doMove52, board.doMove53, board.doMove54, board.doMove55,
        board.doMove56, board.doMove57, board.doMove58, board.doMove59,
        board.doMove60, board.doMove61, board.doMove62, board.doMove63}

    flipped = doMoveFuncs[index]()

    tmp := board.me | flipped | bitset.Bitset(1<<index)

    board.me = board.opp &^ tmp
    board.opp = tmp

    return flipped
}

// Returns a channel that outputs the children of a Board
func (board Board) GenChildren() (ch chan Board) {
    ch = make(chan Board)
    go func() {
        moves := board.Moves()
        for {
            index := moves.FirstBitIndex()
            if index == uint(0) {
                break
            }
            moves &^= bitset.Bitset(1 << index)

            clone := board.Clone()
            clone.DoMove(index)
            ch <- clone
        }
        close(ch)
    }()
    return
}

// Returns a slice with all children of a Board
func (board Board) GetChildren() (children []Board) {
    for child := range board.GenChildren() {
        children = append(children, child)
    }
    return
}

// Does n random moves on a Board
func (board *Board) DoRandomMoves(n uint) {

    skips := 0

    for (n > 0) && (skips < 2) {

        children := board.GetChildren()
        if len(children) == 0 {
            board.SwitchTurn()
            skips++
            continue
        }

        *board = children[rand.Int()%len(children)]
        n--
        skips = 0
    }
}

// Switches turn of a Board
func (board *Board) SwitchTurn() {
    tmp := board.me
    board.me = board.opp
    board.opp = tmp
}
