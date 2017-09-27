package players

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"dots/board"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func clamp(x, low, high int) int {
	return max(low, min(x, high))
}

// Heuristic is a function that estimates how promising a Board is.
type Heuristic func(board.Board) int

// BotHeuristic is a bot that uses a Heuristic for choosing its moves
type BotHeuristic struct {
	heuristic   Heuristic
	searchDepth int
	exactDepth  int
	writer      io.Writer
	resultChan  chan SearchResult
	stats       SearchStats
	startTime   time.Time
}

// NewBotHeuristic creates a new BotHeuristic
func NewBotHeuristic(heuristic Heuristic,
	searchDepth, exactDepth int, writer io.Writer) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:   heuristic,
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer,
		resultChan:  make(chan SearchResult, 32),
		stats: SearchStats{
			nodes:  0,
			timeNs: 0}}
	return
}

func fmtBig(n uint64) string {

	if n < 10000 {
		return fmt.Sprintf(" %4d", n)
	}

	suffixes := "KMGTPE"

	n /= 1000
	suffixIndex := 0

	for (n > 10000) && (suffixIndex < len(suffixes)-1) {
		n /= 1000
		suffixIndex++
	}

	return fmt.Sprintf("%4d%c", n, suffixes[suffixIndex])
}

func fmtNs(n uint64) string {

	suffixIndex := 0

	suffixes := []string{"n", "μ", "m", " "}

	for (n > 10000) && (suffixIndex < len(suffixes)-1) {
		n /= 1000
		suffixIndex++
	}

	return fmt.Sprintf("%5d%ss", n, suffixes[suffixIndex])
}

func (bot *BotHeuristic) logChildEvaluation(heur, alpha int,
	childStats SearchStats) {

	str := "      | "
	buff := bytes.NewBufferString(str)
	if heur > alpha {
		buff.WriteString(fmt.Sprintf("%5d || ", heur))
	} else {
		buff.WriteString(fmt.Sprintf("      || "))
	}

	safeDiv := func(num, den uint64) uint64 {
		if den == 0 {
			return 0
		}
		return num / den
	}

	avgSpeed := safeDiv(1000000000*bot.stats.nodes, bot.stats.timeNs)
	childSpeed := safeDiv(1000000000*childStats.nodes, childStats.timeNs)

	buff.WriteString(fmt.Sprintf("%s | %s | %s || %s | %s | %s |\n",
		fmtBig(childStats.nodes), fmtNs(childStats.timeNs), fmtBig(childSpeed),
		fmtBig(bot.stats.nodes), fmtNs(bot.stats.timeNs), fmtBig(avgSpeed)))
	bot.writer.Write(buff.Bytes())
}

func (bot *BotHeuristic) processResult(result SearchResult, alpha *int,
	afterwards *board.Board) {

	childStats := result.stats
	bot.stats.nodes += childStats.nodes
	bot.stats.timeNs = uint64(time.Since(bot.startTime).Nanoseconds())

	bot.logChildEvaluation(result.heur, *alpha, *childStats)

	if result.heur > *alpha {
		*alpha = result.heur
		*afterwards = result.query.board
	}

}

// DoMove does a move
func (bot *BotHeuristic) DoMove(b board.Board) (afterwards board.Board) {

	children := b.GetChildren()

	if len(children) == 0 {
		return b
	}

	// prevent returning empty board when bot cannot prevent losing all discs
	afterwards = children[0]

	if len(children) == 1 {
		buff := bytes.NewBufferString("Only one move. Skipping evaluation.\n")
		bot.writer.Write(buff.Bytes())
		return
	}

	var alpha, beta int
	var depth int

	if b.CountEmpties() <= bot.exactDepth {
		alpha = board.MinScore
		beta = board.MaxScore
		depth = b.CountEmpties()
	} else {
		depth = bot.searchDepth

		// HACK: stumbling upon an exact solution
		// takes forever to compute. we set limits to solve that for now.
		alpha = -100
		beta = 100
	}

	header := "      | heuri || child |  child  | child || total |  total  |  avg  |\n"
	header += " move | stic  || nodes |   time  | speed || nodes |   time  | speed |\n"
	header += "------|-------||-------|---------|-------||-------|---------|-------|\n"

	bot.writer.Write(bytes.NewBufferString(header).Bytes())

	bot.startTime = time.Now()
	bot.stats.nodes = 0
	bot.stats.timeNs = 0

	child := b
	gen := board.NewGenerator(&child, 1)

	for i := 0; gen.Next(); i++ {

		query := SearchQuery{
			board:      child,
			lowerBound: alpha,
			upperBound: beta,
			depth:      depth,
			guess:      0,
			heuristic:  bot.heuristic}

		query.Run(bot.resultChan)

		//if i == 0 {
		bot.processResult(<-bot.resultChan, &alpha, &afterwards)
		//}
	}

	/*for i := 1; i < bits.OnesCount64(b.Moves()); i++ {
		bot.processResult(<-bot.resultChan, &alpha, &afterwards)
	}*/

	bot.writer.Write(bytes.NewBufferString("\n\n").Bytes())
	return
}

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

func (thread *SearchThread) doMtdf(alpha, depth int) (heur int) {

	thread.stats.nodes++
	b := thread.state.board

	if depth == 0 {
		return mtdfPolish(thread.query.heuristic(b), alpha)
	}

	if depth >= 5 {

		if lookup, ok := thread.state.transpositionTable[b]; ok {
			if lookup.high < alpha {
				return alpha
			}
			if lookup.low > alpha+1 {
				return alpha + 1
			}
		}

		defer func() {
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

		}()
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
