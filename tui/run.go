package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	p := tea.NewProgram(tuiModule(), tea.WithAltScreen())
	err := p.Start()
	return err
}
