package othello

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
)

var xotBoards []Board

type xotBoardModel struct {
	Me  string `json:"me"`
	Opp string `json:"opp"`
}

// LoadXot loads the xot boards into memory.
// When calling this again after a successful call, this function does nothing.
func LoadXot() error {

	var bytes []byte
	var err error

	if bytes, err = ioutil.ReadFile("./assets/xot.json"); err != nil {
		return errors.Wrap(err, "failed to initialize xot")
	}

	var xotModels []xotBoardModel
	if err = json.Unmarshal(bytes, &xotModels); err != nil {
		return errors.Wrap(err, "failed to parse xot file")
	}

	for _, xotModel := range xotModels {

		var me, opp uint64

		if me, err = strconv.ParseUint(xotModel.Me, 0, 64); err != nil {
			return errors.Wrap(err, "processing xot file failed")
		}

		if opp, err = strconv.ParseUint(xotModel.Opp, 0, 64); err != nil {
			return errors.Wrap(err, "processing xot file failed")
		}

		board := *NewCustomBoard(BitSet(me), BitSet(opp))
		xotBoards = append(xotBoards, board)
	}

	return nil
}

// NewXotBoard returns a new xot board
// http://berg.earthlingz.de/xot/aboutxot.php?lang=en
func NewXotBoard() *Board {

	if len(xotBoards) == 0 {
		panic("xot boards not loaded")
	}

	return &xotBoards[rand.Intn(len(xotBoards))]
}
