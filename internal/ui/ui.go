package ui

import "github.com/gizak/termui"

// Init initializes the ui.
func Init() error {
	return termui.Init()
}

// Close closes the ui.
func Close() {
	termui.Close()
}
