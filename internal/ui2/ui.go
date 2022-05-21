package ui2

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"

	"github.com/mbuechmann/terminalblaster/internal/library"
)

var screen tcell.Screen

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

func OpenLoadScreen(ch chan *library.Track) error {
	var counter int

	w, h := screen.Size()

	st := tcell.StyleDefault.
		Background(tcell.NewRGBColor(0, 0, 0)).
		Foreground(tcell.NewRGBColor(255, 255, 255))

	format := "│ %4d titles loaded │"
	y := h / 2
	for range ch {
		counter++
		str := fmt.Sprintf(format, counter)
		x := (w - utf8.RuneCount([]byte(str))) / 2

		border := strings.Repeat("─", utf8.RuneCount([]byte(str))-2)
		renderString("┌"+border+"┐", st, x, y-1)
		renderString(str, st, x, y)
		renderString("└"+border+"┘", st, x, y+1)
		screen.Show()
	}

	return nil
}

func renderString(str string, st tcell.Style, x, y int) {
	var i int
	for _, r := range str {
		screen.SetCell(x+i, y, st, r)
		i++
	}
}

func OpenLibraryScreen() error {
	screen.Clear()
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyRune && ev.Rune() == 'Q' {
				screen.Fini()
				os.Exit(0)
			}
		}
	}
	return nil
}

func Close() {
	screen.Fini()
}
