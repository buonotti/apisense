package theme

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/util"
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

func Location() string {
	path := "~/.config/apisense/themes"
	util.ExpandHome(&path)
	return path
}

type External struct {
	IsDark        bool   `toml:"dark"`
	Background    string `toml:"background"`
	Foreground    string `toml:"foreground"`
	ForegroundAlt string `toml:"foreground_alt"`
	Red           string `toml:"red"`
	Blue          string `toml:"blue"`
	Green         string `toml:"green"`
	Yellow        string `toml:"yellow"`
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

func Load() ([]Theme, error) {
	// TODO
	return nil, nil
}
