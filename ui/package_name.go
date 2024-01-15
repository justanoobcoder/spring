package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) updatePackageName(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.PackageName = m.textInput.Value()
			m.state = choosePackaging
			m.list = NewNormalListModel("Choose Packaging", getPackagingOptions(m.springboot),
				m.springboot.Packaging.Default, m.list.Width(), m.list.Height())
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewPackageName() string {
	return fmt.Sprintf(
		"Enter project package name:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
