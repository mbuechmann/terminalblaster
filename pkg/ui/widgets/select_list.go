package widgets

import (
	"fmt"

	"github.com/gizak/termui"
)

// SelectList represents a list of items that can be selected.
type SelectList struct {
	Source    []string
	index     int
	offset    int
	list      *termui.List
	itemCount int
}

// NewSelectList returns a new SelectList for the given source of items.
func NewSelectList(source []string, x, y, w, h int) SelectList {
	sl := SelectList{
		Source: source,
	}

	list := termui.NewList()
	list.ItemFgColor = termui.ColorBlack
	list.Bg = termui.ColorWhite
	list.ItemBgColor = termui.ColorWhite
	list.Border = false
	list.Height = h
	list.Width = w
	list.X = 0
	list.Y = 0

	sl.list = list
	sl.itemCount = len(source)
	if sl.itemCount > h {
		sl.itemCount = h
	}

	sl.fillList()

	return sl
}

func (sl *SelectList) fillList() {
	strs := make([]string, sl.itemCount)
	for i := 0; i < sl.itemCount; i++ {
		if i == sl.index {
			strs[i] = fmt.Sprintf("[%-39s](fg-white,bg-black)", sl.Source[i+sl.offset])
		} else {
			strs[i] = sl.Source[i+sl.offset]
		}
	}
	sl.list.Items = strs

}

// Next increments the index of the active item.
func (sl *SelectList) Next() {
	if sl.offset+sl.index < len(sl.Source)-1 {
		if sl.index < sl.itemCount-1 {
			sl.index++
		} else {
			sl.offset++
		}
		sl.fillList()
		sl.Render()
	}
}

// Prev decrements the index of the active item.
func (sl *SelectList) Prev() {
	if sl.index+sl.offset > 0 {
		if sl.index > 0 {
			sl.index--
		} else {
			sl.offset--
		}
		sl.fillList()
		sl.Render()
	}
}

// NextPage decrements the index of the active item by one page.
func (sl *SelectList) NextPage() {
	sl.offset += sl.itemCount
	if sl.offset > len(sl.Source)-sl.itemCount {
		sl.offset = len(sl.Source) - sl.itemCount
		sl.index = sl.itemCount - 1
	}
	sl.fillList()
	sl.Render()
}

// PrevPage decrements the index of the active item by one page.
func (sl *SelectList) PrevPage() {
	sl.offset -= sl.itemCount
	if sl.offset < 0 {
		sl.offset = 0
		sl.index = 0
	}
	sl.fillList()
	sl.Render()
}

// Render renders the SelectList on the screen.
func (sl SelectList) Render() {
	termui.Render(sl.list)
}

// CurrentItem returns the item for the current index.
func (sl SelectList) CurrentItem() string {
	return sl.Source[sl.index+sl.offset]
}
