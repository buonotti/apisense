package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/buonotti/apisense/theme"
)

// BASE STYLE THAT GOES AROUND THE WHOLE TUI
var docStyle = lipgloss.NewStyle().Margin(1, 2)

// BASE STYLE FOR KEYMAP DESCRIPTIONS
var styleHotkey = styleHelp.Copy().
	Bold(true).
	MarginLeft(2)

// BASE STYLES FOR MENU RENDERING
var styleBase = lipgloss.NewStyle()
var styleBold = lipgloss.NewStyle().Bold(true)

// BASE STYLES FOR THE CONFIG VIEWPORT RENDERING
var titleStyle = func() lipgloss.Style {
	b := lipgloss.RoundedBorder()
	b.Right = "├"
	return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
}()

var infoStyle = func() lipgloss.Style {
	b := lipgloss.RoundedBorder()
	b.Left = "┤"
	return titleStyle.Copy().BorderStyle(b)
}()

// COLORS
var styleHelp = theme.Overlay0().S()
var stylePrimary = theme.Red().S()
var styleInfo = theme.Blue().S()
var styleSuccess = theme.Green().S()

// CENTER
var styleContentCenter = lipgloss.NewStyle().
	Align(lipgloss.Center)
