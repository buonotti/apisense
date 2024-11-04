package daemon

import (
	"os"
	"strconv"

	"github.com/buonotti/apisense/v2/errors"
)

// Stop stops the daemon. If there is no daemon running it return an *errors.DaemonNotRunningError.
func Stop() error {
	pid, err := Pid()
	if err != nil {
		return err
	}

	if pid == -1 {
		return errors.DaemonNotRunningError.New("daemon is not running")
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return errors.CannotGetProcessInfoError.Wrap(err, "cannot get process info for pid "+strconv.Itoa(pid))
	}

	err = proc.Signal(os.Kill)
	if err != nil {
		return errors.CannotSendSignalError.Wrap(err, "cannot send signal SIGINT to process "+strconv.Itoa(pid))
	}

	return nil
}
