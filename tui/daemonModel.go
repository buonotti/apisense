package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"

	"github.com/buonotti/odh-data-monitor/daemon"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	running bool
)

type daemonModel struct {
	status  string
	pid     int
	eStatus errMsg
	ePid    errMsg
	keymap  keymap
}

func DaemonModel() tea.Model {

	p, pe := daemon.Pid()
	s, se := daemon.Status()
	running = s == daemon.UP

	return daemonModel{
		status:  string(s),
		pid:     p,
		eStatus: se,
		ePid:    pe,
		keymap:  DefaultKeyMap,
	}
}

func (d daemonModel) Init() tea.Cmd {
	return nil
}

func (d daemonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	p, pe := daemon.Pid()
	s, se := daemon.Status()
	running = s == daemon.UP

	d.pid = p
	d.ePid = pe
	d.status = string(s)
	d.eStatus = se

	return d, nil
}

func (d daemonModel) View() string {
	sPid := ""
	sStatus := ""
	if d.ePid == nil {
		sPid = "pid:    " + styleInfo.Render(fmt.Sprintf("%d", d.pid))
	} else {
		sPid = "pid:    " + stylePrimary.Render("unknown")
	}
	if d.eStatus == nil {
		if d.status == string(daemon.UP) {
			sStatus = "status: " + styleSuccess.Render(d.status)
		} else {
			sStatus = "status: " + stylePrimary.Render(d.status)
		}

	} else {
		sStatus = "status: " + stylePrimary.Render("unknown")
	}

	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		PaddingTop(2).
		PaddingLeft(4).
		PaddingRight(4).
		Render(sPid+"\n\n"+sStatus+"\n\n") + "\n"
}
