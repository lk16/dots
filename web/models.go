package web

type boardState struct {
	White []int `json:"white"`
	Black []int `json:"black"`
	Turn  int   `json:"turn"`
}

type wsMessage struct {
	Event      string      `json:"event"`
	Click      *clickEvent `json:"click"`
	ClickReply *clickReply `json:"click_reply"`
}

type clickEvent struct {
	Cell  int        `json:"cell"`
	State boardState `json:"state"`
}

type clickReply struct {
	NewState boardState `json:"state"`
}
