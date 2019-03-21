package othello

import (
	"math/bits"
)

// UnsortedChildGenerator generates children in no particular order
type UnsortedChildGenerator struct {
	movesLeft   uint64
	lastMove    uint64
	lastFlipped uint64
	child       *Board
}

// NewUnsortedChildGenerator returns a child generator for a parent Board
func NewUnsortedChildGenerator(board *Board) UnsortedChildGenerator {

	return UnsortedChildGenerator{
		movesLeft:   board.Moves(),
		lastMove:    0,
		lastFlipped: 0,
		child:       board}

}

// HasMoves returns whether the parent Board has moves
func (gen *UnsortedChildGenerator) HasMoves() bool {
	return gen.movesLeft != 0
}

// Next attempts to generate a child of a Board
// After generating all children the parent state is restored
// If no children are left, false is returned. Otherwise true is returned.
func (gen *UnsortedChildGenerator) Next() bool {

	if gen.lastFlipped != 0 {
		gen.RestoreParent()
	}

	if gen.movesLeft == 0 {
		return false
	}

	index := bits.TrailingZeros64(gen.movesLeft)

	gen.lastMove = uint64(1) << uint(index)
	gen.lastFlipped = gen.child.DoMove(index)
	gen.movesLeft &^= gen.lastMove

	return true
}

// RestoreParent restores the parent state
func (gen *UnsortedChildGenerator) RestoreParent() {
	gen.child.UndoMove(gen.lastMove, gen.lastFlipped)
}
