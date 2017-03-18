package board

import "testing"

func genTestBoards() (ch chan Board) {
    ch = make(chan Board)
    go func() {
        ch <- *NewBoard()
        // TODO
        close(ch)
    }()
    return
}

func (board *Board) doMove(index uint) bitset {
    if (board.me | board.opp) & (bitset(1) << index) != 0 {
        return 0
    }
    flipped := bitset(0)
    for dx:=-1; dx<=1; dx++ {
        for dy:=-1; dy<=1; dy++ {
            if dx == 0 && dy == 0 {
                continue
            }
            s := 1
            for {
                curx := int(index%8) + (dx*s)
                cury := int(index/8) + (dy*s)
                cur := uint(8*cury + curx)
                if curx < 0 || curx >= 8 || cury < 0 || cury >= 8 {
                    break
                }
                if board.opp.TestBit(cur) {
                    s++
                } else {
                    if board.me.TestBit(cur) && (s >= 2) {
                        for p:=1; p<s; p++ {
                            flipped |= bitset(1) << uint(int(index) + (8*dy*p) + (dx*p))
                        }
                    }
                    break
                }
            }
        }
    }
    board.me |= flipped | (bitset(1) << index)
    board.opp &= ^board.me
    board.opp,board.me = board.me,board.opp
    return flipped
}

func (board *Board) moves() bitset {
    moves := bitset(0)
    for i:=uint(0); i<64; i++ {
        clone := board.Clone()
        if clone.DoMove(i) != bitset(0) {
            moves |= bitset(1) << i
        }
    }
    return moves
}

func TestMoves(t *testing.T) {
    for board := range genTestBoards() {
        expected := board.moves()
        got := board.Moves()
        if expected != got {
            t.Errorf("") // TODO
        }
    }
}

func TestDoMove(t *testing.T) {
    for board := range genTestBoards() {
        moves := board.Moves()
        for i:=uint(0); i<64; i++ {
            // board.DoMove() should not be called for invalid moves
            if moves.TestBit(i) {
                clone := board.Clone()
                expected_return_val := clone.doMove(i)
                expected_board_val := clone
                got_return_val := board.DoMove(i)
                got_board_val := board
                if got_return_val != expected_return_val {
                    t.Errorf("Got %d, expected %d.\n",got_return_val,expected_return_val) // TODO
                } 
                if got_board_val != expected_board_val {
                    t.Errorf("") // TODO
                }
            }
        }
    }
}
