package board

import (
    "bytes"
    "fmt"
    "strings"
    "testing"

    "dots/bitset"
)

func genTestBoards() (ch chan Board) {
    ch = make(chan Board)
    go func() {
        ch <- *NewBoard()
        // TODO
        close(ch)
    }()
    return
}

func TestBoardClone(t *testing.T) {
    board := Board{
        me: 1,
        opp: 2}
    clone := board.Clone()
    clone.me = 3
    if board.me != 1 {
        t.Errorf("Board.Clone() does not make a deep copy!\n")
    }     
}


func (board *Board) doMove(index uint) bitset.Bitset {
    if (board.me | board.opp) & bitset.Bitset(1 << index) != 0 {
        return 0
    }
    flipped := bitset.Bitset(0)
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
                            f := uint(int(index) + (8*dy*p) + (dx*p))
                            flipped |= bitset.Bitset(1 << f)
                        }
                    }
                    break
                }
            }
        }
    }
    board.me |= flipped | bitset.Bitset(1 << index)
    board.opp &= ^board.me
    board.opp,board.me = board.me,board.opp
    return flipped
}

func TestBoardDoMove(t *testing.T) {
    for board := range genTestBoards() {
        moves := board.Moves()
        for i:=uint(0); i<64; i++ {
            if !moves.TestBit(i) {
                // board.DoMove() should not be called for invalid moves
                continue
            }
            clone := board.Clone()
            
            expected_return_val := clone.doMove(i)
            got_return_val := board.DoMove(i)
            if got_return_val != expected_return_val {
                t.Errorf("Got %d, expected %d.\n",got_return_val,expected_return_val) // TODO
            } 
            
            expected_board_val := clone
            got_board_val := board
            if got_board_val != expected_board_val {
                t.Errorf("") // TODO
            }
        }
    }
}

func (board *Board) moves() bitset.Bitset {
    moves := bitset.Bitset(0)
    for i:=uint(0); i<64; i++ {
        clone := board.Clone()
        if clone.DoMove(i) != bitset.Bitset(0) {
            moves |= bitset.Bitset(1 << i)
        }
    }
    return moves
}

func TestBoardMoves(t *testing.T) {
    for board := range genTestBoards() {
        clone := board
        expected := board.moves()
        got := board.Moves()
        if expected != got {
            t.Errorf("board.Moves() failed: expected %d, got %d\n",expected,got)
        }
        if clone != board {
            t.Errorf("board.Moves() changed the board!\n")
        }
    }
}

func (board *Board) getChildren() (children []Board) {
    children = make([]Board,10)
    for i:=uint(0); i<64; i++ {
        clone := board.Clone()
        if clone.doMove(i) != bitset.Bitset(0) {
            children = append(children,clone)
        }
    }
    return
}

func TestBoardGetChildren(t *testing.T) {
    for board := range genTestBoards() {

        expected := board.getChildren()
        expected_set := make(map[Board]struct{},10)
        for _,e := range expected {
            expected_set[e] = struct{}{}
        }

        got := board.GetChildren()
        got_set := make(map[Board]struct{},10)
        for _,g := range got {
            got_set[g] = struct{}{}
        }

        if len(got_set) != len(expected_set) {
            t.Errorf("Expected %d children, got %d.\n",len(expected_set),len(got_set))
        }

        equal_sets := true

        for e,_ := range expected_set {
            if _,ok := got_set[e]; !ok {
                equal_sets = false
            }
        }

        if !equal_sets {
            t.Errorf("Children sets are unequal.\n")
        } 
    }
}

func TestBoardAsciiArt(t *testing.T) {
    for board := range genTestBoards() {

        moves := board.Moves()

        ascii_art := board.AsciiArt()

        lines := strings.Split(ascii_art,"\n")

        expected := "+-a-b-c-d-e-f-g-h-+"
        if lines[0] != expected {
            t.Errorf("At lines[0]: expected '%s', got '%s'\n",expected,lines[0])
        }

        for y:=uint(0); y<8; y++ {

            expected_buf := new(bytes.Buffer)
            expected_buf.WriteString(fmt.Sprintf("%d ",y+1))

            for x:=uint(0); x<8; x++ {

                if board.me.TestBit(8*y + x) {
                    expected_buf.WriteString("○ ")
                } else if board.opp.TestBit(8*y + x) {
                    expected_buf.WriteString("● ")
                } else if moves.TestBit(8*y + x) {
                    expected_buf.WriteString("- ")
                } else {
                    expected_buf.WriteString("  ")
                }
            }

            expected_buf.WriteString("|")

            got := expected_buf.String()
            if lines[y+1] != got {
                t.Errorf("At lines[%d]: expected '%s', got '%s'\n",y+1,lines[y+1],got)
            }

        }

        expected = "+-----------------+"
        if lines[9] != expected {
            t.Errorf("At lines[9]: expected '%s', got '%s'\n",expected,lines[9])
        }

    }
}