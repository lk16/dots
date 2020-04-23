package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/lk16/dots/internal/othello"
	"github.com/lk16/dots/internal/treesearch"
)

type moveWebSocket struct {
	ws                *websocket.Conn
	writeLock         sync.Mutex
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
	defer func() {
		err := mws.ws.Close()
		if err != nil {
			log.Printf("Error closing websocket: %s", err)
		}
	}()
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
	evaluated := newBoardState(board, turn)

	for i := range analyzedChildren {

		move, ok := board.GetMoveField(children[i])
		if !ok {
			// TODO handle better
			log.Printf("warning: GetMoveField failed")
		}

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

			bot := treesearch.NewPvs(treesearch.Squared)

			if mws.getAnalyzedBoard() != board {
				return
			}

			analysis := analyzeMoveReply{
				Board:     evaluated,
				Depth:     depth,
				Move:      analyzedChildren[i].analysis.Move,
				Heuristic: bot.Search(analyzedChildren[i].child, treesearch.MinHeuristic, treesearch.MaxHeuristic, depth)}

			message := newWsMessage(&analysis)

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

func (mws *moveWebSocket) handleAnalyzeMoveRequest(arg interface{}) (err error) {

	request, ok := arg.(analyzeMoveRequest)
	if !ok {
		return fmt.Errorf("unexpected type %T in handler, expected analyzeMoveRequest", request)
	}

	board, turn, err := request.State.getBoard()
	if err != nil {
		return err
	}

	mws.setAnalyzedBoard(*board)
	go mws.analyze(*board, turn)

	return nil
}

func (mws *moveWebSocket) handlebotMoveRequest(arg interface{}) error {

	request, ok := arg.(botMoveRequest)
	if !ok {
		return fmt.Errorf("unexpected type %T in handler, expected botMoveRequest", request)
	}

	board, _, err := request.State.getBoard()
	if err != nil {
		return err
	}

	if board.Moves() == 0 {
		return fmt.Errorf("no moves available")
	}

	mws.killAnalysis()

	turn := request.State.Turn
	go mws.sendBotMoveReply(*board, turn)

	return nil
}

func (mws *moveWebSocket) sendBotMoveReply(board othello.Board, turn int) {
	bot := treesearch.NewBot(os.Stdout, 8, 14, treesearch.NewPvs(treesearch.Squared))

	bestMove, err := bot.DoMove(board)
	if err != nil {
		log.Printf("sendBotMoveReply(): %s", err)
		return
	}

	nextTurn := 1 - turn
	if board.Moves() == 0 {
		nextTurn = turn
		board.SwitchTurn()
	}

	message := newWsMessage(&botMoveReply{
		State: newBoardState(*bestMove, nextTurn)})

	err = mws.send(message)
	if err != nil {
		log.Printf("sendBotMoveReply(): %s", err)
	}
}

func (mws *moveWebSocket) handleAnalyzeStopRequest(_ interface{}) error {
	mws.killAnalysis()
	return nil
}

func (mws *moveWebSocket) sendGetXotReply() {
	if err := othello.LoadXot(); err != nil {
		log.Printf("error loading xot boards: %s", err.Error())
		return
	}

	board := othello.NewXotBoard()

	message := newWsMessage(&xotReply{
		State: newBoardState(*board, 0)})

	err := mws.send(message)
	if err != nil {
		log.Printf("sendGetXotReply(): %s", err)
	}
}

func (mws *moveWebSocket) handleGetXotEvent(_ interface{}) error {
	mws.killAnalysis()
	go mws.sendGetXotReply()
	return nil
}

func (mws *moveWebSocket) handleMessage(message wsMessage) error {

	handlerMap := map[string]func(interface{}) error{
		"bot_move_request":     mws.handlebotMoveRequest,
		"analyze_move_request": mws.handleAnalyzeMoveRequest,
		"analyze_stop_request": mws.handleAnalyzeStopRequest,
		"xot_request":          mws.handleGetXotEvent}

	handler, ok := handlerMap[message.Event]
	if !ok {
		return fmt.Errorf("unhandled message of event %s", message.Event)
	}

	return handler(message.Data)
}
