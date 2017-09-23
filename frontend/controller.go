package frontend

import (
	"io"

	"dots/board"
	"dots/players"

	"math/bits"
)

const (
	SWITCH_TURN_MASK = 1
	BLACK            = 0
	WHITE            = BLACK | SWITCH_TURN_MASK
	MOVE_BLACK       = 2
	MOVE_WHITE       = MOVE_BLACK | SWITCH_TURN_MASK
	EMPTY            = 4
)

type GameState struct {
	board board.Board
	turn  int
}

// Returns field value
func (state *GameState) GetFieldValue(field uint) (field_value int) {

	if state.turn != 1 && state.turn != 0 {
		panic("state.turn has impossible value")
	}

	mask := uint64(1) << field

	if state.board.Me()&mask != 0 {
		field_value = BLACK | state.turn
		return
	}

	if state.board.Opp()&mask != 0 {
		field_value = WHITE ^ state.turn
		return
	}

	if state.board.Moves()&mask != 0 {
		field_value = MOVE_BLACK | state.turn
		return
	}

	field_value = EMPTY
	return
}

type Controller struct {
	players  [2]players.Player
	history  []GameState
	state_id uint
	redo_max uint
	frontend Frontend
}

// Returns a new Controller
func NewController(black, white players.Player, writer io.Writer,
	frontend Frontend) (control *Controller) {
	control = &Controller{
		players:  [2]players.Player{black, white},
		frontend: frontend,
		history:  make([]GameState, 100),
		state_id: 0,
		redo_max: 0}
	return
}

func (control *Controller) GetState() (state GameState) {
	state = control.history[control.state_id]
	return
}

func (control *Controller) setChild(child GameState) {

	control.state_id++
	control.history[control.state_id] = child
	control.redo_max = control.state_id
}

// Skips a turn (for when a player has no moves)
func (control *Controller) skipTurn() {

	child := control.GetState()
	child.board.SwitchTurn()
	child.turn = 1 - child.turn
	control.setChild(child)
}

// Lets the player to move do a move
func (control *Controller) doMove() {

	child := control.GetState()
	player := control.players[child.turn]

	if player == nil {
		child.board = control.frontend.OnHumanMove(child)
	} else {
		child.board = player.DoMove(child.board)
	}
	child.turn = 1 - child.turn
	control.setChild(child)
}

// Returns whether the player to move can do a move
func (control *Controller) canMove() (can_move bool) {
	moves_count := bits.OnesCount64(control.GetState().board.Moves())
	can_move = (moves_count != 0)
	return
}

// Resets for a new game
func (control *Controller) reset() {
	control.state_id = 0
	control.redo_max = 0
	control.history[0] = GameState{
		board: *board.NewBoard(),
		turn:  0}
	control.frontend.OnUpdate(control.GetState())
}

// Checks if a game is running
func (control *Controller) gameRunning() bool {
	board := control.GetState().board
	return board.Moves() != 0 || board.OpponentMoves() != 0
}

// Undo last move
func (control *Controller) Undo() {
	if control.state_id != 0 {
		control.state_id--
	}
	control.frontend.OnUpdate(control.GetState())
}

// Undo last undo
func (control *Controller) Redo() {
	if control.state_id != control.redo_max {
		control.state_id++
	}
	control.frontend.OnUpdate(control.GetState())
}

// Runs the game
func (control *Controller) Run() {
	for {
		control.reset()
		for control.gameRunning() {

			if control.canMove() {
				control.doMove()
			} else {
				control.skipTurn()
			}
			control.frontend.OnUpdate(control.GetState())

		}
		control.frontend.OnGameEnd(control.GetState())
	}
}
