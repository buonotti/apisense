package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// mochaTheme is inspired by the mocha flavour of the catppuccin theme.
type mochaTheme struct {
}

func (t *mochaTheme) IsDark() bool {
	return true
}

func (t *mochaTheme) Base() lipgloss.Color {
	return "#1e1e2e"
}

func (t *mochaTheme) Text() lipgloss.Color {
	return "#cdd6f4"
}

func (t *mochaTheme) Overlay0() lipgloss.Color {
	return "#6c7086"
}

func (t *mochaTheme) Red() lipgloss.Color {
	return "#f38ba8"
}

func (t *mochaTheme) Blue() lipgloss.Color {
	return "#89b4fa"
}

func (t *mochaTheme) Green() lipgloss.Color {
	return "#a6e3a1"
}

func (t *mochaTheme) Yellow() lipgloss.Color {
	return "#f9e2af"
}
