package ui

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func (m Model) updateDownloadFile(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("download...: %s, %s, %s, %s, %s, %s, %s, %s, %s, %s\n%v",
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
	)
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
	springboot.CreateProject(body)
	m.quitting = true
	return m, tea.Quit
}

func (m Model) viewDownloadFile() string {
	return "Downloading file...\n"
}
