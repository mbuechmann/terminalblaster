package ui

import (
	"fmt"

	"github.com/gizak/termui"
	lib "github.com/mbuechmann/terminalblaster/pkg/library"
	"github.com/mbuechmann/terminalblaster/pkg/ui/widgets"
)

var artistList widgets.SelectList

// OpenLibraryScreen shows the library.
func OpenLibraryScreen() {
	termui.Clear()

	termui.Handle("/sys/kbd/Q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/<down>", func(e termui.Event) {
		artistList.Next()
	})
	termui.Handle("/sys/kbd/<up>", func(e termui.Event) {
		artistList.Prev()
	})
	termui.Handle("/sys/kbd/<right>", func(e termui.Event) {
		fmt.Println(artistList.CurrentItem())
	})
	// termui.Handle("/sys/kbd", func(e termui.Event) {
	// 	fmt.Printf("%+v\n", e)
	// 	// <right>, <previous>, <next>, <escape>, <tab>
	// })

	width := 40
	height := termui.TermHeight() - 2

	artistList = widgets.NewSelectList(lib.ArtistNames, 0, 0, width, height)
	artistList.Render()

	termui.Loop()
}
