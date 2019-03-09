package web

import (
	"dots/othello"
	"dots/treesearch"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{}

func newState(board othello.Board, turn int) boardState {

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
		return boardState{
			Black: me,
			White: opp,
			Turn:  0}
	}

	return boardState{
		White: me,
		Black: opp,
		Turn:  1}
}

func (s *boardState) getBoard() (*othello.Board, error) {

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

func ws(w http.ResponseWriter, r *http.Request) {
	mws, err := newMoveWebSocket(w, r)
	if err != nil {
		log.Printf("error creating MoveWebSocket: %s", err)
		return
	}
	mws.loop()
}

func root(w http.ResponseWriter, _ *http.Request) {
	buff, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		log.Printf("error opening file: %s", err)
		return
	}
	w.Write(buff)
}

func svgGenerator(w http.ResponseWriter, r *http.Request) {
	svgTemplate := `<?xml version="1.0" encoding="UTF-8" ?>
<svg width="64" height="64" xmlns="http://www.w3.org/2000/svg">	
  <rect x="0" y="0" width="64" height="64" fill="green" stroke-width="1" stroke="black" />
  <text text-anchor="middle" dominant-baseline="central" %s font-family="Arial" font-size="%d" x="32" y="32">%s</text>
</svg>`

	query := r.URL.Query()

	text := ""
	if textParam, ok := query["text"]; ok {
		text = textParam[0]
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	textInt, err := strconv.Atoi(text)

	var svg string
	if err != nil {
		svg = fmt.Sprintf(svgTemplate, "", 25, "???")
	} else if textInt%treesearch.ExactScoreFactor == 0 && textInt != 0 {
		svg = fmt.Sprintf(svgTemplate, "font-weight=\"bold\"", 38, fmt.Sprintf("%d", textInt/treesearch.ExactScoreFactor))
	} else {
		svg = fmt.Sprintf(svgTemplate, "", 25, text)
	}

	w.Write([]byte(svg))
}

func Main() {
	http.HandleFunc("/ws", ws)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/svg/", svgGenerator)
	http.HandleFunc("/", root)
	addr := "localhost:8080"
	log.Printf("Server running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
