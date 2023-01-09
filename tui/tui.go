package tui

import (
	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/common-nighthawk/go-figure"
)

type errMsg error

type model struct {
	help             help.Model
	keymap           keymap
	flexbox          *stickers.FlexBox
	quitting         bool
	err              error
	listMainMenu     list.Model
	listConfigMenu   list.Model
	choiceMainMenu   string
	choiceConfigMenu string
}

func tuiModule() model {
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

	return model{
		keymap:           DefaultKeyMap,
		help:             help.New(),
		flexbox:          stickers.NewFlexBox(0, 0).SetStyle(styleContentCenter.Copy().BorderStyle(lipgloss.RoundedBorder())),
		listMainMenu:     listMainMenu,
		listConfigMenu:   listConfigMenu,
		choiceMainMenu:   "",
		choiceConfigMenu: "",
	}
}

func (m model) Init() tea.Cmd {
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
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	//Cmds that force rerendering of given menu
	var cmdMainMenu tea.Cmd
	var cmdConfigMenu tea.Cmd
	m.listMainMenu, cmdMainMenu = m.listMainMenu.Update(msg)
	m.listConfigMenu, cmdConfigMenu = m.listConfigMenu.Update(msg)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keymap.up):
			m.listMainMenu.CursorUp()
			m.listConfigMenu.CursorUp()
			return m, tea.Batch(cmdMainMenu, cmdConfigMenu)
		case key.Matches(msg, m.keymap.up):
			m.listMainMenu.CursorDown()
			m.listConfigMenu.CursorDown()
			return m, tea.Batch(cmdMainMenu, cmdConfigMenu)
		case key.Matches(msg, m.keymap.choose):
			i, okMainMenu := m.listMainMenu.SelectedItem().(item)
			j, okConfigMenu := m.listConfigMenu.SelectedItem().(item)
			if okMainMenu && m.choiceMainMenu == "" {
				m.choiceMainMenu = string(i.title)
				m.listConfigMenu.ResetSelected()
			} else if okConfigMenu && m.choiceConfigMenu == "" {
				m.choiceConfigMenu = string(j.title)
			}
			return m, tea.Batch(cmdMainMenu, cmdConfigMenu)
		case key.Matches(msg, m.keymap.back):
			if m.choiceMainMenu != "" && m.choiceConfigMenu != "" {
				m.choiceConfigMenu = ""
				m.listConfigMenu.ResetSelected()
			} else if m.choiceMainMenu != "" {
				m.choiceMainMenu = ""
				m.listMainMenu.ResetSelected()
			}
			return m, tea.Batch(cmdMainMenu, cmdConfigMenu)
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

	default:
		return m, tea.Batch(cmdMainMenu, cmdConfigMenu)
	}
}

func (m model) View() string {

	//Render Title
	title := figure.NewFigure("ODM - TUI", "", true)
	m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(title.String()))
	//For testing purposes
	//m.flexbox.Row(0).Cell(0).SetContent(stylePrimary.Render(m.choiceMainMenu + " | " + m.choiceConfigMenu))

	//Act based one main menu changes
	switch m.choiceMainMenu {
	case "Start":
		//Render start option
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.choiceMainMenu))
	case "State":
		//Render state option
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.choiceMainMenu))
	case "Report":
		//Render report option
		m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
		m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.choiceMainMenu))
	case "Config":
		//Act based one config menu changes
		switch m.choiceConfigMenu {
		case "Daemon":
			//Render daemon option
			m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
			m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.choiceConfigMenu))
		case "TUI":
			//Render tui option
			m.flexbox.Row(1).Cell(0).SetStyle(styleContentCenter.Copy().MarginTop(5).MarginLeft(5))
			m.flexbox.Row(1).Cell(0).SetContent(docStyle.Render(m.choiceConfigMenu))
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

//Config von Tui bearbeiten
//Config von Deamon bearbeiten
//Report auslesen
//Status von Deamon abfragen
//Deamon starten??
