package theme

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

type Color lipgloss.Color

func (c Color) S() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(c))
}

func (c Color) Terminal() lipgloss.TerminalColor {
	return lipgloss.Color(c)
}

func hasTrueColorSupport() bool {
	return lipgloss.ColorProfile() == termenv.TrueColor
}

func convAnsiColor(tc termenv.ANSIColor) Color {
	c := strconv.Itoa(int(tc))
	return Color(c)
}

func Base() Color {
	if hasTrueColorSupport() {
		return Color(t().Base())
	}
	return convAnsiColor(termenv.ANSIBlack)
}

func Text() Color {
	if hasTrueColorSupport() {
		return Color(t().Text())
	}
	return convAnsiColor(termenv.ANSIWhite)
}

func Overlay0() Color {
	if hasTrueColorSupport() {
		return Color(t().Overlay0())
	}
	return convAnsiColor(termenv.ANSIWhite)
}

func Red() Color {
	if hasTrueColorSupport() {
		return Color(t().Red())
	}
	return convAnsiColor(termenv.ANSIRed)
}

func Blue() Color {
	if hasTrueColorSupport() {
		return Color(t().Blue())
	}
	return convAnsiColor(termenv.ANSIBlue)
}

func Green() Color {
	if hasTrueColorSupport() {
		return Color(t().Green())
	}
	return convAnsiColor(termenv.ANSIGreen)
}

func Yellow() Color {
	if hasTrueColorSupport() {
		return Color(t().Yellow())
	}
	return convAnsiColor(termenv.ANSIYellow)
}
