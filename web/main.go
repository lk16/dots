package web

import (
	"dots/othello"
	"dots/treesearch"
	"fmt"
	"github.com/ajstarks/svgo"
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

func (s *boardState) getBoard() (*othello.Board, int, error) {

	white := uint64(0)
	black := uint64(0)

	for _, w := range s.White {
		if w < 0 || w >= 64 {
			return nil, 0, fmt.Errorf("invalid white field value %d", w)
		}
		white |= uint64(1 << uint(w))
	}

	for _, b := range s.Black {
		if b < 0 || b >= 64 {
			return nil, 0, fmt.Errorf("invalid black field value %d", b)
		}
		black |= uint64(1 << uint(b))
	}

	if white&black != 0 {
		return nil, 0, fmt.Errorf("white (%+v) and black (%+v) overlap", white, black)
	}

	switch s.Turn {
	case 0:
		return othello.CustomBoard(black, white), s.Turn, nil
	case 1:
		return othello.CustomBoard(white, black), s.Turn, nil
	default:
		return nil, 0, fmt.Errorf("invalid turn value %d", s.Turn)
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	mws, err := newMoveWebSocket(w, r)
	if err != nil {
		log.Printf("error creating MoveWebSocket: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
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

func svgField(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "image/svg+xml")
	size := 64

	query := r.URL.Query()

	text := query.Get("text")
	textInt, err := strconv.Atoi(text)
	if text != "" && err != nil {
		text = "???"
	}

	disc := query.Get("disc")
	move := query.Get("move")
	textColor := query.Get("textcolor")

	canvas := svg.New(w)
	canvas.Start(size, size)
	canvas.Rect(0, 0, size, size, "fill='green' stroke-width='1' stroke='black'")

	textStyleAttrs := []string{
		"text-anchor='middle'",
		"dominant-baseline='central'",
		"font-family='Arial'"}

	if textInt%treesearch.ExactScoreFactor == 0 && textInt != 0 {
		text = fmt.Sprintf("%d", textInt/treesearch.ExactScoreFactor)
		textStyleAttrs = append(textStyleAttrs, []string{
			"font-size='40'",
			"font-weight='bold'"}...)
	} else {
		textStyleAttrs = append(textStyleAttrs, "font-size='25'")
		if disc == "black" || textColor == "white" {
			textStyleAttrs = append(textStyleAttrs, "fill='white'")
		}
	}

	switch disc {
	case "white":
		canvas.Circle(size/2, size/2, 25, "fill='white'")
	case "black":
		canvas.Circle(size/2, size/2, 25, "fill='black'")
	}

	switch move {
	case "white":
		canvas.Circle(size/2, size/2, 6, "fill='white'")
	case "black":
		canvas.Circle(size/2, size/2, 6, "fill='black'")
	}

	canvas.Text(size/2, size/2, text, textStyleAttrs...)
	canvas.End()
}

func svgIcon(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	size := 64

	canvas := svg.New(w)
	canvas.Start(size, size)
	canvas.Rect(0, 0, size, size, "fill='green' stroke-width='1' stroke='black'")
	canvas.Circle(3*size/10, size/2, size/5, "fill='white'")
	canvas.Circle(7*size/10, size/2, size/5, "fill='black'")
	canvas.End()
}

func Main() {
	http.HandleFunc("/ws", ws)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/svg/field/", svgField)
	http.HandleFunc("/svg/icon/", svgIcon)
	http.HandleFunc("/", root)
	addr := "localhost:8080"
	log.Printf("Server running at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
