package tui

import (
	"github.com/76creates/stickers"
	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
	"time"
)

var (
	fileUpdate         = false
	choiceMainMenu     string
	choiceConfigMenu   string
	choiceDaemonButton string
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
	listConfigMenu   list.Model
	listDaemonButton list.Model
	daemonModel      tea.Model
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
		listConfigMenu:   listConfigMenu,
		listDaemonButton: listDaemonButton,
		daemonModel:      daemonModule(),
		elapsedTrigger:   stopwatch.NewWithInterval(time.Second),
	}
}

func (m Model) Init() tea.Cmd {
	m.daemonModel.Init()
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
				stickers.NewFlexBoxCell(1, 3).SetStyle(styleContentCenter.Copy().MarginLeft(1).MarginRight(1).BorderStyle(lipgloss.RoundedBorder())),
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
	var cmdConfigMenu tea.Cmd
	var cmdDaemonButton tea.Cmd
	var cmdElapsedTrigger tea.Cmd
	m.listDaemonButton, cmdDaemonButton = m.listDaemonButton.Update(msg)
	m.listMainMenu, cmdMainMenu = m.listMainMenu.Update(msg)
	m.listConfigMenu, cmdConfigMenu = m.listConfigMenu.Update(msg)
	m.elapsedTrigger, cmdElapsedTrigger = m.elapsedTrigger.Update(msg)

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
			j, okConfigMenu := m.listConfigMenu.SelectedItem().(item)
			k, okDaemonButton := m.listDaemonButton.SelectedItem().(option)
			if okMainMenu && choiceMainMenu == "" {
				choiceMainMenu = i.title
				m.listConfigMenu.ResetSelected()
			} else {
				if okConfigMenu && choiceConfigMenu == "" {
					choiceConfigMenu = j.title
				}
				if okDaemonButton && choiceDaemonButton == "" {
					choiceDaemonButton = k.option
				}
			}
		case key.Matches(msg, m.keymap.up):
			return m, tea.Batch(cmdMainMenu, cmdConfigMenu, cmdDaemonButton)
		case key.Matches(msg, m.keymap.down):
			return m, tea.Batch(cmdMainMenu, cmdConfigMenu, cmdDaemonButton)
		case key.Matches(msg, m.keymap.back):
			if choiceMainMenu != "" && choiceConfigMenu != "" {
				choiceConfigMenu = ""
				m.listConfigMenu.ResetSelected()
			} else if choiceMainMenu != "" {
				choiceMainMenu = ""
				m.listMainMenu.ResetSelected()
			}
		case msg.Type == tea.KeyF5:
			return m, nil
		default:
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.flexbox.SetWidth(msg.Width)
		m.flexbox.SetHeight(msg.Height)
		m.help.Width = msg.Width
		h, v := docStyle.GetFrameSize()
		m.listMainMenu.SetSize(msg.Width-h, msg.Height-v)
		m.listConfigMenu.SetSize(msg.Width-h, msg.Height-v)
		return m, nil

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
	return m, cmdElapsedTrigger

}

func (m Model) View() string {

	if m.err != nil {
		errors.HandleError(errors.UnknownError.Wrap(m.err, "Unknown error"))
	}

	//Render Title
	title := figure.NewFigure("ODM - TUI", "", true)
	m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(title.String()))
	//m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(m.elapsedTrigger.View()))

	//Act based one main menu changes
	switch choiceMainMenu {
	case "Daemon":
		//Render state option
		switch choiceDaemonButton {
		case "Start daemon":
			if !running {
				err := daemon.Start(true)
				errors.HandleError(err)
			}
		case "Stop daemon":
			if running {
				err := daemon.Stop()
				errors.HandleError(err)
			}
		}
		choiceDaemonButton = ""
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.daemonModel.View() + docStyle.Render(m.listDaemonButton.View())))
	case "Report":
		//Render report option
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(choiceMainMenu))
	case "Config":
		//Act based one config menu changes
		switch choiceConfigMenu {
		case "Daemon":
			//Render daemon option
			m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
			m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(choiceConfigMenu))
		case "TUI":
			//Render tui option
			m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
			m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(choiceConfigMenu))
		default:
			//Render config menu
			m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(7))
			m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.listConfigMenu.View()))
		}
	default:
		//Render main menu
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(10))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.listMainMenu.View()))
	}

	//Render help menu
	m.flexbox.Row(2).Cell(0).SetContent(m.help.View(m.keymap))
	return m.flexbox.Render()

}
