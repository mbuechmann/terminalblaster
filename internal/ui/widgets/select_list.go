package widgets

import (
	"github.com/gizak/termui"
)

// SelectList represents a list of items that can be selected.
type SelectList struct {
	items          []*SelectItem
	flattenedItems []*SelectItem
	index          int
	offset         int
	list           *termui.List
	minItemCount   int
	maxItemCount   int
	itemCount      int
}

// NewSelectList returns a new SelectList for the given source of items.
func NewSelectList(source []*SelectItem, x, y, w, h int) SelectList {
	sl := SelectList{
		items: source,
	}
	sl.flattenedItems = sl.items[:]

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
	sl.minItemCount = sl.itemCount
	sl.maxItemCount = h

	if len(sl.flattenedItems) > 0 {
		sl.flattenedItems[0].focussed = true
	}

	sl.fillList()

	return sl
}

func (sl *SelectList) fillList() {
	strs := make([]string, sl.itemCount)
	for i := 0; i < sl.itemCount; i++ {
		strs[i] = sl.flattenedItems[i+sl.offset].String()
	}
	sl.list.Items = strs
	sl.Render()
}

// Next increments the index of the active item.
func (sl *SelectList) Next() {
	if sl.offset+sl.index < len(sl.flattenedItems)-1 {
		i := sl.CurrentItem()
		i.focussed = false

		if sl.index < sl.itemCount-1 {
			sl.index++
		} else {
			sl.offset++
		}

		i = sl.CurrentItem()
		i.focussed = true

		sl.fillList()
	}
}

// Prev decrements the index of the active item.
func (sl *SelectList) Prev() {
	if sl.index+sl.offset > 0 {
		i := sl.CurrentItem()
		i.focussed = false

		if sl.index > 0 {
			sl.index--
		} else {
			sl.offset--
		}

		i = sl.CurrentItem()
		i.focussed = true

		sl.fillList()
	}
}

// NextPage decrements the index of the active item by one page.
func (sl *SelectList) NextPage() {
	i := sl.CurrentItem()
	i.focussed = false

	sl.offset += sl.itemCount
	if sl.offset > len(sl.flattenedItems)-sl.itemCount {
		sl.offset = len(sl.flattenedItems) - sl.itemCount
		sl.index = sl.itemCount - 1
	}

	i = sl.CurrentItem()
	i.focussed = true

	sl.fillList()
}

// PrevPage decrements the index of the active item by one page.
func (sl *SelectList) PrevPage() {
	i := sl.CurrentItem()
	i.focussed = false

	sl.offset -= sl.itemCount
	if sl.offset < 0 {
		sl.offset = 0
		sl.index = 0
	}

	i = sl.CurrentItem()
	i.focussed = true

	sl.fillList()
}

// OpenItem opens the children of the current item.
func (sl *SelectList) OpenItem() {
	i := sl.CurrentItem()
	if i.Openable() {
		i.Open = true
		sl.flattenedItems = append(
			sl.flattenedItems[:sl.index+sl.offset+1],
			append(
				sl.CurrentItem().Children,
				sl.flattenedItems[sl.index+sl.offset+1:]...,
			)...,
		)
		sl.itemCount += len(sl.CurrentItem().Children)
		if sl.itemCount > sl.maxItemCount {
			sl.itemCount = sl.maxItemCount
		}
	}
	sl.fillList()
}

// CloseItem closes the current item if it is open. If closed and if the current
// item has a parent it closes the parent.
func (sl *SelectList) CloseItem() {
	i := sl.CurrentItem()
	if i.Closable() {
		i.Open = false
		sl.flattenedItems = append(
			sl.flattenedItems[:sl.index+sl.offset+1],
			sl.flattenedItems[sl.index+sl.offset+1+len(sl.CurrentItem().Children):]...,
		)
		sl.itemCount -= len(sl.CurrentItem().Children)
		if sl.itemCount < sl.minItemCount {
			sl.itemCount = sl.minItemCount
		}
	}
	sl.fillList()
}

// Render renders the SelectList on the screen.
func (sl SelectList) Render() {
	termui.Render(sl.list)
}

// CurrentItem returns the item for the current index.
func (sl SelectList) CurrentItem() *SelectItem {
	return sl.flattenedItems[sl.index+sl.offset]
}
