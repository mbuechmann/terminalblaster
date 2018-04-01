package ui

import (
	"fmt"

	"github.com/gizak/termui"
	"github.com/mbuechmann/terminalblaster/internal/audio"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
	"github.com/mbuechmann/terminalblaster/internal/ui/widgets"
)

const (
	artistWidth = 40
)

var artistList *widgets.SelectList
var trackList *widgets.SelectList
var currentList *widgets.SelectList

// OpenLibraryScreen shows the library.
func OpenLibraryScreen() {
	termui.Clear()

	height := termui.TermHeight() - 2
	trackWidth := termui.TermWidth() - artistWidth

	buildArtistList(artistWidth, height)
	buildTrackList([]*lib.Track{}, trackWidth, height)

	setCurrentList(artistList)

	termui.Handle("/sys/kbd/Q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/<down>", func(e termui.Event) {
		currentList.Next()
	})
	termui.Handle("/sys/kbd/<up>", func(e termui.Event) {
		currentList.Prev()
	})
	termui.Handle("/sys/kbd/<next>", func(e termui.Event) {
		currentList.NextPage()
	})
	termui.Handle("/sys/kbd/<previous>", func(e termui.Event) {
		currentList.PrevPage()
	})
	termui.Handle("/sys/kbd/<right>", func(e termui.Event) {
		currentList.OpenItem()
		item := currentList.CurrentItem()
		switch v := item.Value.(type) {
		case *lib.Album:
			buildTrackList(v.Tracks, trackWidth, height)
			setCurrentList(trackList)
			trackList.Render()
		default:
		}
	})
	termui.Handle("/sys/kbd/<left>", func(e termui.Event) {
		if currentList == artistList {
			currentList.CloseItem()
		} else {
			setCurrentList(artistList)
		}
	})
	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		item := currentList.CurrentItem()
		switch v := item.Value.(type) {
		case *lib.Album:
			buildTrackList(v.Tracks, trackWidth, height)
			setCurrentList(trackList)
		case *lib.Track:
			go func() {
				err := audio.SetTracks(v.Album.Tracks, v.Album.TrackIndex(v))
				if err != nil {
					panic(err)
				}
				audio.Play()
			}()
		default:
		}
	})
	termui.Handle("/sys/kbd/<space>", func(e termui.Event) {
		audio.Toggle()
	})

	// termui.Handle("/sys/kbd", func(e termui.Event) {
	// 	fmt.Printf("%+v\n", e)
	// 	// <right>, <previous>, <next>, <escape>, <tab>
	// })

	termui.Loop()
}

func setCurrentList(list *widgets.SelectList) {
	artistList.SetFocussed(artistList == list)
	trackList.SetFocussed(trackList == list)
	artistList.Render()
	trackList.Render()

	currentList = list
}

func buildArtistList(width, height int) {
	items := make([]*widgets.SelectItem, len(lib.Artists))
	for i, a := range lib.Artists {
		item := widgets.NewSelectItem(a.Name, a)
		children := make([]*widgets.SelectItem, len(a.Albums))
		for j, alb := range a.Albums {
			children[j] = &widgets.SelectItem{Name: alb.Title, Value: alb}
		}
		item.SetChildren(children)

		items[i] = &item
	}

	artistList = widgets.NewSelectList(items, 0, 0, width, height)
}

func buildTrackList(tracks []*lib.Track, width, height int) {
	items := make([]*widgets.SelectItem, len(tracks))
	for i, t := range tracks {
		title := fmt.Sprintf("%2d. %s", t.TrackNumber, t.Title)
		items[i] = &widgets.SelectItem{Name: title, Value: t}
	}
	trackList = widgets.NewSelectList(items, artistWidth+1, 0, width, height)
}
