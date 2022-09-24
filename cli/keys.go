package main

import "github.com/charmbracelet/bubbles/key"

var (
	Quit = key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	)
	TabAddNewLine = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "add new line"),
	)
	ShiftTabRemoveLine = key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "remove focused line"),
	)
)
