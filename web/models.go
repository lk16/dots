package web

import (
	"encoding/json"
	"fmt"
)

type boardState struct {
	White []int `json:"white"`
	Black []int `json:"black"`
	Turn  int   `json:"turn"`
}

type wsReply interface {
	GetEventName() string
}

type wsMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func newWsMessage(reply wsReply) *wsMessage {
	return &wsMessage{
		Event: reply.GetEventName(),
		Data:  reply}
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

type botMoveRequest struct {
	State boardState `json:"state"`
}

type botMoveReply struct {
	State boardState `json:"state"`
}

func (reply botMoveReply) GetEventName() string {
	return "bot_move_reply"
}

type analyzeMoveRequest struct {
	State boardState `json:"state"`
}

type analyzeMoveReply struct {
	Board     boardState `json:"board"`
	Move      int        `json:"move"`
	Depth     int        `json:"depth"`
	Heuristic int        `json:"heuristic"`
}

func (reply analyzeMoveReply) GetEventName() string {
	return "analyze_move_reply"
}

type analyzeStopRequest struct{}

type xotReply struct {
	State boardState `json:"state"`
}

func (reply xotReply) GetEventName() string {
	return "xot_reply"
}
