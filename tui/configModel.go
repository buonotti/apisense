package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/buonotti/apisense/config"
	"github.com/buonotti/apisense/errors"
)

const useHighPerformanceRenderer = false

type configModel struct {
	keymap   keymap
	textarea textarea.Model
	content  string
	err      error
}

func ConfigModel() tea.Model {

	s, err := os.ReadFile(config.FullPath)
	errors.CheckErr(err)

	ti := textarea.New()
	ti.SetHeight(80)
	ti.SetWidth(80)
	ti.SetValue(string(s))
	ti.Focus()

	return configModel{
		keymap:   DefaultKeyMap,
		content:  string(s),
		textarea: ti,
		err:      nil,
	}
}

func (c configModel) Init() tea.Cmd {
	return nil
}

func (c configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, c.keymap.back):
			if c.textarea.Focused() {
				c.textarea.Blur()
			}
		case key.Matches(msg, c.keymap.quit):
			return c, tea.Quit
		case key.Matches(msg, c.keymap.choose):
			if !c.textarea.Focused() {
				cmd = c.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		c.err = msg
		errors.CheckErr(c.err)
	}

	c.textarea, cmd = c.textarea.Update(msg)
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}

func (c configModel) View() string {
	return fmt.Sprintf(
		"Edit Config\n\n%s\n\n%s",
		c.textarea.View(),
		"(INSERT INFO HERE)",
	) + "\n\n"
}
