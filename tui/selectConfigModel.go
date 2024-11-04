package tui

import (
	"fmt"
	"github.com/buonotti/apisense/v2/log"
	"sort"
	"strconv"
	strings2 "strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

var (
	updateEditConfigField = true
	allowConfigSelection  bool
	selectedField         = ""
	sortedStrings         = make([][]string, 5)
)

type selectConfigModel struct {
	keymap          keymap
	err             error
	table           table.Model
	editConfigModel tea.Model
}

func SelectConfigModel() tea.Model {
	t := table.New(
		table.WithColumns(getSelectConfigColumns()),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#F38BA8")).
		Background(lipgloss.Color("#1e1e2e")).
		Bold(false)
	t.SetStyles(s)

	return selectConfigModel{
		keymap:          DefaultKeyMap,
		err:             nil,
		table:           t,
		editConfigModel: EditConfigModel(),
	}
}

func (s selectConfigModel) Init() tea.Cmd {
	return nil
}

func (s selectConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmdModel tea.Cmd

	if choiceConfigModel != "configModel" {
		if updateSelectConfigRows {
			t := table.New(
				table.WithColumns(getSelectConfigColumns()),
				table.WithRows(getSelectConfigRows()),
				table.WithFocused(true),
				table.WithHeight(7),
			)
			sty := table.DefaultStyles()
			sty.Selected = sty.Selected.
				Foreground(lipgloss.Color("#F38BA8")).
				Background(lipgloss.Color("#1e1e2e")).
				Bold(false)
			t.SetStyles(sty)
			s.table = t
			updateSelectConfigRows = false
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, s.keymap.back):
				if choiceConfigModel == "selectConfigModel" {
					choiceConfigModel = "configModel"
				}
			case key.Matches(msg, s.keymap.quit):
				return s, tea.Quit
			case key.Matches(msg, s.keymap.choose):
				if allowConfigSelection {
					i, err := strconv.Atoi(s.table.SelectedRow()[0])
					if err != nil {
						log.TuiLogger().Fatal(err)
					}

					if choiceConfigModel == "selectConfigModel" {
						selectedField = getSelectedFieldName(i)
					}

					if choiceConfigModel != "selectConfigModel" {
						s.editConfigModel, cmdModel = s.editConfigModel.Update(msg)
						return s, tea.Batch(cmd, cmdModel)
					}
					updateEditConfigField = true
					choiceConfigModel = "editConfigModel"
					s.table, cmd = s.table.Update(msg)
				}
				return s, cmd
			}

		case errMsg:
			s.err = msg
			if s.err != nil {
				log.TuiLogger().Fatal(s.err)
			}
		}

		s.table, cmd = s.table.Update(msg)
		s.editConfigModel, cmdModel = s.editConfigModel.Update(msg)
		return s, cmd
	}
	return s, nil
}

func (s selectConfigModel) View() string {
	if choiceConfigModel != "selectConfigModel" {
		return s.editConfigModel.View() + "\n"
	}
	return s.table.View() + "\n"
}

func getSelectConfigColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Type", Width: 40},
		{Title: "Value", Width: 32},
	}
}

func getSelectedFieldName(i int) string {
	return sortedStrings[selectedConfig][i]
}

func getSelectConfigRows() []table.Row {
	rows := make([]table.Row, 0)
	strings := viper.AllKeys()
	sort.Strings(strings)

	sortedStrings = make([][]string, 5)

	for _, s := range strings {
		if strings2.HasPrefix(s, "daemon.") {
			sortedStrings[1] = append(sortedStrings[1], s)
		} else if strings2.HasPrefix(s, "ssh.") {
			sortedStrings[2] = append(sortedStrings[2], s)
		} else if strings2.HasPrefix(s, "api.") {
			sortedStrings[3] = append(sortedStrings[3], s)
		} else if strings2.HasPrefix(s, "tui.") {
			sortedStrings[4] = append(sortedStrings[4], s)
		} else {
			sortedStrings[0] = append(sortedStrings[0], s)
		}
	}

	for i, configKey := range sortedStrings[selectedConfig] {
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), configKey, viper.GetString(configKey)})
	}

	allowConfigSelection = true
	if len(rows) < 1 {
		rows = append(rows, table.Row{"", "No config fields found", ""})
		allowConfigSelection = false
	}

	return rows
}
