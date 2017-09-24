package frontend

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"dots/board"
)

// GtkFrontend is a frontend for GTK
type GtkFrontend struct {
	window    *gtk.Window
	pixBuffs  map[int]*gdk.Pixbuf
	state     chan GameState
	images    [8][8]*gtk.Image
	humanMove chan uint
}

func cachePixbuf() (cache map[int]*gdk.Pixbuf) {

	cache = make(map[int]*gdk.Pixbuf)

	addToCache := func(field_value int, image_path string) {
		image, _ := gtk.ImageNewFromFile(image_path)
		cache[field_value] = image.GetPixbuf()
	}

	addToCache(Black, "resources/black.png")
	addToCache(White, "resources/white.png")
	addToCache(Empty, "resources/empty.png")
	addToCache(MoveBlack, "resources/move_black.png")
	addToCache(MoveWhite, "resources/move_white.png")

	return
}

// Timeout function for updating the user interface
func timeoutCallback(gtkf *GtkFrontend) bool {
	select {
	case update := <-gtkf.state:
		gtkf.updateFields(update)
	default:
		//fmt.Print("timeout received nothing\n")
	}

	// HACK: add itself to timeout loop again since return value is being ignored
	// this bug has been reported to lib authors, PR pending since 2014
	glib.TimeoutAdd(uint(30), timeoutCallback, gtkf)

	// HACK: true is supposed to be returned but since call back return value is ignored
	// it doesn't matter. This is for forward compatibility.
	return false
}

// NewkGtkFrontend returns a new GtkFrontend
func NewkGtkFrontend() Frontend {

	gtk.Init(nil)

	gtkf := &GtkFrontend{
		window:    nil,
		state:     make(chan GameState),
		pixBuffs:  cachePixbuf(),
		humanMove: make(chan uint)}

	mainWindow, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	mainWindow.SetTitle("Dots")
	mainWindow.Connect("destroy", func() {
		gtk.MainQuit()
	})

	mainVbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	mainWindow.Add(mainVbox)

	boardContainer, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	mainVbox.Add(boardContainer)

	for y := uint(0); y < 8; y++ {
		boardRow, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		boardContainer.Add(boardRow)

		for x := uint(0); x < 8; x++ {
			// create a unique int value for each cell
			fieldID := new(uint)
			*fieldID = 8*y + x

			boardCell, _ := gtk.EventBoxNew()
			boardCell.Connect("button-press-event", func() {
				gtkf.humanMove <- *fieldID
			})

			boardRow.Add(boardCell)

			pixBuff, _ := gtkf.pixBuffs[Empty]
			image, _ := gtk.ImageNewFromPixbuf(pixBuff)
			boardCell.Add(image)

			gtkf.images[y][x] = image
		}
	}

	gtkf.updateFields(GameState{
		board: *board.NewBoard(),
		turn:  0})

	gtkf.window = mainWindow
	gtkf.window.ShowAll()

	glib.TimeoutAdd(uint(30), timeoutCallback, gtkf)

	go gtk.Main()

	return gtkf
}

// Updates fields of user interface to represent a GameState
func (gtkf *GtkFrontend) updateFields(state GameState) {

	for f := uint(0); f < 64; f++ {
		image := gtkf.images[f/8][f%8]
		fieldValue := state.GetFieldValue(f)
		image.SetFromPixbuf(gtkf.pixBuffs[fieldValue])
	}

}

// OnUpdate updates the Gtk user interface
func (gtkf *GtkFrontend) OnUpdate(state GameState) {
	gtkf.state <- state
}

// OnGameEnd updates the Gtk user interface on game end
func (gtkf *GtkFrontend) OnGameEnd(state GameState) {
	gtkf.state <- state

	// click for new game
	<-gtkf.humanMove
}

// OnHumanMove waits for a human to move
func (gtkf *GtkFrontend) OnHumanMove(state GameState) (afterwards board.Board) {
	moves := state.board.Moves()
	for {
		fieldID := <-gtkf.humanMove
		if moves&(uint64(1)<<fieldID) != 0 {
			state.board.DoMove(int(fieldID))
			state.turn = 1 - state.turn
			gtkf.state <- state
			afterwards = state.board
			return
		}
	}
}
