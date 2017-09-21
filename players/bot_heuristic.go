package players

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"dots/board"
)

const (
	Max_exact_heuristic = 64
	Min_exact_heuristic = -Max_exact_heuristic
	Exact_score_factor  = 1000
	Max_heuristic       = Exact_score_factor * Max_exact_heuristic
	Min_heuristic       = Exact_score_factor * Min_exact_heuristic
)

type Heuristic func(board board.Board) (heur int)

type Interface interface {
	Search(board board.Board, depth_left uint, heuristic Heuristic, alpha int) (heur int)
	ExactSearch(board board.Board, alpha int) (heur int)
	Name() (name string)
	Nodes() (nodes uint64)
	ComputeTimeNs() (ns uint64)
}

type SortedBoard struct {
	board board.Board
	heur  int
}

type BotHeuristic struct {
	heuristic    Heuristic
	search_depth uint
	exact_depth  uint
	writer       io.Writer
	result_chan  chan SearchResult
}

// Creates a new BotHeuristic
func NewBotHeuristic(heuristic Heuristic,
	search_depth, exact_depth uint, writer io.Writer) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:    heuristic,
		search_depth: search_depth,
		exact_depth:  exact_depth,
		writer:       writer,
		result_chan:  make(chan SearchResult, 32)}
	return
}

func fmt_big(n uint64) string {

	if n < 10000 {
		return fmt.Sprintf(" %4d", n)
	}

	suffixes := "KMGTPE"

	n /= 1000
	suffix_index := 0

	for (n > 10000) && (suffix_index < len(suffixes)-1) {
		n /= 1000
		suffix_index++
	}

	return fmt.Sprintf("%4d%c", n, suffixes[suffix_index])
}

func fmt_ns(n uint64) string {

	suffix_index := 0

	suffixes := []string{"n", "μ", "m", " "}

	for (n > 10000) && (suffix_index < len(suffixes)-1) {
		n /= 1000
		suffix_index++
	}

	return fmt.Sprintf("%5d%ss", n, suffixes[suffix_index])
}

func (bot *BotHeuristic) logChildEvaluation(child_id, heur, alpha int,
	child_stats, total_stats SearchStats) {

	str := fmt.Sprintf("%5d | ", child_id+1)
	buff := bytes.NewBufferString(str)
	if heur > alpha {
		buff.WriteString(fmt.Sprintf("%5d || ", heur))
	} else {
		buff.WriteString(fmt.Sprintf("      || "))
	}

	safe_div := func(num, den uint64) uint64 {
		if den == 0 {
			return 0
		}
		return num / den
	}

	avg_speed := safe_div(1000000000*total_stats.nodes, total_stats.time_ns)
	child_speed := safe_div(1000000000*child_stats.nodes, child_stats.time_ns)

	buff.WriteString(fmt.Sprintf("%s | %s | %s || %s | %s | %s |\n",
		fmt_big(child_stats.nodes), fmt_ns(child_stats.time_ns), fmt_big(child_speed),
		fmt_big(total_stats.nodes), fmt_ns(total_stats.time_ns), fmt_big(avg_speed)))
	bot.writer.Write(buff.Bytes())
}

// Does the best move according to heuristic and minimax algorithm
func (bot *BotHeuristic) DoMove(b board.Board) (afterwards board.Board) {

	children := []SortedBoard{}

	for _, child := range b.GetChildren() {
		children = append(children, SortedBoard{board: child, heur: 0})
	}

	if len(children) == 0 {
		panic("Cannot do move, because there are no moves.")
	}

	// prevent returning empty board when bot cannot prevent losing all discs
	afterwards = children[0].board

	if len(children) == 1 {
		buff := bytes.NewBufferString("Only one move. Skipping evaluation.\n")
		bot.writer.Write(buff.Bytes())
		return
	}

	var alpha, beta int
	var depth uint

	if b.CountEmpties() <= bot.exact_depth {
		alpha = Min_exact_heuristic
		beta = Max_exact_heuristic
		depth = b.CountEmpties()
	} else {
		depth = bot.search_depth

		// HACK: stumbling upon an exact solution
		// takes forever to compute. we set limits to solve that for now.
		alpha = -100
		beta = 100
	}

	header := "      | heuri || child |  child  | child || total |  total  |  avg  |\n"
	header += " move | stic  || nodes |   time  | speed || nodes |   time  | speed |\n"
	header += "------|-------||-------|---------|-------||-------|---------|-------|\n"

	bot.writer.Write(bytes.NewBufferString(header).Bytes())

	total_stats := SearchStats{
		nodes:   0,
		time_ns: 0}

	start_time := time.Now()

	processResult := func(child_id int) {
		result := <-bot.result_chan

		child_stats := result.stats
		total_stats.nodes += child_stats.nodes
		total_stats.time_ns = uint64(time.Since(start_time).Nanoseconds())

		bot.logChildEvaluation(child_id, result.heur, alpha, *child_stats, total_stats)
		if result.heur > alpha {
			alpha = result.heur
			afterwards = result.query.board
		}

	}

	for i, child := range children {

		query := SearchQuery{
			board:       child.board,
			lower_bound: alpha,
			upper_bound: beta,
			depth:       depth,
			guess:       child.heur,
			heuristic:   bot.heuristic}

		query.Run(bot.result_chan)

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

type SearchQuery struct {
	board       board.Board
	lower_bound int
	upper_bound int
	depth       uint
	guess       int
	heuristic   Heuristic
}

type SearchResult struct {
	query *SearchQuery
	stats *SearchStats
	heur  int
}

type SearchStats struct {
	nodes   uint64
	time_ns uint64
}

type SearchState struct {
	depth uint
	board board.Board
}

type SearchThread struct {
	query *SearchQuery
	state *SearchState
	stats *SearchStats
}

func (query *SearchQuery) Run(ch chan SearchResult) {

	thread := &SearchThread{
		query: query,
		state: &SearchState{
			depth: query.depth,
			board: query.board},
		stats: &SearchStats{
			nodes:   0,
			time_ns: 0}}

	go thread.Run(ch)
}

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
	high := thread.query.upper_bound
	low := thread.query.lower_bound

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

	thread.stats.nodes += 1

	if thread.state.depth == 0 {
		heur = mtdf_polish(thread.query.heuristic(thread.state.board), alpha)
		return
	}

	gen := board.NewChildGen(&thread.state.board)

	if gen.HasMoves() {
		thread.state.depth--
		heur = alpha
		for gen.Next() {
			child_heur := -thread.doMtdf(-(alpha + 1))
			if child_heur > alpha {
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

	heur = mtdf_polish(Exact_score_factor*thread.state.board.ExactScore(), alpha)
	return
}

func (thread *SearchThread) doMtdfExact(alpha int) (heur int) {

	thread.stats.nodes += 1

	gen := board.NewChildGen(&thread.state.board)

	if gen.HasMoves() {
		heur = alpha
		for gen.Next() {
			child_heur := -thread.doMtdfExact(-(alpha + 1))
			if child_heur > alpha {
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

	heur = mtdf_polish(thread.state.board.ExactScore(), alpha)
	return
}

func mtdf_polish(heur, alpha int) (outheur int) {
	if heur > alpha {
		outheur = alpha + 1
		return
	}
	outheur = alpha
	return
}
