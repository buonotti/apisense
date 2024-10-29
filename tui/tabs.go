package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	activeTabStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Background(lipgloss.Color("#5e81ac")).Padding(0, 1)
	inactiveTabStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#4c566a")).Padding(0, 1)
	tabRowStyle      = lipgloss.NewStyle().BorderBottom(true).BorderForeground(lipgloss.Color("#5e81ac"))
)

type tabModel struct {
	tabs        []string // List of tab names
	activeTab   int      // Index of the currently active tab
	tabContents []string // Content for each tab
	keymap      keymap
	reports     reportModel
}

func TabModule() tabModel {
	return tabModel{
		tabs: []string{
			"Reports", "Definitions", "Daemon", "TUI",
		},
		activeTab: 0,
		tabContents: []string{
			"Reports go here.",
			"Definitions go here",
			"Daemon settings go here.",
			"TUI settings go here.",
		},
		keymap:  DefaultKeyMap,
		reports: ReportModel(),
	}
}

func (t tabModel) Init() tea.Cmd {
	return nil
}

func (t tabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, t.keymap.left):
			if t.activeTab > 0 {
				t.activeTab--
			}
			updateActiveView(t)
		case key.Matches(msg, t.keymap.right):
			if t.activeTab < len(t.tabs)-1 {
				t.activeTab++
			}
			updateActiveView(t)
		case key.Matches(msg, t.keymap.quit):
			return t, tea.Quit
		}
	}
	var reportsCmd tea.Cmd
	updatedReportsModel, reportsCmd := t.reports.Update(msg)
	t.reports = updatedReportsModel.(reportModel)
	return t, reportsCmd
}

func (t tabModel) View() string {
	t.tabContents[0] = t.reports.View()
	var tabRow string
	for i, tab := range t.tabs {
		if i == t.activeTab {
			tabRow += activeTabStyle.Render(tab)
		} else {
			tabRow += inactiveTabStyle.Render(tab)
		}
		if i < len(t.tabs)-1 {
			tabRow += " "
		}
	}

	content := fmt.Sprintf("\n%s\n\n%s", tabRowStyle.Render(tabRow), t.tabContents[t.activeTab])

	return content
}

func updateActiveView(t tabModel) {
	publicActiveTab = t.activeTab
}
