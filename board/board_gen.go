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

// NewChildGen returns a child generator for a parent Board
func NewChildGen(board *Board) (gen *UnsortedChildGenerator) {
	gen = &UnsortedChildGenerator{
		movesLeft:   board.Moves(),
		lastMove:    0,
		lastFlipped: 0,
		child:       board}
	return
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

// NewChildGenSorted returns a new SortedChildGenerator
func NewChildGenSorted(board *Board,
	heuristic func(Board) int) (gen *SortedChildGenerator) {

	gen = &SortedChildGenerator{
		parent:     *board,
		child:      board,
		children:   []sortedBoard{},
		childIndex: 0}

	child := *board

	unsortedGen := NewChildGen(&child)
	for unsortedGen.Next() {
		gen.children = append(gen.children, sortedBoard{
			board: child,
			heur:  heuristic(child),
		})
	}

	sort.Slice(gen.children, func(i, j int) bool {
		return gen.children[i].heur > gen.children[j].heur
	})

	return
}

// RestoreParent restores the state of the parent
func (gen *SortedChildGenerator) RestoreParent() {
	*gen.child = gen.parent
}

// Next attempts to generate a child of a Board
// After generating all children the parent state is restored
// If no children are left, false is returned. Otherwise true is returned.
func (gen *SortedChildGenerator) Next() (ok bool) {
	if gen.childIndex == len(gen.children) {
		gen.RestoreParent()
		ok = false
		return
	}

	*gen.child = gen.children[gen.childIndex].board
	gen.childIndex++
	ok = true
	return
}

// HasMoves returns whether the parent Board has moves
func (gen *SortedChildGenerator) HasMoves() bool {
	return len(gen.children) != 0
}
