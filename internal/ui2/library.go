package ui2

import (
	"os"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell"

	"github.com/mbuechmann/terminalblaster/internal/library"
)

type menuItem struct {
	artist *library.Artist
	album  *library.Album
	open   bool
}

var menuPosition int
var menuItems []*menuItem

func OpenLibraryScreen(artists library.ArtistList) error {
	menuItems = make([]*menuItem, len(artists))
	for i, artist := range artists {
		menuItems[i] = &menuItem{artist: artist}
	}

	renderScreen()
	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			renderScreen()
			screen.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyRune:
				if ev.Rune() == 'Q' {
					screen.Fini()
					os.Exit(0)
				}
			case tcell.KeyDown:
				menuPosition++
				// TODO: minimize rendering
				renderScreen()
			case tcell.KeyUp:
				menuPosition--
				// TODO: minimize rendering
				renderScreen()
			case tcell.KeyEnter:
				toggleItem()
				renderScreen()
			}
		}
	}
	return nil
}

func toggleItem() {
	item := menuItems[menuPosition]

	if item.artist != nil {
		if item.open {
			cut := len(item.artist.Albums)
			menuItems = append(menuItems[:menuPosition+1], menuItems[menuPosition+cut+1:]...)
		} else {
			insert := make([]*menuItem, len(item.artist.Albums))
			for i, album := range item.artist.Albums {
				insert[i] = &menuItem{
					album: album,
				}
			}
			menuItems = append(menuItems[:menuPosition+1], append(insert, menuItems[menuPosition+1:]...)...)
		}

		item.open = !item.open
	}
}

func renderScreen() {
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
		style := styleRegular
		if i == menuPosition {
			style = styleCursor
		}

		item := menuItems[i]

		var line string

		switch {
		case item.artist != nil:
			line = " " + item.artist.Name
		case item.album != nil:
			line = "   " + item.album.Title
		default:
			line = ""
		}

		// if too long cap after panelWidth
		if utf8.RuneCount([]byte(line)) > panelWidth {
			line = string([]rune(line)[0:panelWidth])
		}

		renderString(line, style, 0, y)
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
