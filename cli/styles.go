package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle        = lipgloss.NewStyle().Margin(1, 0, 1, 2).Bold(true).Underline(true).Foreground(lipgloss.Color("#FFFF00"))
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(1).Bold(true)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Bold(true).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(2).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 2).Bold(true).Foreground(lipgloss.Color("#006400"))
	ErrStyle          = lipgloss.NewStyle().Margin(1, 0, 1, 2).Bold(true).Underline(true).Foreground(lipgloss.Color("#B22222"))
	CursorStyle       = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#228B22"))
	ExitErrStyle      = ErrStyle.Copy().UnsetMargins().Margin(1, 0, 0, 2)

	NonStyle       = lipgloss.NewStyle()
	LmarginStyle   = lipgloss.NewStyle().MarginLeft(2)
	FocusedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurredStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	FocusedButton  = FocusedStyle.Copy().MarginLeft(2).Render("[ Submit ]")
	BlurredButton  = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
	THelpStyle     = HelpStyle.Copy().Foreground(lipgloss.Color("240"))
	ResultKeyStyle = lipgloss.NewStyle().MarginLeft(2).Bold(true).Underline(true).Foreground(lipgloss.Color("#228B22"))
	ResultStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{
		Light: "#000000",
		Dark:  "#FFFFFF",
	})
)
