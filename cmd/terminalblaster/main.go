package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
	"github.com/mbuechmann/terminalblaster/internal/ui"
)

var cmd = &cobra.Command{
	Use:   "terminalblaster",
	Short: "Blast your music from the terminal",
	Long: `A lightweight music player for your terminal.
		Play all your music from the command line.`,
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func main() {
	cmd.Execute()
}
