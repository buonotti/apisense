package daemon

import (
	lf "github.com/nightlyone/lockfile"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/external"
	"github.com/buonotti/apisense/validation/validators"
)

// Start starts the daemon. If the daemon is already running it returns an
// *errors.CannotLockFileError because the already running daemon has the lock on
// the file.
func Start(runOnStart bool) error {
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

	pipeline, err := NewPipeline()
	if err != nil {
		return err
	}

	d := daemon{
		Pipeline: pipeline,
	}
	return d.run(runOnStart)
}

func NewPipeline() (*validation.Pipeline, error) {
	pipeline, err := validation.NewPipelineWithValidators(
		validators.All()...,
	)

	externalValidators, err := external.Parse()
	if err != nil {
		return nil, err
	}

	for _, externalValidator := range externalValidators {
		pipeline.AddValidator(validators.NewExternalValidator(externalValidator))
	}
	return &pipeline, err
}
