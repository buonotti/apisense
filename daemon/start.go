package daemon

import (
	"os"

	"github.com/nightlyone/lockfile"

	"github.com/buonotti/odh-data-monitor/errors"
)

func Start() error {
	if _, err := os.Stat(Directory); os.IsNotExist(err) {
		err = os.Mkdir(Directory, 0755)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "Cannot create daemon directory")
		}
	}

	lock, err := Lockfile()
	if err != nil {
		return errors.CannotReadLockFileError.Wrap(err, "Cannot create lock file")
	}
	err = lock.TryLock()
	if err != nil {
		return errors.CannotLockFileError.Wrap(err, "Cannot acquire lock file")
	}
	defer func(lock lockfile.Lockfile) {
		err := lock.Unlock()
		if err != nil {
			err = errors.CannotUnlockFileError.Wrap(err, "Cannot unlock lock file")
			errors.HandleError(err)
		}
	}(lock)
	return run()
}
