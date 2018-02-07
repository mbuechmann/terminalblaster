package widgets

import (
	"github.com/gizak/termui"
)

// Modal displays text on top of the ui.
type Modal struct {
	text string
	par  *termui.Par
}

// NewModal returns a Modal with the given text.
func NewModal(text string) Modal {
	m := Modal{
		text: text,
		par:  termui.NewPar(text),
	}

	m.par.Height = 3
	x := (termui.TermWidth() - m.par.Width) / 2
	m.par.X = x
	y := (termui.TermHeight() - m.par.Height) / 2
	m.par.Y = y

	m.par.TextFgColor = termui.ColorBlack

	return m
}

// SetText sets the text of the Modal. It adapts its width.
func (m Modal) SetText(t string) {
	oldWidth := len(m.par.Text) + 2
	width := len(t) + 2

	if oldWidth != width {
		m.par.Width = width
		x := (termui.TermWidth() - m.par.Width) / 2
		m.par.X = x

	}
	m.par.Text = t
}

// Render renders the Modal.
func (m Modal) Render() {
	termui.Render(m.par)
}
