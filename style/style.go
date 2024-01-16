package style

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle        = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
	ChoosenTitleStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.AdaptiveColor{Light: "#84f5a2", Dark: "#3cc962"}).
				Foreground(lipgloss.AdaptiveColor{Light: "#61c77d", Dark: "#61c77d"}).
				Padding(0, 0, 0, 1)
	ChoosenDescStyle = ChoosenTitleStyle.Copy().
				Foreground(lipgloss.AdaptiveColor{Light: "#84f5a2", Dark: "#84f5a2"})
	DocStyle = lipgloss.NewStyle().Margin(1, 2)
)
