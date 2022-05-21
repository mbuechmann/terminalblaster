package ui2

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var screen tcell.Screen

var (
	colorDarkBlue  = tcell.NewRGBColor(20, 33, 61)
	colorLightGrey = tcell.NewRGBColor(229, 229, 229)
	colorOrange    = tcell.NewRGBColor(252, 163, 17)

	styleRegular = tcell.StyleDefault.
		Background(colorDarkBlue).
		Foreground(colorLightGrey)
	styleHeadline = tcell.StyleDefault.
		Background(colorOrange).
		Foreground(colorLightGrey)
)

func Init() error {
	encoding.Register()

	var err error
	screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := screen.Init(); err != nil {
		return err
	}

	screen.SetStyle(styleRegular)
	screen.Clear()

	return nil
}

func Close() {
	screen.Fini()
}

func renderString(str string, st tcell.Style, x, y int) {
	var i int
	for _, r := range str {
		screen.SetCell(x+i, y, st, r)
		i++
	}
}
