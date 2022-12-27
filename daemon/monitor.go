package daemon

import (
	"os"

	"github.com/nightlyone/lockfile"

	"github.com/buonotti/odh-data-monitor/config"
	"github.com/buonotti/odh-data-monitor/errors"
)

var Directory = config.Directory + "/daemon"
var PidFile = Directory + "/daemon.pid"
var StatusFile = Directory + "/daemon.status"

type State string

const (
	UP   State = "up"
	DOWN State = "down"
)

func Status() (State, error) {
	statusString, err := os.ReadFile(StatusFile)
	if err != nil {
		return DOWN, errors.CannotReadStatusError.Wrap(err, "Cannot read status file")
	}
	return State(statusString), nil
}

func Lockfile() (lockfile.Lockfile, error) {
	return lockfile.New(Directory + "/daemon.lock")
}
