package treesearch

import (
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/lk16/dots/internal/othello"
)

var dummyBoard *othello.Board

func BenchmarkMtdf(b *testing.B) {
	rand.Seed(0)

	var boards []othello.Board

	if err := othello.LoadXotBoards(); err != nil {
		b.Error(err)
		b.FailNow()
	}

	for i := 0; i < 10; i++ {
		boards = append(boards, *othello.NewXotBoard())
	}

	bot := NewBot(ioutil.Discard, 12, 18, NewMtdf(nil, Squared))

	for i := 0; i < b.N; i++ {
		board := *othello.NewXotBoard()
		dummyBoard, _ = bot.DoMove(board)
	}

}

func BenchmarkMtdfCached(b *testing.B) {
	rand.Seed(0)

	var boards []othello.Board

	if err := othello.LoadXotBoards(); err != nil {
		b.Error(err)
		b.FailNow()
	}

	for i := 0; i < 10; i++ {
		boards = append(boards, *othello.NewXotBoard())
	}

	cache := NewMemoryCache()

	bot := NewBot(ioutil.Discard, 12, 18, NewMtdf(cache, Squared))

	for i := 0; i < b.N; i++ {
		board := *othello.NewXotBoard()
		dummyBoard, _ = bot.DoMove(board)
	}
}
