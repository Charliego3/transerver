package main

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type IModel interface {
	Update(tea.Msg) tea.Cmd
	View() string
	Callback(Program)
}

type Program struct {
	*tea.Program
	models   []IModel
	current  int
	quitting bool
	done     bool
}

func NewProgram(models ...IModel) *tea.Program {
	if len(models) == 0 {
		println(ErrStylr.Render("No model specified!"))
		return nil
	}

	pg := &Program{models: models}
	p := tea.NewProgram(pg)
	pg.Program = p
	return p
}

func (pg Program) Init() tea.Cmd {
	return textinput.Blink
}

func (pg Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			pg.quitting = true
			return pg, tea.Quit

		case tea.KeyEnter:
			c := pg.Current()
			if _, ok := c.(*Input); ok {
				cmd := c.Update(msg)
				if cmd == nil {
					return pg, cmd
				}
			}

			go func() {
				c.Callback(pg)
			}()
			pg.current++
			if pg.nonNext() {
				time.Sleep(time.Millisecond)
				pg.done = true
				return pg, tea.Quit
			}
			return pg, nil
		}
	}
	return pg, pg.Current().Update(msg)
}

func (pg Program) View() string {
	if pg.done {
		return ""
	}
	if pg.quitting {
		return QuitTextStyle.Render("Bye-Bye!")
	}
	return pg.Current().View()
}

func (pg *Program) nonNext() bool {
	current := pg.current
	return current < 0 || current > len(pg.models)-1
}

func (pg Program) Current() IModel {
	if pg.nonNext() {
		return &empty{}
	}
	return pg.models[pg.current]
}

type empty struct{}

func (empty) Update(tea.Msg) tea.Cmd { return nil }
func (empty) View() string           { return "" }
func (empty) Callback(Program)       {}
