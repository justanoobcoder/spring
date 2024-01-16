package springlist

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/justanoobcoder/spring/style"
	"github.com/muesli/reflow/truncate"
)

type FilteredListItem struct {
	Id, Name, Category, Desc string
	Selected                 bool
}

func (i FilteredListItem) Title() string       { return i.Name }
func (i FilteredListItem) Description() string { return i.Desc }
func (i FilteredListItem) FilterValue() string { return i.Name }

type FilteredListItemDelegate struct{}

func (d FilteredListItemDelegate) Height() int                             { return 2 }
func (d FilteredListItemDelegate) Spacing() int                            { return 1 }
func (d FilteredListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d FilteredListItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	var (
		title, desc  string
		choosen      bool
		matchedRunes []int
		s            = list.NewDefaultItemStyles()
	)

	if i, ok := item.(FilteredListItem); ok {
		title = i.Title() + " - " + i.Category
		desc = i.Description()
		choosen = i.Selected
	} else {
		return
	}

	if m.Width() <= 0 {
		// short-circuit
		return
	}

	// Prevent text from exceeding list width
	textwidth := uint(m.Width() - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight())
	title = truncate.StringWithTail(title, textwidth, "...")
	var lines []string
	for i, line := range strings.Split(desc, "\n") {
		if i >= d.Height()-1 {
			break
		}
		lines = append(lines, truncate.StringWithTail(line, textwidth, "..."))
	}
	desc = strings.Join(lines, "\n")

	// Conditions
	var (
		isSelected  = index == m.Index()
		emptyFilter = m.FilterState() == list.Filtering && m.FilterValue() == ""
		isFiltered  = m.FilterState() == list.Filtering || m.FilterState() == list.FilterApplied
	)

	if isFiltered && index < len(m.VisibleItems()) {
		// Get indices of matched characters
		matchedRunes = m.MatchesForItem(index)
	}

	if emptyFilter {
		title = s.DimmedTitle.Render(title)
		desc = s.DimmedDesc.Render(desc)
	} else if isSelected && m.FilterState() != list.Filtering {
		if isFiltered {
			// Highlight matches
			unmatched := s.SelectedTitle.Inline(true)
			matched := unmatched.Copy().Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.SelectedTitle.Render(title)
		desc = s.SelectedDesc.Render(desc)
	} else {
		if isFiltered {
			// Highlight matches
			unmatched := s.NormalTitle.Inline(true)
			matched := unmatched.Copy().Inherit(s.FilterMatch)
			title = lipgloss.StyleRunes(title, matchedRunes, matched, unmatched)
		}
		title = s.NormalTitle.Render(title)
		desc = s.NormalDesc.Render(desc)
	}
	if choosen {
		title = style.ChoosenTitleStyle.Render(fmt.Sprintf("%s %s", title, "✓"))
		desc = style.ChoosenDescStyle.Render(desc)
	} else {
		title = s.DimmedTitle.Render(title)
		desc = s.DimmedDesc.Render(desc)
	}

	fmt.Fprintf(w, "%s\n%s", title, desc)
}
