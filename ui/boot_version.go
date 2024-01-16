package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
	springlist "github.com/justanoobcoder/spring/springlist"
)

func getBootVersions(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.BootVersion.Values {
		items = append(items, springlist.NormalListItem{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	return items
}

func (m Model) updateBootVersion(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(springlist.NormalListItem)
			if ok {
				m.BootVersion = i.Id
				m.state = inputGroupId
			}
			return m, nil
		}
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 1)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewBootVersion() string {
	return "\n" + m.list.View()
}
