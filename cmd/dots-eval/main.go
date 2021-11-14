package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/lk16/dots/internal/othello"
	"github.com/lk16/dots/internal/treesearch"
)

type ErrorOutput struct {
	Error string `json:"error"`
}

type Output struct {
	BestMove int `json:"move"`
}

func main() {
	myDiscs := flag.Uint64("me", 0, "Bitset of player to move")
	oppDiscs := flag.Uint64("opp", 0, "Bitset of opponent of player to move")
	searchDepth := flag.Int("depth", 0, "Search depth for heuristic evaluation")
	exactSearchDepth := flag.Int("exact", 0, "Search depth for exact solution")
	verbose := flag.Bool("verbose", false, "Print bot and debug output")

	flag.Parse()

	me := othello.BitSet(*myDiscs)
	opp := othello.BitSet(*oppDiscs)
	board := othello.NewCustomBoard(me, opp)

	var botWriter io.Writer = ioutil.Discard

	if *verbose {
		log.Printf("Input board:\n%s", board)
		log.Printf("Search depth: %d\n", *searchDepth)
		log.Printf("Exact depth: %d\n", *exactSearchDepth)
		botWriter = os.Stderr
	}

	bot := treesearch.NewBot(botWriter, *searchDepth, *exactSearchDepth, treesearch.NewPvs(treesearch.Squared))

	bestChild, err := bot.DoMove(*board)

	if err != nil {
		output := ErrorOutput{Error: err.Error()}
		json.NewEncoder(os.Stdout).Encode(output)
		return
	}

	discsBefore := board.Me() | board.Opp()
	discsAfter := bestChild.Me() | bestChild.Opp()

	newDiscMask := discsBefore ^ discsAfter

	if newDiscMask.Count() != 1 {
		output := ErrorOutput{Error: fmt.Sprintf("disc difference of one move is %d", newDiscMask.Count())}
		json.NewEncoder(os.Stdout).Encode(output)
		return
	}

	best_move := newDiscMask.Lowest()

	output := Output{BestMove: best_move}
	json.NewEncoder(os.Stdout).Encode(output)
}
