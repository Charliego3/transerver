package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gookit/goutil/strutil"
)

type Select struct {
	list.Model
	title          string
	width          int
	heigh          int
	showStatusBar  bool
	showFilter     bool
	showPagination bool
	showHelp       bool
	filterEnable   bool

	callback func(list.Model) string
}

func NewSelect(items []list.Item, opts ...SelectOpt) IModel {
	sl := &Select{}
	for _, opt := range opts {
		opt(sl)
	}

	width := sl.width
	if width == 0 {
		width = 20
	}
	heigh := sl.heigh
	if heigh == 0 {
		if !sl.showPagination {
			heigh = len(items) + 2
		} else {
			heigh = 10
		}
	}
	l := list.New(items, selectItemDelegate{}, width, heigh)
	l.SetShowStatusBar(sl.showStatusBar)
	l.SetFilteringEnabled(sl.filterEnable)
	l.SetShowFilter(sl.showFilter)
	l.SetShowPagination(sl.showPagination)
	l.SetShowHelp(sl.showHelp)
	l.SetShowTitle(false)
	l.KeyMap.Quit.SetEnabled(false)
	l.KeyMap.ShowFullHelp.SetEnabled(false)
	l.Styles.PaginationStyle = PaginationStyle
	l.Styles.HelpStyle = HelpStyle
	sl.Model = l
	return sl
}

func (l *Select) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	l.Model, cmd = l.Model.Update(msg)
	return nil, cmd
}

func (l *Select) View() string {
	var s string
	if strutil.IsNotBlank(l.title) {
		s += TitleStyle.Render(l.title)
	}
	return s + "\n" + l.Model.View()
}

func (l *Select) Callback(pg Program) {
	if l.callback == nil {
		return
	}
	pg.Println(l.callback(l.Model))
}

type SelectItem string

func (i SelectItem) FilterValue() string { return "" }

type selectItemDelegate struct{}

func (d selectItemDelegate) Height() int                               { return 1 }
func (d selectItemDelegate) Spacing() int                              { return 0 }
func (d selectItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d selectItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(SelectItem)
	if !ok {
		return
	}

	cursor := " "
	content := ItemStyle.Render(string(i))
	if index == m.Index() {
		cursor = "◉"
		content = SelectedItemStyle.Render(string(i))
	}

	_, _ = fmt.Fprint(w, CursorStyle.Render(cursor)+content)
}

type SelectOpt func(*Select)

func WithTitle(title string) SelectOpt {
	return func(s *Select) {
		s.title = title
	}
}

func WithWidth(width int) SelectOpt {
	return func(s *Select) {
		s.width = width
	}
}

func WithHeigh(heigh int) SelectOpt {
	return func(s *Select) {
		s.heigh = heigh
	}
}

func WithShowStatusBar() SelectOpt {
	return func(s *Select) {
		s.showStatusBar = true
	}
}

func WithShowFilter() SelectOpt {
	return func(s *Select) {
		s.showFilter = true
	}
}

func WithShowPagination() SelectOpt {
	return func(s *Select) {
		s.showPagination = true
	}
}

func WithShowHelp() SelectOpt {
	return func(s *Select) {
		s.showHelp = true
	}
}

func WithFilteringEnable() SelectOpt {
	return func(s *Select) {
		s.filterEnable = true
	}
}

func WithCallback(callback func(list.Model) string) SelectOpt {
	return func(s *Select) {
		s.callback = callback
	}
}
