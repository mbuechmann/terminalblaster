package widgets

import "github.com/gizak/termui"

// ProgressBar displays a progress with a label.
type ProgressBar struct {
	gauge *termui.Gauge
}

// NewProgressBar returns a new ProgressBar.
func NewProgressBar(label string, x, y, w, h int) *ProgressBar {
	gauge := termui.NewGauge()

	gauge.Percent = 0
	gauge.Border = false
	gauge.X = x
	gauge.Y = y
	gauge.Width = w
	gauge.Height = h
	gauge.Label = label
	gauge.PercentColor = termui.ColorYellow
	gauge.BarColor = termui.ColorGreen
	gauge.PercentColorHighlighted = termui.ColorBlack
	gauge.LabelAlign = termui.AlignLeft

	bar := ProgressBar{gauge: gauge}

	return &bar
}

// Render renders the ProgressBar.
func (pb *ProgressBar) Render() {
	termui.Render(pb.gauge)
}

// SetLabel sets the label of the ProgressBar.
func (pb *ProgressBar) SetLabel(label string) {
	pb.gauge.Label = label
	termui.Render(pb.gauge)
}
