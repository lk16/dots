package playok

import (
	"math/rand"
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

func (t table) countPlayers() int {
	var count int
	for _, player := range t.players {
		if player != "" {
			count++
		}
	}
	return count
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

func (s state) getShuffledTableIDs() []int {
	IDs := make([]int, len(s.tables))
	i := 0
	for ID := range s.tables {
		IDs[i] = ID
		i++
	}

	rand.Shuffle(len(IDs), func(i, j int) { IDs[i], IDs[j] = IDs[j], IDs[i] })
	return IDs
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
