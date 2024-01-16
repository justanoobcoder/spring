package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

type errMsg struct{ error }

func (e errMsg) Error() string { return e.error.Error() }

type statusMsg int

func (m Model) downloadFileMsg() tea.Msg {
	request := springboot.Request{
		Type:         m.Type,
		BootVersion:  m.BootVersion,
		Language:     m.Language,
		GroupId:      m.GroupId,
		ArtifactId:   m.ArtifactId,
		Name:         m.Name,
		Description:  m.Description,
		PackageName:  m.PackageName,
		Version:      m.Version,
		Packaging:    m.Packaging,
		JavaVersion:  m.JavaVersion,
		Dependencies: strings.Join(m.Dependencies, ","),
	}
	var filename string
	for _, t := range m.springBoot.Type.Values {
		if t.ID == m.Type {
			filename = strings.TrimPrefix(t.Action, "/")
			break
		}
	}
	statusCode, err := springboot.Download(request, filename)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(statusCode)
}

func (m Model) updateDownloadFile(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		if msg == 200 {
			m.message = "Downloaded file successfully"
			m.quitting = true
			return m, tea.Quit
		}
		return m, tea.Quit
	case errMsg:
		m.failed = true
		m.message = msg.Error()
		m.quitting = true
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) viewDownloadFile() string {
	return "Downloading file...\n"
}
