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
		return Color(Current().Base())
	}
	return convAnsiColor(termenv.ANSIBlack)
}

func BaseS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With()
	}
	return lipgloss.NewStyle().Background(lipgloss.Color(convAnsiColor(termenv.ANSIBlack)))
}

func Text() Color {
	if hasTrueColorSupport() {
		return Color(Current().Text())
	}
	return convAnsiColor(termenv.ANSIWhite)
}

func TextS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With().Foreground(Current().Text())
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(convAnsiColor(termenv.ANSIWhite)))
}

func TextDark() Color {
	if hasTrueColorSupport() {
		return Color(Current().TextDark())
	}
	return convAnsiColor(termenv.ANSIWhite)
}

func TextDarkS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With().Foreground(Current().TextDark())
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(convAnsiColor(termenv.ANSIWhite)))
}

func Red() Color {
	if hasTrueColorSupport() {
		return Color(Current().Error())
	}
	return convAnsiColor(termenv.ANSIRed)
}

func RedS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With().Foreground(Current().Error())
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(convAnsiColor(termenv.ANSIRed)))
}

func Blue() Color {
	if hasTrueColorSupport() {
		return Color(Current().Info())
	}
	return convAnsiColor(termenv.ANSIBlue)
}

func BlueS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With().Foreground(Current().Info())
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(convAnsiColor(termenv.ANSIBlue)))
}

func Green() Color {
	if hasTrueColorSupport() {
		return Color(Current().Ok())
	}
	return convAnsiColor(termenv.ANSIGreen)
}

func GreenS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With().Foreground(Current().Ok())
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(convAnsiColor(termenv.ANSIGreen)))
}

func Yellow() Color {
	if hasTrueColorSupport() {
		return Color(Current().Warning())
	}
	return convAnsiColor(termenv.ANSIYellow)
}

func YellowS() lipgloss.Style {
	if hasTrueColorSupport() {
		return Style().With().Foreground(Current().Warning())
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color(convAnsiColor(termenv.ANSIYellow)))
}
