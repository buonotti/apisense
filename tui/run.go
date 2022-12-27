package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	p := tea.NewProgram(loadingModule(), tea.WithAltScreen())
	err := p.Start()
	return err
}
