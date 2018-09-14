package players

import (
	"bytes"
	"dots/treesearch"
	"fmt"
	"io"
	"time"

	"dots/othello"
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
type Heuristic func(othello.Board) int

// BotHeuristic is a bot that uses a Heuristic for choosing its moves
type BotHeuristic struct {
	heuristic      Heuristic
	searchDepth    int
	exactDepth     int
	writer         io.Writer
	resultChan     chan SearchResult
	stats          SearchStats
	startTime      time.Time
	parallelSearch bool
}

// NewBotHeuristic creates a new BotHeuristic
func NewBotHeuristic(heuristic Heuristic,
	searchDepth, exactDepth int, writer io.Writer,
	parallelSearch bool) *BotHeuristic {

	return &BotHeuristic{
		heuristic:   heuristic,
		searchDepth: searchDepth,
		exactDepth:  exactDepth,
		writer:      writer,
		resultChan:  make(chan SearchResult, 32),
		stats: SearchStats{
			nodes:  0,
			timeNs: 0},
		parallelSearch: parallelSearch}
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
	afterwards *othello.Board) {

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
func (bot *BotHeuristic) DoMove(b othello.Board) (afterwards othello.Board) {

	children := b.GetChildren()

	if len(children) == 0 {
		return b
	}

	// prevent returning empty othello when bot cannot prevent losing all discs
	afterwards = children[0]

	if len(children) == 1 {
		buff := bytes.NewBufferString("Only one move. Skipping evaluation.\n")
		bot.writer.Write(buff.Bytes())
		return
	}

	var alpha, beta int
	var depth int

	if b.CountEmpties() <= bot.exactDepth {
		alpha = treesearch.MinScore
		beta = treesearch.MaxScore
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
	gen := othello.NewGenerator(&child, 1)

	for i := 0; gen.Next(); i++ {

		query := SearchQuery{
			board:      child,
			lowerBound: alpha,
			upperBound: beta,
			depth:      depth,
			guess:      0,
			heuristic:  bot.heuristic}

		query.Run(bot.resultChan)

		if i == 0 || !bot.parallelSearch {
			bot.processResult(<-bot.resultChan, &alpha, &afterwards)
		}
	}

	if bot.parallelSearch {
		for i := 1; i < len(children); i++ {
			bot.processResult(<-bot.resultChan, &alpha, &afterwards)
		}
	}

	bot.writer.Write(bytes.NewBufferString("\n\n").Bytes())
	return
}
