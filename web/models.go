package web

type boardState struct {
	White []int `json:"white"`
	Black []int `json:"black"`
	Turn  int   `json:"turn"`
}

type wsMessage struct {
	Event        string        `json:"event"`
	BotMove      *botMoveEvent `json:"bot_move"`
	BotMoveReply *botMoveReply `json:"bot_move_reply"`
}

type botMoveEvent struct {
	State boardState `json:"state"`
}

type botMoveReply struct {
	State boardState `json:"state"`
}
