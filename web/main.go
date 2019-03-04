package web

import (
	"dots/othello"
	"dots/players"
	"dots/treesearch"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"math/bits"
	"net/http"
	"sort"
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

func analyze(board othello.Board, quitCh <-chan struct{}, analyzeCh chan<- analyzeMoveReply) {

	type analyzedChild struct {
		child    othello.Board
		analysis analyzeMoveReply
	}

	children := board.GetChildren()
	analyzedChildren := make([]analyzedChild, len(children))

	for i := range analyzedChildren {

		move := bits.TrailingZeros64(board.Me() | board.Opp() ^ (children[i].Me() | children[i].Opp()))

		analyzedChildren[i] = analyzedChild{
			child: children[i],
			analysis: analyzeMoveReply{
				Depth:     0,
				Heuristic: 0,
				Move:      move}}
	}

	depth := 4

	for {
		sort.Slice(analyzedChildren, func(i, j int) bool {
			return analyzedChildren[i].analysis.Heuristic > analyzedChildren[j].analysis.Heuristic
		})

		for i := range analyzedChildren {

			bot := treesearch.NewMtdf(treesearch.MinHeuristic, treesearch.MaxHeuristic)

			analysis := analyzeMoveReply{
				Depth:     depth,
				Move:      analyzedChildren[i].analysis.Move,
				Heuristic: bot.Search(analyzedChildren[i].child, depth)}

			analyzeCh <- analysis
			analyzedChildren[i].analysis = analysis
		}

		depth++
	}
}

func handleAnalyzeMoveEvent(analyzeMoveEvent *analyzeMove, ws *websocket.Conn) (*wsMessage, error) {
	if analyzeMoveEvent == nil {
		return nil, fmt.Errorf("analyzeMoveEvent is nil")
	}

	board, err := analyzeMoveEvent.State.getBoard()
	if err != nil {
		return nil, err
	}

	analyzeCh := make(chan analyzeMoveReply)
	go analyze(*board, nil, analyzeCh)

	go func() {
		for analysis := range analyzeCh {

			rawMessage := &wsMessage{
				Event:            "analyze_move_reply",
				AnalyzeMoveReply: &analysis}

			rawReply, err := json.Marshal(rawMessage)
			err = ws.WriteMessage(websocket.TextMessage, rawReply)
			if err != nil {
				log.Printf("write error: %s", err)
				continue
			}
		}
	}()

	return nil, nil
}

func handleBotMoveEvent(botMoveEvent *botMoveEvent, _ *websocket.Conn) (*wsMessage, error) {

	if botMoveEvent == nil {
		return nil, fmt.Errorf("botMoveEvent is nil")
	}

	board, err := botMoveEvent.State.getBoard()
	if err != nil {
		return nil, err
	}

	if board.Moves() == 0 {
		return nil, fmt.Errorf("no moves available")
	}

	bot := players.NewBotHeuristic(ioutil.Discard, 8, 16)
	bestMove := bot.DoMove(*board)

	nextTurn := 1 - botMoveEvent.State.Turn
	if board.Moves() == 0 {
		nextTurn = botMoveEvent.State.Turn
		board.SwitchTurn()
	}

	reply := &wsMessage{
		Event: "bot_move_reply",
		BotMoveReply: &botMoveReply{
			State: newState(bestMove, nextTurn)}}

	return reply, nil
}

func handleMessage(message wsMessage, ws *websocket.Conn) (*wsMessage, error) {
	switch message.Event {
	case "bot_move":
		return handleBotMoveEvent(message.BotMove, ws)
	case "analyze_move":
		return handleAnalyzeMoveEvent(message.AnalyzeMove, ws)
	default:
		return nil, fmt.Errorf("unhandled message of event %s", message.Event)
	}
}

func ws(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %s", err)
		return
	}
	defer ws.Close()
	for {
		messageType, rawMessage, err := ws.ReadMessage()

		if err != nil {
			switch err.(type) {
			case *websocket.CloseError:
				// don't print error
			default:
				log.Printf("Unexpected read error %T: %s", err, err)
			}
			break
		}

		var message wsMessage
		err = json.Unmarshal(rawMessage, &message)
		if err != nil {
			log.Printf("json decode error: %s", err)
			continue
		}
		reply, err := handleMessage(message, ws)
		if err != nil {
			log.Printf("message handling error: %s", err)
			continue
		}
		if reply == nil {
			continue
		}
		rawReply, err := json.Marshal(reply)
		err = ws.WriteMessage(messageType, rawReply)
		if err != nil {
			log.Printf("write rror: %s", err)
			continue
		}
	}
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
  <text text-anchor="middle" dominant-baseline="central" font-family="Arial" font-size="25" x="32" y="32">%s</text>
</svg>`

	query := r.URL.Query()

	text := ""
	if textParam, ok := query["text"]; ok {
		text = textParam[0]
	}
	svg := fmt.Sprintf(svgTemplate, text)

	w.Header().Set("Content-Type", "image/svg+xml")
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
