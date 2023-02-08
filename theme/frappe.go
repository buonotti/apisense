package theme

import (
	"github.com/charmbracelet/lipgloss"
)

// frappeTheme is inspired by the macchiato flavour of the catppuccin theme.
type frappeTheme struct {
}

func (t *frappeTheme) IsDark() bool {
	return true
}

func (t *frappeTheme) Base() lipgloss.Color {
	return "#303446"
}

func (t *frappeTheme) Text() lipgloss.Color {
	return "#c6d0f5"
}

func (t *frappeTheme) Overlay0() lipgloss.Color {
	return "#737994"
}

func (t *frappeTheme) Red() lipgloss.Color {
	return "#e78284"
}

func (t *frappeTheme) Blue() lipgloss.Color {
	return "#8caaee"
}

func (t *frappeTheme) Green() lipgloss.Color {
	return "#a6d189"
}

func (t *frappeTheme) Yellow() lipgloss.Color {
	return "#e5c890"
}
