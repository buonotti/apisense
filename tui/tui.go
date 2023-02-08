package tui

import (
	"os/exec"
	"time"

	"github.com/common-nighthawk/go-figure"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

var (
	fileUpdate         = false
	choiceMainMenu     string
	choiceDaemonButton string
	choiceReportModel  string
	choiceConfigModel  string
	reports            []validation.Report
)

type errMsg error

type Model struct {
	help             help.Model
	keymap           keymap
	flexbox          *stickers.FlexBox
	elapsedTrigger   stopwatch.Model
	quitting         bool
	err              error
	listMainMenu     list.Model
	listDaemonButton list.Model
	daemonModel      tea.Model
	reportModel      tea.Model
	configModel      tea.Model
	daemonCmd        *exec.Cmd
}

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
		keymap:           DefaultKeyMap,
		help:             help.New(),
		flexbox:          stickers.NewFlexBox(0, 0).SetStyle(styleContentCenter.Copy().BorderStyle(lipgloss.RoundedBorder())),
		listMainMenu:     listMainMenu,
		listDaemonButton: listDaemonButton,
		daemonModel:      DaemonModel(),
		reportModel:      ReportModel(),
		elapsedTrigger:   stopwatch.NewWithInterval(time.Second),
		daemonCmd:        nil,
		configModel:      ConfigModel(),
	}
}

func (m Model) Init() tea.Cmd {
	m.daemonModel.Init()
	m.reportModel.Init()
	m.configModel.Init()
	watcher := NewFileWatcher()
	err := watcher.AddFile(daemon.PidFile)
	if err != nil {
		errors.HandleError(errors.WatcherError.Wrap(err, "Failed to add file to Watcher"))
	}

	go func() {
		err := watcher.Start()
		if err != nil {
			errors.HandleError(errors.WatcherError.Wrap(err, "Failed to start Watcher"))
		}
	}()

	go func() {
		for {
			fileUpdate = <-watcher.Events
		}
	}()

	m.flexbox.AddRows([]*stickers.FlexBoxRow{
		m.flexbox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 3).SetStyle(styleContentCenter.Copy().MarginLeft(1).MarginRight(1).MarginTop(3)),
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//Cmds that force rendering of given component
	var cmdMainMenu tea.Cmd
	var cmdDaemonButton tea.Cmd
	var cmdReportModel tea.Cmd
	var cmdElapsedTrigger tea.Cmd
	var cmdConfigModel tea.Cmd
	m.listDaemonButton, cmdDaemonButton = m.listDaemonButton.Update(msg)
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
			j, okDaemonButton := m.listDaemonButton.SelectedItem().(option)
			if okMainMenu && choiceMainMenu == "" {
				choiceMainMenu = i.title
			} else {
				if okDaemonButton && choiceDaemonButton == "" && choiceMainMenu == "Daemon" {
					choiceDaemonButton = j.option
					switch choiceDaemonButton {
					case "Start daemon":
						if !running {
							daemonCmd, err := daemon.Start(true, viper.GetBool("daemon.validate-on-startup"))
							m.daemonCmd = daemonCmd
							errors.HandleError(err)
						}
					case "Stop daemon":
						if running {
							err := daemon.Stop()
							errors.HandleError(err)
							if m.daemonCmd != nil {
								err = m.daemonCmd.Wait()
								errors.HandleError(err)
								m.daemonCmd = nil
							}
						}
					}
				}
			}
		case key.Matches(msg, m.keymap.up):
			return m, tea.Batch(cmdMainMenu, cmdDaemonButton)
		case key.Matches(msg, m.keymap.down):
			return m, tea.Batch(cmdMainMenu, cmdDaemonButton)
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
			r, err := validation.Reports()
			errors.HandleError(err)
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

func (m Model) View() string {

	if m.err != nil {
		errors.HandleError(errors.UnknownError.Wrap(m.err, "Unknown error"))
	}

	//Render Title
	title := figure.NewFigure("API SENSE", "", true)
	m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(title.String()))

	//Act based one main menu changes
	switch choiceMainMenu {
	case "Daemon":
		choiceDaemonButton = ""
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.daemonModel.View() + docStyle.Render(m.listDaemonButton.View())))
	case "Report":
		//Render report option
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.reportModel.View()))
	case "Config":
		//Act based one config menu changes
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.configModel.View()))
	default:
		//Render main menu
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(10))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.listMainMenu.View()))
	}

	//Render help menu
	m.flexbox.Row(2).Cell(0).SetContent(m.help.View(m.keymap))
	return m.flexbox.Render()

}
