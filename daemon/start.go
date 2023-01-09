package daemon

import (
	"github.com/nightlyone/lockfile"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/buonotti/odh-data-monitor/validation/external"
	"github.com/buonotti/odh-data-monitor/validation/validators"
)

func Start() error {
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
	pipeline, err := validation.NewPipelineV(
		validators.NewStatusValidator(),
		validators.NewSchemaValidator(),
		validators.NewRangeValidator(),
	)
	externalValidators, err := external.Parse()
	if err != nil {
		return err
	}
	for _, externalValidator := range externalValidators {
		pipeline.AddValidator(validators.NewExternalValidator(externalValidator))
	}
	if err != nil {
		return err
	}
	d := daemon{
		Pipeline: &pipeline,
	}
	return d.run()
}
