package input

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	aurora "github.com/logrusorgru/aurora/v4"
)

type model struct {
	textInput textinput.Model
	message   string
	err       error
}

func (m *model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			os.Exit(0)
		case tea.KeyEnter:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	return fmt.Sprintf(
		"%s: %s\n",
		aurora.Cyan(m.message),
		m.textInput.View(),
	) + "\n"
}

func NewInput(message string, placeholder string) (string, error) {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 100
	md := model{
		textInput: ti,
		message:   message,
		err:       nil,
	}
	p := tea.NewProgram(&md)
	if _, err := p.Run(); err != nil {
		return "", err
	}
	return md.textInput.Value(), nil
}
