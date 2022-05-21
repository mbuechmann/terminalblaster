package ui2

import (
	"os"

	"github.com/gdamore/tcell"

	"github.com/mbuechmann/terminalblaster/internal/library"
)

func OpenLibraryScreen(artists library.ArtistList) error {
	renderLibrary(artists)
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

func renderLibrary(artists library.ArtistList) {
	screen.Clear()

	_, h := screen.Size()

	var i int
	for y := 0; y < h; y++ {
		renderString(artists[i].Name, style, 0, y)
		i++
	}
}
