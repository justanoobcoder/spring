package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/style"
)

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
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
	case inputMetaData:
		return m.updateMetaData(msg)
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
		if m.failed {
			return fmt.Sprintf("%s %s", style.ErrorStyle.Render("Error:"), m.message) + "\n"
		}
		return style.SuccessStyle.Render(m.message) + "\n"
	}
	var s string
	switch m.state {
	case chooseProjectType:
		s = m.viewProjectType()
	case chooseLanguage:
		s = m.viewLanguage()
	case chooseBootVersion:
		s = m.viewBootVersion()
	case inputMetaData:
		s = m.viewMetaData()
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
