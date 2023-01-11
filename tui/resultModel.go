package tui

import (
	"fmt"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type resultModel struct {
	keymap   keymap
	table    table.Model
	selected string
}

func ResultModel() tea.Model {

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

	return resultModel{
		keymap:   DefaultKeyMap,
		table:    t,
		selected: "",
	}
}

func (r resultModel) Init() tea.Cmd {
	return nil
}

func (r resultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

func (r resultModel) View() string {
	if r.selected == "" {
		return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(r.table.View() + "\n")
	}
	return r.selected
}

func getResultRows(results []validation.Result) []table.Row {
	rows := make([]table.Row, 0)
	for i, result := range results {
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), result.Url})
	}
	return rows
}

func getResultColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Url", Width: 7},
	}
}
