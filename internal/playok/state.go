package playok

import (
	"sync"

	"github.com/lk16/dots/internal/othello"
)

type table struct {
	timeLimit   int // minutes
	xot         bool
	rated       bool
	players     [2]string
	minRatingID int // TODO
}

type currentTable struct {
	table
	ID           int
	viewers      []string
	op           string
	allowUndo    bool
	board        othello.BoardWithTurn
	playerToMove int
}

type state struct {
	sync.RWMutex
	userName     string
	rating       int
	tables       map[int]table
	players      map[string]player
	currentTable currentTable
}

type player struct {
	rating int
}

func newState() *state {
	return &state{
		tables:  make(map[int]table),
		players: make(map[string]player),
	}
}
