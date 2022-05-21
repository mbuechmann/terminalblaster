package ui2

import (
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var screen tcell.Screen
var style = tcell.StyleDefault.
	Background(tcell.NewRGBColor(0, 0, 0)).
	Foreground(tcell.NewRGBColor(255, 255, 255))

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
