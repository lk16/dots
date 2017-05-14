package players

import (
	"bytes"
	"fmt"
	"io"

	"dots/board"
	"dots/minimax"
)

type BotHeuristic struct {
	heuristic    minimax.Heuristic
	minimax      minimax.Interface
	search_depth uint
	exact_depth  uint
	writer       io.Writer
}

// Creates a new BotHeuristic
func NewBotHeuristic(heuristic minimax.Heuristic, minimax minimax.Interface,
	search_depth, exact_depth uint, writer io.Writer) (bot *BotHeuristic) {
	bot = &BotHeuristic{
		heuristic:    heuristic,
		minimax:      minimax,
		search_depth: search_depth,
		exact_depth:  exact_depth,
		writer:       writer}
	return
}

func (bot *BotHeuristic) logChildEvaluation(child_id, heur, alpha int, nodes,
	total_nodes, time_ns, total_time_ns uint64) {

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

	avg_speed := safe_div(1000000000*total_nodes, total_time_ns)
	child_speed := safe_div(1000000000*nodes, time_ns)

	buff.WriteString(fmt.Sprintf("%s | %5dms | %s || %s | %5dms | %s |\n",
		fmt_big(nodes), time_ns/1000000, fmt_big(child_speed),
		fmt_big(total_nodes), total_time_ns/1000000, fmt_big(avg_speed)))
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

	var child_eval func(board.Board, uint, minimax.Heuristic, int) int
	var alpha int

	if b.CountEmpties() <= bot.exact_depth {
		alpha = minimax.Min_exact_heuristic
		// wrapper function to achieve same prototype for ExactSearch() as Search()
		// it just drops arguments depth and heuristic
		child_eval = func(child board.Board, depth uint,
			heuristic minimax.Heuristic, alpha int) int {
			return bot.minimax.ExactSearch(child, alpha)
		}
	} else {
		alpha = minimax.Min_heuristic
		child_eval = bot.minimax.Search
	}

	header := "      | heuri || child |  child  | child || total |  total  |  avg  |\n"
	header += " move | stic  || nodes |   time  | speed || nodes |   time  | speed |\n"
	header += "------|-------||-------|---------|-------||-------|---------|-------|\n"

	bot.writer.Write(bytes.NewBufferString(header).Bytes())

	total_nodes := uint64(0)
	total_time_ns := uint64(0)

	for i, child := range children {

		heur := child_eval(child, bot.search_depth, bot.heuristic, alpha)

		nodes := bot.minimax.Nodes()
		total_nodes += nodes
		ns := bot.minimax.ComputeTimeNs()
		total_time_ns += ns

		bot.logChildEvaluation(i, heur, alpha, nodes, total_nodes, ns, total_time_ns)

		if heur > alpha {
			alpha = heur
			afterwards = child
		}
	}

	bot.writer.Write(bytes.NewBufferString("\n\n").Bytes())
	return
}
