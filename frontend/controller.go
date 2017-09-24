package frontend

import (
	"io"

	"dots/board"
	"dots/players"
)

const (

	// Black represents a field with a black disc
	Black = 0

	// White represents a field with a white disc
	White = 1

	switchTurnMask = Black ^ White

	// MoveBlack represents a field where black can place a disc
	MoveBlack = 2

	// MoveWhite represents a field where white can place a disc
	MoveWhite = MoveBlack | switchTurnMask

	// Empty represents a field without a disc
	Empty = 4
)

// GameState is a Board combined with the color of the player to move
type GameState struct {
	board board.Board
	turn  int
}

// GetFieldValue returns the value of a field
func (state *GameState) GetFieldValue(field uint) int {

	if state.turn != 1 && state.turn != 0 {
		panic("state.turn has impossible value")
	}

	mask := uint64(1) << field

	if state.board.Me()&mask != 0 {
		return Black ^ state.turn
	}

	if state.board.Opp()&mask != 0 {
		return White ^ state.turn
	}

	if state.board.Moves()&mask != 0 {
		return MoveBlack ^ state.turn
	}

	return Empty
}

// Controller is a game controller for an othello game
type Controller struct {
	players  [2]players.Player
	history  []GameState
	stateID  uint
	redoMax  uint
	frontend Frontend
}

// NewController returns a new Controller
func NewController(black, white players.Player, writer io.Writer,
	frontend Frontend) (control *Controller) {
	control = &Controller{
		players:  [2]players.Player{black, white},
		frontend: frontend,
		history:  make([]GameState, 100),
		stateID:  0,
		redoMax:  0}
	return
}

// GetState returns the current state of a Controller
func (control *Controller) GetState() (state GameState) {
	state = control.history[control.stateID]
	return
}

func (control *Controller) setChild(child GameState) {
	control.stateID++
	control.history[control.stateID] = child
	control.redoMax = control.stateID
}

func (control *Controller) skipTurn() {
	child := control.GetState()
	child.board.SwitchTurn()
	child.turn = 1 - child.turn
	control.setChild(child)
}

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

func (control *Controller) canMove() bool {
	return control.GetState().board.Moves() != 0
}

func (control *Controller) reset() {
	control.stateID = 0
	control.redoMax = 0
	control.history[0] = GameState{
		board: *board.NewBoard(),
		turn:  0}
	control.frontend.OnUpdate(control.GetState())
}

func (control *Controller) gameRunning() bool {
	board := control.GetState().board
	return board.Moves() != 0 || board.OpponentMoves() != 0
}

// Undo undoes the last move
func (control *Controller) Undo() {
	if control.stateID != 0 {
		control.stateID--
	}
	control.frontend.OnUpdate(control.GetState())
}

// Redo does the last move again
func (control *Controller) Redo() {
	if control.stateID != control.redoMax {
		control.stateID++
	}
	control.frontend.OnUpdate(control.GetState())
}

// Run is the main loop of the Controller
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
