package frontend

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"dots/board"
)

type GtkController struct {
	window     *gtk.Window
	pix_buffs  map[int]*gdk.Pixbuf
	state      chan GameState
	images     [8][8]*gtk.Image
	human_move chan uint
}

// Cache pixel buffers
func cachePixbuf() (cache map[int]*gdk.Pixbuf) {

	var image *gtk.Image
	cache = make(map[int]*gdk.Pixbuf)

	addToCache := func(field_value int, image_path string) {
		image, _ = gtk.ImageNewFromFile(image_path)
		cache[field_value] = image.GetPixbuf()
	}

	addToCache(BLACK, "resources/black.png")
	addToCache(WHITE, "resources/white.png")
	addToCache(EMPTY, "resources/empty.png")
	addToCache(MOVE_BLACK, "resources/move_black.png")
	addToCache(MOVE_WHITE, "resources/move_white.png")

	return
}

// Timeout function for updating the user interface
func timeoutCallback(gtkc *GtkController) bool {
	select {
	case update := <-gtkc.state:
		gtkc.updateFields(update)
	default:
		//fmt.Print("timeout received nothing\n")
	}

	// HACK: add itself to timeout loop again since return value is being ignored
	// this bug has been reported to lib authors, PR pending since 2014
	glib.TimeoutAdd(uint(30), timeoutCallback, gtkc)

	// HACK: true is supposed to be returned but since call back return value is ignored
	// it doesn't matter. This is for forward compatibility.
	return false
}

// Creates a new GtkController
func NewGtk() (gtkc *GtkController) {

	gtk.Init(nil)

	gtkc = &GtkController{
		window:     nil,
		state:      make(chan GameState),
		pix_buffs:  cachePixbuf(),
		human_move: make(chan uint)}

	main_window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	main_window.SetTitle("Dots")
	main_window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	main_vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	main_window.Add(main_vbox)

	board_container, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	main_vbox.Add(board_container)

	for y := uint(0); y < 8; y++ {
		board_row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		board_container.Add(board_row)

		for x := uint(0); x < 8; x++ {
			// create a unique int value for each cell
			field_id := new(uint)
			*field_id = 8*y + x

			board_cell, _ := gtk.EventBoxNew()
			board_cell.Connect("button-press-event", func() {
				gtkc.human_move <- *field_id
			})

			board_row.Add(board_cell)

			pix_buff, _ := gtkc.pix_buffs[EMPTY]
			image, _ := gtk.ImageNewFromPixbuf(pix_buff)
			board_cell.Add(image)

			gtkc.images[y][x] = image
		}
	}

	gtkc.updateFields(GameState{
		board: *board.NewBoard(),
		turn:  0})

	gtkc.window = main_window
	gtkc.window.ShowAll()

	glib.TimeoutAdd(uint(30), timeoutCallback, gtkc)

	go gtk.Main()

	return
}

// Updates fields of user interface to represent a GameState
func (gtkc *GtkController) updateFields(state GameState) {

	for f := uint(0); f < 64; f++ {
		image := gtkc.images[f/8][f%8]
		field_value := state.GetFieldValue(f)
		image.SetFromPixbuf(gtkc.pix_buffs[field_value])
	}

}

// Update user interface on GameState change
func (gtkc *GtkController) OnUpdate(state GameState) {
	gtkc.state <- state
}

// Update user interface on game end
func (gtkc *GtkController) OnGameEnd(state GameState) {
	gtkc.state <- state

	// click for new game
	<-gtkc.human_move
}

// Wait for a human to move
func (gtkc *GtkController) OnHumanMove(state GameState) (afterwards board.Board) {
	moves := state.board.Moves()
	for {
		field_id := <-gtkc.human_move
		if moves.TestBit(field_id) {
			state.board.DoMove(field_id)
			state.turn = 1 - state.turn
			gtkc.state <- state
			afterwards = state.board
			return
		}
	}
}
