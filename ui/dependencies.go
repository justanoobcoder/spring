package ui

import (
	"slices"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/springboot"
)

func getDependencies(sp springboot.SpringBoot) []list.Item {
	var items []list.Item
	for _, v := range sp.Dependencies.Values {
		for _, v2 := range v.Values {
			items = append(items, FilteredListItem{
				id:       v2.ID,
				title:    v2.Name,
				category: v.Name,
				desc:     v2.Description,
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
			selected := m.list.SelectedItem().(FilteredListItem)
			if !selected.selected {
				m.Dependencies = append(m.Dependencies, selected.id)
			} else {
				idx := slices.Index(m.Dependencies, selected.id)
				m.Dependencies = append(m.Dependencies[:idx], m.Dependencies[idx+1:]...)
			}
			var newList []list.Item
			for _, v := range m.list.Items() {
				if v.(FilteredListItem).id != selected.id {
					newList = append(newList, v)
				} else {
					i := v.(FilteredListItem)
					newList = append([]list.Item{
						FilteredListItem{
							id:       i.id,
							title:    i.title,
							category: i.category,
							desc:     i.desc,
							selected: !i.selected,
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
	return docStyle.Render(m.list.View())
}
