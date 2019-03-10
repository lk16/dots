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
	"sync"
)

type moveWebSocket struct {
	ws                *websocket.Conn
	writeLock         sync.Mutex
	analyzeQuitCh     chan struct{}
	analyzedBoard     othello.Board
	analyzedBoardLock sync.Mutex
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

func (mws *moveWebSocket) getAnalyzedBoard() othello.Board {
	mws.analyzedBoardLock.Lock()
	defer mws.analyzedBoardLock.Unlock()
	return mws.analyzedBoard
}

func (mws *moveWebSocket) setAnalyzedBoard(board othello.Board) {
	mws.analyzedBoardLock.Lock()
	defer mws.analyzedBoardLock.Unlock()
	mws.analyzedBoard = board
}

func (mws *moveWebSocket) killAnalysis() {
	mws.setAnalyzedBoard(othello.Board{})
}

func (mws *moveWebSocket) send(message *wsMessage) error {

	if message == nil {
		return nil
	}

	mws.writeLock.Lock()
	defer mws.writeLock.Unlock()

	rawMessage, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("JSON ecoding error: %s", err)
	}

	return mws.ws.WriteMessage(websocket.TextMessage, rawMessage)
}

func (mws *moveWebSocket) loop() {
	defer mws.ws.Close()
	for {
		_, rawMessage, err := mws.ws.ReadMessage()
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

		log.Printf("mws received %s", message.Event)

		err = mws.handleMessage(message)
		if err != nil {
			log.Printf("message handling error: %s", err)
			continue
		}
	}
}

func (mws *moveWebSocket) analyze(board othello.Board, turn int) {

	type analyzedChild struct {
		child    othello.Board
		analysis analyzeMoveReply
	}

	children := board.GetChildren()
	analyzedChildren := make([]analyzedChild, len(children))
	evaluated := newState(board, turn)

	for i := range analyzedChildren {

		move := bits.TrailingZeros64(board.Me() | board.Opp() ^ (children[i].Me() | children[i].Opp()))

		analyzedChildren[i] = analyzedChild{
			child: children[i],
			analysis: analyzeMoveReply{
				Board:     evaluated,
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

			if mws.getAnalyzedBoard() != board {
				return
			}

			analysis := analyzeMoveReply{
				Board:     evaluated,
				Depth:     depth,
				Move:      analyzedChildren[i].analysis.Move,
				Heuristic: bot.Search(analyzedChildren[i].child, depth)}

			message := &wsMessage{
				Event:            "analyze_move_reply",
				AnalyzeMoveReply: &analysis}

			if mws.getAnalyzedBoard() != board {
				return
			}

			err := mws.send(message)
			if err != nil {
				if err != websocket.ErrCloseSent {
					log.Printf("Unexpected write error %T: %s", err, err)

				}
				return
			}

			analyzedChildren[i].analysis = analysis
		}

		depth++
	}
}

func (mws *moveWebSocket) handleAnalyzeMoveEvent(analyzeMoveEvent *analyzeMoveEvent) (err error) {

	if analyzeMoveEvent == nil {
		return fmt.Errorf("analyzeMoveEvent is nil")
	}

	board, turn, err := analyzeMoveEvent.State.getBoard()
	if err != nil {
		return err
	}

	mws.setAnalyzedBoard(*board)
	go mws.analyze(*board, turn)

	return nil
}

func (mws *moveWebSocket) handleBotMoveEvent(botMoveEvent *botMoveEvent) error {

	if botMoveEvent == nil {
		return fmt.Errorf("botMoveEvent is nil")
	}

	board, _, err := botMoveEvent.State.getBoard()
	if err != nil {
		return err
	}

	if board.Moves() == 0 {
		return fmt.Errorf("no moves available")
	}

	mws.killAnalysis()

	turn := botMoveEvent.State.Turn
	go mws.sendBotMoveReply(*board, turn)

	return nil
}

func (mws *moveWebSocket) sendBotMoveReply(board othello.Board, turn int) {
	bot := players.NewBotHeuristic(ioutil.Discard, 8, 16)
	bestMove := bot.DoMove(board)

	nextTurn := 1 - turn
	if board.Moves() == 0 {
		nextTurn = turn
		board.SwitchTurn()
	}

	message := &wsMessage{
		Event: "bot_move_reply",
		BotMoveReply: &botMoveReply{
			State: newState(bestMove, nextTurn)}}

	mws.send(message)

}

func (mws *moveWebSocket) handleAnalyzeStopEvent(_ *analyzeStopEvent) error {
	mws.killAnalysis()
	return nil
}

func (mws *moveWebSocket) sendGetXotReply() {
	board := othello.NewXotBoard()

	message := &wsMessage{
		Event: "get_xot_reply",
		GetXotReply: &getXotReply{
			State: newState(board, 0)}}

	mws.send(message)
}

func (mws *moveWebSocket) handleGetXotEvent() error {
	mws.killAnalysis()
	go mws.sendGetXotReply()
	return nil
}

func (mws *moveWebSocket) handleMessage(message wsMessage) error {
	switch message.Event {
	case "bot_move":
		return mws.handleBotMoveEvent(message.BotMove)
	case "analyze_move":
		return mws.handleAnalyzeMoveEvent(message.AnalyzeMove)
	case "analyze_stop":
		return mws.handleAnalyzeStopEvent(message.AnalyzeStop)
	case "get_xot":
		return mws.handleGetXotEvent()
	default:
		return fmt.Errorf("unhandled message of event %s", message.Event)
	}
}
