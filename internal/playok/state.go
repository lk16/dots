package playok

import (
	"bytes"
	"fmt"
	"sort"
)

type table struct {
	rules   string
	players [2]string
}

type state struct {
	userName string
	rating   int
	tables   map[int]table
	players  map[string]player
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

// String gives a string representation of the state
func (state state) String() string {

	var tableIDs []int
	for ID := range state.tables {
		tableIDs = append(tableIDs, ID)
	}

	sort.Ints(tableIDs)

	var buff bytes.Buffer

	for _, ID := range tableIDs {
		table := state.tables[ID]

		buff.WriteString(fmt.Sprintf("%10s%5d%15s%15s\n",
			table.rules,
			ID,
			table.players[0],
			table.players[1],
		))
	}

	return buff.String()
}
