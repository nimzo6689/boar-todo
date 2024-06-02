package modals

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
	"github.com/nimzo6689/boar-todo/config"
	"github.com/rivo/tview"
)

const (
	logo = "" +
		`
	____                      ______          __
	/ __ )____  ____ ______   /_  __/___  ____/ /___
   / __  / __ \/ __  / ___/    / / / __ \/ __  / __ \
  / /_/ / /_/ / /_/ / /       / / / /_/ / /_/ / /_/ /
 /_____/\____/\__,_/_/       /_/  \____/\__,_/\____/
 `
)

type Help struct {
	*tview.TextView
	doneFunc func()
	visible  bool

	page       int
	totalPages int
}

func (h *Help) SetDoneFunc(doneFunc func()) {
	h.doneFunc = doneFunc
}

func (h *Help) SetVisible(visible bool) {
	h.visible = visible
}

func (h *Help) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		key := event.Key()
		if key == tcell.KeyEscape {
			h.doneFunc()
		} else if key == tcell.KeyLeft {
			if h.page > 0 {
				h.page -= 1
				h.setContent()
			}
		} else if key == tcell.KeyRight {
			if h.page < h.totalPages-1 {
				h.page += 1
				h.setContent()
			}
		} else {
			h.TextView.InputHandler()(event, setFocus)
		}
	}
}

func (h *Help) setContent() {
	title := ""
	got := ""
	switch h.page {
	case 0:
		got = h.mainPage()
		title = "About"
	case 1:
		got = h.shortcutsPage()
		title = "Usage"
	default:
	}

	if title != "" {
		title = "[yellow::b]" + title + "[-::-]"
	}

	if got != "" {
		h.Clear()
		text := fmt.Sprintf("< %d / %d > %s \n\n", h.page+1, h.totalPages, title)
		text += got
		h.SetText(text)
		h.ScrollToBeginning()
	}
}

func NewHelp() *Help {
	h := &Help{
		TextView: tview.NewTextView(),
		doneFunc: nil,
		visible:  false,
	}

	colors := config.Configuration.Colors
	h.TextView.SetBorder(true)
	h.TextView.SetBorderColor(colors.BorderFocus)
	h.TextView.SetBackgroundColor(colors.Help.Background)
	h.TextView.SetTextColor(colors.Help.Text)
	h.TextView.SetBorderPadding(1, 1, 2, 2)
	h.TextView.SetDynamicColors(true)
	h.TextView.SetWordWrap(true)

	h.totalPages = 2
	h.setContent()
	return h
}

func (h *Help) mainPage() string {
	text := fmt.Sprintf("%s\n[yellow]v%s[-]\n\n", logo, config.Version)
	text += "License: MIT, https://opensource.org/license/mit"
	return text
}

func (h *Help) shortcutsPage() string {
	return `[yellow]Movement[-]:
 * Up/Down: Key up / down
 * VIM-like keys:
	 * Up / Down: J / K
	 * Top / Bottom: g / G
	 * Page Up / Down: Ctrl+F / Ctrl+B
 * Switch panels: Tab

 [yellow]Forms[-]:
 * Tab / Shift-Tab moves between form fields

 [yellow]Search[-]:
 * Open search panel: Ctrl-D
 * Search: Enter
 * Cancel: Escape

 [yellow]Task[-]:
 * Ctrl-space opens task viewer for selected task

 [yellow]Sorting[-]:
 * Navigate to any column header and press enter to sort either ascending or descending
 `
}
