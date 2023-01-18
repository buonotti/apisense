package daemon

import (
	"os"
	"strconv"

	"golang.org/x/sys/unix"

	"github.com/buonotti/apisense/errors"
)

// Stop stops the daemon. If there is no daemon running it return an *errors.DaemonNotRunningError.
func Stop() error {
	pid, err := Pid()
	if err != nil {
		return err
	}
	if pid == -1 {
		return errors.DaemonNotRunningError.New("Daemon is not running")
	} else {
		proc, err := os.FindProcess(pid)
		if err != nil {
			return errors.CannotGetProcessInfoError.Wrap(err, "Cannot get process info for pid "+strconv.Itoa(pid))
		}
		err = proc.Signal(unix.SIGINT)
		if err != nil {
			return errors.CannotSendSignalError.Wrap(err, "Cannot send signal SIGINT to process "+strconv.Itoa(pid))
		}
	}
	return nil
}
