package players

import(
	"dots/board"
)


// SearchQuery is a query for searching the best child of a Board
type SearchQuery struct {
	board      board.Board
	lowerBound int
	upperBound int
	depth      int
	guess      int
	heuristic  Heuristic
}

// SearchResult is a result of a SearchQuery
type SearchResult struct {
	query *SearchQuery
	stats *SearchStats
	heur  int
}

// SearchStats contains statistics of a SearchQuery
type SearchStats struct {
	nodes  uint64
	timeNs uint64
}

// value type for SearchState.transpositionTable
type tptValue struct {
	high int
	low  int
}

// SearchState is the state of a SearchQuery
type SearchState struct {
	board              board.Board
	skipped            bool
	transpositionTable map[board.Board]tptValue
}

// SearchThread contains all thread data for a SearchQuery
type SearchThread struct {
	query *SearchQuery
	state *SearchState
	stats *SearchStats
}

// Run runs a SearchQuery using a new SearchThread
func (query *SearchQuery) Run(ch chan SearchResult) {

	thread := &SearchThread{
		query: query,
		state: &SearchState{
			board:              query.board,
			skipped:            false,
			transpositionTable: make(map[board.Board]tptValue, 50000)},
		stats: &SearchStats{
			nodes:  0,
			timeNs: 0}}

	go thread.Run(ch)
}

// Run runs a SearchThread and sends its SearchResult over a channel
func (thread *SearchThread) Run(ch chan SearchResult) {
	result := SearchResult{
		query: thread.query,
		heur:  0,
		stats: thread.stats}

	// copy values because thread.query should not be modified
	high := thread.query.upperBound
	low := thread.query.lowerBound

	f := thread.query.guess

	var step int
	if thread.query.board.CountEmpties() > thread.query.depth {
		step = 1
	} else {
		step = 2
	}

	// prevent odd results for exact search
	f -= (f % step)

	f = clamp(f, low, high)

	for high-low >= step {
		var bound int
		if thread.query.board.CountEmpties() > thread.query.depth {
			bound = -thread.doMtdf(-(f + 1), thread.query.depth)
		} else {
			bound = -thread.doMtdfExact(-(f + 1))
		}

		if f == bound {
			f -= step
			high = bound
		} else {
			f += step
			low = bound
		}
	}
	result.heur = high

	ch <- result
}


func (thread *SearchThread) updateTranspositionTable(heur,alpha int) {

	b := thread.state.board

	entry, ok := thread.state.transpositionTable[b]

	if !ok {
		entry = tptValue{
			low:  board.MinHeuristic,
			high: board.MaxHeuristic}
	}

	if heur == alpha {
		entry.high = min(alpha, entry.high)
	} else {
		entry.low = max(alpha+1, entry.low)
	}
	thread.state.transpositionTable[b] = entry
}

func (thread *SearchThread) checkTranspositionTable(alpha int) (cutOff int,ok bool) {

	lookup, ok := thread.state.transpositionTable[thread.state.board]

	if !ok {
		return 0,false
	}
	if lookup.high < alpha {
		return alpha, true
	}
	if lookup.low > alpha+1 {
		return alpha + 1,true
	}

	return 0, false

}

func (thread *SearchThread) doMtdf(alpha, depth int) (heur int) {

	thread.stats.nodes++
	b := thread.state.board

	if depth == 0 {
		return mtdfPolish(thread.query.heuristic(b), alpha)
	}

	if depth >= 5 {
		if cutOff, ok := thread.checkTranspositionTable(alpha); ok {
			return cutOff
		}

		defer thread.updateTranspositionTable(heur, alpha)
	}

	// BUG: board.NewGenerator is broken with lookAhead > 0
	gen := board.NewGenerator(&thread.state.board, 0) //depth/3)

	if !gen.HasMoves() {

		if thread.state.skipped {
			return mtdfPolish(board.ExactScoreFactor*b.ExactScore(), alpha)
		}

		thread.state.skipped = true
		b.SwitchTurn()
		heur = -thread.doMtdf(-(alpha + 1), depth)
		b.SwitchTurn()
		return
	}

	thread.state.skipped = false

	for gen.Next() {
		childHeur := -thread.doMtdf(-(alpha + 1), depth-1)
		if childHeur > alpha {
			gen.RestoreParent()
			return alpha + 1
		}
	}
	return alpha
}

func (thread *SearchThread) doMtdfExact(alpha int) (heur int) {

	thread.stats.nodes++

	// BUG: board.NewGenerator is broken with lookAhead > 0
	// emptiesCount := thread.state.board.CountEmpties()
	gen := board.NewGenerator(&thread.state.board, 0) //emptiesCount/4)

	if gen.HasMoves() {
		heur = alpha
		for gen.Next() {
			childHeur := -thread.doMtdfExact(-(alpha + 1))
			if childHeur > alpha {
				heur = alpha + 1
				gen.RestoreParent()
				break
			}
		}
		return
	}

	if thread.state.board.OpponentMoves() != 0 {
		thread.state.board.SwitchTurn()
		heur = -thread.doMtdfExact(-(alpha + 1))
		thread.state.board.SwitchTurn()
		return
	}

	heur = mtdfPolish(thread.state.board.ExactScore(), alpha)
	return
}

func mtdfPolish(heur, alpha int) int {
	if heur > alpha {
		return alpha + 1
	}
	return alpha
}
