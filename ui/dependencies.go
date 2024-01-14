package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func getDependencies(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Dependencies.Values {
		for _, v2 := range v.Values {
			items = append(items, filteritem{
				id:    v2.ID,
				title: v2.Name,
				desc:  v2.Description,
			})
		}
	}

	return items
}

func (m Model) updateDependencies(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			selected := m.list.SelectedItem().(filteritem)
			m.dependencies = append(m.dependencies, selected.id)
			var newList []list.Item
			for _, v := range m.list.Items() {
				if v.(filteritem).id != selected.id {
					newList = append(newList, v)
				}
			}
			m.list.SetItems(newList)
			m.list.ResetFilter()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewDependencies() string {
	return docStyle.Render(m.list.View())
}
