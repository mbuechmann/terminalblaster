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
	termui.Handle("/sys/kbd/<next>", func(e termui.Event) {
		artistList.NextPage()
	})
	termui.Handle("/sys/kbd/<previous>", func(e termui.Event) {
		artistList.PrevPage()
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

	items := make([]widgets.SelectItem, len(lib.Artists))
	for i, a := range lib.Artists {
		items[i] = widgets.SelectItem{Name: a.Name, Value: a}
	}

	artistList = widgets.NewSelectList(items, 0, 0, width, height)
	artistList.Render()

	termui.Loop()
}
