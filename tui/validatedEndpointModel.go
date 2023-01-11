package tui

import (
	"fmt"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type validationEndpointModel struct {
	keymap   keymap
	table    table.Model
	selected string
}

func ValidationEndpointModel() tea.Model {

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

	return validationEndpointModel{
		keymap:   DefaultKeyMap,
		table:    t,
		selected: "",
	}
}

func (v validationEndpointModel) Init() tea.Cmd {
	return nil
}

func (v validationEndpointModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return v, nil
}

func (v validationEndpointModel) View() string {
	if v.selected == "" {
		return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(v.table.View() + "\n")
	}
	return v.selected
}

func getValidatedEndpointRows(validatedEndpoint validation.Report) []table.Row {
	rows := make([]table.Row, 0)
	for i, point := range validatedEndpoint.Results {
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), point.EndpointName})
	}
	return rows
}

func getValidatedEndpointColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "EndpointName", Width: 7},
	}
}
