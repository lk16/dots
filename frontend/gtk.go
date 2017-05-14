package frontend

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"dots/board"
)

type GtkUpdater struct {
	window     *gtk.Window
	pix_buffs  map[int]*gdk.Pixbuf
	ch         chan GameState
	images     [8][8]*gtk.Image
	human_move chan uint
}

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

func timeoutCallback(updater *GtkUpdater) bool {
	select {
	case update := <-updater.ch:
		updater.updateFields(update)
	default:
		//fmt.Print("timeout received nothing\n")
	}

	// HACK: add itself to timeout loop again since return value is being ignored
	// this bug has been reported to lib authors, PR pending since 2014
	glib.TimeoutAdd(uint(30), timeoutCallback, updater)

	// HACK: true is supposed to be returned but since call back return value is ignored
	// it doesn't matter. This is for forward compatibility.
	return false
}

func NewGtk() (updater *GtkUpdater) {

	gtk.Init(nil)

	updater = &GtkUpdater{
		window:     nil,
		ch:         make(chan GameState),
		pix_buffs:  cachePixbuf(),
		human_move: make(chan uint)}

	main_window, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	main_window.SetTitle("Dots")
	main_window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	main_vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	main_window.Add(main_vbox)

	//menu, _ := gtk.MenuNew()
	//main_vbox.Add(menu)

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
				updater.human_move <- *field_id
			})

			board_row.Add(board_cell)

			pix_buff, _ := updater.pix_buffs[EMPTY]
			image, _ := gtk.ImageNewFromPixbuf(pix_buff)
			board_cell.Add(image)

			updater.images[y][x] = image
		}
	}

	updater.updateFields(GameState{
		board: *board.NewBoard(),
		turn:  0})

	updater.window = main_window
	updater.window.ShowAll()

	glib.TimeoutAdd(uint(30), timeoutCallback, updater)

	go gtk.Main()

	return
}

func (updater *GtkUpdater) updateFields(state GameState) {

	for f := uint(0); f < 64; f++ {
		image := updater.images[f/8][f%8]
		field_value := state.GetFieldValue(f)
		image.SetFromPixbuf(updater.pix_buffs[field_value])
	}

}

func (updater *GtkUpdater) OnUpdate(state GameState) {
	updater.ch <- state
}

func (updater *GtkUpdater) OnGameEnd(state GameState) {
	updater.ch <- state

	// click for new game
	<-updater.human_move
}

func (updater *GtkUpdater) OnHumanMove(state GameState) (afterwards board.Board) {
	moves := state.board.Moves()
	for {
		field_id := <-updater.human_move
		if moves.TestBit(field_id) {
			state.board.DoMove(field_id)
			state.turn = 1 - state.turn
			updater.ch <- state
			afterwards = state.board
			return
		}
	}
}
