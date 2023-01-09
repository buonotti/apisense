package daemon

import (
	"os"
	"strconv"

	"github.com/buonotti/odh-data-monitor/errors"
)

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
		err = proc.Signal(SIGINT)
		if err != nil {
			return errors.CannotSendSignalError.Wrap(err, "Cannot send signal SIGINT to process "+strconv.Itoa(pid))
		}
	}
	return nil
}
