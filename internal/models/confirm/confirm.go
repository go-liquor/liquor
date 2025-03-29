package confirm

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/logrusorgru/aurora/v4"
)

type model struct {
	cursor  int
	choice  string
	choices []string
	text    string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			os.Exit(0)

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m *model) View() string {
	s := strings.Builder{}
	s.WriteString(aurora.Cyan(m.text + "\n\n").String())

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(✳️) ")
		} else {
			s.WriteString("( ) ")
		}
		text := aurora.Green("Yes").String()
		if m.choices[i] == "No" {
			text = aurora.Red("No").String()
		}
		s.WriteString(text)
		s.WriteString("\n\n")
	}

	return s.String()
}

func NewConfirm(text string) (bool, error) {
	md := model{
		text:    text,
		choices: []string{"Yes", "No"},
	}
	p := tea.NewProgram(&md)
	m, err := p.Run()
	if err != nil {
		return false, err
	}
	if m, ok := m.(*model); ok && m.choice != "" {
		if m.choice == "Yes" {
			return true, nil
		}
		return false, nil
	}
	return false, fmt.Errorf("you need choice")

}
