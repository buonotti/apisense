package theme

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
)

type Theme interface {
	Name() string
	Base() lipgloss.Color
	Text() lipgloss.Color
	TextDark() lipgloss.Color
	Error() lipgloss.Color
	Info() lipgloss.Color
	Ok() lipgloss.Color
	Warning() lipgloss.Color
}

type LipglossStyle struct {
	Theme Theme
}

func (t LipglossStyle) With() lipgloss.Style {
	return lipgloss.NewStyle().Background(t.Theme.Base())
}

func Location() string {
	path := "~/.config/apisense/themes"
	util.ExpandHome(&path)
	return path
}

func Current() Theme {
	return themeMap[viper.GetString("tui.theme")]
}

func Style() *LipglossStyle {
	return &LipglossStyle{Theme: Current()}
}

var themeMap map[string]Theme

func Setup() error {
	themes, err := load()
	if err != nil {
		return err
	}
	themeMap = themes
	return nil
}

func load() (map[string]Theme, error) {
	themes := make(map[string]Theme, 0)
	files, err := os.ReadDir(Location())
	if err != nil {
		return nil, errors.CannotReadDirectoryError.New("cannot read directory: " + Location())
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		data, err := os.ReadFile(file.Name())
		if err != nil {
			return nil, errors.CannotReadFileError.New("cannot read file: " + file.Name())
		}
		var theme External
		err = toml.Unmarshal(data, &theme)
		if err != nil {
			return nil, errors.CannotUnmarshalThemeError.New("cannot unmarshal file: " + file.Name())
		}
		themes[theme.Name()] = theme
	}
	return themes, nil
}
