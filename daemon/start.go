package daemon

import (
	"net/http"
	"os"

	lf "github.com/nightlyone/lockfile"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/files"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/pipeline"
	"github.com/buonotti/apisense/validation/validators"
)

// Start starts the daemon. If the daemon is already running it returns an
// *errors.CannotLockFileError because the already running daemon has the lock on
// the file.
func Start(runOnStart bool) error {
	file, err := os.Create(files.DaemonStatusFile())
	if err != nil {
		return errors.CannotCreateFileError.Wrap(err, "cannot create status file")
	}

	err = file.Close()
	if err != nil {
		return errors.CannotCloseFileError.Wrap(err, "cannot close status file")
	}

	lock, err := lockfile()
	if err != nil {
		return errors.CannotReadLockFileError.Wrap(err, "cannot create lock file")
	}

	err = lock.TryLock()
	if err != nil {
		return errors.CannotLockFileError.Wrap(err, "cannot acquire lock file")
	}

	defer func(lock lf.Lockfile) {
		err := lock.Unlock()
		if err != nil {
			err = errors.CannotUnlockFileError.Wrap(err, "cannot unlock lock file")
			errors.CheckErr(err)
		}
	}(lock)

	pipe, err := NewPipeline()
	if err != nil {
		return err
	}

	d := daemon{
		Pipeline: pipe,
	}

	if viper.GetBool("daemon.rpc") {
		go func() {
			err := startRpcServer(&d)
			if err != nil && err != http.ErrServerClosed {
				errors.CheckErr(err)
			}
		}()
	}

	return d.run(runOnStart)
}

func NewPipeline() (*pipeline.Pipeline, error) {
	pipelineWithValidators, err := pipeline.NewPipelineWithValidators(
		validators.Without(viper.GetStringSlice("validation.excluded_builtin_validators")...)...,
	)
	if err != nil {
		return nil, err
	}

	externalValidators, err := validators.LoadExternalValidators()
	if err != nil {
		return nil, err
	}

	for _, externalValidator := range externalValidators {
		log.DaemonLogger().Info("Loading external validator", "name", externalValidator.Name())
		pipelineWithValidators.AddValidator(externalValidator)
	}

	return &pipelineWithValidators, err
}
