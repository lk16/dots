package web

type boardState struct {
	White []int `json:"white"`
	Black []int `json:"black"`
	Turn  int   `json:"turn"`
}

type wsMessage struct {
	Event            string            `json:"event"`
	BotMove          *botMoveEvent     `json:"bot_move"`
	BotMoveReply     *botMoveReply     `json:"bot_move_reply"`
	AnalyzeMove      *analyzeMoveEvent `json:"analyze_move"`
	AnalyzeMoveReply *analyzeMoveReply `json:"analyze_move_reply"`
	AnalyzeStop      *analyzeStopEvent `json:"analyze_stop"`
}

type botMoveEvent struct {
	State boardState `json:"state"`
}

type botMoveReply struct {
	State boardState `json:"state"`
}

type analyzeMoveEvent struct {
	State boardState `json:"state"`
}

type analyzeMoveReply struct {
	Board     boardState `json:"board"`
	Move      int        `json:"move"`
	Depth     int        `json:"depth"`
	Heuristic int        `json:"heuristic"`
}

type analyzeStopEvent struct {
	State boardState `json:"state"`
}
