package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

const useHighPerformanceRenderer = false

type configModel struct {
	keymap   keymap
	viewport viewport.Model
	content  string
	ready    bool
}

func ConfigModel() tea.Model {

	return configModel{
		keymap:  DefaultKeyMap,
		content: "TEST CONTENT",
	}
}

func (c configModel) Init() tea.Cmd {
	return nil
}

func (c configModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, c.keymap.quit) {
			return c, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(c.headerView())
		footerHeight := lipgloss.Height(c.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !c.ready {

			//Wait for window dimensions
			c.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			c.viewport.YPosition = headerHeight
			c.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			c.viewport.SetContent(c.content)
			c.ready = true

			// for high performance rendering
			c.viewport.YPosition = headerHeight + 1
		} else {
			c.viewport.Width = msg.Width
			c.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(c.viewport))
		}
	}
	c.viewport, cmd = c.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}

func (c configModel) View() string {
	return "ConfigModel works"
}

func (c configModel) headerView() string {
	title := titleStyle.Render("Config")
	line := strings.Repeat("─", max(0, c.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (c configModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", c.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, c.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
