package tui

import (
	"fmt"

	"github.com/buonotti/odh-data-monitor/daemon"
	tea "github.com/charmbracelet/bubbletea"
)

type daemonModel struct {
	status  string
	pid     int
	eStatus errMsg
	ePid    errMsg
}

func daemonModule() tea.Model {

	p, pe := daemon.Pid()
	s, se := daemon.Status()

	return daemonModel{
		status:  string(s),
		pid:     p,
		eStatus: se,
		ePid:    pe,
	}
}

func (d daemonModel) Init() tea.Cmd {
	return nil
}

func (d daemonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	p, pe := daemon.Pid()
	s, se := daemon.Status()

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

	return sPid + "\n\n" + sStatus + "\n\n"
}
