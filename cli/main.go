package main

import "github.com/charmbracelet/bubbles/list"

var (
	action = NewSelect([]list.Item{
		SelectItem("Create"),
		SelectItem("Add"),
		SelectItem("Remove Module"),
		SelectItem("Remove Service"),
	}, WithTitle("What to do?"), WithShowHelp(), WithCallback(func(m list.Model) string {
		return "Select mode: " + string(m.SelectedItem().(SelectItem))
	}))
)

func main() {
	pg := NewProgram(action, NewSelect([]list.Item{
		SelectItem("Test 1"),
		SelectItem("Test 2"),
		SelectItem("Test 3"),
	}, WithTitle("What test for?"), WithShowHelp(), WithCallback(func(m list.Model) string {
		return "Test select: " + string(m.SelectedItem().(SelectItem))
	})))
	if pg != nil {
		pg.Start()
	}
}
