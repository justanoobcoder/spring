package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateDescription(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.description = m.textInput.Value()
			m.state = inputPackageName
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewDescription() string {
	return fmt.Sprintf(
		"Enter project description:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
