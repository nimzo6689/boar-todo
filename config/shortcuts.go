package config

import (
	tcell "github.com/gdamore/tcell/v2"
)

type Shortcuts struct {
	NavBar NavBar
}

func defaultShortcuts() Shortcuts {
	return Shortcuts{
		NavBar: defaultNavBar(),
	}
}

type NavBar struct {
	Help    tcell.Key
	New     tcell.Key
	Current tcell.Key
	Next    tcell.Key
	Stats   tcell.Key
	Quit    tcell.Key
}

func defaultNavBar() NavBar {
	return NavBar{
		Help:    tcell.KeyF1,
		New:     tcell.KeyF2,
		Current: tcell.KeyF3,
		Next:    tcell.KeyF4,
		Stats:   tcell.KeyF5,
		Quit:    tcell.KeyF6,
	}
}
