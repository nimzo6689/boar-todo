package ui

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Window struct {
	app *tview.Application

	layout *ModalLayout
	grid   *tview.Grid
}

func (w *Window) Draw(screen tcell.Screen) {
	w.grid.Draw(screen)
}

func (w *Window) GetRect() (int, int, int, int) {
	return w.grid.GetRect()
}

func (w *Window) SetRect(x, y, width, height int) {
	w.grid.SetRect(x, y, width, height)
}

func (w *Window) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		w.grid.InputHandler()(event, setFocus)
	}
}

func (w *Window) Focus(delegate func(p tview.Primitive)) {
	w.grid.Focus(delegate)
}

func (w *Window) HasFocus() bool {
	return w.grid.HasFocus()
}

func (w *Window) Blur() {
	w.grid.Blur()
}

func (w *Window) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return w.grid.MouseHandler()
}

func (w *Window) PasteHandler() func(text string, setFocus func(p tview.Primitive)) {
	return w.grid.PasteHandler()
}

func NewWindow() *Window {
	w := &Window{
		app:    tview.NewApplication(),
		layout: NewModalLayout(),
		grid:   tview.NewGrid(),
	}

	w.app.SetRoot(w, true)
	w.app.SetInputCapture(w.inputCapture)

	return w
}

func (w *Window) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	switch key {
	default:
		return event
	}
}

func (w *Window) Run() error {
	return w.app.Run()
}
