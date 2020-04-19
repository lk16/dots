package othello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testDataFolder = "/app/assets/testdata/xot/"

func TestNewXotBoard(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		xotBoards = []Board{*NewBoard()}
		board := *NewXotBoard()
		assert.Equal(t, xotBoards[0], board)
	})

	t.Run("NotLoaded", func(t *testing.T) {
		xotBoards = []Board{}
		assert.Panics(t, func() {
			NewXotBoard()
		})
	})
}

func TestLoadXot(t *testing.T) {

	t.Run("OK", func(t *testing.T) {
		xotBoards = nil
		assert.Nil(t, LoadXot())
		assert.Equal(t, 10784, len(xotBoards))
	})

	t.Run("AlreadyLoaded", func(t *testing.T) {
		xotBoards = []Board{*NewBoard()}
		assert.Nil(t, LoadXot())
		assert.Equal(t, 1, len(xotBoards))
	})

	xotDataPathCopy := xotDataPath

	restoreXotDataPath := func() {
		xotDataPath = xotDataPathCopy
	}

	t.Run("CustomPath", func(t *testing.T) {
		xotBoards = nil
		xotDataPath = testDataFolder + "valid.json"
		defer restoreXotDataPath()

		assert.Nil(t, LoadXot())
		assert.Equal(t, 1, len(xotBoards))
	})

	t.Run("FileNotFound", func(t *testing.T) {
		xotBoards = nil
		xotDataPath = testDataFolder + "nonexistent.json"
		defer restoreXotDataPath()

		err := LoadXot()

		assert.Contains(t, err.Error(), "failed to load xot file")
		assert.Equal(t, 0, len(xotBoards))
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		xotBoards = nil
		xotDataPath = testDataFolder + "invalid.json"
		defer restoreXotDataPath()

		err := LoadXot()

		assert.Contains(t, err.Error(), "failed to parse xot file")
		assert.Equal(t, 0, len(xotBoards))
	})

	t.Run("BrokenMeField", func(t *testing.T) {
		xotBoards = nil
		xotDataPath = testDataFolder + "broken_me_field.json"
		defer restoreXotDataPath()

		err := LoadXot()

		assert.Contains(t, err.Error(), "processing xot file failed")
		assert.Equal(t, 0, len(xotBoards))
	})

	t.Run("BrokenOppField", func(t *testing.T) {
		xotBoards = nil
		xotDataPath = testDataFolder + "broken_opp_field.json"
		defer restoreXotDataPath()

		err := LoadXot()

		assert.Contains(t, err.Error(), "processing xot file failed")
		assert.Equal(t, 0, len(xotBoards))
	})
}
