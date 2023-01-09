package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

type keymap struct {
	help   key.Binding
	quit   key.Binding
	up     key.Binding
	down   key.Binding
	left   key.Binding
	right  key.Binding
	choose key.Binding
	format key.Binding
	back   key.Binding
}

func (k keymap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.help,
		k.quit,
	}
}

func (k keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.help, k.choose, k.up, k.format},
		{k.quit, k.back, k.down, k.format},
	}
}

var DefaultKeyMap = keymap{
	help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp(styleHotkey.Render("?"), styleHelp.Render("help")),
	),
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp(styleHotkey.Render("q"), styleHelp.Render("quit")),
	),
	up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp(styleHotkey.Render("↑|k"), styleHelp.Render("up")),
	),
	down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp(styleHotkey.Render("↓|j"), styleHelp.Render("down")),
	),
	left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp(styleHotkey.Render("←|h"), styleHelp.Render("left")),
	),
	right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp(styleHotkey.Render("→|l"), styleHelp.Render("right")),
	),
	choose: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp(styleHotkey.Render("↲"), styleHelp.Render("choose")),
	),
	back: key.NewBinding(
		key.WithKeys("b"),
		key.WithHelp(styleHotkey.Render("b"), styleHelp.Render("back")),
	),
	format: key.NewBinding(
		key.WithKeys(""),
		key.WithHelp(styleHotkey.Render(""), styleHelp.Render("")),
	),
}
