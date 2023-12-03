package table

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tbl "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

type tableArea struct {
	isOK        bool
	selectedRow Row
}

func New() *tableArea {
	return &tableArea{}
}

func (ta *tableArea) IsOK() bool {
	return ta.isOK
}

func (ta *tableArea) SelectedRow() Row {
	return ta.selectedRow
}

type Column struct {
	Title string
	Width int
}

type Row struct {
	Index int
	Body  []string
}

func (ta *tableArea) Open(title string, cols []Column, rs []Row) error {
	fmt.Println("")
	fmt.Printf(" %s ", title)
	fmt.Println("")
	columns := lo.Map(cols, func(p Column, _ int) tbl.Column { return tbl.Column{Title: p.Title, Width: p.Width} })
	rows := lo.Map(rs, func(r Row, _ int) tbl.Row { return tbl.Row(r.Body) })

	t := tbl.New(
		tbl.WithColumns(columns),
		tbl.WithRows(rows),
		tbl.WithFocused(true),
		tbl.WithHeight(len(rows)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#ffffff")).
		BorderBottom(true).
		Padding(0, 1).
		Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#037d9d")).
		Bold(true)
	t.SetStyles(s)
	p := tea.NewProgram(model{table: t})
	m, err := p.Run()
	if err != nil {
		return err
	}

	ta.isOK = m.(model).isOK
	if ta.isOK {
		row := m.(model).table.SelectedRow()
		ta.selectedRow = Row{Index: m.(model).table.Cursor(), Body: row}
	}
	return nil
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
	isOK  bool
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.isOK = true

			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}
