package main

import (
	"fmt"
	"os"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
	"github.com/mbuechmann/terminalblaster/internal/ui"
)

func main() {
	checkArgs()
	loadLibrary()
}

func checkArgs() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Provide directory to be loaded")
		os.Exit(1)
	}
}

func loadLibrary() {
	if err := ui.Init(); err != nil {
		fmt.Println("Could not open ui")
		os.Exit(1)
	}

	path := os.Args[1]
	c, err := lib.Load(path)
	if err != nil {
		panic(err)
	}

	if err := ui.OpenLoadScreen(c); err != nil {
		panic(err)
	}

	ui.OpenLibraryScreen()

	ui.Close()
}
