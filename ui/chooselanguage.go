package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func getLanguages(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Language.Values {
		items = append(items, item{
			id:   v.ID,
			name: v.Name,
		})
	}
	return items
}

func (m Model) updateLanguage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.language = i.id
				m.state = chooseBootVersion
				m.list = NewList("Choose Spring Boot Version", getBootVersions(m.springboot),
					m.springboot.Version.Default, listHeight)
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewLanguage() string {
	return "\n" + m.list.View()
}
