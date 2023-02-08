package theme

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func Ansi2Color(tc termenv.ANSIColor) lipgloss.Color {
	c := strconv.Itoa(int(tc))
	return lipgloss.Color(c)
}
