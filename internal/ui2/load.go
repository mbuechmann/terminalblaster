package ui2

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mbuechmann/terminalblaster/internal/library"
)

func OpenLoadScreen(ch chan *library.Track) error {
	var counter int

	w, h := screen.Size()

	format := "│ %4d titles loaded │"
	y := h / 2
	for range ch {
		counter++
		str := fmt.Sprintf(format, counter)
		x := (w - utf8.RuneCount([]byte(str))) / 2

		border := strings.Repeat("─", utf8.RuneCount([]byte(str))-2)
		renderString("┌"+border+"┐", styleRegular, x, y-1)
		renderString(str, styleRegular, x, y)
		renderString("└"+border+"┘", styleRegular, x, y+1)
		screen.Show()
	}

	return nil
}
