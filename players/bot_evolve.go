package players

import (
	"dots/board"
	"io/ioutil"
	"math/bits"
	"math/rand"
)

// Parameters has parameters for the heuristic of BotEvolve
type Parameters struct {
	PositionValue [10]int
}

// RandomParameters returns random parameters within given range
func RandomParameters(min, max int) (params Parameters) {
	for i := range params.PositionValue {
		params.PositionValue[i] = min + (rand.Int() % (max - min))
	}
	return
}

// BotEvolve is a parametrised bot designed to work with Evolution
type BotEvolve struct {
	*BotHeuristic
}

// NewBotEvolve creates a new BotEvolve
func NewBotEvolve(searchDepth, exactDepth int, params Parameters) (bot *BotEvolve) {
	botheuristic := NewBotHeuristic(evolveHeuristic, searchDepth,
		exactDepth, ioutil.Discard, false)
	bot = &BotEvolve{
		BotHeuristic: botheuristic}
	bot.BotHeuristic.heuristicParams = &params
	return
}

func evolveHeuristic(b board.Board, object interface{}) (heur int) {
	params := object.(*Parameters)

	for i, value := range params.PositionValue {
		heur += value * bits.OnesCount64(b.Me()&board.PositionMasks[i])
		heur -= value * bits.OnesCount64(b.Opp()&board.PositionMasks[i])
	}
	return
}
