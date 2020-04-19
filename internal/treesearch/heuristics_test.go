package treesearch

import (
	"testing"

	"github.com/lk16/dots/internal/othello"
	"github.com/stretchr/testify/assert"
)

func TestSquared(t *testing.T) {
	assert.Equal(t, 0, Squared(*othello.NewBoard()))
	assert.Equal(t, 0, Squared(othello.Board{}))
	assert.Equal(t, 3, Squared(*othello.NewCustomBoard(0x1, 0x0)))
	assert.Equal(t, 4, Squared(*othello.NewCustomBoard(0x1, 0x6)))
}

func TestFastHeuristic(t *testing.T) {
	assert.Equal(t, 0, FastHeuristic(*othello.NewBoard()))
	assert.Equal(t, 0, FastHeuristic(othello.Board{}))
	assert.Equal(t, 2, FastHeuristic(*othello.NewCustomBoard(0x1, 0x0)))
	assert.Equal(t, 8, FastHeuristic(*othello.NewCustomBoard(0x1, 0x6)))
}

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
