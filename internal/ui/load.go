package ui

import (
	"fmt"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
	"github.com/mbuechmann/terminalblaster/internal/ui/widgets"
)

var tracksLoaded = 0
var m widgets.Modal

// OpenLoadScreen sets up the ui for a loading view.
func OpenLoadScreen(c chan *lib.Track) (err error) {
	m = widgets.NewModal("")

	for range c {
		tracksLoaded++
		renderMsg()
	}

	return
}

func renderMsg() {
	msg := fmt.Sprintf(" Loading tracks: %d ", tracksLoaded)
	m.SetText(msg)
	m.Render()
}
