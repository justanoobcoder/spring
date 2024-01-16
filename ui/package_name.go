package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	springlist "github.com/justanoobcoder/spring/springlist"
)

func (m Model) updatePackageName(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			input := strings.TrimSpace(m.textInput.Value())
			if input != "" {
				m.PackageName = input
			}
			m.state = choosePackaging
			m.list = springlist.NewNormalListModel(
				"Packaging",
				getPackagingOptions(m.springBoot),
				m.Packaging,
				m.list.Width(), m.list.Height(),
			)
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) viewPackageName() string {
	return fmt.Sprintf(
		"Package Name:\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
