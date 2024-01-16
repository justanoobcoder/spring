package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updateGroupId(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			input := strings.TrimSpace(m.textInput.Value())
			if input != "" {
				m.GroupId = input
			}
			m.state = inputArtifactId
			m.textInput.Placeholder = m.ArtifactId
			m.textInput.SetValue(m.ArtifactId)
			m.textInput.CursorEnd()
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewGroupId() string {
	return fmt.Sprintf(
		"Group ID:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
