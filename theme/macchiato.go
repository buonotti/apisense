package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// macchiatoTheme is inspired by the macchiato flavour of the catppuccin theme.
type macchiatoTheme struct {
}

func (t *macchiatoTheme) IsDark() bool {
	return true
}

func (t *macchiatoTheme) Base() lipgloss.Color {
	return "#24273a"
}

func (t *macchiatoTheme) Text() lipgloss.Color {
	return "#cad3f5"
}

func (t *macchiatoTheme) TextDark() lipgloss.Color {
	return "#6e738d"
}

func (t *macchiatoTheme) Error() lipgloss.Color {
	return "#ed8796"
}

func (t *macchiatoTheme) Info() lipgloss.Color {
	return "#8aadf4"
}

func (t *macchiatoTheme) Ok() lipgloss.Color {
	return "#a6da95"
}

func (t *macchiatoTheme) Warning() lipgloss.Color {
	return "#eed49f"
}
