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
	"golang.org/x/sys/unix"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/tui"
	"github.com/buonotti/apisense/validation"
)

// host returns the host value from the config
func host() string {
	return viper.GetString("ssh.host")
}

// port returns the port value from the config
func port() int {
	return viper.GetInt("ssh.port")
}

// Start starts the ssh server that listens on the predefined host and port.
func Start() error {
	fsHandler := scp.NewFileSystemHandler(validation.ReportLocation())
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host(), port())),
		wish.WithHostKeyPath(os.Getenv("HOME")+"/.ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			log.WishMiddleware(),
			activeterm.Middleware(),
			scp.Middleware(fsHandler, nil),
		),
	)
	if err != nil {
		err = errors.CannotCreateSSHServerError.Wrap(err, "cannot create SSH server")
		return err
	}

	log.SSHLogger.Infof("starting SSH server on %s:%d", host(), port())

	done := make(chan os.Signal, 1)
	signal.Notify(done, unix.SIGINT, unix.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.SSHLogger.Warn(err.Error())
		}
	}()

	log.SSHLogger.Infof("SSH server started")

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.SSHLogger.Infof("stopping SSH server")

	if err := s.Shutdown(ctx); err != nil {
		err = errors.CannotStopSSHServerError.Wrap(err, "cannot stop ssh server")
	}
	return nil
}

// teaHandler returns a tea.Model and a []tea.ProgramOption to pass to the bubbletea.Middleware function
// The tea.Model is the default tui.Model
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		return nil, nil
	}
	return tui.TuiModule(), []tea.ProgramOption{tea.WithAltScreen()}
}
