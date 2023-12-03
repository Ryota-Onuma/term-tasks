package dialog

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type dialog struct{}

func New() *dialog {
	return &dialog{}
}

func (d *dialog) Open(msg string, secForClose int) error {
	if secForClose <= 0 {
		secForClose = 5
	}

	p := tea.NewProgram(model{secForClose: secForClose, msg: msg})
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

type model struct {
	secForClose int
	msg         string
}

func (m model) Init() tea.Cmd {
	return tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tickMsg:
		m.secForClose--
		if m.secForClose <= 0 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

func (m model) View() string {
	return m.msg
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Second)
	return tickMsg{}
}
