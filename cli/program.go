package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/goutil/strutil"
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
	ch       chan string
}

func NewProgram(mfs ...func() IModel) *Program {
	pg := &Program{models: mfs, ch: make(chan string)}
	p := tea.NewProgram(pg)
	pg.Program = p
	return pg
}

func (pg *Program) AddModel(mfs ...func() IModel) {
	pg.models = append(pg.models, mfs...)
}

func (pg *Program) Output(s string, v ...any) {
	pg.ch <- fmt.Sprintf(s, v...)
}

func (pg *Program) NewLine() {
	pg.Output(" ")
}

func (pg *Program) Start() {
	if len(pg.models) == 0 {
		println(ErrStyle.Render("No model specified!"))
		return
	}

	go func() {
		for s := range pg.ch {
			if !strutil.IsEmpty(s) {
				pg.Program.Println(LmarginStyle.Render(s))
			}
		}
	}()

	pg.Next()
	err := pg.Program.Start()
	if err != nil {
		println(ErrStyle.Render(err.Error()))
	}
}

func (pg *Program) Init() tea.Cmd {
	return nil
}

func (pg *Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			if _, ok := pg.current.(*Generator); ok {
				return pg, nil
			}

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

			if _, ok := c.(*Generator); ok {
				return pg, c.Update(msg)
			}

			content, exit := c.Callback(pg)
			pg.ch <- content
			pg.Next()
			if _, ok := pg.current.(*empty); ok || exit {
				time.Sleep(time.Millisecond)
				pg.done = true
				return pg, tea.Quit
			}

			var cmd tea.Cmd
			if _, ok := pg.current.(*Input); ok {
				cmd = textinput.Blink
			} else if _, ok := pg.current.(*Generator); ok {
				cmd = spinner.Tick
			}
			return pg, cmd
		}
	}
	return pg, pg.current.Update(msg)
}

func (pg *Program) View() string {
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

func (empty) Update(tea.Msg) tea.Cmd           { return tea.Quit }
func (empty) View() string                     { return "" }
func (empty) Callback(*Program) (string, bool) { return "", true }
