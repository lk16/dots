package treesearch

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/lk16/dots/internal/othello"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewBot(t *testing.T) {
	writer := &bytes.Buffer{}
	searchDepth := 7
	exactDepth := 14
	searcher := NewPvs()

	bot := NewBot(writer, searchDepth, exactDepth, searcher)

	assert.Equal(t, writer, bot.writer)
	assert.Equal(t, searchDepth, bot.searchDepth)
	assert.Equal(t, exactDepth, bot.exactDepth)
	assert.Equal(t, searcher, searcher)
}

type failWriter struct{}

var _ io.Writer = (*failWriter)(nil)

func (writer *failWriter) Write([]byte) (int, error) {
	return 0, errors.New("failed to write")
}

func TestBotWrite(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		var buff bytes.Buffer
		bot := &Bot{writer: &buff}
		bot.write("%s %d", "foo", 3)

		assert.Equal(t, "foo 3", buff.String())
	})

	t.Run("Fail", func(t *testing.T) {
		var logBuff bytes.Buffer
		log.SetOutput(&logBuff)
		defer log.SetOutput(os.Stderr)

		bot := &Bot{writer: &failWriter{}}
		bot.write("%s %d", "foo", 3)

		assert.Contains(t, logBuff.String(), "Bot write() error:")
	})
}

func TestBotDoMove(t *testing.T) {

	depth := 7
	exactDepth := 14

	t.Run("NoMoves", func(t *testing.T) {
		bot := NewBot(ioutil.Discard, depth, exactDepth, nil)
		board := othello.Board{}

		move, err := bot.DoMove(board)

		assert.Nil(t, move)
		assert.Equal(t, ErrNoMoves, err)
	})

	t.Run("OneMove", func(t *testing.T) {
		board := *othello.NewCustomBoard(0x1, 0x2)
		bot := NewBot(ioutil.Discard, depth, exactDepth, nil)

		move, err := bot.DoMove(board)

		assert.Nil(t, err)
		assert.Equal(t, *othello.NewCustomBoard(0x0, 0x7), *move)
	})

	t.Run("MultipleMoves", func(t *testing.T) {
		board := *othello.NewBoard()

		searcher := &mockSearcher{}
		searcher.On("Search", mock.Anything, mock.Anything, mock.Anything, 1).Times(4).Return(0)
		searcher.On("GetStats").Return(Stats{})
		searcher.On("ResetStats").Return(Stats{})

		bot := NewBot(ioutil.Discard, 1, exactDepth, searcher)

		move, err := bot.DoMove(board)
		assert.Contains(t, board.GetChildren(), *move)
		assert.Nil(t, err)

		searcher.AssertExpectations(t)
	})

	t.Run("MultipleMovesSorted", func(t *testing.T) {
		board := *othello.NewBoard()

		searcher := &mockSearcher{}
		searcher.On("Search", mock.Anything, mock.Anything, mock.Anything, 6).Times(4).Return(0)
		searcher.On("Search", mock.Anything, mock.Anything, mock.Anything, 8).Times(4).Return(0)
		searcher.On("GetStats").Return(Stats{})
		searcher.On("ResetStats").Return(Stats{})

		bot := NewBot(ioutil.Discard, 8, exactDepth, searcher)

		move, err := bot.DoMove(board)
		assert.Contains(t, board.GetChildren(), *move)
		assert.Nil(t, err)

		searcher.AssertExpectations(t)
	})

	t.Run("MultipleExact", func(t *testing.T) {
		board := *othello.NewBoard()

		searcher := &mockSearcher{}
		searcher.On("ExactSearch", mock.Anything, mock.Anything, mock.Anything).Times(4).Return(0)
		searcher.On("GetStats").Return(Stats{})
		searcher.On("ResetStats").Return(Stats{})

		bot := NewBot(ioutil.Discard, 8, 60, searcher)

		move, err := bot.DoMove(board)
		assert.Contains(t, board.GetChildren(), *move)
		assert.Nil(t, err)

		searcher.AssertExpectations(t)
	})

}
