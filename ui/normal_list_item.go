package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/ui/style"
)

type normalListItem struct {
	id   string
	name string
}

func (i normalListItem) FilterValue() string { return "" }

type normalListItemDelegate struct{}

func (d normalListItemDelegate) Height() int                             { return 1 }
func (d normalListItemDelegate) Spacing() int                            { return 0 }
func (d normalListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d normalListItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(normalListItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.name)

	fn := style.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return style.SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func NewNormalListModel(title string, items []list.Item, def string, width, height int) list.Model {
	l := list.New(items, normalListItemDelegate{}, width, height)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	for i := range items {
		if items[i].(normalListItem).id == def {
			l.Select(i)
			break
		}
	}
	return l
}
