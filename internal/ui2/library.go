package ui2

import (
	"os"
	"strings"
	"unicode/utf8"

	"github.com/gdamore/tcell"

	"github.com/mbuechmann/terminalblaster/internal/audio"
	"github.com/mbuechmann/terminalblaster/internal/library"
)

const (
	focusMenu = iota
	focusAlbum
)

type menuItem struct {
	artist *library.Artist
	album  *library.Album
	open   bool
}

var (
	focus int = focusMenu

	menuPosition int
	menuItems    []*menuItem

	selectedAlbum *library.Album
	albumPosition int

	selectedTrack *library.Track
)

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
				keyDown()
				// TODO: minimize rendering
				renderScreen()
			case tcell.KeyUp:
				keyUp()
				renderScreen()
			case tcell.KeyEnter:
				keyEnter()
				renderScreen()
			}
		}
	}
	return nil
}

func keyEnter() {
	if focus == focusMenu {
		toggleItem()
	} else {
		go func() {
			_ = audio.Play(selectedTrack)
		}()
	}
}

func keyUp() {
	if focus == focusMenu {
		if menuPosition > 0 {
			menuPosition--
		}
	} else {
		if albumPosition > 0 {
			albumPosition--
		}
		selectedTrack = selectedAlbum.Tracks[albumPosition]
	}
}

func keyDown() {
	if focus == focusMenu {
		_, h := screen.Size()
		if menuPosition < h-4 {
			menuPosition++
		}
	} else {
		if albumPosition < len(selectedAlbum.Tracks)-1 {
			albumPosition++
		}
		selectedTrack = selectedAlbum.Tracks[albumPosition]
	}
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

	if item.album != nil {
		selectedAlbum = item.album
		albumPosition = 0
		focus = focusAlbum
		selectedTrack = selectedAlbum.Tracks[0]
	}
}

func renderScreen() {
	screen.Clear()

	// render top bar
	// TODO: check for negative repeat counts
	top := " Artist / Title"
	top += strings.Repeat(" ", 52-utf8.RuneCount([]byte(top)))
	top += "Track"
	top += strings.Repeat(" ", screenWidth()-utf8.RuneCount([]byte(top)))
	renderString(top, styleHeadline, 0, 0)

	// render list of artists
	var i int
	for y := 1; y < menuHeight(); y++ {
		style := styleRegular
		if i == menuPosition {
			if focus == focusMenu {
				style = styleActiveCursor
			} else {
				style = styleInActiveCursor
			}
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

		// if too long cap after menuWidth
		if utf8.RuneCount([]byte(line)) > menuWidth() {
			line = string([]rune(line)[0:menuWidth()])
		}

		// fill with spaces for better highlighting
		line += strings.Repeat(" ", menuWidth()-utf8.RuneCount([]byte(line)))

		renderString(line, style, 0, y)

		i++
	}

	// render album's songs
	if selectedAlbum != nil {
		for i, track := range selectedAlbum.Tracks {
			style := styleRegular
			if i == albumPosition {
				if focus == focusAlbum {
					style = styleActiveCursor
				} else {
					style = styleInActiveCursor
				}
			}
			renderString(track.Title, style, menuWidth()+1, i+1)
		}
	}

	// render current title and play data
	current := " Artist - Album - X. Title"
	play := "00:00 / 00:00 "
	// TODO: check for negative repeat count
	bottom := current + strings.Repeat(" ", screenWidth()-utf8.RuneCount([]byte(current))-utf8.RuneCount([]byte(play))) + play
	renderString(bottom, styleHeadline, 0, menuHeight())

	// render divider
	for y := 1; y < menuHeight(); y++ {
		renderString("│", styleRegular, menuWidth(), y)
	}

	screen.Sync()
}

func menuHeight() int {
	_, h := screen.Size()
	return h - 2
}

func menuWidth() int {
	// TODO: Make dynamic
	return 50
}

func screenWidth() int {
	w, _ := screen.Size()
	return w
}
