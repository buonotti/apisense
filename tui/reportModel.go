package tui

import (
	"encoding/json"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var (
	columns = []table.Column{
		{Title: "Id", Width: 3},
		{Title: "Report", Width: 15},
	}
)

type reportModel struct {
	keymap   keymap
	table    table.Model
	selected string
}

func ReportModel() tea.Model {

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(getReports()),
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

func getReports() []table.Row {
	files, err := os.ReadDir(validation.ReportLocation())
	errors.HandleError(err)

	//reports := make([]table.Row, 0)
	//reports = append(reports, table.Row{fmt.Sprintf("%v", i), file.Name()})
	for _, file := range files {
		if !file.IsDir() {
			content, err := os.ReadFile(file.Name())
			errors.HandleError(err)
			var report validation.Report
			err = json.Unmarshal(content, &report)
			errors.HandleError(err)

		}
	}
	return nil
}
