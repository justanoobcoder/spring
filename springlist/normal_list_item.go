package springlist

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/style"
)

type NormalListItem struct {
	Id   string
	Name string
}

func (i NormalListItem) FilterValue() string { return "" }

type NormalListItemDelegate struct{}

func (d NormalListItemDelegate) Height() int                             { return 1 }
func (d NormalListItemDelegate) Spacing() int                            { return 0 }
func (d NormalListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d NormalListItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(NormalListItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	fn := style.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return style.SelectedNormalListItemStyle.
				Render(style.SpringBootIcon + " " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func NewNormalListModel(title string, items []list.Item, def string, width, height int) list.Model {
	l := list.New(items, NormalListItemDelegate{}, width, height)
	l.Title = title
	l.Styles.Title = style.TitleStyle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	for i := range items {
		if items[i].(NormalListItem).Id == def {
			l.Select(i)
			break
		}
	}
	return l
}
