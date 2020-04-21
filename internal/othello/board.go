// Package othello contains the implementation of the othello rules including boards and computing available moves.
package othello

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

const (
	cornerMask   = BitSet(1<<0 | 1<<7 | 1<<56 | 1<<63)
	xSquareMask  = BitSet(1<<9 | 1<<14 | 1<<49 | 1<<54)
	cSquareMask  = BitSet(1<<1 | 1<<6 | 1<<8 | 1<<15 | 1<<48 | 1<<55 | 1<<57 | 1<<62)
	bSquareMask  = BitSet(1<<3 | 1<<4 | 1<<24 | 1<<31 | 1<<32 | 1<<39 | 1<<59 | 1<<60)
	aSquareMask  = BitSet(1<<2 | 1<<5 | 1<<16 | 1<<23 | 1<<40 | 1<<47 | 1<<58 | 1<<61)
	center16Mask = BitSet(0x00003C3C3C3C0000)
	ring16Mask   = BitSet(0x003C424242423C00)

	startDiscsMask = BitSet(1<<27 | 1<<28 | 1<<35 | 1<<36)
)

var (
	// ErrInvalidDiscAmount is used when an unexpected amount of discs is requested
	ErrInvalidDiscAmount = errors.New("cannot create board with requested amount of discs")

	xotBoards  []Board
	assetsPath = os.Getenv("DOTS_ASSETS_PATH")
)

// Board represents the state of an othello othello game.
// It does not keep track which discs are white or black.
// Instead it keeps track which discs are owned by the player to move.
type Board struct {
	me, opp BitSet
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

// NewCustomBoard returns a Board with a custom state
func NewCustomBoard(me, opp BitSet) (board *Board) {
	return &Board{
		me:  me,
		opp: opp}
}

// NewRandomBoard returns a random Board with a given number of discs
func NewRandomBoard(discs int) (*Board, error) {
	if discs < 4 || discs > 64 {
		return nil, ErrInvalidDiscAmount
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

		if board.Moves() == 0 {
			skips++
			board.SwitchTurn()
			continue
		}

		skips = 0
		board.DoRandomMove()
	}

	return board, nil
}

// UnmarshalJSON loads a board from JSON data
func (board *Board) UnmarshalJSON(bytes []byte) error {
	type boardModel struct {
		Me  string `json:"me"`
		Opp string `json:"opp"`
	}

	var (
		model   boardModel
		err     error
		me, opp uint64
	)

	if err = json.Unmarshal(bytes, &model); err != nil {
		return err
	}

	if me, err = strconv.ParseUint(model.Me, 0, 64); err != nil {
		return err
	}

	if opp, err = strconv.ParseUint(model.Opp, 0, 64); err != nil {
		return err
	}

	*board = *NewCustomBoard(BitSet(me), BitSet(opp))
	return nil
}

// GetSortableChildren returns a slice with all children of a Board
// such that they can easily be sorted
func (board Board) GetSortableChildren() []SortableBoard {
	moves := board.Moves()
	children := make([]SortableBoard, moves.Count())

	for i := range children {
		moveBit := moves & (-moves)
		moves &^= moveBit

		children[i].Board = board
		children[i].Board.DoMove(moveBit)
		children[i].Heur = 0
	}
	return children
}

// GetMoveField computes the index of the move given a child and a parent
func (board Board) GetMoveField(child Board) (int, bool) {
	moveBit := board.Me() | board.Opp() ^ (child.Me() | child.Opp())
	if moveBit.Count() != 1 {
		return 0, false
	}
	return moveBit.Lowest(), true
}

// Normalize returns a normalized othello with regards to symmetry
func (board Board) Normalize() Board {
	mirrorHor := func(bitset BitSet) BitSet {
		result := bitset
		result = (result&0x00000000FFFFFFFF)<<32 | (result&0xFFFFFFFF00000000)>>32
		result = (result&0x0000FFFF0000FFFF)<<16 | (result&0xFFFF0000FFFF0000)>>16
		result = (result&0x00FF00FF00FF00FF)<<8 | (result&0xFF00FF00FF00FF00)>>8
		return result
	}

	mirrorVer := func(bitset BitSet) BitSet {
		result := bitset
		result = (result&0x0F0F0F0F0F0F0F0F)<<4 | (result&0xF0F0F0F0F0F0F0F0)>>4
		result = (result&0x3333333333333333)<<2 | (result&0xCCCCCCCCCCCCCCCC)>>2
		result = (result&0x5555555555555555)<<1 | (result&0xAAAAAAAAAAAAAAAA)>>1
		return result
	}

	mirrorDia := func(bitset BitSet) BitSet {
		var tmp BitSet
		result := bitset
		k1 := BitSet(0xaa00aa00aa00aa00)
		k2 := BitSet(0xcccc0000cccc0000)
		k4 := BitSet(0xf0f0f0f00f0f0f0f)
		tmp = result ^ (result << 36)
		result ^= k4 & (tmp ^ (result >> 36))
		tmp = k2 & (result ^ (result << 18))
		result ^= tmp ^ (tmp >> 18)
		tmp = k1 & (result ^ (result << 9))
		result ^= tmp ^ (tmp >> 9)
		return result
	}

	less := func(lhs, rhs Board) bool {
		if lhs.me < rhs.me {
			return true
		}
		return lhs.me == rhs.me && lhs.opp < rhs.opp
	}

	lowest := board

	board = Board{
		me:  mirrorHor(board.me),
		opp: mirrorHor(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	board = Board{
		me:  mirrorVer(board.me),
		opp: mirrorVer(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	board = Board{
		me:  mirrorHor(board.me),
		opp: mirrorHor(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	board = Board{
		me:  mirrorDia(board.me),
		opp: mirrorDia(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	board = Board{
		me:  mirrorHor(board.me),
		opp: mirrorHor(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	board = Board{
		me:  mirrorVer(board.me),
		opp: mirrorVer(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	board = Board{
		me:  mirrorHor(board.me),
		opp: mirrorHor(board.opp),
	}

	if less(board, lowest) {
		lowest = board
	}

	return lowest
}

// String returns an ASCII-art string representation of a board
func (board Board) String() string {
	buffer := new(bytes.Buffer)
	_, _ = buffer.WriteString("+-a-b-c-d-e-f-g-h-+\n")

	moves := board.Moves()

	for y := uint(0); y < 8; y++ {
		_, _ = buffer.WriteString(fmt.Sprintf("%d ", y+1))

		for x := uint(0); x < 8; x++ {
			mask := BitSet(1) << (y*8 + x)
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

	_, _ = buffer.WriteString("+-----------------+\n")
	_, _ = buffer.WriteString("To move: ○\n")
	_, _ = buffer.WriteString("Raw: " + fmt.Sprintf("%#v", board) + "\n")

	return buffer.String()
}

// Moves returns a bitset of valid moves for a Board
func (board Board) Moves() BitSet {
	return moves(board.me, board.opp)
}

// OpponentMoves returns a bitset with all valid moves for the opponent
func (board Board) OpponentMoves() BitSet {
	return moves(board.opp, board.me)
}

func moves(me, opp BitSet) BitSet {
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

	movesSet &^= me | opp | startDiscsMask
	return movesSet
}

// GetChildren returns a slice with all children of a Board
func (board Board) GetChildren() []Board {
	moves := board.Moves()
	children := make([]Board, moves.Count())

	for i := range children {
		moveBit := moves & (-moves)
		moves &^= moveBit

		children[i] = board
		children[i].DoMove(moveBit)
	}
	return children
}

// UndoMove undoes a move
func (board *Board) UndoMove(moveBit, flipped BitSet) {
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
	return (board.me | board.opp).Count()
}

// CornerCountDifference returns the corner count difference.
// Positive result means player to move has more corners.
func (board Board) CornerCountDifference() int {
	return (board.me & cornerMask).Count() - (board.opp & cornerMask).Count()
}

// XsquareCountDifference returns the x-square count difference
// X-squares are fields fields diagonal to a corner.
// Positive result means player to move has more x-squares.
func (board Board) XsquareCountDifference() int {
	return (board.me & xSquareMask).Count() - (board.opp & xSquareMask).Count()
}

// CsquareCountDifference returns the c-square count difference
// C-squares are fields on side of the board next to a corner.
// Positive result means player to move has more c-squares.
func (board Board) CsquareCountDifference() int {
	return (board.me & cSquareMask).Count() - (board.opp & cSquareMask).Count()
}

func potentialMoves(me, opp BitSet) BitSet {
	const (
		leftMask  = 0x7F7F7F7F7F7F7F7F
		rightMask = 0xFEFEFEFEFEFEFEFE
	)

	oppSurrounded := BitSet(0)
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
	mePotentialMoveCount := potentialMoves(board.me, board.opp).Count()
	oppPotentialMoveCount := potentialMoves(board.opp, board.me).Count()
	return mePotentialMoveCount - oppPotentialMoveCount
}

// CountEmpties returns the number of empty fields on a Board
func (board Board) CountEmpties() int {
	return 64 - board.CountDiscs()
}

// ExactScore returns the final score of a Board
func (board Board) ExactScore() int {
	meCount := board.me.Count()
	oppCount := board.opp.Count()

	if meCount > oppCount {
		return 64 - (2 * oppCount)
	}
	if meCount < oppCount {
		return -64 + (2 * meCount)
	}
	return 0
}

// Me returns a bitset with the discs of the player to move
func (board Board) Me() BitSet {
	return board.me
}

// Opp returns a bitset with the discs of the opponent of the player to move
func (board Board) Opp() BitSet {
	return board.opp
}

// LoadXotBoards loads the xot boards into memory.
// When calling this again after a successful call, this function does nothing.
func LoadXotBoards() error {
	if len(xotBoards) != 0 {
		return nil
	}

	var bytes []byte
	var err error

	xotDataPath := assetsPath + "xot.json"

	if bytes, err = ioutil.ReadFile(xotDataPath); err != nil {
		return errors.Wrap(err, "failed to load xot file")
	}

	if err = json.Unmarshal(bytes, &xotBoards); err != nil {
		// json.Unmarshal may set xotBoards to a non-empty slice when it returns an error
		xotBoards = nil

		return errors.Wrap(err, "failed to parse xot file")
	}

	return nil
}

// NewXotBoard returns a random xot board
// http://berg.earthlingz.de/xot/aboutxot.php?lang=en
func NewXotBoard() *Board {
	if len(xotBoards) == 0 {
		panic("xot boards are not loaded")
	}

	return &xotBoards[rand.Intn(len(xotBoards))]
}

// ChildGenerator generates children of a board in no particular order
type ChildGenerator struct {
	movesLeft   BitSet
	lastMove    BitSet
	lastFlipped BitSet
	child       *Board
}

// NewChildGenerator returns an child generator for a parent Board.
// Generated children are not sorted
func NewChildGenerator(board *Board) ChildGenerator {
	return ChildGenerator{
		movesLeft:   board.Moves(),
		lastMove:    0,
		lastFlipped: 0,
		child:       board,
	}
}

// HasMoves returns whether the parent Board has moves
func (gen *ChildGenerator) HasMoves() bool {
	return gen.movesLeft != 0
}

// Next attempts to generate a child of a Board
// After generating all children the parent state is restored
// The returned value indicates if more children are available.
func (gen *ChildGenerator) Next() bool {
	if gen.lastFlipped != 0 {
		gen.RestoreParent()
	}

	if gen.movesLeft == 0 {
		return false
	}

	gen.lastMove = gen.movesLeft & (-gen.movesLeft)
	gen.lastFlipped = gen.child.DoMove(gen.lastMove)
	gen.movesLeft &^= gen.lastMove
	return true
}

// RestoreParent restores the parent state
func (gen *ChildGenerator) RestoreParent() {
	gen.child.UndoMove(gen.lastMove, gen.lastFlipped)
}
