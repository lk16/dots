package web

import (
	"encoding/json"
	"fmt"

	"github.com/lk16/dots/internal/othello"
)

func newWsMessage(reply wsReply) *wsMessage {
	return &wsMessage{
		Event: reply.GetEventName(),
		Data:  reply}
}

func (reply botMoveReply) GetEventName() string {
	return "bot_move_reply"
}

func (reply *analyzeMoveReply) GetEventName() string {
	return "analyze_move_reply"
}

func (reply xotReply) GetEventName() string {
	return "xot_reply"
}

func (m *wsMessage) UnmarshalJSON(data []byte) error {
	// alias to prevent infinite recursion
	type alias wsMessage
	var aliased alias
	err := json.Unmarshal(data, &aliased)
	if err != nil {
		return err
	}

	m.Event = aliased.Event

	switch m.Event {
	case "bot_move_request":
		wrapper := struct {
			Data botMoveRequest `json:"data"`
		}{}

		err := json.Unmarshal(data, &wrapper)
		if err != nil {
			return err
		}
		m.Data = wrapper.Data

	case "analyze_move_request":
		wrapper := struct {
			Data analyzeMoveRequest `json:"data"`
		}{}

		err := json.Unmarshal(data, &wrapper)
		if err != nil {
			return err
		}
		m.Data = wrapper.Data

	// events without data
	case "xot_request":
	case "analyze_stop_request":

	default:
		return fmt.Errorf("event \"%s\" is not handled for json decoding", m.Event)
	}

	return nil
}

func newBoardState(board othello.Board, turn int) boardState {
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
	white := othello.BitSet(0)
	black := othello.BitSet(0)

	for _, w := range s.White {
		if w < 0 || w >= 64 {
			return nil, 0, fmt.Errorf("invalid white field value %d", w)
		}
		white.Set(uint(w))
	}

	for _, b := range s.Black {
		if b < 0 || b >= 64 {
			return nil, 0, fmt.Errorf("invalid black field value %d", b)
		}
		black.Set(uint(b))
	}

	if white&black != 0 {
		return nil, 0, fmt.Errorf("white (%+v) and black (%+v) overlap", white, black)
	}

	switch s.Turn {
	case 0:
		return othello.NewCustomBoard(black, white), s.Turn, nil
	case 1:
		return othello.NewCustomBoard(white, black), s.Turn, nil
	default:
		return nil, 0, fmt.Errorf("invalid turn value %d", s.Turn)
	}
}
