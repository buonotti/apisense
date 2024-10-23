package tui

import (
	"fmt"
	"github.com/buonotti/apisense/log"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation/pipeline"
)

var (
	selectedResult            pipeline.TestCaseResult
	validatorOutputRows       []table.Row
	updateValidatorOutputRows = false
)

type resultModel struct {
	keymap               keymap
	table                table.Model
	validatorOutputModel tea.Model
}

func ResultModel() tea.Model {
	t := table.New(
		table.WithColumns(getResultColumns()),
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
		keymap:               DefaultKeyMap,
		table:                t,
		validatorOutputModel: ValidatorOutputModel(),
	}
}

func (r resultModel) Init() tea.Cmd {
	return nil
}

func (r resultModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmdModel tea.Cmd
	if choiceReportModel != "validatedEndpointModel" {
		if updateResultRows {
			t := table.New(
				table.WithColumns(getResultColumns()),
				table.WithRows(resultRows),
				table.WithFocused(true),
				table.WithHeight(7),
			)
			s := table.DefaultStyles()
			s.Selected = s.Selected.
				Foreground(lipgloss.Color("#F38BA8")).
				Background(lipgloss.Color("#1e1e2e")).
				Bold(false)
			t.SetStyles(s)
			r.table = t
			updateResultRows = false
		}
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, r.keymap.back):
				if choiceReportModel == "resultModel" {
					choiceReportModel = "validatedEndpointModel"
				}
			case key.Matches(msg, r.keymap.quit):
				return r, tea.Quit
			case key.Matches(msg, r.keymap.choose):
				i, err := strconv.Atoi(r.table.SelectedRow()[0])
				if err != nil {
					log.TuiLogger().Fatal(err)
				}
				res, err := getSelectedResult(selectedValidatedEndpoint, i)
				if err != nil {
					log.TuiLogger().Fatal(err)
				}
				selectedResult = res
				validatorOutputRows = getValidatorOutputRows(selectedResult.ValidatorResults)
				if choiceReportModel != "resultModel" {
					r.validatorOutputModel, cmdModel = r.validatorOutputModel.Update(msg)
					r.table, cmd = r.table.Update(msg)
					return r, tea.Batch(cmd, cmdModel)
				}
				choiceReportModel = "validatorOutputModel"
				updateValidatorOutputRows = true
				r.table, cmd = r.table.Update(msg)
				return r, tea.Batch(cmd, cmdModel)
			}
		}
		r.validatorOutputModel, cmdModel = r.validatorOutputModel.Update(msg)
		r.table, cmd = r.table.Update(msg)
		return r, tea.Batch(cmd, cmdModel)
	}
	return r, nil
}

func (r resultModel) View() string {
	if choiceReportModel != "resultModel" {
		return lipgloss.NewStyle().Render(r.validatorOutputModel.View())
	}
	return lipgloss.NewStyle().Render(r.table.View())
}

func getResultRows(results []pipeline.TestCaseResult) []table.Row {
	rows := make([]table.Row, 0)
	queries := make([][]string, 0)
	queriesToRender := make([]string, 0)
	for _, result := range results {
		u, err := url.Parse(result.Url)
		if err != nil {
			log.TuiLogger().Fatal(err)
		}
		query := make([]string, 0)
		for value := range u.Query() {
			query = append(query, value+"="+u.Query().Get(value))
			sort.Strings(query)
		}
		queries = append(queries, query)
	}
	rQueries := util.Transpose(queries)
	for _, query := range rQueries {
		for i, result := range query {
			if i > 0 {
				if strings.Split(query[i-1], "=")[1] != strings.Split(result, "=")[1] {
					queriesToRender = append(queriesToRender, strings.Split(result, "=")[0])
					break
				}
			}
		}
	}

	for i, result := range queries {
		s := ""
		for _, query := range result {
			if util.Contains(queriesToRender, strings.Split(query, "=")[0]) {
				s += query + ", "
			}
		}
		rows = append(rows, table.Row{fmt.Sprintf("%v", i), s})
	}
	return rows
}

func getResultColumns() []table.Column {
	return []table.Column{
		{Title: "", Width: 3},
		{Title: "Url", Width: 73},
	}
}

func getSelectedResult(validatedEndpoint pipeline.ValidatedEndpoint, index int) (pipeline.TestCaseResult, error) {
	if index > len(validatedEndpoint.TestCaseResults) || index < 0 {
		return pipeline.TestCaseResult{}, errors.ModelError.New("Index out of range")
	}
	return validatedEndpoint.TestCaseResults[index], nil
}
