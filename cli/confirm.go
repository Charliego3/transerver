package main

import (
	"github.com/charmbracelet/bubbles/list"
)

type Confirm struct {
	*Select
}

func NewConfirm(title string, callback func(bool) string) *Confirm {
	items := []list.Item{
		SelectItem("Yes"),
		SelectItem("No"),
	}
	s := NewSelect(items, WithTitle(title), WithShowHelp(), WithCallback(func(m list.Model) string {
		return callback(m.Index() == 0)
	}))
	return &Confirm{Select: s.(*Select)}
}
