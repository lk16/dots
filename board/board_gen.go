package board

import (
	"math/bits"
	"sort"
)

type ChildGenerator interface {
	HasMoves() bool
	Next() bool
	RestoreParent()
}

type UnsortedChildGenerator struct {
	moves_left   uint64
	last_move    uint64
	last_flipped uint64
	child        *Board
}

// Returns a child generator for a Board
func NewChildGen(board *Board) (gen *UnsortedChildGenerator) {
	gen = &UnsortedChildGenerator{
		moves_left:   board.Moves(),
		last_move:    0,
		last_flipped: 0,
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

	index := bits.TrailingZeros64(gen.moves_left)

	gen.last_move = uint64(1) << uint(index)
	gen.last_flipped = gen.child.DoMove(index)
	gen.moves_left &^= gen.last_move

	ok = true
	return
}

// Force restore parent state
// This is usefull when not visting all children
func (gen *UnsortedChildGenerator) RestoreParent() {
	gen.child.UndoMove(gen.last_move, gen.last_flipped)
}

type sortedBoard struct {
	board Board
	heur  int
}

type SortedChildGenerator struct {
	parent      Board
	child       *Board
	child_index int
	children    []sortedBoard
}

func NewChildGenSorted(board *Board, heuristic func(Board) int) (gen *SortedChildGenerator) {
	gen = &SortedChildGenerator{
		parent:      *board,
		child:       board,
		children:    []sortedBoard{},
		child_index: 0}

	child := *board

	unsorted_gen := NewChildGen(&child)
	for unsorted_gen.Next() {
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

func (gen *SortedChildGenerator) RestoreParent() {
	*gen.child = gen.parent
}

func (gen *SortedChildGenerator) Next() (ok bool) {
	if gen.child_index == len(gen.children) {
		gen.RestoreParent()
		ok = false
		return
	}

	*gen.child = gen.children[gen.child_index].board
	gen.child_index++
	ok = true
	return
}

func (gen *SortedChildGenerator) HasMoves() (has_moves bool) {
	has_moves = len(gen.children) != 0
	return
}
