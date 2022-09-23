package main

import (
	"sync/atomic"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type IModel interface {
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
	Callback(Program)
}

type Program struct {
	*tea.Program
	models   []IModel
	current  atomic.Int32
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
		case tea.KeyCtrlC, tea.KeyCtrlQ:
			pg.quitting = true
			return pg, tea.Quit

		case tea.KeyEnter:
			if pg.nonNext() {
				pg.done = true
				return pg, tea.Quit
			}

			c := pg.Current()
			go func() {
				c.Callback(pg)
			}()
			pg.current.Add(1)
			if pg.nonNext() {
				time.Sleep(time.Millisecond)
				pg.done = true
				return pg, tea.Quit
			}
			return pg, nil
		}
	}
	m, cmd := pg.Current().Update(msg)
	if m == nil {
		m = pg
	}
	return m, cmd
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
	current := pg.current.Load()
	return current < 0 || int(current) > len(pg.models)-1
}

func (pg Program) Current() IModel {
	if pg.nonNext() {
		return &empty{}
	}
	return pg.models[pg.current.Load()]
}

type empty struct{}

func (empty) Update(tea.Msg) (tea.Model, tea.Cmd) { return nil, nil }
func (empty) View() string                        { return "" }
func (empty) Callback(Program)                    {}
