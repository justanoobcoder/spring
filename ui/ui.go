package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/ui/style"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	switch m.state {
	case chooseProjectType:
		return m.updateProjectType(msg)
	case chooseLanguage:
		return m.updateLanguage(msg)
	case chooseBootVersion:
		return m.updateBootVersion(msg)
	case inputGroupId:
		return m.updateInputGroupId(msg)
	case inputArtifactId:
		return m.updateArtifactId(msg)
	case inputName:
		return m.updateName(msg)
	case inputDescription:
		return m.updateDescription(msg)
	}
	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return style.QuitTextStyle.Render(
			fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
				m.typE,
				m.language,
				m.bootVersion,
				m.groupId,
				m.artifactId,
				m.name,
				m.description,
			),
		)
	}
	var s string
	switch m.state {
	case chooseProjectType:
		s = m.viewProjectType()
	case chooseLanguage:
		s = m.viewLanguage()
	case chooseBootVersion:
		s = m.viewBootVersion()
	case inputGroupId:
		s = m.viewInputGroupId()
	case inputArtifactId:
		s = m.viewArtifactId()
	case inputName:
		s = m.viewName()
	case inputDescription:
		s = m.viewDescription()
	}
	return s
}
