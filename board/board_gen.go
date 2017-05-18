package board

import (
	"dots/bitset"
)

type ChildGenerator interface {
	HasMoves() bool
	Next() bool
	RestoreParent()
}

type UnsortedChildGenerator struct {
	moves_left   bitset.Bitset
	last_move    bitset.Bitset
	last_flipped bitset.Bitset
	child        *Board
}

// Returns a child generator for a Board
func NewChildGen(board *Board) (gen *UnsortedChildGenerator) {
	gen = &UnsortedChildGenerator{
		moves_left:   board.Moves(),
		last_move:    bitset.Bitset(0),
		last_flipped: bitset.Bitset(0),
		child:        board}
	return
}

func (gen *UnsortedChildGenerator) HasMoves() (has_moves bool) {
	has_moves = (gen.moves_left != 0)
	return
}

// Generate next child of a Board
// After generating all children the parent state is restored
func (gen *UnsortedChildGenerator) Next() (ok bool) {

	if gen.last_flipped != 0 {
		gen.RestoreParent()
	}

	if gen.moves_left == 0 {
		ok = false
		return
	}

	index := gen.moves_left.FirstBitIndex()
	gen.moves_left.ResetBit(index)

	gen.last_flipped = gen.child.DoMove(index)
	gen.last_move = bitset.Bitset(1 << index)

	ok = true
	return
}

// Force restore parent state
// This is usefull when not visting all children
func (gen *UnsortedChildGenerator) RestoreParent() {
	gen.child.UndoMove(gen.last_move, gen.last_flipped)
}
