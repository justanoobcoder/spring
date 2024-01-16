package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
	springlist "github.com/justanoobcoder/spring/springlist"
	"github.com/justanoobcoder/spring/style"
)

func getJavaVersion(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.JavaVersion.Values {
		items = append(items, springlist.NormalListItem{
			Id:   v.ID,
			Name: v.Name,
		})
	}
	return items
}

func (m Model) updateJavaVersion(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(springlist.NormalListItem)
			if ok {
				m.JavaVersion = i.Id
				m.state = chooseDependencies
				m.list = list.New(
					getDependencies(m.springBoot),
					springlist.FilteredListItemDelegate{},
					m.list.Width()/2, m.list.Height()-3,
				)
				m.list.Title = "Dependencies"
				m.list.Styles.Title = style.TitleStyle
				m.list.SetShowStatusBar(false)
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

func (m Model) viewJavaVersion() string {
	return "\n  " + m.list.View()
}
