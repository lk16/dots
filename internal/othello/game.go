package othello

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

var (
	pgnAttributeRegex  = regexp.MustCompile(`\[(.*) "(.*)"\]`)
	pgnFinalScoreRegex = regexp.MustCompile(`(\d+)-(\d+)`)
)

const (
	// PassMoveID is used when a player has to pass because there are no moves
	PassMoveID = 99
)

// PGNParseError is used when a PGN fails to parse
type PGNParseError struct {
	LineNumber int
	Message    string
}

func (err PGNParseError) Error() string {
	return fmt.Sprintf("pgn failed to parse line %d: %s", err.LineNumber, err.Message)
}

// GamePlayer has details on a Player as found in PGNs
type GamePlayer struct {
	Name   string
	Rating int
}

// Game represents an entire othello game
type Game struct {
	Site         string
	Date         time.Time
	Black, White GamePlayer
	Xot          bool
	Moves        []uint
}

// Verify checks if the sequence of moves is a valid game
func (game Game) Verify() bool {
	board := NewBoard()

	for _, move := range game.Moves {
		validMoves := board.Moves()

		if move == PassMoveID {
			if validMoves != 0 {
				// incorrectly skipping turn
				return false
			}
			board.SwitchTurn()
			continue
		}

		if !validMoves.Test(move) {
			return false
		}
		board.DoMove(BitSet(1 << move))
	}

	return board.Moves() == 0 && board.OpponentMoves() == 0
}

func fieldToIndex(field string) (uint, bool) {
	if field == "--" {
		return PassMoveID, true
	}

	if len(field) != 2 {
		return 0, false
	}

	col := field[0] - 'a'
	row := field[1] - '1'

	if row > 7 || col > 7 {
		return 0, false
	}

	return uint((8 * row) + col), true
}

// LoadGamesFromPGN loads one or more othello games from Portable Game Notation data
func LoadGamesFromPGN(bytes []byte) ([]Game, error) {
	return newPgnParser(bytes).parse()
}

type pgnParser struct {
	lines  []string
	offset int
}

func newPgnParser(bytes []byte) *pgnParser {
	return &pgnParser{
		lines:  strings.Split(string(bytes), "\n"),
		offset: 0,
	}
}

func (parser *pgnParser) parse() ([]Game, error) {
	var (
		games []Game
		game  *Game
		err   error
	)

	for parser.offset < len(parser.lines) {
		line := parser.lines[parser.offset]

		if len(line) == 0 {
			parser.offset++
			continue
		}

		if game, err = parser.parseGame(); err != nil {
			return nil, err
		}

		if ok := game.Verify(); !ok {
			err := &PGNParseError{
				LineNumber: parser.offset,
				Message:    "game verification failed",
			}

			return nil, err
		}

		games = append(games, *game)
	}

	return games, nil
}

func (parser *pgnParser) parseGame() (*Game, error) {
	var (
		game *Game
		err  error
	)

	if game, err = parser.parseAttributes(); err != nil {
		return nil, err
	}

	if err = parser.parseMoves(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (parser *pgnParser) parseAttributes() (*Game, error) {
	var game Game

	for parser.offset < len(parser.lines) {
		line := parser.lines[parser.offset]

		if len(line) == 0 {
			parser.offset++
			continue
		}

		if line[0] != '[' {
			break
		}

		matches := pgnAttributeRegex.FindStringSubmatch(line)

		if len(matches) != 3 {
			err := &PGNParseError{
				LineNumber: parser.offset,
				Message:    "parsing game attributes failed",
			}

			return nil, err
		}

		key := matches[1]
		value := matches[2]

		if err := parser.parseAttribute(&game, key, value); err != nil {
			return nil, err
		}

		parser.offset++
	}

	return &game, nil
}

func (parser *pgnParser) parseAttribute(game *Game, key, value string) error {
	switch key {
	case "Site":
		game.Site = value
	case "Date":
		date, err := time.Parse("2006.01.02", value)
		if err != nil {
			pgnErr := &PGNParseError{
				LineNumber: parser.offset,
				Message:    fmt.Sprintf("date parsing failed: %s", err.Error()),
			}
			return pgnErr
		}

		game.Date = date
	case "Black":
		game.Black.Name = value
	case "White":
		game.White.Name = value

	case "BlackRating":
		fallthrough
	case "BlackElo":
		rating, err := strconv.Atoi(value)
		if err != nil {
			pgnErr := &PGNParseError{
				LineNumber: parser.offset,
				Message:    fmt.Sprintf("failed to parse black rating: %s", err.Error()),
			}
			return pgnErr
		}
		game.Black.Rating = rating

	case "WhiteRating":
		fallthrough
	case "WhiteElo":
		rating, err := strconv.Atoi(value)
		if err != nil {
			pgnErr := &PGNParseError{
				LineNumber: parser.offset,
				Message:    fmt.Sprintf("failed to parse white rating: %s", err.Error()),
			}
			return pgnErr
		}
		game.White.Rating = rating

	case "Variant":
		if value == "xot" {
			game.Xot = true
		}
	}
	return nil
}

func (parser *pgnParser) parseMoves(game *Game) error {
	for parser.offset < len(parser.lines) {
		line := parser.lines[parser.offset]

		for _, moveField := range strings.Split(line, " ") {
			if strings.HasSuffix(moveField, ".") {
				continue
			}

			if moveID, ok := fieldToIndex(moveField); ok {
				game.Moves = append(game.Moves, moveID)
				continue
			}

			matches := pgnFinalScoreRegex.FindStringSubmatch(moveField)

			discCountErr := &PGNParseError{
				LineNumber: parser.offset,
				Message:    "parsing disc counts at end of game failed",
			}

			if len(matches) != 3 {
				return discCountErr
			}

			// TODO process final dics count

			// final disc counts indicate the end of the game
			parser.offset++
			return nil
		}

		parser.offset++
	}

	// files should not end before the "discs at end of game" part
	return errors.New("unexpected file end")
}
