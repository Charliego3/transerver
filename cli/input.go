package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/goutil/strutil"
)

var keys = []key.Binding{
	TabAddNewLine,
	ShiftTabRemoveLine,
	Quit,
}

type Input struct {
	inputs      []textinput.Model
	help        help.Model
	title       string
	placeholder string
	focusIndex  int
	callback    func([]string) string
}

func NewInput(title, placeholder string, callback func([]string) string) *Input {
	t := newInput(placeholder)
	t.Focus()
	t.SetCursorMode(textinput.CursorBlink)
	return &Input{
		inputs:      []textinput.Model{t},
		help:        help.New(),
		title:       title,
		focusIndex:  0,
		placeholder: placeholder,
		callback:    callback,
	}
}

func newInput(placeholder string) textinput.Model {
	t := textinput.New()
	t.CursorStyle = FocusedStyle.Copy()
	t.CharLimit = 32
	t.Placeholder = placeholder
	t.PromptStyle = FocusedStyle
	t.TextStyle = FocusedStyle
	return t
}

func (m *Input) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "shift+tab":
			if len(m.inputs) == 1 || m.focusIndex == len(m.inputs) {
				return nil
			}
			m.inputs = append(m.inputs[:m.focusIndex], m.inputs[m.focusIndex+1:]...)
			return m.update(msg.String())
		case "tab":
			m.inputs = append(m.inputs, newInput(m.placeholder))
			fallthrough
		case "enter", "up", "down":
			return m.update(msg.String())
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)
	return cmd
}

func (m *Input) update(key string) tea.Cmd {
	if key == "enter" && m.focusIndex == len(m.inputs) {
		return func() tea.Msg { return nil }
	}

	if key == "up" || key == "shift+tab" {
		m.focusIndex--
	} else {
		m.focusIndex++
	}

	if m.focusIndex > len(m.inputs) {
		m.focusIndex = 0
	} else if m.focusIndex < 0 {
		m.focusIndex = len(m.inputs)
	}

	cmds := make([]tea.Cmd, len(m.inputs))
	for i := 0; i <= len(m.inputs)-1; i++ {
		if i == m.focusIndex {
			cmds[i] = m.inputs[i].Focus()
			m.inputs[i].PromptStyle = FocusedStyle
			m.inputs[i].TextStyle = FocusedStyle
			continue
		}
		m.inputs[i].Blur()
		m.inputs[i].PromptStyle = NonStyle
		m.inputs[i].TextStyle = NonStyle
	}

	return tea.Batch(cmds...)
}

func (m *Input) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *Input) View() string {
	var b strings.Builder
	if strutil.IsNotBlank(m.title) {
		b.WriteString(TitleStyle.Render(m.title))
		b.WriteRune('\n')
	}

	for i := range m.inputs {
		b.WriteString(LmarginStyle.Render(m.inputs[i].View()))
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	btn := LmarginStyle.Render(BlurredButton)
	if m.focusIndex == len(m.inputs) {
		btn = FocusedButton
	}
	b.WriteString("\n\n")
	b.WriteString(btn)
	b.WriteByte('\n')
	b.WriteString(THelpStyle.Render(m.help.ShortHelpView(keys)))
	return b.String()
}

func (m *Input) Callback(pg Program) {
	if m.callback == nil {
		return
	}

	var values []string
	for _, i := range m.inputs {
		if strutil.IsNotBlank(i.Value()) {
			values = append(values, i.Value())
		}
	}
	pg.Println(m.callback(values))
}
