package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func getProjectTypes(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Type.Values {
		items = append(items, item{
			id:   v.ID,
			name: v.Name,
		})
	}
	return items
}

func (m Model) updateProjectType(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.typE = i.id
				m.state = chooseLanguage
				m.list = NewList("Choose Language", getLanguages(m.springboot),
					m.springboot.Language.Default, listHeight)
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewProjectType() string {
	return "\n" + m.list.View()
}
