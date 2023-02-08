package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type External struct {
	ThemeName     string `toml:"name"`
	Background    string `toml:"background"`
	Foreground    string `toml:"foreground"`
	ForegroundAlt string `toml:"foreground_alt"`
	Red           string `toml:"red"`
	Blue          string `toml:"blue"`
	Green         string `toml:"green"`
	Yellow        string `toml:"yellow"`
}

func (e External) Name() string {
	return e.ThemeName
}

func (e External) Base() lipgloss.Color {
	return lipgloss.Color(e.Background)
}

func (e External) Text() lipgloss.Color {
	return lipgloss.Color(e.Foreground)
}

func (e External) TextDark() lipgloss.Color {
	return lipgloss.Color(e.ForegroundAlt)
}

func (e External) Error() lipgloss.Color {
	return lipgloss.Color(e.Red)
}

func (e External) Info() lipgloss.Color {
	return lipgloss.Color(e.Blue)
}

func (e External) Ok() lipgloss.Color {
	return lipgloss.Color(e.Green)
}

func (e External) Warning() lipgloss.Color {
	return lipgloss.Color(e.Yellow)
}
