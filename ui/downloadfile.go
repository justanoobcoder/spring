package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func (m Model) updateDownloadFile(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	springboot.Download(body, filename)
	m.quitting = true
	return m, tea.Quit
}

func (m Model) viewDownloadFile() string {
	return "Downloading file...\n"
}
