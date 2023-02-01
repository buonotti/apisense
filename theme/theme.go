package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

type Theme interface {
	IsDark() bool
	Base() lipgloss.Color
	Text() lipgloss.Color
	Overlay0() lipgloss.Color
	Red() lipgloss.Color
	Blue() lipgloss.Color
	Green() lipgloss.Color
	Yellow() lipgloss.Color
}

var defaultTheme Theme = &mochaTheme{}

func t() Theme {
	theme := viper.GetString("tui.theme")
	switch theme {
	case "dark":
		return defaultTheme
	case "light":
		return &latteTheme{}
	case "latte":
		return &latteTheme{}
	case "frappe":
		return &frappeTheme{}
	case "macchiato":
		return &macchiatoTheme{}
	case "mocha":
		return &mochaTheme{}
	default:
		return defaultTheme
	}
}
