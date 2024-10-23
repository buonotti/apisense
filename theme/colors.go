package theme

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Ansi2Color converts a termenv.Ansi2Color to a lipgloss.Color
func Ansi2Color(tc termenv.ANSIColor) lipgloss.Color {
	c := strconv.Itoa(int(tc))
	return lipgloss.Color(c)
}
