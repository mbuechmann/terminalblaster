package ui2

import (
	"os"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell"

	"github.com/mbuechmann/terminalblaster/internal/library"
)

func OpenLibraryScreen(artists library.ArtistList) error {
	renderScreen(artists)
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			renderScreen(artists)
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

func renderScreen(artists library.ArtistList) {
	screen.Clear()

	w, h := screen.Size()
	// TODO: Make panel width dynamic
	panelWidth := 50

	// render top bar
	// TODO: check for negative repeat counts
	top := " Artist / Title"
	top += strings.Repeat(" ", 52-utf8.RuneCount([]byte(top)))
	top += "Track"
	top += strings.Repeat(" ", w-utf8.RuneCount([]byte(top)))
	renderString(top, styleHeadline, 0, 0)

	// render list of artists
	var i int
	for y := 1; y < h-2; y++ {
		// TODO: Limit length of string
		renderString(" "+artists[i].Name, styleRegular, 0, y)
		i++
	}

	// render current title and play data
	current := " Artist - Album - X. Title"
	play := "00:00 / 00:00 "
	// TODO: check for negative repeat count
	bottom := current + strings.Repeat(" ", w-utf8.RuneCount([]byte(current))-utf8.RuneCount([]byte(play))) + play
	renderString(bottom, styleHeadline, 0, h-2)

	// render divider
	for y := 1; y < h-2; y++ {
		renderString("│", styleRegular, panelWidth, y)
	}

	screen.Sync()
}
