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

type moveWebSocket struct {
	ws            *websocket.Conn
	analyzeQuitCh chan struct{}
}

func newMoveWebSocket(w http.ResponseWriter, r *http.Request) (*moveWebSocket, error) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %s", err)
		return nil, err
	}

	mws := &moveWebSocket{
		ws: ws}

	return mws, nil
}

func (mws *moveWebSocket) loop() {
	defer mws.ws.Close()
	for {
		messageType, rawMessage, err := mws.ws.ReadMessage()
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
		log.Printf("%+v", message)
		if err != nil {
			log.Printf("json decode error: %s", err)
			continue
		}
		reply, err := mws.handleMessage(message)
		if err != nil {
			log.Printf("message handling error: %s", err)
			continue
		}
		if reply == nil {
			continue
		}
		rawReply, err := json.Marshal(reply)
		err = mws.ws.WriteMessage(messageType, rawReply)
		if err != nil {
			log.Printf("write rror: %s", err)
			continue
		}
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

	for depth <= board.CountEmpties() {
		sort.Slice(analyzedChildren, func(i, j int) bool {
			return analyzedChildren[i].analysis.Heuristic > analyzedChildren[j].analysis.Heuristic
		})

		for i := range analyzedChildren {

			bot := treesearch.NewMtdf(treesearch.MinHeuristic, treesearch.MaxHeuristic)

			analysis := analyzeMoveReply{
				Depth:     depth,
				Move:      analyzedChildren[i].analysis.Move,
				Heuristic: bot.Search(analyzedChildren[i].child, depth)}

			select {
			case <-quitCh:
				return
			default:
			}

			analyzeCh <- analysis
			analyzedChildren[i].analysis = analysis
		}

		depth++
	}
}

func (mws *moveWebSocket) handleAnalyzeMoveEvent(analyzeMoveEvent *analyzeMove) (err error) {
	if analyzeMoveEvent == nil {
		return fmt.Errorf("analyzeMoveEvent is nil")
	}

	board, err := analyzeMoveEvent.State.getBoard()
	if err != nil {
		return err
	}

	analyzeCh := make(chan analyzeMoveReply)
	mws.analyzeQuitCh = make(chan struct{})

	go analyze(*board, mws.analyzeQuitCh, analyzeCh)

	go func() {
		for analysis := range analyzeCh {

			rawMessage := &wsMessage{
				Event:            "analyze_move_reply",
				AnalyzeMoveReply: &analysis}

			rawReply, err := json.Marshal(rawMessage)
			err = mws.ws.WriteMessage(websocket.TextMessage, rawReply)
			if err != nil {
				log.Printf("write error: %s", err)
				continue
			}
		}
	}()

	return nil
}

func (mws *moveWebSocket) handleBotMoveEvent(botMoveEvent *botMoveEvent) (*wsMessage, error) {

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

func (mws *moveWebSocket) handleMessage(message wsMessage) (*wsMessage, error) {
	switch message.Event {
	case "bot_move":
		return mws.handleBotMoveEvent(message.BotMove)
	case "analyze_move":
		return nil, mws.handleAnalyzeMoveEvent(message.AnalyzeMove)
	case "analyze_stop":
		if mws.analyzeQuitCh != nil {
			mws.analyzeQuitCh <- struct{}{}
		}
		return nil, nil
	default:
		return nil, fmt.Errorf("unhandled message of event %s", message.Event)
	}
}
