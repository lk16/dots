package players

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"time"

	"dots/board"
	"dots/minimax"
)

type SortedBoard struct {
	board board.Board
	heur  int
}

type BotHeuristic struct {
	heuristic    minimax.Heuristic
	search_depth uint
	exact_depth  uint
	writer       io.Writer
}

// Creates a new BotHeuristic
func NewBotHeuristic(heuristic minimax.Heuristic,
	search_depth, exact_depth uint, writer io.Writer) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:    heuristic,
		search_depth: search_depth,
		exact_depth:  exact_depth,
		writer:       writer}
	return
}

func (bot *BotHeuristic) logChildEvaluation(child_id, heur, alpha int,
	child_stats, total_stats SearchStats) {

	fmt_big := func(n uint64) string {

		if n < 10000 {
			return fmt.Sprintf(" %4d", n)
		}

		n /= 1000
		suffix_index := 0

		for n > 10000 {
			n /= 1000
			suffix_index++
		}

		suffixes := "KMGTPE"

		return fmt.Sprintf("%4d%c", n, suffixes[suffix_index])
	}

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

	buff.WriteString(fmt.Sprintf("%s | %5dms | %s || %s | %5dms | %s |\n",
		fmt_big(child_stats.nodes), child_stats.time_ns/1000000, fmt_big(child_speed),
		fmt_big(total_stats.nodes), total_stats.time_ns/1000000, fmt_big(avg_speed)))
	bot.writer.Write(buff.Bytes())
}

// Does the best move according to heuristic and minimax algorithm
func (bot *BotHeuristic) DoMove(b board.Board) (afterwards board.Board) {

	children := b.GetChildren()

	if len(children) == 0 {
		panic("Cannot do move, because there are no moves.")
	}

	// prevent returning empty board when bot cannot prevent losing all discs
	afterwards = children[0]

	if len(children) == 1 {
		buff := bytes.NewBufferString("Only one move. Skipping evaluation.\n")
		bot.writer.Write(buff.Bytes())
		return
	}

	var alpha, beta int
	var depth uint

	if b.CountEmpties() <= bot.exact_depth {
		alpha = minimax.Min_exact_heuristic
		beta = minimax.Max_exact_heuristic
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

	result_chan := make(chan SearchResult)

	total_stats := SearchStats{
		nodes:   0,
		time_ns: 0}

	sorted_children := []SortedBoard{}

	for _, child := range children {
		sorted_children = append(sorted_children, SortedBoard{board: child, heur: 0})
	}

	for d := depth % 2; d <= depth; d += 2 {

		for i, child := range sorted_children {

			query := SearchQuery{
				board:       child.board,
				lower_bound: alpha,
				upper_bound: beta,
				depth:       d,
				guess:       child.heur,
				heuristic:   bot.heuristic}

			go RunQuery(query, result_chan)
			result := <-result_chan

			child_stats := result.stats
			total_stats.nodes += child_stats.nodes
			total_stats.time_ns += child_stats.time_ns

			sorted_children[i].heur = result.heur

			if d == depth {
				bot.logChildEvaluation(i, result.heur, alpha, child_stats, total_stats)
				if result.heur > alpha {
					alpha = result.heur
				}
			}

		}

		sort.Slice(sorted_children, func(i, j int) bool {
			return sorted_children[i].heur > sorted_children[j].heur
		})
	}

	afterwards = sorted_children[0].board
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
	heuristic   minimax.Heuristic
}

type SearchResult struct {
	query SearchQuery
	stats SearchStats
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
	start time.Time
	query SearchQuery
	state SearchState
	stats SearchStats
}

func RunQuery(query SearchQuery, ch chan SearchResult) {

	stats := SearchStats{
		nodes:   0,
		time_ns: 0}

	state := SearchState{
		depth: query.depth,
		board: query.board}

	thread := &SearchThread{
		query: query,
		start: time.Now(),
		state: state,
		stats: stats}

	var heur int
	if thread.query.board.CountEmpties() > thread.query.depth {
		heur = thread.loop(1, thread.doMtdf)
	} else {
		heur = thread.loop(2, thread.doMtdfExact)
	}

	result := SearchResult{
		query: thread.query,
		heur:  heur,
		stats: thread.stats}

	result.stats.time_ns = uint64(time.Since(thread.start).Nanoseconds())

	ch <- result

}

func (thread *SearchThread) loop(step int, call func(int) int) (heur int) {

	// copy values because thread.query should not be modified
	high := thread.query.upper_bound
	low := thread.query.lower_bound

	// prevent odd results for exact search
	f := thread.query.guess & ^int(2)

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

	heur = mtdf_polish(minimax.Exact_score_factor*thread.state.board.ExactScore(), alpha)
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
