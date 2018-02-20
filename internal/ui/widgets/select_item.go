package widgets

import (
	"fmt"
	"strings"
)

// NewSelectItem returns a new SelectItem for the given name and value.
func NewSelectItem(name string, value interface{}) SelectItem {
	return SelectItem{Name: name, Value: value}
}

// SelectItem is a item of a SelectList.
type SelectItem struct {
	Name     string
	Value    interface{}
	Open     bool
	Parent   *SelectItem
	Children []*SelectItem
	focussed bool
}

// SetChildren sets the children of the selectitems and the parent of all
// children.
func (si *SelectItem) SetChildren(children []*SelectItem) {
	si.Children = children
	for _, c := range children {
		c.Parent = si
	}
}

func (si *SelectItem) String() string {
	spaces := strings.Repeat("  ", si.depth())
	format := "%s%s"
	if si.focussed {
		format = fmt.Sprintf("[%%s%%-%ds](fg-white,bg-black)", 39-len(spaces))
	}
	return fmt.Sprintf(format, spaces, si.Name)
}

// Openable returns if the item can be opened.
func (si *SelectItem) Openable() bool {
	return !si.Open && len(si.Children) > 0
}

// Closable returns if the item can be closed.
func (si *SelectItem) Closable() bool {
	return si.Open && len(si.Children) > 0
}

func (si *SelectItem) depth() int {
	res := 0
	current := si
	for current.Parent != nil {
		res++
		current = current.Parent
	}
	return res
}
