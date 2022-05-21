package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
	ui "github.com/mbuechmann/terminalblaster/internal/ui2"
)

var cmd = &cobra.Command{
	Use:   "terminalblaster",
	Short: "Blast your music from the terminal",
	Long: `A lightweight music player for your terminal.
		Play all your music from the command line.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("accepts 1 arg, received %d", len(args))
		}

		path := args[0]
		fileInfo, err := os.Stat(path)
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("directory %s does not exist", path)
		}
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			return fmt.Errorf("%s is not a directory", path)
		}

		return nil
	},
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

		ui.OpenLibraryScreen(lib.Artists)

		ui.Close()
	},
}

func main() {
	cmd.Execute()
}
