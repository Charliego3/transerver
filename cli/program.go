package main

import (
	"github.com/gookit/goutil/strutil"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type IModel interface {
	Update(tea.Msg) tea.Cmd
	View() string
	Callback(*Program) (string, bool)
}

type Program struct {
	*tea.Program
	models   []func() IModel
	current  IModel
	quitting bool
	done     bool
}

func NewProgram(mfs ...func() IModel) *Program {
	pg := &Program{models: mfs}
	p := tea.NewProgram(pg)
	pg.Program = p
	return pg
}

func (pg *Program) AddModel(mfs ...func() IModel) {
	pg.models = append(pg.models, mfs...)
}

func (pg *Program) Start() {
	if len(pg.models) == 0 {
		println(ErrStyle.Render("No model specified!"))
		return
	}

	pg.Next()
	err := pg.Program.Start()
	if err != nil {
		println(ErrStyle.Render(err.Error()))
	}
}

func (pg Program) Init() tea.Cmd {
	return textinput.Blink
}

func (pg *Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			pg.quitting = true
			return pg, tea.Quit

		case tea.KeyEnter:
			c := pg.current
			if _, ok := c.(*Input); ok {
				cmd := c.Update(msg)
				if cmd == nil {
					return pg, cmd
				}
			}

			content, exit := c.Callback(pg)
			if strutil.IsNotBlank(content) {
				go func() {
					pg.Println(content)
				}()
			}
			pg.Next()
			if _, ok := pg.current.(*empty); ok || exit {
				time.Sleep(time.Millisecond)
				pg.done = true
				return pg, tea.Quit
			}
			return pg, nil
		}
	}
	return pg, pg.current.Update(msg)
}

func (pg Program) View() string {
	if pg.done {
		return "\n"
	}
	if pg.quitting {
		return QuitTextStyle.Render("Bye-Bye!")
	}
	return pg.current.View()
}

func (pg *Program) Next() {
	if len(pg.models) == 0 {
		pg.current = &empty{}
		return
	}

	pg.current = pg.models[0]()
	pg.models = pg.models[1:]
}

type empty struct{}

func (empty) Update(tea.Msg) tea.Cmd           { return nil }
func (empty) View() string                     { return "" }
func (empty) Callback(*Program) (string, bool) { return "", true }
