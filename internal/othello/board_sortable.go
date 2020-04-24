package othello

// SortableBoard is a board with associated heuristic estimation suitable for sorting
type SortableBoard struct {
	Board Board
	Heur  int
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
