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
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/scp"
	"github.com/spf13/viper"

	"github.com/buonotti/odh-data-monitor/daemon"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/log"
	"github.com/buonotti/odh-data-monitor/tui"
	"github.com/buonotti/odh-data-monitor/validation"
)

func host() string {
	return viper.GetString("ssh.host")
}

func port() int {
	return viper.GetInt("ssh.port")
}

func Start(startDaemon bool) error {
	if startDaemon {
		err := daemon.Start(true)
		if err != nil {
			return err
		}
	} else {
		log.SSHLogger.Warnf("Daemon not started by user")
	}
	fsHandler := scp.NewFileSystemHandler(validation.ReportLocation())
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host(), port())),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			log.WishMiddleware(),
			activeterm.Middleware(),
			scp.Middleware(fsHandler, fsHandler),
		),
	)
	if err != nil {
		err = errors.CannotCreateSSHServer.Wrap(err, "Cannot create SSH server")
		return err
	}
	log.SSHLogger.Infof("Starting SSH server on %s:%d", host(), port())
	done := make(chan os.Signal, 1)
	signal.Notify(done, daemon.SIGINT, daemon.SIGTERM)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.SSHLogger.Warn(err.Error())
		}
	}()
	log.SSHLogger.Infof("SSH server started")
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.SSHLogger.Infof("SSH server stopped")
	if startDaemon {
		err := daemon.Stop()
		if err != nil {
			return err
		}
	}
	log.SSHLogger.Infof("Stopping SSH server")
	if err := s.Shutdown(ctx); err != nil {
		err = errors.CannotStopSSHServer.Wrap(err, "Cannot stop SSH server")
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
