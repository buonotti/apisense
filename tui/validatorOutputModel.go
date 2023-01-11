package tui

import (
	"fmt"
	"github.com/buonotti/odh-data-monitor/validation"
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
		table.WithColumns(getReportColumns()),
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
	return v, nil
}

func (v validatorOutputModel) View() string {
	if v.selected == "" {
		return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(v.table.View() + "\n")
	}
	return v.selected
}

func getValidatorOutputRows(validatorOutputs []validation.ValidatorOutput) []table.Row {
	rows := make([]table.Row, 0)
	for i, output := range validatorOutputs {
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), output.Validator, output.Error, output.Status})
	}
	return rows
}

func getValidatorOutputColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Validator", Width: 7},
		{Title: "Error", Width: 7},
		{Title: "Status", Width: 7},
	}
}
