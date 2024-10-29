package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"

	"github.com/buonotti/apisense/theme"
)

// BASE STYLE THAT GOES AROUND THE WHOLE TUI
var docStyle = lipgloss.NewStyle().Margin(1, 2)

// BASE STYLE FOR KEYMAP DESCRIPTIONS
var styleHotkey = styleHelp.Copy().
	Bold(true)

// BASE STYLES FOR MENU RENDERING

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
	stylePrimary = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIWhite))
	styleError   = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIRed))
	styleSuccess = lipgloss.NewStyle().Foreground(theme.Ansi2Color(termenv.ANSIGreen))
)

// CENTER
var styleContentCenter = lipgloss.NewStyle().Align(lipgloss.Center)
var styleContentCenterT1 = lipgloss.NewStyle().Align(lipgloss.Center).Background(theme.Ansi2Color(termenv.ANSICyan))
var styleContentCenterT2 = lipgloss.NewStyle().Align(lipgloss.Center).Background(theme.Ansi2Color(termenv.ANSIRed))
var styleContentCenterT3 = lipgloss.NewStyle().Align(lipgloss.Center).Background(theme.Ansi2Color(termenv.ANSIGreen))
var styleContentCenterT4 = lipgloss.NewStyle().Align(lipgloss.Center).Background(theme.Ansi2Color(termenv.ANSIBlue))
var styleContentCenterT5 = lipgloss.NewStyle().Align(lipgloss.Center).Background(theme.Ansi2Color(termenv.ANSIYellow))
var styleActive = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#5e81ac"))
var styleBase = lipgloss.NewStyle()

// TODO: Change once Flexboxlibrary adds optional footers
var styleFooter = lipgloss.NewStyle().Align(lipgloss.Right).PaddingTop(1000)
var styleBold = lipgloss.NewStyle().Bold(true)
