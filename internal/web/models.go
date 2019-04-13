package web

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

type botMoveRequest struct {
	State boardState `json:"state"`
}

type botMoveReply struct {
	State boardState `json:"state"`
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

type xotReply struct {
	State boardState `json:"state"`
}
