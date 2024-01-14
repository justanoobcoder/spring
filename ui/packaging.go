package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func getPackagingOptions(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Packaging.Values {
		items = append(items, item{
			id:   v.ID,
			name: v.Name,
		})
	}
	return items
}

func (m Model) updatePackaging(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.packaging = i.id
				m.state = chooseJavaVersion
				m.quitting = true
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewPackaging() string {
	return "\n" + m.list.View()
}
