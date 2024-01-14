package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateInputGroupId(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.groupId = m.textInput.Value()
			m.state = inputArtifactId
			m.textInput.Placeholder = "Artifact ID"
			m.textInput.SetValue(m.springboot.ArtifactID.Default)
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewInputGroupId() string {
	return fmt.Sprintf(
		"Enter project Group ID:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
