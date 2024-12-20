package daemon

import (
	"os"

	"github.com/buonotti/apisense/v2/errors"
	"github.com/buonotti/apisense/v2/filesystem/locations/files"
	"github.com/buonotti/apisense/v2/log"
	"github.com/buonotti/apisense/v2/validation/pipeline"
	"github.com/buonotti/apisense/v2/validation/validators"
	lf "github.com/nightlyone/lockfile"
)

// Setup creates the necessary files for the daemon. Is called also on daemon start
func Setup() error {
	file, err := os.Create(files.DaemonStatusFile())
	if err != nil {
		return errors.CannotCreateFileError.Wrap(err, "cannot create status file")
	}
	defer file.Close()

	pFile, err := os.Create(files.DaemonPidFile())
	if err != nil {
		return errors.CannotCreateFileError.Wrap(err, "cannot create pid file")
	}
	defer pFile.Close()

	return nil
}

// Start starts the daemon. If the daemon is already running it returns an
// *errors.CannotLockFileError because the already running daemon has the lock on
// the file.
func Start(runOnStart bool) error {
	err := Setup()
	if err != nil {
		return err
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
		lockErr := lock.Unlock()
		if lockErr != nil {
			lockErr = errors.CannotUnlockFileError.Wrap(lockErr, "cannot unlock lock file")
			log.DaemonLogger().Fatal(lockErr)
		}
	}(lock)

	pipe, err := NewPipeline()
	if err != nil {
		return err
	}

	d := daemon{
		Pipeline: pipe,
	}

	return d.run(runOnStart)
}

// NewPipeline creates a new validation pipeline
func NewPipeline() (*pipeline.Pipeline, error) {
	pipelineWithValidators, err := pipeline.NewPipelineWithValidators(validators.NewSchemaValidator())
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
