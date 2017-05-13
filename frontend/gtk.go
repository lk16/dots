package frontend

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"dots/board"
	"fmt"
)

type GtkUpdater struct {
	window    *gtk.Window
	pix_buffs map[string]*gdk.Pixbuf
	ch        chan GameState
	images    [8][8]*gtk.Image
}

type WindowUpdate struct {
	state    GameState
	game_end bool
}

func cachePixbuf() (cache map[string]*gdk.Pixbuf) {

	var image *gtk.Image
	cache = make(map[string]*gdk.Pixbuf)

	image, _ = gtk.ImageNewFromFile("resources/black.png")
	cache["black"] = image.GetPixbuf()

	image, _ = gtk.ImageNewFromFile("resources/white.png")
	cache["white"] = image.GetPixbuf()

	image, _ = gtk.ImageNewFromFile("resources/empty.png")
	cache["empty"] = image.GetPixbuf()

	image, _ = gtk.ImageNewFromFile("resources/move_white.png")
	cache["move_white"] = image.GetPixbuf()

	image, _ = gtk.ImageNewFromFile("resources/move_black.png")
	cache["move_black"] = image.GetPixbuf()

	return
}

func timeoutCallback(updater *GtkUpdater) bool {
	select {
	case update := <-updater.ch:
		updater.updateFields(update)
	default:
		fmt.Print("timeout received nothing\n")
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
		window:    nil,
		ch:        make(chan GameState),
		pix_buffs: cachePixbuf()}

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

	for y := 0; y < 8; y++ {
		board_row, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
		board_container.Add(board_row)

		for x := 0; x < 8; x++ {
			board_cell, _ := gtk.EventBoxNew()
			board_row.Add(board_cell)

			pix_buff, _ := updater.pix_buffs["empty"]
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

	for y := uint(0); y < 8; y++ {
		for x := uint(0); x < 8; x++ {
			image := updater.images[y][x]

			switch state.GetField(x, y) {
			case BLACK:
				image.SetFromPixbuf(updater.pix_buffs["black"])
			case WHITE:
				image.SetFromPixbuf(updater.pix_buffs["white"])
			case MOVE_BLACK:
				image.SetFromPixbuf(updater.pix_buffs["move_black"])
			case MOVE_WHITE:
				image.SetFromPixbuf(updater.pix_buffs["move_white"])
			default:
				image.SetFromPixbuf(updater.pix_buffs["empty"])
			}
		}
	}

}

func (updater *GtkUpdater) OnUpdate(state GameState) {
	updater.ch <- state
}

func (updater *GtkUpdater) OnGameEnd(state GameState) {
	updater.ch <- state
}

func (updater *GtkUpdater) OnHumanMove(state GameState) (afterwards board.Board) {
	afterwards = state.board
	return
}
