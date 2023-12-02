package textinput

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textInput struct {
	txt string
}

func New() *textInput {
	return &textInput{}
}

func (t *textInput) Text() string {
	return t.txt
}

func (t *textInput) Open(textInputTitle string) error {
	if textInputTitle == "" {
		textInputTitle = "Enter text: "
	}
	// 全画面表示
	p := tea.NewProgram(initialModel(textInputTitle), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return err
	}
	if m.(model).shouldSave {
		t.txt = m.(model).resultText
	} else {
		return errors.New("canceled. task name must be set")
	}
	return nil
}

type (
	errMsg error
)

type model struct {
	textInput      textinput.Model
	textInputTitle string
	resultText     string
	shouldSave     bool
	err            error
}

func initialModel(textInputTitle string) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput:      ti,
		textInputTitle: textInputTitle,
		err:            nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlS:
			m.shouldSave = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	m.resultText = m.textInput.Value()
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s%s\n\n%s",
		m.textInputTitle,
		m.textInput.View(),
		"(Enter or Ctrl+s to save)\n(Ctrl+c to quit)",
	) + "\n"
}
