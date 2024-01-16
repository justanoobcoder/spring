package ui

import (
	"slices"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	springlist "github.com/justanoobcoder/spring/list"
	"github.com/justanoobcoder/spring/springboot"
	"github.com/justanoobcoder/spring/style"
)

func getDependencies(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Dependencies.Values {
		for _, v2 := range v.Values {
			items = append(items, springlist.FilteredListItem{
				Id:       v2.ID,
				Name:     v2.Name,
				Category: v.Name,
				Desc:     v2.Description,
			})
		}
	}

	return items
}

func (m Model) updateDependencies(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch key {
		case "enter", " ":
			selected := m.list.SelectedItem().(springlist.FilteredListItem)
			if !selected.Selected {
				m.Dependencies = append(m.Dependencies, selected.Id)
			} else {
				idx := slices.Index(m.Dependencies, selected.Id)
				m.Dependencies = append(m.Dependencies[:idx], m.Dependencies[idx+1:]...)
			}
			var newList []list.Item
			for _, v := range m.list.Items() {
				if v.(springlist.FilteredListItem).Id != selected.Id {
					newList = append(newList, v)
				} else {
					i := v.(springlist.FilteredListItem)
					newList = append([]list.Item{
						springlist.FilteredListItem{
							Id:       i.Id,
							Name:     i.Name,
							Category: i.Category,
							Desc:     i.Desc,
							Selected: !i.Selected,
						},
					},
						newList...)
				}
			}
			m.list.SetItems(newList)
			m.list.ResetFilter()
			return m, nil
		case "ctrl+s":
			m.state = downloadFileState
			return m, m.downloadFile
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) viewDependencies() string {
	return style.DocStyle.Render(m.list.View())
}
