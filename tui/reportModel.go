package tui

import (
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
)

var (
	validatedEndpointRows []table.Row
	selectedReport        validation.Report
)

type reportModel struct {
	keymap                  keymap
	table                   table.Model
	validationEndpointModel tea.Model
}

func ReportModel() tea.Model {

	r, err := validation.Reports()
	errors.HandleError(err)
	reports = r

	t := table.New(
		table.WithColumns(getReportColumns()),
		table.WithRows(getReportRows(reports)),
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
		keymap:                  DefaultKeyMap,
		table:                   t,
		validationEndpointModel: ValidationEndpointModel(),
	}
}

func (r reportModel) Init() tea.Cmd {
	return nil
}

func (r reportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmdModel tea.Cmd
	if choiceReportModel != "" {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, r.keymap.back):
				if choiceReportModel == "reportModel" {
					choiceReportModel = ""
					choiceMainMenu = ""
				}
			case key.Matches(msg, r.keymap.quit):
				return r, tea.Quit
			case key.Matches(msg, r.keymap.choose):
				i, err := strconv.Atoi(r.table.SelectedRow()[0])
				errors.HandleError(err)
				rep, err := getSelectedReport(reports, i)
				errors.HandleError(err)
				selectedReport = rep
				validatedEndpointRows = getValidatedEndpointRows(selectedReport)
				if choiceReportModel != "reportModel" {
					r.validationEndpointModel, cmdModel = r.validationEndpointModel.Update(msg)
					r.table, cmd = r.table.Update(msg)
					return r, tea.Batch(cmd, cmdModel)
				}
				choiceReportModel = "validatedEndpointModel"
				r.table, cmd = r.table.Update(msg)
				return r, cmd
			}
		}
		r.validationEndpointModel, cmdModel = r.validationEndpointModel.Update(msg)
		r.table, cmd = r.table.Update(msg)
		return r, tea.Batch(cmd, cmdModel)
	}
	return r, nil

}

func (r reportModel) View() string {
	if choiceReportModel != "reportModel" {
		return styleContentCenter.Copy().MarginLeft(1).MarginRight(1).BorderStyle(lipgloss.RoundedBorder()).Render(r.validationEndpointModel.View() + "\n")
	}
	return styleContentCenter.Copy().MarginLeft(1).MarginRight(1).BorderStyle(lipgloss.RoundedBorder()).Render(r.table.View() + "\n")
}

func getReportRows(reports []validation.Report) []table.Row {
	rows := make([]table.Row, 0)
	for i, report := range reports {
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), report.Id, fmt.Sprintf("%v", time.Time(report.Time).Format("2006-01-02 15:04:05"))})
	}
	return rows
}

func getReportColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Id", Width: 10},
		{Title: "Timestamp", Width: 62},
	}
}

func getSelectedReport(reports []validation.Report, index int) (validation.Report, error) {
	if index > len(reports) || index < 0 {
		return validation.Report{}, errors.ModelError.New("Index out of range")
	}
	return reports[index], nil
}
