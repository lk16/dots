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

// ErrorOutput is an error output model
type ErrorOutput struct {
	Error string `json:"error"`
}

// Output is an output model that reports success
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
		if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
			log.Printf("Failed to write to stdout: %s", err.Error())
		}
		return
	}

	discsBefore := board.Me() | board.Opp()
	discsAfter := bestChild.Me() | bestChild.Opp()

	newDiscMask := discsBefore ^ discsAfter

	if newDiscMask.Count() != 1 {
		output := ErrorOutput{Error: fmt.Sprintf("disc difference of one move is %d", newDiscMask.Count())}
		if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
			log.Printf("Failed to write to stdout: %s", err.Error())
		}
		return
	}

	bestMove := newDiscMask.Lowest()

	output := Output{BestMove: bestMove}
	if err := json.NewEncoder(os.Stdout).Encode(output); err != nil {
		log.Printf("Failed to write to stdout: %s", err.Error())
	}
}
