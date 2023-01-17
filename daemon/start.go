package daemon

import (
	"os/exec"

	lf "github.com/nightlyone/lockfile"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/external"
	"github.com/buonotti/apisense/validation/validators"
)

// Start starts the daemon. If the daemon is already running it returns an
// *errors.CannotLockFileError because the already running daemon has the lock on
// the file. The background parameters controls whether the daemon should be run
// in the foreground or not.
func Start(background bool, runOnStart bool) (*exec.Cmd, error) {
	// If the background flag is set start a new process which runs the daemon
	// without the --bg flag which calls this function with background = false
	if background {
		cmd := exec.Command("apisense", "daemon", "start")
		return cmd, cmd.Start()
	}

	// Get the lockfile and try to lock it. If the lock cannot be acquired return an error else
	// defer a call to the unlock function
	lock, err := lockfile()
	if err != nil {
		return nil, errors.CannotReadLockFileError.Wrap(err, "Cannot create lock file")
	}

	err = lock.TryLock()
	if err != nil {
		return nil, errors.CannotLockFileError.Wrap(err, "Cannot acquire lock file")
	}

	defer func(lock lf.Lockfile) {
		err := lock.Unlock()
		if err != nil {
			err = errors.CannotUnlockFileError.Wrap(err, "Cannot unlock lock file")
			errors.HandleError(err)
		}
	}(lock)

	// Create a new pipeline with the status schema and range validators, then load
	// all external validators from the config file
	pipeline, err := validation.NewPipelineV(
		validators.NewStatusValidator(),
		validators.NewSchemaValidator(),
		validators.NewRangeValidator(),
	)

	externalValidators, err := external.Parse()
	if err != nil {
		return nil, err
	}

	for _, externalValidator := range externalValidators {
		pipeline.AddValidator(validators.NewExternalValidator(externalValidator))
	}
	if err != nil {
		return nil, err
	}

	// Create the daemon with the pipeline then run the daemon
	d := daemon{
		Pipeline: &pipeline,
	}
	return nil, d.run(runOnStart)
}
