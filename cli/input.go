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
	callbacks   func([]string) (string, bool)
	callback    func(string) (string, bool)
	multi       bool
}

func NewInput(title string, opts ...InputOpt) *Input {
	i := &Input{
		title: title,
		help:  help.New(),
	}
	for _, opt := range opts {
		opt(i)
	}

	t := i.newInput()
	t.Focus()
	t.SetCursorMode(textinput.CursorBlink)

	i.inputs = []textinput.Model{t}
	return i
}

func (m *Input) newInput() textinput.Model {
	t := textinput.New()
	t.CursorStyle = FocusedStyle.Copy()
	t.CharLimit = 32
	t.Placeholder = m.placeholder
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
			if !m.multi {
				return nil
			}
			m.inputs = append(m.inputs, m.newInput())
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
	if key == "enter" && (!m.multi || m.focusIndex == len(m.inputs)) {
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

	if m.multi {
		btn := LmarginStyle.Render(BlurredButton)
		if m.focusIndex == len(m.inputs) {
			btn = FocusedButton
		}
		b.WriteString("\n\n")
		b.WriteString(btn)
	}
	b.WriteByte('\n')
	ks := keys
	if !m.multi {
		ks = ks[2:]
	}
	b.WriteString(THelpStyle.Render(m.help.ShortHelpView(ks)))
	return b.String()
}

func (m *Input) Callback(pg *Program) (string, bool) {
	var content string
	var exit bool
	if m.multi && m.callbacks != nil {
		var values []string
		for _, i := range m.inputs {
			if strutil.IsNotBlank(i.Value()) {
				values = append(values, i.Value())
			}
		}
		content, exit = m.callbacks(values)
	} else if m.callback != nil {
		content, exit = m.callback(m.inputs[0].Value())
	}
	return content, exit
}

type InputOpt func(*Input)

func WithPlaceholder(placeholder string) InputOpt {
	return func(i *Input) {
		i.placeholder = placeholder
	}
}

func WithMulti() InputOpt {
	return func(i *Input) {
		i.multi = true
	}
}

func WithInputCallbacks(fn func([]string) (string, bool)) InputOpt {
	return func(i *Input) {
		i.callbacks = fn
	}
}

func WithInputCallback(fn func(string) (string, bool)) InputOpt {
	return func(i *Input) {
		i.callback = fn
	}
}
