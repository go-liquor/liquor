package choice

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/logrusorgru/aurora/v4"
)

type ChoiceOption struct {
	Title string
	Value string
	Color func(arg interface{}) aurora.Value
}

type model struct {
	cursor  int
	choice  string
	choices []ChoiceOption
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
			m.choice = m.choices[m.cursor].Value
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
			s.WriteString("(✅) ")
		} else {
			s.WriteString("( ) ")
		}

		text := m.choices[i].Title
		if m.choices[i].Color != nil {
			text = m.choices[i].Color(text).String()
		}

		s.WriteString(text)
		s.WriteString("\n\n")
	}

	return s.String()
}

func NewChoice(text string, choices []ChoiceOption) (string, error) {
	md := model{
		text:    text,
		choices: choices,
	}
	p := tea.NewProgram(&md)
	m, err := p.Run()
	if err != nil {
		return "", err
	}
	if m, ok := m.(*model); ok && m.choice != "" {
		return m.choice, nil
	}
	return "", fmt.Errorf("you need choice")

}
