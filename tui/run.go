package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	p := tea.NewProgram(TuiModule(), tea.WithAltScreen(), tea.WithMouseCellMotion())
	err := p.Start()
	return err
}
