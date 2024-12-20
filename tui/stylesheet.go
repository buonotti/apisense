package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/buonotti/apisense/v2/theme"
)

// BASE STYLE THAT GOES AROUND THE WHOLE TUI
var docStyle = lipgloss.NewStyle().Margin(1, 2)

// BASE STYLE FOR KEYMAP DESCRIPTIONS
var styleHotkey = styleHelp.Copy().
	Bold(true)

// BASE STYLES FOR MENU RENDERING
var (
	styleBase = lipgloss.NewStyle()
	styleBold = lipgloss.NewStyle().Bold(true)
)

// BASE STYLES FOR THE CONFIG VIEWPORT RENDERING
func titleStyle() lipgloss.Style {
	b := lipgloss.RoundedBorder()
	b.Right = "├"
	return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
}

func infoStyle() lipgloss.Style {
	b := lipgloss.RoundedBorder()
	b.Left = "┤"
	return titleStyle().Copy().BorderStyle(b)
}

// COLORS
var (
	styleHelp    = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIWhite))
	stylePrimary = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIRed))
	styleInfo    = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIBlue))
	styleSuccess = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIGreen))
)

// CENTER
var styleContentCenter = lipgloss.NewStyle().
	Align(lipgloss.Center)
