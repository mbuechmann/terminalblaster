package ui2

import (
	"os"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell"

	"github.com/mbuechmann/terminalblaster/internal/library"
)

func OpenLibraryScreen(artists library.ArtistList) error {
	renderArtistList(artists)
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

func renderArtistList(artists library.ArtistList) {
	screen.Clear()

	w, h := screen.Size()

	// render top bar
	top := " Artist / Title"
	top += strings.Repeat(" ", w-utf8.RuneCount([]byte(top)))
	renderString(top, styleHeadline, 0, 0)

	// render list of artists
	var i int
	for y := 1; y < h; y++ {
		renderString(artists[i].Name, styleRegular, 0, y)
		i++
	}
}
