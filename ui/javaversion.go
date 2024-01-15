package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func getJavaVersion(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.JavaVersion.Values {
		items = append(items, item{
			id:   v.ID,
			name: v.Name,
		})
	}
	return items
}

func (m Model) updateJavaVersion(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.javaVersion = i.id
				m.state = chooseDependencies
				m.list = list.New(getDependencies(m.springboot), filterItemDelegate{}, 100, 30)
				m.list.Title = "Choose Dependencies"
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewJavaVersion() string {
	return "\n" + m.list.View()
}
