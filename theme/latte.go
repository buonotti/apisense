package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// latteTheme is inspired by the latte flavour of the catppuccin theme.
type latteTheme struct {
}

func (t *latteTheme) IsDark() bool {
	return false
}

func (t *latteTheme) Base() lipgloss.Color {
	return "#eff1f5"
}

func (t *latteTheme) Text() lipgloss.Color {
	return "#4c4f69"
}

func (t *latteTheme) TextDark() lipgloss.Color {
	return "#9ca0b0"
}

func (t *latteTheme) Error() lipgloss.Color {
	return "#d20f39"
}

func (t *latteTheme) Info() lipgloss.Color {
	return "#1e66f5"
}

func (t *latteTheme) Ok() lipgloss.Color {
	return "#40a02b"
}

func (t *latteTheme) Warning() lipgloss.Color {
	return "#df8e1d"
}
