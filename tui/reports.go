package tui

import (
	"fmt"
	"strconv"
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

type reportModel struct {
	views   []*tableView
	keymap  keymap
	flexBox *flexbox.FlexBox
	reports []pipeline.Report
}

func ReportModel() reportModel {
	r, err := pipeline.Reports()
	if err != nil {
		log.TuiLogger().Fatal(err)
	}
	reports := r

	reportTableView := createReportTableView(reports)

	f := flexbox.New(0, 0).SetStyle(styleContentCenter)
	mainRow := f.NewRow().AddCells(
		flexbox.NewCell(0, 12),
		flexbox.NewCell(10, 12),
		flexbox.NewCell(0, 12),
	)
	f.AddRows([]*flexbox.Row{mainRow})

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
		r.flexBox.SetWidth(msg.Width)
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
					if len(r.views) == 1 { // In the reports view
						selectedIdx, _ := strconv.Atoi(r.activeView().table.GetCursorValue())
						r.pushView(createValidatedEndpointTableView(r.reports[selectedIdx]))
					} else if len(r.views) == 2 { // In the validated endpoints view
						//selectedValidatedEndpoint, _ := strconv.Atoi(r.activeView().table.GetCursorValue())
						// TODO: Display or handle the selected validated endpoint
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
	row := r.flexBox.GetRow(0)
	cell := row.GetCell(1)
	r.activeView().table.SetWidth(cell.GetWidth())
	r.activeView().table.SetHeight(cell.GetHeight())
	cell.SetContent(r.activeView().table.Render())

	ret := r.flexBox.Render()

	return ret
}

// Function to push new view onto stack
func (r *reportModel) pushView(view *tableView) {
	r.views = append(r.views, view)
}

// Function to pop top view from stack
func (r *reportModel) popView() {
	if len(r.views) > 1 {
		r.views = r.views[:len(r.views)-1]
	}
}

// Function to get active view from stack
func (r *reportModel) activeView() *tableView {
	return r.views[len(r.views)-1]
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
		rows = append(rows, []string{"", "No reports found!", ""})
		return rows, false
	}
	return rows, true
}

func getValidatedEndpointRows(validatedEndpoint pipeline.Report) ([][]string, bool) {
	rows := make([][]string, 0)
	for i, endpoint := range validatedEndpoint.Endpoints {
		row := []string{
			fmt.Sprintf("%d", i),
			endpoint.EndpointName,
		}
		rows = append(rows, row)
	}
	if len(rows) == 0 {
		rows = append(rows, []string{"", "No endpoints found"})
		return rows, false
	}

	return rows, true
}

func getResultRows(results []pipeline.TestCaseResult) ([][]string, bool) {
	rows := make([][]string, 0)
	for i, result := range results {
		result.ValidatorResults
	}
}

func createReportTableView(reports []pipeline.Report) *tableView {
	headers := []string{"", "", ""}
	rows, allowSelection := getReportRows(reports)
	ratio := []int{2, 5, 20}
	minSize := []int{4, 10, 15}
	t := table.NewTableSingleType[string](0, 0, headers)
	t.SetRatio(ratio).SetMinWidth(minSize)
	t.AddRows(rows).SetStylePassing(true)
	t.SetStyles(map[table.TableStyleKey]lipgloss.Style{
		table.TableRowsCursorStyleKey:     styleActive,
		table.TableRowsSubsequentStyleKey: styleBase,
		table.TableRowsStyleKey:           styleBase,
		table.TableHeaderStyleKey:         styleBase,
		table.TableFooterStyleKey:         styleFooter,
		table.TableCellCursorStyleKey:     styleActive,
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
		table.TableRowsCursorStyleKey:     styleActive,
		table.TableRowsSubsequentStyleKey: styleBase,
		table.TableRowsStyleKey:           styleBase,
		table.TableHeaderStyleKey:         styleBase,
		table.TableFooterStyleKey:         styleFooter,
		table.TableCellCursorStyleKey:     styleActive,
	})
	return &tableView{table: t, allowSelect: allowSelection}
}
