package tui

import (
	"time"

	"github.com/76creates/stickers"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type model struct {
	spinner   spinner.Model
	stopwatch stopwatch.Model
	help      help.Model
	keymap    KeyMap
	flexbox   *stickers.FlexBox
	quitting  bool
	err       error
}

func loadingModule() model {
	s := spinner.New()
	s.Spinner = spinner.Line
	s.Style = stylePrimary
	return model{
		spinner:   s,
		keymap:    DefaultKeyMap,
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
		flexbox:   stickers.NewFlexBox(0, 0).SetStyle(styleContent),
	}
}

func (m model) Init() tea.Cmd {
	m.flexbox.AddRows([]*stickers.FlexBoxRow{
		m.flexbox.NewRow().AddCells(
			[]*stickers.FlexBoxCell{
				stickers.NewFlexBoxCell(1, 1),
				stickers.NewFlexBoxCell(1, 1),
			},
		),
	})
	return tea.Batch(m.spinner.Tick, m.stopwatch.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keymap.help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		default:
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.flexbox.SetWidth(msg.Width)
		m.flexbox.SetHeight(msg.Height)
		return m, nil
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var spinnerCmd tea.Cmd
		m.spinner, spinnerCmd = m.spinner.Update(msg)
		var stopwatchCmd tea.Cmd
		m.stopwatch, stopwatchCmd = m.stopwatch.Update(msg)

		return m, tea.Batch(spinnerCmd, stopwatchCmd)
	}
}

func (m model) View() string {

	// if m.err != nil {
	// 	if err, ok := m.err.(*errorx.Error); ok {
	// 		return err.Error()
	// 	}
	// 	return errors.UnknownError.Wrap(m.err, "Unknown error").Error()
	// }

	m.flexbox.Row(0).Cell(0).SetContent(m.spinner.View() + "Loading..")
	m.flexbox.Row(0).Cell(1).SetContent(m.stopwatch.View())

	return m.flexbox.Render()

	// helpView := m.help.View(DefaultKeyMap)

}
