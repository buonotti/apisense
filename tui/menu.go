package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/buonotti/apisense/errors"
)

type item struct {
	title, desc string
}

type itemDelegate struct{}

const (
	defaultWidth      int = 10
	defaultListHeight int = 10
)

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("  %s \n  %s \n", styleBold.Render(i.title), i.desc)

	fn := styleBase.Render
	if index == m.Index() {
		fn = func(s string) string {
			//Add cursor by modifying the format string
			return stylePrimary.Render(fmt.Sprintf("│ %s \n│ %s \n", styleBold.Render(i.title), i.desc))
		}
	}

	_, err := fmt.Fprint(w, fn(str))
	errors.CheckErr(err)
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Main menu
var itemsMainMenu = []list.Item{
	item{title: "Daemon", desc: "Control daemon activity"},
	item{title: "Report", desc: "Look into available reports"},
	item{title: "Config", desc: "Edit daemon / TUI config"},
}
var listMainMenu = list.New(itemsMainMenu, itemDelegate{}, defaultWidth, defaultListHeight)

// Config menu
var itemsConfigMenu = []list.Item{
	item{title: "Daemon", desc: "Edit daemon config"},
	item{title: "TUI", desc: "Edit TUI config"},
}
var listConfigMenu = list.New(itemsConfigMenu, itemDelegate{}, defaultWidth, defaultListHeight)
