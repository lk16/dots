package othello

// UnsortedChildGenerator generates children in no particular order
type UnsortedChildGenerator struct {
	movesLeft   BitSet
	lastMove    BitSet
	lastFlipped BitSet
	child       *Board
}

// NewUnsortedChildGenerator returns a child generator for a parent Board
func NewUnsortedChildGenerator(board *Board) UnsortedChildGenerator {
	return UnsortedChildGenerator{
		movesLeft:   board.Moves(),
		lastMove:    0,
		lastFlipped: 0,
		child:       board,
	}
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

	gen.lastMove = gen.movesLeft & (-gen.movesLeft)
	gen.lastFlipped = gen.child.DoMove(gen.lastMove)
	gen.movesLeft &^= gen.lastMove
	return true
}

// RestoreParent restores the parent state
func (gen *UnsortedChildGenerator) RestoreParent() {
	gen.child.UndoMove(gen.lastMove, gen.lastFlipped)
}
