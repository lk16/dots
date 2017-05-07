package frontend

import (
	"io"

	"dots/board"
	"dots/players"
)

type GameState struct {
	board board.Board
	turn  int
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
	frontend Frontend) (cont *Controller) {
	cont = &Controller{
		players:  [2]players.Player{black, white},
		frontend: frontend,
		history:  make([]GameState, 100),
		state_id: 0,
		redo_max: 0}

	if white == nil || black == nil {
		panic("A player cannot be nil!")
	}

	return
}

func (cont *Controller) GetState() (state GameState) {
	state = cont.history[cont.state_id]
	return
}

func (cont *Controller) setChild(child GameState) {

	cont.state_id++
	cont.history[cont.state_id] = child
	cont.redo_max = cont.state_id
}

// Skips a turn (for when a player has no moves)
func (cont *Controller) skipTurn() {

	child := cont.GetState()
	child.board.SwitchTurn()
	child.turn = 1 - child.turn
	cont.setChild(child)
}

// Lets the player to move do a move
func (cont *Controller) doMove() {

	child := cont.GetState()
	player := cont.players[child.turn]

	child.board = player.DoMove(child.board)
	child.turn = 1 - child.turn
	cont.setChild(child)
}

// Returns whether the player to move can do a move
func (cont *Controller) canMove() (can_move bool) {
	moves_count := cont.GetState().board.Moves().Count()
	can_move = (moves_count != 0)
	return
}

// Resets for a new game
func (cont *Controller) reset() {
	cont.state_id = 0
	cont.redo_max = 0
	cont.history[0] = GameState{
		board: *board.NewBoard(),
		turn:  0}
	cont.frontend.OnUpdate(cont.GetState())
}

// Checks if a game is running
func (cont *Controller) gameRunning() (running bool) {
	return !cont.GetState().board.IsLeaf()
}

func (cont *Controller) Undo() {
	if cont.state_id != 0 {
		cont.state_id--
	}
	cont.frontend.OnUpdate(cont.GetState())
}

func (cont *Controller) Redo() {
	if cont.state_id != cont.redo_max {
		cont.state_id++
	}
	cont.frontend.OnUpdate(cont.GetState())
}

// Runs the game
func (cont *Controller) Run() {

	cont.reset()
	for cont.gameRunning() {

		if cont.canMove() {
			cont.doMove()
		} else {
			cont.skipTurn()
		}
		cont.frontend.OnUpdate(cont.GetState())

	}
	cont.frontend.OnGameEnd(cont.GetState())
}
