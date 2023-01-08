package daemon

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/nightlyone/lockfile"

	"github.com/buonotti/odh-data-monitor/errors"
)

var pidFile = Directory() + "/daemon.pid"
var statusFile = Directory() + "/daemon.status"

type State string

const (
	UP   State = "up"
	DOWN State = "down"
)

func Status() (State, error) {
	statusString, err := os.ReadFile(statusFile)
	if err != nil {
		return DOWN, errors.CannotReadStatusError.Wrap(err, "Cannot read status file")
	}
	return State(statusString), nil
}

func Pid() (int, error) {
	pidString, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, errors.CannotReadPidError.Wrap(err, "Cannot read pid file")
	}
	pid, err := strconv.Atoi(string(pidString))
	return pid, nil
}

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

func writeStatus(state State) error {
	err := os.WriteFile(statusFile, []byte(state), 0644)
	if err != nil {
		return errors.CannotWriteStatusFileError.Wrap(err, "Cannot write status file")
	}
	return nil
}

func writePid(pid int) error {
	err := os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return errors.CannotWritePidFileError.Wrap(err, "Cannot write pid file")
	}
	return nil
}

func Lockfile() (lockfile.Lockfile, error) {
	return lockfile.New(Directory() + string(filepath.Separator) + "daemon.lock")
}
