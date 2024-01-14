package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type filteritem struct {
	id, title, desc string
}

func (i filteritem) Title() string       { return i.title }
func (i filteritem) Description() string { return i.desc }
func (i filteritem) FilterValue() string { return i.title }
