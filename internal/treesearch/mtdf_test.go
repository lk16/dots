package treesearch

import (
	"math/rand"
	"testing"

	"github.com/lk16/dots/internal/othello"
)

var dummyInt int

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

	mtdf := NewMtdf(Squared)

	for i := 0; i < b.N; i++ {
		dummyInt = mtdf.Search(boards[i%10], MinHeuristic, MaxHeuristic, 10)
	}
}
