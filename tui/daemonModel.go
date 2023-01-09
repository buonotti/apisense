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

func daemonModule() daemonModel {

	s, se := daemon.Status()
	p, pe := daemon.Pid()

	return daemonModel{
		status:  string(s),
		pid:     p,
		eStatus: se,
		ePid:    pe,
	}
}

func (s daemonModel) Init() tea.Cmd {
	return nil
}

func (s daemonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s daemonModel) View() string {
	sPid := ""
	sStatus := ""
	if s.ePid == nil {
		sPid = "pid:    " + stylePrimary.Render(fmt.Sprintf("%d", s.pid))
	} else {
		sPid = "pid:    " + stylePrimary.Render("unknown")
	}
	if s.eStatus == nil {
		sStatus = "status: " + stylePrimary.Render(s.status)
	} else {
		sStatus = "status: " + stylePrimary.Render("unknown")
	}

	return sPid + "\n\n" + sStatus
}
