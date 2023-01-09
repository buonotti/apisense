package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var styleHotkey = styleHelp.Copy().
	Bold(true).
	MarginLeft(2)
var styleBase = lipgloss.NewStyle()
var styleHelp = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#536878"))
var stylePrimary = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#F38BA8"))
var styleBold = lipgloss.NewStyle().Bold(true)

// CENTER
var styleContentCenter = lipgloss.NewStyle().
	Align(lipgloss.Center)

// LEFT
var styleContentLeft = lipgloss.NewStyle().
	Align(lipgloss.Center)

// RIGHT
var styleContentRight = lipgloss.NewStyle().
	Align(lipgloss.Center)
