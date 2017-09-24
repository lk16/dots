package players

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"dots/board"
)

const (
	// MaxScore is the highest game result score possible
	MaxScore = 64

	// MinScore is the lowest game result score possible
	MinScore = -MaxScore

	// ExactScoreFactor is the multiplication.
	// This is used when a non exact search runs into an exact result
	ExactScoreFactor = 1000

	// MaxHeuristic is the highest heuristic value possible
	MaxHeuristic = ExactScoreFactor * MaxScore

	// MinHeuristic is the lowest heuristic value possible
	MinHeuristic = ExactScoreFactor * MinScore
)

// Heuristic is a function that estimates how promising a Board is.
type Heuristic func(board.Board) int

// BotHeuristic is a bot that uses a Heuristic for choosing its moves
type BotHeuristic struct {
	heuristic   Heuristic
	searchDepth int
	exactDepth  int
	writer      io.Writer
	resultChan  chan SearchResult
}

// NewBotHeuristic creates a new BotHeuristic
func NewBotHeuristic(heuristic Heuristic,
	searchDepth, exactDepth int, writer io.Writer) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:   heuristic,
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer,
		resultChan:  make(chan SearchResult, 32)}
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

func (bot *BotHeuristic) logChildEvaluation(childID, heur, alpha int,
	childStats, totalStats SearchStats) {

	str := fmt.Sprintf("%5d | ", childID+1)
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

	avgSpeed := safeDiv(1000000000*totalStats.nodes, totalStats.timeNs)
	childSpeed := safeDiv(1000000000*childStats.nodes, childStats.timeNs)

	buff.WriteString(fmt.Sprintf("%s | %s | %s || %s | %s | %s |\n",
		fmtBig(childStats.nodes), fmtNs(childStats.timeNs), fmtBig(childSpeed),
		fmtBig(totalStats.nodes), fmtNs(totalStats.timeNs), fmtBig(avgSpeed)))
	bot.writer.Write(buff.Bytes())
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
		alpha = MinScore
		beta = MaxScore
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

	totalStats := SearchStats{
		nodes:  0,
		timeNs: 0}

	startTime := time.Now()

	processResult := func(childID int) {
		result := <-bot.resultChan

		childStats := result.stats
		totalStats.nodes += childStats.nodes
		totalStats.timeNs = uint64(time.Since(startTime).Nanoseconds())

		bot.logChildEvaluation(childID, result.heur, alpha, *childStats, totalStats)
		if result.heur > alpha {
			alpha = result.heur
			afterwards = result.query.board
		}

	}

	for i, child := range children {

		query := SearchQuery{
			board:      child,
			lowerBound: alpha,
			upperBound: beta,
			depth:      depth,
			guess:      0,
			heuristic:  bot.heuristic}

		query.Run(bot.resultChan)

		// evaluate first child before launching other child threads
		if i == 0 {
			processResult(0)
		}

	}

	for i := int(1); i < len(children); i++ {
		processResult(i)
	}

	bot.writer.Write(bytes.NewBufferString("\n\n").Bytes())
	return
}

func clamp(x, min, max int) int {
	if x > max {
		return max
	}
	if x < min {
		return min
	}
	return x
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

// SearchState is the state of a SearchQuery
type SearchState struct {
	depth int
	board board.Board
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
			depth: query.depth,
			board: query.board},
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

	if thread.query.board.CountEmpties() > thread.query.depth {
		result.heur = thread.loop(1, thread.doMtdf)
	} else {
		result.heur = thread.loop(2, thread.doMtdfExact)
	}
	ch <- result
}

func (thread *SearchThread) loop(step int, call func(int) int) (heur int) {

	// copy values because thread.query should not be modified
	high := thread.query.upperBound
	low := thread.query.lowerBound

	f := thread.query.guess

	// prevent odd results for exact search
	f -= (f % step)

	f = clamp(f, low, high)

	for high-low >= step {
		bound := -call(-(f + 1))
		if f == bound {
			f -= step
			high = bound
		} else {
			f += step
			low = bound
		}
	}
	heur = high
	return
}

func (thread *SearchThread) doMtdf(alpha int) (heur int) {

	thread.stats.nodes++

	if thread.state.depth == 0 {
		heur = mtdfPolish(thread.query.heuristic(thread.state.board), alpha)
		return
	}

	gen := board.NewChildGen(&thread.state.board)

	if gen.HasMoves() {
		thread.state.depth--
		heur = alpha
		for gen.Next() {
			childHeur := -thread.doMtdf(-(alpha + 1))
			if childHeur > alpha {
				heur = alpha + 1
				gen.RestoreParent()
				break
			}
		}
		thread.state.depth++
		return
	}

	if thread.state.board.OpponentMoves() != 0 {
		thread.state.board.SwitchTurn()
		heur = -thread.doMtdf(-(alpha + 1))
		thread.state.board.SwitchTurn()
		return
	}

	heur = mtdfPolish(ExactScoreFactor*thread.state.board.ExactScore(), alpha)
	return
}

func (thread *SearchThread) doMtdfExact(alpha int) (heur int) {

	thread.stats.nodes++

	gen := board.NewChildGen(&thread.state.board)

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
