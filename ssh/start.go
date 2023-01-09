package ssh

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/tui"
)

const (
	host = "10.216.220.251"
	port = 23234
)

func Start() error {
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		return err // TODO custom err
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, daemon.SIGINT, daemon.SIGTERM)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			errors.HandleError(err) // TODO custom err
		}
	}()
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		return err // TODO custom err
	}
	return nil
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		return nil, nil
	}
	return tui.TuiModule(), []tea.ProgramOption{tea.WithAltScreen()}
}
