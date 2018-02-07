package main

import (
	"fmt"
	"os"

	lib "github.com/mbuechmann/terminalblaster/pkg/library"
	"github.com/mbuechmann/terminalblaster/pkg/ui"
)

// var file = "/Users/maltebuchmann/Music/Queen/Greatest Hits/01 Bohemian Rhapsody (1993 Digital Remaster).mp3"
// var file = "/Users/maltebuchmann/Music/Led Zeppelin/Early days - The Best Of Led Zeppelin, Volume 1/13 Led Zeppelin - Stairway To Heaven.flac"
// var file = "/Users/maltebuchmann/Music/A Fine Frenzy/Bomb in a Birdcage/04 A Fine Frenzy - Blow Away.flac"

// var file = "/Users/maltebuchmann/Music/A Fine Frenzy/Bomb in a Birdcage/10 A Fine Frenzy - Stood Up.flac"
// var file = "/Users/maltebuchmann/Music/A Fine Frenzy/One Cell in the Sea/01 A Fine Frenzy - Come On, Come Out.flac"

var trackCount = 0

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
