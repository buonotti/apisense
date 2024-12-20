package tui

import (
	"fmt"
	"strings"

	"github.com/buonotti/apisense/v2/validation/pipeline"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type validatorOutputModel struct {
	keymap   keymap
	table    table.Model
	selected string
}

func ValidatorOutputModel() tea.Model {
	t := table.New(
		table.WithColumns(getValidatorOutputColumns()),
		table.WithRows(nil),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#F38BA8")).
		Background(lipgloss.Color("#1e1e2e")).
		Bold(false)
	t.SetStyles(s)

	return validatorOutputModel{
		keymap:   DefaultKeyMap,
		table:    t,
		selected: "",
	}
}

func (v validatorOutputModel) Init() tea.Cmd {
	return nil
}

func (v validatorOutputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if choiceReportModel != "resultModel" {
		if updateValidatorOutputRows {
			t := table.New(
				table.WithColumns(getValidatorOutputColumns()),
				table.WithRows(validatorOutputRows),
				table.WithFocused(true),
				table.WithHeight(7),
			)
			s := table.DefaultStyles()
			s.Selected = s.Selected.
				Foreground(lipgloss.Color("#F38BA8")).
				Background(lipgloss.Color("#1e1e2e")).
				Bold(false)
			t.SetStyles(s)
			v.table = t
			updateValidatorOutputRows = false
		}
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, v.keymap.back):
				if choiceReportModel == "validatorOutputModel" {
					choiceReportModel = "resultModel"
				}
			case key.Matches(msg, v.keymap.quit):
				return v, tea.Quit
			}
			v.table, cmd = v.table.Update(msg)
			return v, cmd
		}
	}
	return v, nil
}

func (v validatorOutputModel) View() string {
	return lipgloss.NewStyle().Render(v.table.View())
}

func getValidatorOutputRows(validatorOutputs []pipeline.ValidatorResult) []table.Row {
	rows := make([]table.Row, 0)
	for i, output := range validatorOutputs {
		s := strings.Split(output.Message, ": ")
		if len(s) > 1 {
			q := strings.Join(s[1:], "")
			rows = append(rows, table.Row{fmt.Sprintf("%v", i), output.Name, q, string(output.Status)})
		} else {
			rows = append(rows, table.Row{fmt.Sprintf("%v", i), output.Name, "", string(output.Status)})
		}
	}
	return rows
}

func getValidatorOutputColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Validator", Width: 10},
		{Title: "Error", Width: 50},
		{Title: "Status", Width: 10},
	}
}
