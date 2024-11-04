package cmd

import (
	"github.com/buonotti/apisense/v2/theme"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func greyedOutStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIWhite))
} // grey
func blueStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIBlue))
} // blue
func redStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIRed))
} // red
func greenStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIGreen))
} // green
func yellowStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIYellow))
} // yellow
