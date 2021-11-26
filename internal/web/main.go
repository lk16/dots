// Package web contains all code for the frontend to work
package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
	"github.com/gorilla/websocket"
	"github.com/lk16/dots/internal/othello"
	"github.com/lk16/dots/internal/treesearch"
)

const (
	white = "white"
	black = "black"
)

var upgrader = websocket.Upgrader{}

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
	buff, err := ioutil.ReadFile("internal/web/index.html")
	if err != nil {
		log.Printf("error opening file: %s", err)
		return
	}
	_, err = w.Write(buff)
	if err != nil {
		log.Printf("response writer Write() failed: %s", err)
	}
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
		if textColor == white {
			textStyleAttrs = append(textStyleAttrs, "fill='white'")
		}
	} else {
		textStyleAttrs = append(textStyleAttrs, "font-size='25'")
		if disc == black || textColor == white {
			textStyleAttrs = append(textStyleAttrs, "fill='white'")
		}
	}

	switch disc {
	case white:
		canvas.Circle(size/2, size/2, 25, "fill='white'")
	case black:
		canvas.Circle(size/2, size/2, 25, "fill='black'")
	}

	switch move {
	case white:
		canvas.Circle(size/2, size/2, 6, "fill='white'")
	case black:
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

func getHeuristic(w http.ResponseWriter, r *http.Request) {
	var (
		me, opp       uint64
		err           error
		search, exact int64
	)

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)

	me, err = strconv.ParseUint(strings.ReplaceAll(r.URL.Query().Get("me"), "0x", ""), 16, 64)

	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(HeuristicError{Where: "me", Error: err.Error()})
		return
	}

	opp, err = strconv.ParseUint(strings.ReplaceAll(r.URL.Query().Get("opp"), "0x", ""), 16, 64)

	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(HeuristicError{Where: "opp", Error: err.Error()})
		return
	}

	search, err = strconv.ParseInt(r.URL.Query().Get("search"), 10, 64)

	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(HeuristicError{Where: "search", Error: err.Error()})
		return
	}

	exact, err = strconv.ParseInt(r.URL.Query().Get("exact"), 10, 64)

	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(HeuristicError{Where: "exact", Error: err.Error()})
		return
	}

	board := othello.NewCustomBoard(othello.BitSet(me), othello.BitSet(opp))

	bot := treesearch.NewBot(os.Stdout, int(search), int(exact), treesearch.NewPvs(treesearch.Squared))

	bestChild, heur, err := bot.DoMove(*board)

	if err != nil {
		w.WriteHeader(400)
		encoder.Encode(HeuristicError{Where: "do_move", Error: err.Error()})
		return
	}

	discsBefore := board.Me() | board.Opp()
	discsAfter := bestChild.Me() | bestChild.Opp()

	newDiscMask := discsBefore ^ discsAfter
	bestMove := newDiscMask.Lowest()

	response := HeuristicResponse{
		Me:        me,
		Opp:       opp,
		Search:    uint64(search),
		Exact:     uint64(exact),
		BestMove:  bestMove,
		Heuristic: heur,
	}

	encoder.Encode(response)
}

// TODO replace by something better
type logger struct{}

func (logger logger) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.URL.String())
	http.DefaultServeMux.ServeHTTP(w, req)
}

// Main initializes and runs the dots webserver
func Main() {
	http.HandleFunc("/ws", ws)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("internal/web/static"))))
	http.HandleFunc("/svg/field", svgField)
	http.HandleFunc("/svg/icon", svgIcon)
	http.HandleFunc("/heuristic", getHeuristic)
	http.HandleFunc("/", root)
	addr := "0.0.0.0:8080"
	log.Printf("Server running at %s", addr)

	log.Fatal(http.ListenAndServe(addr, &logger{}))
}
