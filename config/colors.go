package config

import (
	tcell "github.com/gdamore/tcell/v2"

	"github.com/nimzo6689/boar-todo/ui/widget"
)

const (
	colorBackground      = tcell.Color235
	colorModalbackground = tcell.Color239
	colorBackgroundLight = tcell.Color239
	colorText            = tcell.Color23
)

type Colors struct {
	Background               tcell.Color
	TextPrimary              tcell.Color
	TextPrimaryLight         tcell.Color
	TextPrimaryDim           tcell.Color
	SelectionBackground      tcell.Color
	SelectionText            tcell.Color
	Border                   tcell.Color
	BorderFocus              tcell.Color
	ButtonBackground         tcell.Color
	ButtonBackgroundSelected tcell.Color
	ButtonLabel              tcell.Color
	ButtonLabelSelected      tcell.Color
	ModalBackground          tcell.Color
	NavBar                   ColorNavBar
	Tasks                    ColorTasks
	Help                     ColorHelp
}

func defaultColors() Colors {
	return Colors{
		Background:               colorBackground,
		TextPrimary:              tcell.Color252,
		TextPrimaryLight:         tcell.Color254,
		TextPrimaryDim:           tcell.Color249,
		SelectionBackground:      tcell.Color23,
		SelectionText:            tcell.Color253,
		Border:                   tcell.Color246,
		BorderFocus:              tcell.Color253,
		ButtonBackground:         tcell.Color241,
		ButtonBackgroundSelected: tcell.Color23,
		ButtonLabel:              tcell.Color254,
		ButtonLabelSelected:      tcell.Color253,
		ModalBackground:          colorModalbackground,
		NavBar:                   defaultColorNavBar(),
		Tasks:                    defaultColorTasks(),
		Help:                     defaultColorHelp(),
	}
}

type ColorNavBar struct {
	Background       tcell.Color
	BackgroundFocus  tcell.Color
	Text             tcell.Color
	TextFocus        tcell.Color
	ButtonBackground tcell.Color
	ButtonFocus      tcell.Color
	Shortcut         tcell.Color
	ShortcutFocus    tcell.Color
}

func defaultColorNavBar() ColorNavBar {
	return ColorNavBar{
		Background:       colorBackground,
		BackgroundFocus:  tcell.Color235,
		Text:             tcell.Color252,
		TextFocus:        tcell.Color253,
		ButtonBackground: colorBackground,
		ButtonFocus:      tcell.Color23,
		Shortcut:         tcell.Color214,
		ShortcutFocus:    tcell.Color214,
	}
}

func (c *ColorNavBar) ToNavBar() *widget.NavBarColors {
	return &widget.NavBarColors{
		Background:            c.Background,
		BackgroundFocus:       c.BackgroundFocus,
		ButtonBackground:      c.ButtonBackground,
		ButtonBackgroundFocus: c.ButtonFocus,
		Text:                  c.Text,
		TextFocus:             c.TextFocus,
		Shortcut:              c.Shortcut,
		ShortcutFocus:         c.ShortcutFocus,
	}
}

type ColorTasks struct {
	Background         tcell.Color
	Background2nd      tcell.Color
	BackgroundSelected tcell.Color
	Text               tcell.Color
	TextSelected       tcell.Color
	HeaderText         tcell.Color
	HeaderBackground   tcell.Color
}

func defaultColorTasks() ColorTasks {
	return ColorTasks{
		Background:         colorBackground,
		Background2nd:      tcell.Color236,
		BackgroundSelected: tcell.Color23,
		Text:               tcell.Color252,
		TextSelected:       tcell.Color253,
		HeaderText:         tcell.Color180,
		HeaderBackground:   tcell.Color235,
	}
}

type ColorHelp struct {
	Background tcell.Color
	Text       tcell.Color
	Headers    tcell.Color
}

func defaultColorHelp() ColorHelp {
	return ColorHelp{
		Background: colorModalbackground,
		Text:       tcell.Color252,
		Headers:    tcell.Color228,
	}
}
