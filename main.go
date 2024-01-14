package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/justanoobcoder/spring/ui"
)

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := tea.NewProgram(ui.NewModel()).Run(); err != nil {
		panic(err)
	}
}
