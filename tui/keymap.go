package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	help  key.Binding
	quit  key.Binding
	test  key.Binding
	test1 key.Binding
	test2 key.Binding
	test3 key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.help,
		k.quit,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.help, k.quit},
		{k.test, k.test1, k.test2, k.test3},
	}
}

var DefaultKeyMap = KeyMap{
	help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp(styleHotkey.Render("?"), styleHelp.Render("help")),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp(styleHotkey.Render("q"), styleHelp.Render("quit")),
	),
	test: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp(styleHotkey.Render("↑|k"), styleHelp.Render("up")),
	),
	test1: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp(styleHotkey.Render("↓|j"), styleHelp.Render("down")),
	),
	test2: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp(styleHotkey.Render("←|h"), styleHelp.Render("left")),
	),
	test3: key.NewBinding(
		key.WithKeys("t"),
		key.WithHelp(styleHotkey.Render("→|l"), styleHelp.Render("right")),
	),
}
