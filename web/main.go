package web

import (
	"dots/othello"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

type state struct {
	White []int `json:"white"`
	Black []int `json:"black"`
	Turn  int   `json:"turn"`
}

func newState(board othello.Board, turn int) state {

	me := make([]int, 0)
	opp := make([]int, 0)

	for i := uint(0); i < 64; i++ {
		if board.Me()&(1<<i) != 0 {
			me = append(me, int(i))
		}
		if board.Opp()&(1<<i) != 0 {
			opp = append(opp, int(i))
		}
	}

	if turn == 0 {
		return state{
			Black: me,
			White: opp,
			Turn:  0}
	}

	return state{
		White: me,
		Black: opp,
		Turn:  1}
}

func (s *state) getBoard() (*othello.Board, error) {

	white := uint64(0)
	black := uint64(0)

	for _, w := range s.White {
		if w < 0 || w >= 64 {
			return nil, fmt.Errorf("invalid white field value %d", w)
		}
		white |= uint64(1 << uint(w))
	}

	for _, b := range s.Black {
		if b < 0 || b >= 64 {
			return nil, fmt.Errorf("invalid black field value %d", b)
		}
		black |= uint64(1 << uint(b))
	}

	if white&black != 0 {
		return nil, fmt.Errorf("white (%+v) and black (%+v) overlap", white, black)
	}

	switch s.Turn {
	case 0:
		return othello.CustomBoard(black, white), nil
	case 1:
		return othello.CustomBoard(white, black), nil
	default:
		return nil, fmt.Errorf("invalid turn value %d", s.Turn)
	}
}

type clickData struct {
	Cell  int   `json:"cell"`
	State state `json:"state"`
}

type wsMessage struct {
	Event string    `json:"event"`
	Data  clickData `json:"data"`
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %s", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read error: %s", err)
			break
		}
		var wsMessage wsMessage
		err = json.Unmarshal(message, &wsMessage)
		if err != nil {
			log.Printf("json decode error: %s", err)
			continue
		}
		board, err := wsMessage.Data.State.getBoard()
		if err != nil {
			log.Printf("error constructing board: %s", err)
			continue
		}
		if wsMessage.Data.Cell < 0 || wsMessage.Data.Cell >= 64 {
			log.Printf("invalid Cell value %d", wsMessage.Data.Cell)
			continue
		}
		if board.Moves()&uint64(1<<uint(wsMessage.Data.Cell)) == 0 {
			log.Printf("invalid move %d", wsMessage.Data.Cell)
			continue
		}
		board.DoMove(wsMessage.Data.Cell)
		updatedState := newState(*board, 1-wsMessage.Data.State.Turn)

		messageOut, err := json.Marshal(updatedState)
		err = c.WriteMessage(mt, messageOut)
		if err != nil {
			log.Printf("write rror: %s", err)
			continue
		}
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	buff, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		log.Printf("error opening file: %s", err)
		return
	}
	w.Write(buff)
}

func Main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/ws", ws)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
