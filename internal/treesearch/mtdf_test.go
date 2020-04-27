package treesearch

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/lk16/dots/internal/othello"
	"github.com/stretchr/testify/assert"
)

var dummyBoard *othello.Board

func BenchmarkSearch(b *testing.B) {

	assert.Nil(b, othello.LoadXotBoards())

	rand.Seed(0)

	type boardSet struct {
		name   string
		boards []othello.Board
	}

	xotBoardset := boardSet{name: "Xot"}
	for i := 0; i < 100; i++ {
		xotBoardset.boards = append(xotBoardset.boards, *othello.NewXotBoard())
	}

	boardSets := []boardSet{xotBoardset}

	for depth := 20; depth <= 45; depth += 5 {
		boardSet := boardSet{name: fmt.Sprintf("depth%d", depth)}
		for i := 0; i < 100; i++ {
			board, err := othello.NewRandomBoard(depth)
			assert.Nil(b, err)
			boardSet.boards = append(boardSet.boards, *board)
		}
		boardSets = append(boardSets, boardSet)
	}

	type namedBot struct {
		name string
		bot  *Bot
	}

	depth := 10
	exactDepth := 0

	namedBots := []namedBot{
		{"MtdfNotCached", NewBot(ioutil.Discard, depth, exactDepth, NewMtdf(nil, FastHeuristic))},
		{"MtdfCached", NewBot(ioutil.Discard, depth, exactDepth, NewMtdf(NewMemoryCache(), FastHeuristic))},
		{"PvsNotCached", NewBot(ioutil.Discard, depth, exactDepth, NewPvs(nil, FastHeuristic))},
		{"PvsCached", NewBot(ioutil.Discard, depth, exactDepth, NewPvs(NewMemoryCache(), FastHeuristic))},
	}

	for _, boardSet := range boardSets {
		for _, namedBot := range namedBots {
			runName := fmt.Sprintf("%s/%s", boardSet.name, namedBot.name)

			b.Run(runName, func(b *testing.B) {

				b.ReportAllocs()

				for i := 0; i < b.N; i++ {
					board := *othello.NewXotBoard()
					dummyBoard, _ = namedBot.bot.DoMove(board)
				}
			})
		}
	}

}
