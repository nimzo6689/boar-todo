package ui

import (
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/nimzo6689/boar-todo/config"
	"github.com/nimzo6689/boar-todo/ui/widget"
)

var navBarLabels = make([]string, 0)
var navBarShortucts = make([]tcell.Key, 0)

type Window struct {
	app *tview.Application

	layout *widget.ModalLayout
	grid   *tview.Grid

	navBar *widget.NavBar
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

func NewWindow(colors config.Colors, shortcuts *config.Shortcuts) *Window {
	w := &Window{
		app:    tview.NewApplication(),
		layout: widget.NewModalLayout(),
		grid:   tview.NewGrid(),
	}

	w.app.SetRoot(w, true)
	w.app.SetInputCapture(w.inputCapture)

	w.layout.SetGridYSize([]int{3, -1, -1, -1, -1, -1, -1, -1, -1, 3})

	col := colors.NavBar.ToNavBar()

	w.navBar = widget.NewNavBar(col, w.navBarClicked)
	navBarLabels = []string{"Help", "New Task", "Current Task", "Next Task", "Stats", "Quit"}

	sc := shortcuts.NavBar
	navBarShortucts = []tcell.Key{sc.Help, sc.NewTask, sc.CurrentTask, sc.NextTask, sc.Stats, sc.Quit}

	for i, v := range navBarLabels {
		btn := tview.NewButton(v)
		w.navBar.AddButton(btn, navBarShortucts[i])
	}

	w.grid.AddItem(w.navBar, 0, 0, 1, 1, 1, 10, false)
	w.grid.AddItem(w.layout, 1, 0, 1, 1, 4, 4, true)

	return w
}

func (w *Window) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	switch key {
	default:
		return event
	}
}

func (w *Window) navBarClicked(label string) {
	// Noop
}

func (w *Window) Run() error {
	return w.app.Run()
}
