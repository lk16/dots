package board

import (
	"math/bits"
	"sort"
)

// ChildGenerator is an interface for child generators
type ChildGenerator interface {
	HasMoves() bool
	Next() bool
	RestoreParent()
}

// UnsortedChildGenerator generates children in no particular order
type UnsortedChildGenerator struct {
	movesLeft   uint64
	lastMove    uint64
	lastFlipped uint64
	child       *Board
}

type sortedBoard struct {
	board Board
	heur  int
}

// SortedChildGenerator generates children sorted by their heuristic values
type SortedChildGenerator struct {
	parent     Board
	child      *Board
	childIndex int
	children   []sortedBoard
}

// NewGenerator returns a child generator for a parent Board
func NewGenerator(board *Board, lookAhead int) ChildGenerator {

	unsortedGen := &UnsortedChildGenerator{
		movesLeft:   board.Moves(),
		lastMove:    0,
		lastFlipped: 0,
		child:       board}

	if lookAhead == 0 {
		return unsortedGen
	}

	var sortedChildren []sortedBoard

	for _, child := range board.GetChildren() {
		sortedChild := sortedBoard{
			board: child,
			heur:  -Negamax(&child, lookAhead)}
		sortedChildren = append(sortedChildren, sortedChild)
	}

	sort.Slice(sortedChildren, func(i, j int) bool {
		return sortedChildren[i].heur > sortedChildren[j].heur
	})

	sortedGen := &SortedChildGenerator{
		parent:     *board,
		child:      board,
		children:   sortedChildren,
		childIndex: 0}

	return sortedGen
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

// RestoreParent restores the state of the parent
func (gen *SortedChildGenerator) RestoreParent() {
	*gen.child = gen.parent
}

// Next attempts to generate a child of a Board
// After generating all children the parent state is restored
// If no children are left, false is returned. Otherwise true is returned.
func (gen *SortedChildGenerator) Next() bool {
	if gen.childIndex == len(gen.children) {
		gen.RestoreParent()
		return false
	}

	*gen.child = gen.children[gen.childIndex].board
	gen.childIndex++
	return true
}

// HasMoves returns whether the parent Board has moves
func (gen *SortedChildGenerator) HasMoves() bool {
	return len(gen.children) != 0
}
