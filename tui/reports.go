package tui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/buonotti/apisense/log"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/buonotti/apisense/validation/pipeline"
)

type tableView struct {
	table       *table.TableSingleType[string]
	allowSelect bool
}

type selection struct {
	index       int
	description string
	item        interface{}
}

type reportModel struct {
	views            []*tableView
	keymap           keymap
	flexBox          *flexbox.FlexBox
	reports          []pipeline.Report
	selectionHistory []selection
}

func ReportModel() reportModel {
	r, err := pipeline.Reports()
	if err != nil {
		log.TuiLogger().Fatal(err)
	}
	reports := r

	reportTableView := createReportTableView(reports)

	f := flexbox.New(0, 0).SetStyle(styleContentCenter)
	pathRow := f.NewRow().AddCells(
		flexbox.NewCell(1, 1).SetStyle(styleContentCenter),
	)
	mainRow := f.NewRow().AddCells(
		flexbox.NewCell(10, 10).SetStyle(styleContentCenter),
	)
	f.AddRows([]*flexbox.Row{pathRow, mainRow})

	return reportModel{
		views:   []*tableView{reportTableView},
		keymap:  DefaultKeyMap,
		flexBox: f,
		reports: reports,
	}
}

func (r reportModel) Init() tea.Cmd {
	return nil
}

func (r reportModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.flexBox.SetWidth(msg.Width / 4 * 2)
		r.flexBox.SetHeight(msg.Height / 5 * 4)
		r.activeView().table.SetWidth(r.flexBox.GetWidth())
		r.activeView().table.SetHeight(r.flexBox.GetHeight())
	case tea.KeyMsg:
		if publicActiveTab == 0 {
			switch {
			case key.Matches(msg, r.keymap.up):
				r.activeView().table.CursorUp()
				return r, nil
			case key.Matches(msg, r.keymap.down):
				r.activeView().table.CursorDown()
				return r, nil
			case key.Matches(msg, r.keymap.choose):
				if r.activeView().allowSelect {
					if len(r.views) == 1 {
						// In the reports view
						selectedIdx, _ := strconv.Atoi(r.activeView().table.GetCursorValue())
						selectedReport := r.reports[selectedIdx]
						description := selectedReport.Id
						r.pushView(createValidatedEndpointTableView(selectedReport), selection{
							index:       selectedIdx,
							description: description,
							item:        selectedReport,
						})
						return r, nil
					}
					if len(r.views) == 2 {
						// In the validated endpoints view
						lastSelection := r.selectionHistory[len(r.selectionHistory)-1]
						selectedReport := lastSelection.item.(pipeline.Report)
						selectedIdx, _ := strconv.Atoi(r.activeView().table.GetCursorValue())
						selectedEndpoint := selectedReport.Endpoints[selectedIdx]
						description := selectedEndpoint.EndpointName
						r.pushView(createResultTableView(selectedEndpoint), selection{
							index:       selectedIdx,
							description: description,
							item:        selectedEndpoint,
						})
						return r, nil
					}
					if len(r.views) == 3 {
						// In the result view
						lastSelection := r.selectionHistory[len(r.selectionHistory)-1]
						selectedValidatedEndpoint := lastSelection.item.(pipeline.ValidatedEndpoint)
						selectedIdx, _ := strconv.Atoi(r.activeView().table.GetCursorValue())
						selectedResult := selectedValidatedEndpoint.TestCaseResults[selectedIdx]
						description := selectedResult.Name
						//TODO: FIX CRASH
						r.pushView(createValidatorResultTableView(selectedResult), selection{
							index:       selectedIdx,
							description: description,
							item:        selectedResult,
						})
						return r, nil
					}
				}
				return r, nil
			case key.Matches(msg, r.keymap.back):
				r.popView()
				return r, nil
			case key.Matches(msg, r.keymap.quit):
				return r, tea.Quit
			}
		}

	}
	return r, nil
}

func (r reportModel) View() string {
	r.flexBox.ForceRecalculate()
	rowPath := r.flexBox.GetRow(0)
	cellPath := rowPath.GetCell(0)
	cellPath.SetContent(r.GetSelectionPath())

	row := r.flexBox.GetRow(1)
	cell := row.GetCell(0)
	r.activeView().table.SetWidth(cell.GetWidth())
	r.activeView().table.SetHeight(cell.GetHeight())
	cell.SetContent(r.activeView().table.Render())
	r.flexBox.ForceRecalculate()
	return r.flexBox.Render()
}

// Function to push new view onto stack
func (r *reportModel) pushView(view *tableView, selection selection) {
	r.views = append(r.views, view)
	r.selectionHistory = append(r.selectionHistory, selection)
}

// Function to pop top view from stack
func (r *reportModel) popView() {
	if len(r.views) > 1 && len(r.selectionHistory) > 0 {
		r.views = r.views[:len(r.views)-1]
		r.selectionHistory = r.selectionHistory[:len(r.selectionHistory)-1]
	}
}

// Function to get active view from stack
func (r *reportModel) activeView() *tableView {
	return r.views[len(r.views)-1]
}

// Get selectionHistory as path
func (r *reportModel) GetSelectionPath() string {
	var pathParts []string
	for _, selection := range r.selectionHistory {
		pathParts = append(pathParts, selection.description)
	}
	if len(pathParts) == 0 {
		return styleAccent.Render("Current") + "@Report" + styleAccent.Render("~") + ">"
	}
	return styleAccent.Render("Current") + "@Report" + styleAccent.Render("~/") + styleAccent.Render(strings.Join(pathParts, "/")) + ">" // Prepend with '/' for filepath style
}

func getReportRows(reports []pipeline.Report) ([][]string, bool) {
	rows := make([][]string, 0)
	for i, report := range reports {
		rows = append(rows, []string{
			fmt.Sprintf("%d", i),
			report.Id,
			time.Time(report.Time).Format("2006-01-02 15:04:05"),
		})
	}
	if len(rows) == 0 {
		rows = append(rows, []string{"-1", "", "No reports found"})
		return rows, false
	}
	return rows, true
}

func getValidatedEndpointRows(validatedEndpoint pipeline.Report) ([][]string, bool) {
	rows := make([][]string, 0)
	for i, endpoint := range validatedEndpoint.Endpoints {
		rows = append(rows, []string{
			fmt.Sprintf("%d", i),
			endpoint.EndpointName,
		})
	}
	if len(rows) == 0 {
		rows = append(rows, []string{"-1", "No endpoints found"})
		return rows, false
	}

	return rows, true
}

func getResultRows(validatedEndpoint pipeline.ValidatedEndpoint) ([][]string, bool) {
	rows := make([][]string, 0)
	for i, result := range validatedEndpoint.TestCaseResults {
		rows = append(rows, []string{
			fmt.Sprintf("%d", i),
			result.Name,
		})
	}
	if len(rows) == 0 {
		rows = append(rows, []string{"-1", "No results found"})
		return rows, false
	}
	return rows, true
}

func getValidatorResultRows(result pipeline.TestCaseResult) ([][]string, bool) {
	rows := make([][]string, 0)
	//WHY TF IS THERE A INDEX OUT OF RANGE IN A FUCKING FOREACH LOOP I HATE MY FUCKING LIFE
	if result.ValidatorResults == nil {
		log.TuiLogger().Fatal("FUCK ME IT NIL")
	}
	for i, validatorResult := range result.ValidatorResults {
		rows = append(rows, []string{
			fmt.Sprintf("%d", i),
			validatorResult.Name,
			validatorResult.Message,
			string(validatorResult.Status),
		})
	}
	if len(rows) == 0 {
		rows = append(rows, []string{"-1", "No validatorResults found"})
		return rows, false
	}
	return rows, true
}

func createReportTableView(reports []pipeline.Report) *tableView {
	headers := []string{"", "", ""}
	rows, allowSelection := getReportRows(reports)
	ratio := []int{5, 12, 40}
	minSize := []int{4, 8, 20}
	t := table.NewTableSingleType[string](0, 0, headers)
	t.SetRatio(ratio).SetMinWidth(minSize)
	t.AddRows(rows).SetStylePassing(true)
	t.SetStyles(map[table.TableStyleKey]lipgloss.Style{
		table.TableRowsCursorStyleKey:     styleActive.Align(lipgloss.Right),
		table.TableRowsSubsequentStyleKey: styleContentRight,
		table.TableRowsStyleKey:           styleContentRight,
		table.TableHeaderStyleKey:         styleContentRight,
		table.TableFooterStyleKey:         styleFooter,
		table.TableCellCursorStyleKey:     styleActive.Align(lipgloss.Right),
	})
	return &tableView{table: t, allowSelect: allowSelection}
}

func createValidatedEndpointTableView(endpoint pipeline.Report) *tableView {
	headers := []string{"", ""}
	rows, allowSelection := getValidatedEndpointRows(endpoint)
	ratio := []int{2, 25}
	minSize := []int{4, 25}
	t := table.NewTableSingleType[string](0, 0, headers)
	t.SetRatio(ratio).SetMinWidth(minSize)
	t.AddRows(rows).SetStylePassing(true)
	t.SetStyles(map[table.TableStyleKey]lipgloss.Style{
		table.TableRowsCursorStyleKey:     styleActive.Align(lipgloss.Right),
		table.TableRowsSubsequentStyleKey: styleContentRight,
		table.TableRowsStyleKey:           styleContentRight,
		table.TableHeaderStyleKey:         styleContentRight,
		table.TableFooterStyleKey:         styleFooter,
		table.TableCellCursorStyleKey:     styleActive.Align(lipgloss.Right),
	})
	return &tableView{table: t, allowSelect: allowSelection}
}

func createResultTableView(validatedEndpoint pipeline.ValidatedEndpoint) *tableView {
	headers := []string{"", ""}
	rows, allowSelection := getResultRows(validatedEndpoint)
	ratio := []int{2, 25}
	minSize := []int{4, 25}
	t := table.NewTableSingleType[string](0, 0, headers)
	t.SetRatio(ratio).SetMinWidth(minSize)
	t.AddRows(rows).SetStylePassing(true)
	t.SetStyles(map[table.TableStyleKey]lipgloss.Style{
		table.TableRowsCursorStyleKey:     styleActive.Align(lipgloss.Right),
		table.TableRowsSubsequentStyleKey: styleContentRight,
		table.TableRowsStyleKey:           styleContentRight,
		table.TableHeaderStyleKey:         styleContentRight,
		table.TableFooterStyleKey:         styleFooter,
		table.TableCellCursorStyleKey:     styleActive.Align(lipgloss.Right),
	})
	return &tableView{table: t, allowSelect: allowSelection}
}

func createValidatorResultTableView(result pipeline.TestCaseResult) *tableView {
	headers := []string{"", ""}
	rows, allowSelection := getValidatorResultRows(result)
	ratio := []int{2, 25}
	minSize := []int{4, 25}
	t := table.NewTableSingleType[string](0, 0, headers)
	t.SetRatio(ratio).SetMinWidth(minSize)
	t.AddRows(rows).SetStylePassing(true)
	t.SetStyles(map[table.TableStyleKey]lipgloss.Style{
		table.TableRowsCursorStyleKey:     styleActive.Align(lipgloss.Right),
		table.TableRowsSubsequentStyleKey: styleContentRight,
		table.TableRowsStyleKey:           styleContentRight,
		table.TableHeaderStyleKey:         styleContentRight,
		table.TableFooterStyleKey:         styleFooter,
		table.TableCellCursorStyleKey:     styleActive.Align(lipgloss.Right),
	})
	return &tableView{table: t, allowSelect: allowSelection}
}
