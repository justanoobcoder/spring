package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateArtifactId(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			input := m.textInput.Value()
			if input != "" {
				m.ArtifactId = input
			}
			m.state = inputApplicationName
			m.textInput.Placeholder = m.Name
			m.textInput.SetValue(m.Name)
			m.textInput.CursorEnd()
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewArtifactId() string {
	return fmt.Sprintf(
		"Artifact ID:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
