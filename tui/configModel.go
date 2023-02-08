package tui

import (
	"github.com/buonotti/apisense/errors"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strconv"
)

var (
	selectedConfig         = 0
	updateSelectConfigRows = true
)

type configModel struct {
	keymap          keymap
	err             error
	table           table.Model
	editConfigModel tea.Model
}

func ConfigModel() tea.Model {

	t := table.New(
		table.WithColumns(getConfigColumns()),
		table.WithRows(getConfigRows()),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#F38BA8")).
		Background(lipgloss.Color("#1e1e2e")).
		Bold(false)
	t.SetStyles(s)

	return configModel{
		keymap:          DefaultKeyMap,
		err:             nil,
		table:           t,
		editConfigModel: SelectConfigModel(),
	}
}

func getConfigColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Config", Width: 30},
		{Title: "Description", Width: 42},
	}

}

func getConfigRows() []table.Row {

	rows := []table.Row{
		{"0", "env", "Edit env values"},
		{"1", "daemon", "Configure daemon"},
		{"2", "ssh", "Configure ssh"},
		{"3", "api", "Configure api"},
		{"4", "tui", "Configure tui"},
	}

	return rows
}

func (c configModel) Init() tea.Cmd {
	return nil
}

func (c configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmdModel tea.Cmd

	if choiceConfigModel != "" {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, c.keymap.back):
				if choiceConfigModel == "configModel" {
					choiceConfigModel = ""
					choiceMainMenu = ""
				}
			case key.Matches(msg, c.keymap.quit):
				return c, tea.Quit
			case key.Matches(msg, c.keymap.choose):
				i, err := strconv.Atoi(c.table.SelectedRow()[0])
				errors.CheckErr(err)
				if choiceConfigModel != "configModel" {
					c.editConfigModel, cmdModel = c.editConfigModel.Update(msg)
					c.table, cmd = c.table.Update(msg)
					return c, tea.Batch(cmd, cmdModel)
				}
				selectedConfig = i
				updateSelectConfigRows = true
				choiceConfigModel = "selectConfigModel"
				c.table, cmd = c.table.Update(msg)
				return c, cmd
			}

		case errMsg:
			c.err = msg
			errors.CheckErr(c.err)
		}

		c.editConfigModel, cmdModel = c.editConfigModel.Update(msg)
		c.table, cmd = c.table.Update(msg)
		return c, tea.Batch(cmd, cmdModel)
	}
	return c, nil
}

func (c configModel) View() string {
	if choiceConfigModel != "configModel" {
		return styleContentCenter.Copy().MarginLeft(1).MarginRight(1).BorderStyle(lipgloss.RoundedBorder()).Render(c.editConfigModel.View() + "\n")
	}
	return styleContentCenter.Copy().MarginLeft(1).MarginRight(1).BorderStyle(lipgloss.RoundedBorder()).Render(c.table.View() + "\n")
}
