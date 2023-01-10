package daemon

import (
	"os"
	"path/filepath"
	"strconv"

	lf "github.com/nightlyone/lockfile"

	"github.com/buonotti/odh-data-monitor/errors"
)

// PidFile is the file where the pid of the daemon is stored
var PidFile = Directory() + "/daemon.pid"

// StatusFile is the file where the status of the daemon is stored
var StatusFile = Directory() + "/daemon.status"

// State represents the possible states of the daemon
type State string

const (
	UP   State = "up"
	DOWN State = "down"
)

// Status returns the status of the daemon. If the status file cannot be read, it returns DOWN with an *errors.CannotReadStatusError.
func Status() (State, error) {
	statusString, err := os.ReadFile(StatusFile)
	if err != nil {
		return DOWN, errors.CannotReadStatusError.Wrap(err, "Cannot read status file")
	}
	return State(statusString), nil
}

// Pid returns the pid of the daemon. If the pid file cannot be read, it returns 0 with an *errors.CannotReadPidError.
// If the daemon is not running the returned pid is -1.
func Pid() (int, error) {
	pidString, err := os.ReadFile(PidFile)
	if err != nil {
		return 0, errors.CannotReadPidError.Wrap(err, "Cannot read pid file")
	}
	pid, err := strconv.Atoi(string(pidString))
	return pid, nil
}

// ReloadDaemon sends a SIGHUP to the daemon to force it to reload its configuration.
// If an error occurs the error will be of type *errors.CannotReloadDaemonError.
func ReloadDaemon() error {
	pid, err := Pid()
	if err != nil {
		return errors.CannotReloadDaemonError.Wrap(err, "Cannot read pid file")
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return errors.CannotReloadDaemonError.Wrap(err, "Cannot find process")
	}
	err = process.Signal(SIGHUP)
	if err != nil {
		return errors.CannotReloadDaemonError.Wrap(err, "Cannot send interrupt signal")
	}
	return nil
}

// writeStatus is a helper function to write a daemon status to file.
// If an error occurs the error will be of type *errors.CannotWriteStatusFileError.
func writeStatus(state State) error {
	err := os.WriteFile(StatusFile, []byte(state), 0644)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "Cannot write status file")
	}
	return nil
}

// writePid is a helper function to write a daemon pid to file.
// If an error occurs the error will be of type *errors.CannotWritePidFileError.
func writePid(pid int) error {
	err := os.WriteFile(PidFile, []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return errors.CannotWriteFileError.Wrap(err, "Cannot write pid file")
	}
	return nil
}

// lockfile returns a lockfile.Lockfile on the lockfile to prevent multiple daemon instances from running at ths same time
func lockfile() (lf.Lockfile, error) {
	return lf.New(Directory() + string(filepath.Separator) + "daemon.lock")
}
