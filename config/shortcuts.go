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
	Help        tcell.Key
	NewTask     tcell.Key
	CurrentTask tcell.Key
	NextTask    tcell.Key
	Stats       tcell.Key
	Quit        tcell.Key
}

func defaultNavBar() NavBar {
	return NavBar{
		Help:        tcell.KeyF1,
		NewTask:     tcell.KeyF2,
		CurrentTask: tcell.KeyF3,
		NextTask:    tcell.KeyF4,
		Stats:       tcell.KeyF5,
		Quit:        tcell.KeyF6,
	}
}
