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
		case "ctrl+c":
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
	case inputPackageName:
		return m.updatePackageName(msg)
	case choosePackaging:
		return m.updatePackaging(msg)
	case chooseJavaVersion:
		return m.updateJavaVersion(msg)
	case chooseDependencies:
		return m.updateDependencies(msg)
	case downloadFile:
		return m.updateDownloadFile(msg)
	}
	return m, nil
}

func (m Model) View() string {
	if m.quitting {
		return style.QuitTextStyle.Render(
			fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s, %s, %s\n%v",
				m.Type,
				m.Language,
				m.BootVersion,
				m.GroupId,
				m.ArtifactId,
				m.Name,
				m.Description,
				m.PackageName,
				m.Packaging,
				m.JavaVersion,
				m.Dependencies,
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
	case inputPackageName:
		s = m.viewPackageName()
	case choosePackaging:
		s = m.viewPackaging()
	case chooseJavaVersion:
		s = m.viewJavaVersion()
	case chooseDependencies:
		s = m.viewDependencies()
	case downloadFile:
		s = m.viewDownloadFile()
	}
	return s
}
