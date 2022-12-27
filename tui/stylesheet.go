package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var styleHotkey = styleHelp.Copy().
	Bold(true).
	MarginLeft(2)
var styleHelp = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#536878"))
var stylePrimary = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#F38BA8"))
var styleContent = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).Align(lipgloss.Center).MarginRight(2)
var styleContentRed = styleContent.Copy().Background(lipgloss.Color("ff0000"))
