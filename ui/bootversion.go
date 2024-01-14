package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
	"github.com/justanoobcoder/spring/ui/style"
)

func getBootVersions(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.BootVersion.Values {
		items = append(items, item{
			id:   v.ID,
			name: v.Name,
		})
	}
	return items
}

func (m Model) updateBootVersion(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.bootVersion = i.id
				m.choice = m.bootVersion
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewBootVersion() string {
	if m.choice != "" {
		return style.QuitTextStyle.Render(fmt.Sprintf("%s, %s, %s", m.typE, m.language, m.choice))
	}
	return "\n" + m.list.View()
}
