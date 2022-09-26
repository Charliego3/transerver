package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"testing"
)

type model struct {
	spinner spinner.Model
	done    bool
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	if m.done {
		return m, tea.Quit
	}
	return m, nil
}

func (m *model) View() string {
	if m.done {
		return ""
	}
	return m.spinner.View() + " Loading..."
}

func TestPrintln(t *testing.T) {
	txt := "Starting with Go 1.16, you can embed files into your Go binaries. It means that you can build and ship to your users a binary that already has all the necessary files from your hard drive. So, there is no need to ship them separately and place them at a certain location on the computer. And the next time you move the binary to another directory, you do not need to update the paths to these files."
	style := lipgloss.NewStyle().MarginLeft(2).Bold(true).Underline(true).Foreground(lipgloss.Color("#B22222"))
	txt = style.Render(txt)
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = SpinnerStyle
	m := &model{s, false}
	pg := tea.NewProgram(m)
	go func() {
		pg.Println(txt)
		m.done = true
	}()
	_ = pg.Start()

	println()
	fmt.Println(txt)
}
