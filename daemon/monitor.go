package daemon

import (
	"os"
	"strconv"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/files"
	lf "github.com/nightlyone/lockfile"
)

// State represents the possible states of the daemon
type State string

const (
	UpStatus   State = "up"
	DownStatus State = "down"
)

// Status returns the status of the daemon. If the status file cannot be read, it returns DOWN with an *errors.CannotReadStatusError.
func Status() (State, error) {
	statusString, err := os.ReadFile(files.DaemonStatusFile())
	if err != nil {
		return DownStatus, errors.CannotReadFileError.Wrap(err, "cannot read status file")
	}

	return State(statusString), nil
}

// Pid returns the pid of the daemon. If the pid file cannot be read, it returns 0 with an *errors.CannotReadPidError.
// If the daemon is not running the returned pid is -1.
func Pid() (int, error) {
	pidString, err := os.ReadFile(files.DaemonPidFile())
	if err != nil {
		return 0, errors.CannotReadFileError.Wrap(err, "cannot read pid file")
	}

	pid, err := strconv.Atoi(string(pidString))
	if err != nil {
		return 0, errors.CannotReadFileError.Wrap(err, "cannot convert pid file to int")
	}

	return pid, nil
}

// writeStatus is a helper function to write a daemon status to file.
// If an error occurs the error will be of type *errors.CannotWriteStatusFileError.
func writeStatus(state State) error {
	err := os.WriteFile(files.DaemonStatusFile(), []byte(state), 0o644)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot write status file")
	}

	return nil
}

// writePid is a helper function to write a daemon pid to file.
// If an error occurs the error will be of type *errors.CannotWritePidFileError.
func writePid(pid int) error {
	err := os.WriteFile(files.DaemonPidFile(), []byte(strconv.Itoa(pid)), 0o644)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "cannot write pid file")
	}

	return nil
}

// lockfile returns a lockfile.Lockfile on the lockfile to prevent multiple daemon instances from running at ths same time
func lockfile() (lf.Lockfile, error) {
	return lf.New(files.DaemonLockFile())
}
