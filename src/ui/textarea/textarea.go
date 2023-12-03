package textarea

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type textArea struct {
	txt string
}

func New() *textArea {
	return &textArea{}
}
func (t *textArea) SetText(txt string) *textArea {
	t.txt = txt
	return t
}

func (t *textArea) Text() string {
	return t.txt
}

func (t *textArea) Open(textAreaTitle string) error {
	if textAreaTitle == "" {
		textAreaTitle = "Enter text below..."
	}
	// 全画面表示
	p := tea.NewProgram(t.initialModel(textAreaTitle), tea.WithAltScreen())

	m, err := p.Run()
	if err != nil {
		return err
	}
	if m.(model).shouldSave {
		t.txt = m.(model).resultText
	}
	return nil
}

type errMsg error

type model struct {
	textarea      textarea.Model
	textAreaTitle string
	resultText    string
	shouldSave    bool
	err           error
}

func (t *textArea) initialModel(textAreaTitle string) model {
	ti := textarea.New()
	ti.ShowLineNumbers = false

	ti.Focus()
	if t.txt != "" {
		ti.SetValue(t.txt)
	}

	return model{
		textAreaTitle: textAreaTitle,
		textarea:      ti,
		err:           nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlS:
			m.shouldSave = true
			return m, tea.Quit
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	m.resultText = m.textarea.Value()
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.textAreaTitle,
		m.textarea.View(),
		"(Ctrl+s to save)\n(Ctrl+c to quit)",
	) + "\n\n"
}
