package fs

import (
	"os"

	"github.com/buonotti/apisense/daemon"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
)

// Setup creates the daemon directory and writes the default files to it.
func Setup() error {
	err := createDirectories()
	if err != nil {
		return err
	}

	err = writeDaemonStatusFile()
	if err != nil {
		return err
	}

	return nil
}

// createDirectories creates the daemon, reports and the definitions directories.
// If any of the directories exist it does nothing.
// If there is an error while creating the directories it returns an *errors.CannotCreateDirectoryError.
func createDirectories() error {
	dataDirectory := os.Getenv("HOME") + "/apisense"

	if _, err := os.Stat(dataDirectory); os.IsNotExist(err) {
		err = os.Mkdir(dataDirectory, os.ModePerm)
		if err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create apisense directory: "+dataDirectory)
		}
	}

	if _, err := os.Stat(daemon.Directory()); os.IsNotExist(err) {
		if err = os.Mkdir(daemon.Directory(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create daemon directory: "+daemon.Directory())
		}
	}

	if _, err := os.Stat(validation.DefinitionsLocation()); os.IsNotExist(err) {
		if err = os.Mkdir(validation.DefinitionsLocation(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create definitions directory: "+validation.DefinitionsLocation())
		}
	}

	if _, err := os.Stat(validation.ReportLocation()); os.IsNotExist(err) {
		if err = os.Mkdir(validation.ReportLocation(), 0755); err != nil {
			return errors.CannotCreateDirectoryError.Wrap(err, "cannot create reports directory: "+validation.ReportLocation())
		}
	}

	return nil
}

// writeDaemonStatusFile writes the daemon status file with the default value of "DOWN".
// If there is an error while writing the file it returns an *errors.CannotWriteFileError.
func writeDaemonStatusFile() error {
	if _, err := os.Stat(daemon.StatusFile); os.IsNotExist(err) {
		err = os.WriteFile(daemon.StatusFile, []byte(daemon.DOWN), os.ModePerm)
		if err != nil {
			return errors.CannotWriteFileError.Wrap(err, "cannot write status file")
		}
	}
	return nil
}
