package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/justanoobcoder/spring/ui/style"
	"github.com/muesli/reflow/truncate"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type filteredListItem struct {
	id, title, category, desc string
	selected                  bool
}

func (i filteredListItem) Title() string       { return i.title }
func (i filteredListItem) Description() string { return i.desc }
func (i filteredListItem) FilterValue() string { return i.title }

type filteredListItemDelegate struct{}

func (d filteredListItemDelegate) Height() int                             { return 2 }
func (d filteredListItemDelegate) Spacing() int                            { return 1 }
func (d filteredListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d filteredListItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	var (
		title, desc  string
		choosen      bool
		matchedRunes []int
		s            = list.NewDefaultItemStyles()
	)

	if i, ok := listItem.(filteredListItem); ok {
		title = i.Title() + " - " + i.category
		desc = i.Description()
		choosen = i.selected
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
