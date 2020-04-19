package treesearch

import (
	"testing"

	"github.com/lk16/dots/internal/othello"
)

var dummy int

func BenchmarkSquared(b *testing.B) {
	board := *othello.NewBoard()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dummy += Squared(board)
	}
	b.StopTimer()
}

func BenchmarkFastHeuristic(b *testing.B) {
	board := *othello.NewBoard()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dummy += FastHeuristic(board)
	}
	b.StopTimer()
}
