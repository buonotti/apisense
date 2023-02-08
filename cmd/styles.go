package cmd

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/buonotti/apisense/theme"
)

func greyedOutStyle() lipgloss.Style { return theme.Overlay0().S() } // grey
func blueStyle() lipgloss.Style      { return theme.Blue().S() }     // blue
func redStyle() lipgloss.Style       { return theme.Red().S() }      // red
func greenStyle() lipgloss.Style     { return theme.Green().S() }    // green
func yellowStyle() lipgloss.Style    { return theme.Yellow().S() }   // yellow
