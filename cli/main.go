package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
)

var (
	action = NewSelect([]list.Item{
		SelectItem("Create Module"),
		SelectItem("Add Service"),
		SelectItem("Remove Module"),
		SelectItem("Remove Service"),
	}, WithTitle("What to do?"), WithShowHelp(), WithCallback(func(m list.Model) string {
		return "Select mode: " + string(m.SelectedItem().(SelectItem))
	}))
)

func main() {
	pg := NewProgram(action, NewConfirm("Want using GRPC optoins?", func(b bool) string {
		return "Using GRPC options: " + strconv.FormatBool(b)
	}), NewInput("What servie want to add?", "please enter a service name", func(s []string) string {
		return "Service names: [" + strings.Join(s, " ") + "]"
	}))
	if pg != nil {
		pg.Start()
	}
}
