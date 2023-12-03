package list

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	l "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

type listInput struct {
	txt string
}

func New() *listInput {
	return &listInput{}
}

func (t *listInput) SetText(txt string) *listInput {
	t.txt = txt
	return t
}

func (t *listInput) Text() string {
	return t.txt
}

func (t *listInput) Open(listTitle string, items []Item) error {
	if len(items) == 0 {
		return errors.New("no items")
	}

	var litems []l.Item
	for _, item := range items {
		litems = append(litems, item)
	}

	const defaultWidth = 20

	l := l.New(litems, itemDelegate{}, defaultWidth, listHeight)
	l.Title = listTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	if t.txt != "" {
		_, index, ok := lo.FindIndexOf(items, func(i Item) bool { return i.Title() == t.txt })
		if !ok {
			return fmt.Errorf("item not found: %s", t.txt)
		}
		l.Select(index)
	}
	p := tea.NewProgram(model{list: l}, tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return err
	}
	if m.(model).shouldSave {
		t.txt = m.(model).resultText
	}

	return nil
}

type Item struct {
	title string
}

func NewItem(title string) Item {
	return Item{title}
}

func (i Item) Title() string       { return i.title }
func (i Item) FilterValue() string { return i.title }

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem l.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Title())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list       list.Model
	resultText string
	shouldSave bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.resultText = m.list.SelectedItem().FilterValue()
			m.shouldSave = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}
