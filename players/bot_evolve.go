package players

import (
	"dots/board"
	"fmt"
	"math/bits"
	"os"
)

// Parameters has parameters for the heuristic of BotEvolve
type Parameters struct {
	count int
}

// BotEvolve is a parametrised bot designed to work with Evolution
type BotEvolve struct {
	*BotHeuristic
}

// NewBotEvolve creates a new BotEvolve
func NewBotEvolve(searchDepth, exactDepth int, params Parameters) (bot *BotEvolve) {
	botheuristic := NewBotHeuristic(evolveHeuristic, searchDepth,
		exactDepth, os.Stdout, false)
	bot = &BotEvolve{
		BotHeuristic: botheuristic}
	bot.BotHeuristic.heuristicParams = &params
	return
}

func evolveHeuristic(b board.Board, object interface{}) int {
	params := object.(*Parameters)
	params.count++
	fmt.Printf("count = %d\n", params.count)
	return bits.OnesCount64(b.Opp()) - bits.OnesCount64(b.Me())
}
