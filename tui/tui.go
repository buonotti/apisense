package tui

import (
	"time"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem"
	"github.com/buonotti/apisense/filesystem/locations/files"
	"github.com/buonotti/apisense/validation/pipeline"
)

var (
	fileUpdate        bool                  = false // fileUpdate True whenever the pid in the daemon.pid file changes
	directoryUpdate   bool                  = false // directoryUpdate True whenever a new report is generated
	choiceMainMenu    string                        // choiceMainMenu Saves the user choice made on the listMainMenu
	choiceReportModel string                        // choiceReportModel Saves the user choice made in the reportModel view
	choiceConfigModel string                        // choiceConfigModel Saves the user choice made in the configModel view
	reports           []pipeline.Report             // reports Existing reports
	terminalHeight    = getTerminalHeight()         // terminalHeight Terminal height, updates whenever a WindowSizeMsg is triggered
)

type errMsg error

type Model struct {
	help           help.Model
	keymap         keymap
	flexbox        *stickers.FlexBox
	elapsedTrigger stopwatch.Model
	quitting       bool
	err            error
	listMainMenu   list.Model
	daemonModel    tea.Model
	reportModel    tea.Model
	configModel    tea.Model
}

// TuiModule Creates the initial parent model and sets the rendering interval based on the tui.refresh config value
func TuiModule() Model {
	listMainMenu.SetShowFilter(false)
	listMainMenu.SetShowTitle(false)
	listMainMenu.SetShowPagination(false)
	listMainMenu.SetShowHelp(false)
	listMainMenu.SetShowStatusBar(false)

	listConfigMenu.SetShowFilter(false)
	listConfigMenu.SetShowTitle(false)
	listConfigMenu.SetShowPagination(false)
	listConfigMenu.SetShowHelp(false)
	listConfigMenu.SetShowStatusBar(false)

	listDaemonButton.SetShowFilter(false)
	listDaemonButton.SetShowTitle(false)
	listDaemonButton.SetShowPagination(false)
	listDaemonButton.SetShowHelp(false)
	listDaemonButton.SetShowStatusBar(false)

	return Model{
		keymap:         DefaultKeyMap,
		help:           help.New(),
		flexbox:        stickers.NewFlexBox(0, 0).SetStyle(styleContentCenter.Copy()),
		listMainMenu:   listMainMenu,
		elapsedTrigger: stopwatch.NewWithInterval(time.Duration(viper.GetInt("tui.refresh")) * time.Millisecond),
		configModel:    ConfigModel(),
		reportModel:    ReportModel(),
		daemonModel:    DaemonModel(),
	}
}

// Init Initializes the fileWatcher which monitors the daemons state and starts the renderingTrigger
func (m Model) Init() tea.Cmd {
	m.reportModel.Init()
	m.configModel.Init()
	fileWatcher := filesystem.NewFileWatcher()
	err := fileWatcher.AddFile(files.DaemonPidFile())
	errors.CheckErr(err)

	directoryWatcher := filesystem.NewDirectoryWatcher()
	err = directoryWatcher.SetDirectory(pipeline.ReportLocation())
	errors.CheckErr(err)
	go func() {
		err := fileWatcher.Start()
		errors.CheckErr(err)
	}()

	go func() {
		err := directoryWatcher.Start()
		errors.CheckErr(err)
	}()

	go func() {
		for {
			fileUpdate = <-fileWatcher.Events
		}
	}()

	go func() {
		for {
			<-directoryWatcher.Events
			directoryUpdate = true
			r, err := pipeline.Reports()
			errors.CheckErr(err)
			reports = r
		}
	}()

	m.flexbox.AddRows([]*stickers.FlexBoxRow{
		m.flexbox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 4).SetStyle(styleContentCenter.Copy().MarginLeft(1).MarginRight(1).MarginTop(3)),
			},
		),
		m.flexbox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 12).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(10)),
			},
		),
		m.flexbox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 2).SetStyle(styleContentCenter.Copy().MarginLeft(3)),
			},
		),
	})
	return m.elapsedTrigger.Start()
}

// Update Reacts to given I/O and updates the ui values accordingly. Also triggers rendering for given sup models based on menu choices
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Cmds that force rendering of given component
	var cmdMainMenu tea.Cmd
	var cmdReportModel tea.Cmd
	var cmdElapsedTrigger tea.Cmd
	var cmdConfigModel tea.Cmd
	m.reportModel, cmdReportModel = m.reportModel.Update(msg)
	m.listMainMenu, cmdMainMenu = m.listMainMenu.Update(msg)
	m.elapsedTrigger, cmdElapsedTrigger = m.elapsedTrigger.Update(msg)
	m.configModel, cmdConfigModel = m.configModel.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keymap.choose):
			i, okMainMenu := m.listMainMenu.SelectedItem().(item)
			if okMainMenu && choiceMainMenu == "" {
				choiceMainMenu = i.title
			}
		case key.Matches(msg, m.keymap.up):
			return m, cmdMainMenu
		case key.Matches(msg, m.keymap.down):
			return m, cmdMainMenu
		case key.Matches(msg, m.keymap.back):
			if choiceMainMenu != "Report" && choiceMainMenu != "Config" {
				if choiceMainMenu != "" {
					choiceMainMenu = ""
					m.listMainMenu.ResetSelected()
				}
			}
		default:
			return m, nil
		}

	case tea.WindowSizeMsg:
		terminalHeight = getTerminalHeight()
		m.flexbox.SetWidth(msg.Width)
		m.flexbox.SetHeight(msg.Height)
		m.help.Width = msg.Width
		h, v := docStyle.GetFrameSize()
		m.listMainMenu.SetSize(msg.Width-h, msg.Height-v)
		return m, cmdConfigModel

	case errMsg:
		m.err = msg
		return m, nil
	}

	if fileUpdate {
		fileUpdate = false
		var cmdDaemonModal tea.Cmd
		m.daemonModel, cmdDaemonModal = m.daemonModel.Update(msg)
		return m, tea.Batch(cmdDaemonModal, cmdElapsedTrigger)
	}
	if choiceMainMenu == "Report" {
		if choiceReportModel == "" {
			r, err := pipeline.Reports()
			errors.CheckErr(err)
			reports = r
			choiceReportModel = "reportModel"
		}
		return m, tea.Batch(cmdElapsedTrigger, cmdReportModel)
	}
	if choiceMainMenu == "Config" {
		if choiceConfigModel == "" {
			choiceConfigModel = "configModel"
		}
		return m, tea.Batch(cmdElapsedTrigger, cmdConfigModel)
	}
	return m, cmdElapsedTrigger

}

// View Redraws the UI after every Update
func (m Model) View() string {

	// Handle ErrorMsg received during Update()
	if m.err != nil {
		errors.CheckErr(errors.UnknownError.Wrap(m.err, "Unknown error"))
	}

	// Render Title
	title := figure.NewFigure("API SENSE", "", true)
	// Resize flexbox based on the current terminalHeight. This value change whenever a WindowSizeMsg is caught
	if terminalHeight > 25 {
		m.flexbox.Row(0).UpdateCellWithIndex(0, stickers.NewFlexBoxCell(1, 4).SetStyle(styleContentCenter.Copy().MarginLeft(1).MarginRight(1).MarginTop(3)))
		m.flexbox.Row(2).UpdateCellWithIndex(0, stickers.NewFlexBoxCell(1, 2).SetStyle(styleContentCenter.Copy().MarginLeft(3)))
		m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(title.String()))
	} else {
		m.flexbox.Row(0).UpdateCellWithIndex(0, stickers.NewFlexBoxCell(1, 1).SetStyle(styleContentCenter.Copy().MarginLeft(1).MarginRight(1).MarginTop(3)))
		m.flexbox.Row(2).UpdateCellWithIndex(0, stickers.NewFlexBoxCell(1, 4).SetStyle(styleContentCenter.Copy().MarginLeft(3)))
		m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(""))
	}

	// Act based one main menu changes
	switch choiceMainMenu {
	case "Daemon":
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(terminalHeight / 8))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.daemonModel.View()))
	case "Report":
		// Render report option
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(terminalHeight / 8))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.reportModel.View()))
	case "Config":
		// Act based one config menu changes
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(terminalHeight / 8))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.configModel.View()))
	default:
		// Render main menu
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(terminalHeight / 8))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.listMainMenu.View()))
	}

	// Render help menu
	m.flexbox.Row(2).Cell(0).SetContent(m.help.View(m.keymap))
	return m.flexbox.Render()

}
