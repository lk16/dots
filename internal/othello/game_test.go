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
				Name:   "lk16",
				Rating: 1558,
			},
			White: GamePlayer{
				Name:   "pizzandsprite",
				Rating: 1346,
			},
			Xot: true,
		}

		// tested elsewhere
		expectedGame.Moves = games[0].Moves

		assert.Equal(t, expectedGame, games[0])
	})

	t.Run("OKPlayokMany", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/playok_many.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(games))

		date, err := time.Parse("2006.01.02", "2020.04.24")
		assert.Nil(t, err)

		expectedGames := []Game{
			Game{
				Site: "PlayOK",
				Date: date,
				Black: GamePlayer{
					Name:   "cgs3898g",
					Rating: 1612,
				},
				White: GamePlayer{
					Name:   "lk16",
					Rating: 1529,
				},
				Xot: true,
			},
			Game{
				Site: "PlayOK",
				Date: date,
				Black: GamePlayer{
					Name:   "lk16",
					Rating: 1542,
				},
				White: GamePlayer{
					Name:   "cgs3898g",
					Rating: 1599,
				},
				Xot: true,
			},
		}

		// tested elsewhere
		expectedGames[0].Moves = games[0].Moves
		expectedGames[1].Moves = games[1].Moves

		assert.Equal(t, expectedGames, games)
	})

	t.Run("OKFlyOrdie", func(t *testing.T) {

		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/flyordie_one.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(games))

		expectedDate, err := time.Parse("2006.01.02", "2019.04.14")
		assert.Nil(t, err)

		expectedGame := Game{
			Site: "www.flyordie.com",
			Date: expectedDate,
			Black: GamePlayer{
				Name:   "wattego",
				Rating: 155,
			},
			White: GamePlayer{
				Name:   "LK16",
				Rating: 325,
			},
			Xot: false,
		}

		// tested elsewhere
		expectedGame.Moves = games[0].Moves

		assert.Equal(t, expectedGame, games[0])
	})

	t.Run("BrokenAttribute", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/broken_attribute.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, games)
		assert.Contains(t, err.Error(), "parsing game attributes failed")
	})

	t.Run("BrokenDate", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/broken_date.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, games)
		assert.Contains(t, err.Error(), "date parsing failed")
	})

	t.Run("BrokenDiscCount", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/broken_disc_count.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, games)
		assert.Contains(t, err.Error(), "parsing disc counts at end of game failed")
	})

	t.Run("BrokenGameInvalidMove", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/broken_invalid_move.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, games)
		assert.Contains(t, err.Error(), "game verification failed")
	})

	t.Run("BrokenBlackRating", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/broken_rating_black.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, games)
		assert.Contains(t, err.Error(), "failed to parse black rating")
	})

	t.Run("BrokenWhiteRating", func(t *testing.T) {
		bytes, err := ioutil.ReadFile(assetsPath + "testdata/pgn/broken_rating_white.txt")
		assert.Nil(t, err)

		games, err := LoadGamesFromPGN(bytes)
		assert.Nil(t, games)
		assert.Contains(t, err.Error(), "failed to parse white rating")
	})

}

func TestFieldToIndex(t *testing.T) {
	type testCase struct {
		field         string
		expectedIndex uint
		expectedOk    bool
	}

	testCases := []testCase{
		{"foo", 0, false},
		{"a", 0, false},
		{"a0", 0, false},
		{"a9", 0, false},
		{"`1", 0, false},
		{"--", PassMoveID, true},
		{"a1", 0, true},
		{"h1", 7, true},
		{"a8", 56, true},
		{"h8", 63, true},
	}

	for _, testCase := range testCases {
		moveID, ok := fieldToIndex(testCase.field)
		assert.Equal(t, testCase.expectedOk, ok)
		assert.Equal(t, testCase.expectedIndex, moveID)
	}

}

func TestGameVerify(t *testing.T) {

	t.Run("OK", func(t *testing.T) {

		shortestGame := []string{"e6", "f4", "e3", "f6", "g5", "d6", "e7", "f5", "c5"}
		var moves []uint

		for _, field := range shortestGame {
			moveID, ok := fieldToIndex(field)
			assert.True(t, ok)
			moves = append(moves, moveID)
		}
		game := &Game{
			Moves: moves,
		}
		assert.True(t, game.Verify())
	})

	t.Run("FailPassWithMoves", func(t *testing.T) {
		game := &Game{
			Moves: []uint{PassMoveID},
		}
		assert.False(t, game.Verify())
	})

	t.Run("FailInvalidMove", func(t *testing.T) {
		game := &Game{
			Moves: []uint{0},
		}
		assert.False(t, game.Verify())
	})

	t.Run("FailGameEndWithMovesLeft", func(t *testing.T) {
		game := &Game{
			Moves: []uint{},
		}
		assert.False(t, game.Verify())
	})

}
