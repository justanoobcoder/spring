package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

type statusMsg int

func (m Model) downloadFile() tea.Msg {
	body := springboot.Request{
		Dependencies: strings.Join(m.Dependencies, ","),
		JavaVersion:  m.JavaVersion,
		Type:         m.Type,
		Version:      m.Version,
		Packaging:    m.Packaging,
		Language:     m.Language,
		BootVersion:  m.BootVersion,
		GroupId:      m.GroupId,
		ArtifactId:   m.ArtifactId,
		Name:         m.Name,
		Description:  m.Description,
		PackageName:  m.PackageName,
	}
	var filename string
	for _, t := range m.springboot.Type.Values {
		if t.ID == m.Type {
			filename = strings.TrimPrefix(t.Action, "/")
			break
		}
	}
	statusCode, err := springboot.Download(body, filename)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(statusCode)
}

func (m Model) updateDownloadFile(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		if msg == 200 {
			m.quitting = true
			return m, tea.Quit
		}
		return m, tea.Quit
	case errMsg:
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m Model) viewDownloadFile() string {
	return "Downloading file...\n"
}
