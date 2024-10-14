package tui

import (
	"fmt"
	"github.com/buonotti/apisense/log"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type option struct {
	option, description string
}

type optionDelegate struct{}

func (d optionDelegate) Height() int                               { return 1 }
func (d optionDelegate) Spacing() int                              { return 0 }
func (d optionDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d optionDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(option)
	if !ok {
		return
	}

	var (
		str string
		fn  func(strings ...string) string
	)

	if terminalHeight < 25 {
		str = fmt.Sprintf("[   ] %s", styleBold.Render(i.option))

		fn = styleBase.Render
		if index == m.Index() {
			fn = func(s ...string) string {
				// Add cursor by modifying the format string
				return fmt.Sprintf("[ %s ] %s", stylePrimary.Render("x"), styleBold.Render(i.option))
			}
		}
	} else {
		str = fmt.Sprintf("╭───╮\n│   │ %s \n╰───╯", styleBold.Render(i.option))

		fn = styleBase.Render
		if index == m.Index() {
			fn = func(s ...string) string {
				// Add cursor by modifying the format string
				return fmt.Sprintf("╭───╮\n│ %s │ %s \n╰───╯", stylePrimary.Render("x"), styleBold.Render(i.option))
			}
		}
	}

	_, err := fmt.Fprint(w, fn(str))
	if err != nil {
		log.TuiLogger().Fatal(err)
	}
}

func (o option) Title() string       { return o.option }
func (o option) Description() string { return o.description }
func (o option) FilterValue() string { return o.option }

// Main menu
var optionsDaemonMenu = []list.Item{
	option{option: "Start daemon", description: "Starts a new daemon"},
	option{option: "Stop daemon", description: "Stops the current daemon"},
}

var listDaemonButton = list.New(optionsDaemonMenu, optionDelegate{}, defaultWidth, defaultListHeight)
