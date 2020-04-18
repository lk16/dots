package othello

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"strconv"
)

type xotBoardModel struct {
	Me  string `json:"me"`
	Opp string `json:"opp"`
}

func init() {
	bytes, err := ioutil.ReadFile("/assets/xot.json")
	if err != nil {
		panic("failed to initialize xot: " + err.Error())
	}

	var xotModels []xotBoardModel
	err = json.Unmarshal(bytes, &xotModels)

	if err != nil {
		panic("failed to parse xot json file: " + err.Error())
	}

	for _, xotModel := range xotModels {

		me, err := strconv.ParseUint(xotModel.Me, 0, 64)
		if err != nil {
			panic("loading xot json file failed: " + err.Error())
		}

		opp, err := strconv.ParseUint(xotModel.Opp, 0, 64)
		if err != nil {
			panic("loading xot json file failed: " + err.Error())
		}

		board := *NewCustomBoard(BitSet(me), BitSet(opp))
		xotBoards = append(xotBoards, board)
	}
}

var xotBoards []Board

// NewXotBoard returns a new xot board
// http://berg.earthlingz.de/xot/aboutxot.php?lang=en
func NewXotBoard() *Board {
	return &xotBoards[rand.Intn(len(xotBoards))]
}
