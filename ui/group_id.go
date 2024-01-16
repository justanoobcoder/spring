package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateGroupId(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.GroupId = m.textInput.Value()
			m.state = inputArtifactId
			m.textInput.Placeholder = "Artifact ID"
			m.textInput.SetValue(m.springBoot.ArtifactID.Default)
			m.textInput.CursorEnd()
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewGroupId() string {
	return fmt.Sprintf(
		"Enter project Group ID:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
