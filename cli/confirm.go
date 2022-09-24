package main

import (
	"github.com/charmbracelet/bubbles/list"
)

type Confirm struct {
	*Select
}

func NewConfirm(title string, callback func(bool) (string, bool)) *Confirm {
	items := []list.Item{
		SelectItem("Yes"),
		SelectItem("No"),
	}
	s := NewSelect(items, WithTitle(title), WithShowHelp(), WithCallback(func(m list.Model) (string, bool) {
		return callback(m.Index() == 0)
	}))
	return &Confirm{Select: s.(*Select)}
}
