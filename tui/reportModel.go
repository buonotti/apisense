package tui

import (
	"fmt"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type reportModel struct {
	keymap   keymap
	table    table.Model
	selected string
}

func ReportModel() tea.Model {

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

	return reportModel{
		keymap:   DefaultKeyMap,
		table:    t,
		selected: "",
	}
}

func (r reportModel) Init() tea.Cmd {
	return nil
}

func (r reportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	r.selected = ""
	var cmd tea.Cmd
	if choiceMainMenu == "Report" {
		r.table.Focus()
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				if r.table.Focused() {
					r.table.Blur()
				} else {
					r.table.Focus()
				}
			case "q", "ctrl+c":
				return r, tea.Quit
			case "enter":
				r.selected = r.table.SelectedRow()[1]
			}
		}
	} else {
		r.table.Blur()
	}
	r.table, cmd = r.table.Update(msg)
	return r, cmd
}

func (r reportModel) View() string {
	if r.selected == "" {
		return lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render(r.table.View() + "\n")
	}
	return r.selected
}

func getReportRows(reports []validation.Report) []table.Row {
	rows := make([]table.Row, 0)
	for i, report := range reports {
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), report.Id, fmt.Sprintf("%v", report.Time)})
	}
	return nil
}

func getReportColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Id", Width: 3},
		{Title: "Timestamp", Width: 7},
	}
}
