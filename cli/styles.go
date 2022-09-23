package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle        = lipgloss.NewStyle().Margin(1, 0, 1, 2).Bold(true).Underline(true).Foreground(lipgloss.Color("#FFFF00"))
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(1)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 2).Bold(true).Foreground(lipgloss.Color("#006400"))
	ErrStylr          = lipgloss.NewStyle().Margin(1, 0, 1, 2).Bold(true).Underline(true).Foreground(lipgloss.Color("#B22222"))
	CursorStyle       = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#228B22"))
)
