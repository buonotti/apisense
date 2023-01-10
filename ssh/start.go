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

// host returns the host value from the config
func host() string {
	return viper.GetString("ssh.host")
}

// port returns the port value from the config
func port() int {
	return viper.GetInt("ssh.port")
}

// Start starts the ssh server that listens on the predefined host and port.
// It also starts the daemon if the startDaemon parameter is true.
func Start(startDaemon bool) error {
	// start the daemon. if not inform the user
	if startDaemon {
		err := daemon.Start(true)
		if err != nil {
			return err
		}
	} else {
		log.SSHLogger.Warnf("Daemon not started by user")
	}

	// create the filesystem handler for scp and create the ssh server
	fsHandler := scp.NewFileSystemHandler(validation.ReportLocation())
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host(), port())), // set the address of the server
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),         // set the path where the keys should be stored
		wish.WithMiddleware( // add middleware to the server
			bm.Middleware(teaHandler),      // add the bubbletea middleware to serve the tui over ssh
			log.WishMiddleware(),           // add the custom logging middleware
			activeterm.Middleware(),        // add the activeterm middleware to enforce an active terminal
			scp.Middleware(fsHandler, nil), // add the scp middleware to allow scp. The CopyFromClientHandler is set to nil to prevent a client to upload files to the server. This is because the root of scp is the reports/ directory, and we use scp only to manage the reports
		),
	)
	if err != nil {
		err = errors.CannotCreateSSHServerError.Wrap(err, "Cannot create SSH server")
		return err
	}

	log.SSHLogger.Infof("Starting SSH server on %s:%d", host(), port())

	// make the ssh server listen to signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, daemon.SIGINT, daemon.SIGTERM)

	// start the server in its own goroutine
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.SSHLogger.Warn(err.Error()) // ListenAndServe always returns an error so we don't panic
		}
	}()

	log.SSHLogger.Infof("SSH server started")

	// wait until the done channel receives a SIGINT or SIGTERM. When this happens, stop the server
	<-done

	// create a context with a timeout of 5 seconds to stop the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.SSHLogger.Infof("SSH server stopped")

	// stop the daemon first
	if startDaemon {
		err := daemon.Stop()
		if err != nil {
			return err
		}
	}
	log.SSHLogger.Infof("Stopping SSH server")

	// shutdown the server
	if err := s.Shutdown(ctx); err != nil {
		err = errors.CannotStopSSHServerError.Wrap(err, "Cannot stop SSH server")
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
