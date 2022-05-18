package ui2

import (
	"os"

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

func OpenLoadScreen(chan *library.Track) error {
	return nil
}

func OpenLibraryScreen() error {
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
