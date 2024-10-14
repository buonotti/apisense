package filesystem

import (
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
)

// Setup creates the daemon directory and writes the default files to it.
func Setup() error {
	if _, err := os.UserHomeDir(); err != nil {
		return errors.CannotCreateDirectoryError.Wrap(nil, "cannot setup environment: $HOME is not set")
	}

	err := createDirectories()
	if err != nil {
		return err
	}

	return nil
}

// createDirectories creates the daemon, reports and the definitions directories.
// If any of the directories exist it does nothing.
// If there is an error while creating the directories it returns an *errors.CannotCreateDirectoryError.
func createDirectories() error {
	err := os.MkdirAll(filepath.FromSlash(directories.AppDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create apisense directory: "+directories.AppDirectory())
	}

	err = os.MkdirAll(filepath.FromSlash(directories.DaemonDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create daemon directory: "+directories.DaemonDirectory())
	}

	err = os.MkdirAll(filepath.FromSlash(directories.DefinitionsDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create definitions directory: "+directories.DefinitionsDirectory())
	}

	err = os.MkdirAll(filepath.FromSlash(directories.ReportsDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create reports directory: "+directories.ReportsDirectory())
	}

	err = os.MkdirAll(filepath.FromSlash(directories.ValidatorsDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create validators directory: "+directories.ValidatorRepoDirectory())
	}

	err = os.MkdirAll(filepath.FromSlash(directories.ValidatorRepoDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create validators repo directory: "+directories.ValidatorRepoDirectory())
	}

	err = os.MkdirAll(filepath.FromSlash(directories.ValidatorCustomDirectory()), os.ModePerm)
	if err != nil {
		return errors.CannotCreateDirectoryError.Wrap(err, "cannot create validators custom directory: "+directories.ValidatorCustomDirectory())
	}

	return nil
}
