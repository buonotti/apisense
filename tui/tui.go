package tui

import (
	"strconv"

	"github.com/76creates/stickers/flexbox"
	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"

	"github.com/buonotti/apisense/log"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	daemonStatus, daemonStatusError = daemon.Status()
	daemonPid, daemonPidError       = daemon.Pid()
)

type watcherMsg struct{}

type tuiModel struct {
	help    help.Model
	err     error
	flexBox *flexbox.FlexBox
	keymap  keymap
	tabs    tabModel
}

func TuiModule() tuiModel {
	flex := flexbox.New(0, 0)
	rows := []*flexbox.Row{
		flex.NewRow().AddCells(
			flexbox.NewCell(3, 8).SetStyle(styleContentCenter),
			flexbox.NewCell(1, 8).SetStyle(styleContentCenter),
		),
		flex.NewRow().AddCells(
			flexbox.NewCell(3, 1).SetStyle(styleContentCenter),
			flexbox.NewCell(1, 1).SetStyle(styleContentCenter),
		),
	}
	flex.AddRows(rows)

	return tuiModel{
		help:    help.New(),
		flexBox: flex,
		keymap:  DefaultKeyMap,
		tabs:    TabModule(),
	}
}

// Init implements tea.Model.
func (t tuiModel) Init() tea.Cmd {
	return t.tabs.Init()
}

// Update implements tea.Model.
func (t tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var tabsCmd tea.Cmd
	updatedTabsModel, tabsCmd := t.tabs.Update(msg)
	t.tabs = updatedTabsModel.(tabModel)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.flexBox.SetWidth(msg.Width)
		t.flexBox.SetHeight(msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.keymap.quit):
			return t, tea.Quit
		case key.Matches(msg, t.keymap.help):
			t.help.ShowAll = !t.help.ShowAll
			return t, nil
		case key.Matches(msg, t.keymap.choose):
			return t, tabsCmd
		case key.Matches(msg, t.keymap.back):
			return t, tabsCmd
		case key.Matches(msg, t.keymap.up):
			return t, tabsCmd
		case key.Matches(msg, t.keymap.down):
			return t, tabsCmd
		case key.Matches(msg, t.keymap.left):
			return t, tabsCmd
		case key.Matches(msg, t.keymap.right):
			return t, tabsCmd
		default:
			return t, nil
		}
	case watcherMsg:
		daemonPid, daemonPidError = daemon.Pid()
		daemonStatus, daemonStatusError = daemon.Status()
		if daemonPidError != nil {
			log.TuiLogger().Fatal(daemonPidError)
		}
		if daemonStatusError != nil {
			log.TuiLogger().Fatal(daemonStatusError)
		}
	}
	return t, nil
}

// View implements tea.Model.
func (t tuiModel) View() string {
	t.flexBox.ForceRecalculate()
	if t.err != nil {
		log.TuiLogger().Fatal(errors.UnknownError.Wrap(t.err, "unknown error"))
	}

	//Help menu render
	rowHelp := t.flexBox.GetRow(1)
	if rowHelp == nil {
		log.TuiLogger().Fatal(errors.RenderError.Wrap(t.err, "could not find the table row"))
	}
	cellHelp := rowHelp.GetCell(0)
	if cellHelp == nil {
		log.TuiLogger().Fatal(errors.RenderError.Wrap(t.err, "could not find the table cell"))
	}

	t.help.Width = cellHelp.GetWidth()
	cellHelp.SetContent(t.help.View(t.keymap))

	//Daemon status render
	rowStatus := t.flexBox.GetRow(1)
	if rowStatus == nil {
		log.TuiLogger().Fatal(errors.RenderError.Wrap(t.err, "could not find the table row"))
	}
	cellStatus := rowStatus.GetCell(1)
	if cellStatus == nil {
		log.TuiLogger().Fatal(errors.RenderError.Wrap(t.err, "could not find the table cell"))
	}

	cellStatus.SetContent("Status: " + string(daemonStatus) + "\nPID: " + strconv.Itoa(daemonPid))

	//Menu render
	rowMenu := t.flexBox.GetRow(0)
	if rowMenu == nil {
		log.TuiLogger().Fatal(errors.RenderError.Wrap(t.err, "could not find the table row"))
	}
	cellMenu := rowMenu.GetCell(0)
	if cellMenu == nil {
		log.TuiLogger().Fatal(errors.RenderError.Wrap(t.err, "could not find the table cell"))
	}

	cellMenu.SetContent(t.tabs.View())
	return t.flexBox.Render()
}
