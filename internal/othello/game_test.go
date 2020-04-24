package othello

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadGamesFromPGN(t *testing.T) {

	t.Run("OKPlayokOne", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/playok_one.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(games))

		expectedDate, err := time.Parse("2006.01.02", "2020.04.23")
		assert.Nil(t, err)

		expectedGame := Game{
			Site: "PlayOK",
			Date: expectedDate,
			Black: GamePlayer{
				Name:        "lk16",
				Rating:      1558,
				ResultDiscs: 17,
			},
			White: GamePlayer{
				Name:        "pizzandsprite",
				Rating:      1346,
				ResultDiscs: 47,
			},
			Xot: true,
		}

		// tested elsewhere
		expectedGame.Moves = games[0].Moves

		assert.Equal(t, expectedGame, games[0])
	})

	t.Run("OKPlayokMany", func(t *testing.T) {
		// TODO
	})

	t.Run("OKFlyOrdie", func(t *testing.T) {
		// TODO
	})

	t.Run("OKFlyOrdieWithPasses", func(t *testing.T) {
		// TODO
	})

	// TODO fail cases
}

func TestGameVerify(t *testing.T) {
	panic("not implemented")
	// TODO
}

func TestFieldToIndex(t *testing.T) {
	panic("not implemented")
	// TODO
}
