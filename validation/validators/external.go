package validators

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/external"
)

// NewExternalValidator creates a new external validator based on the given definition and returns a validation.Validator
func NewExternalValidator(definition external.ValidatorDefinition) validation.Validator {
	return externalValidator{
		Definition: definition,
	}
}

// externalValidator is a validator that was defined in the configuration file
type externalValidator struct {
	Definition external.ValidatorDefinition // Definition is the external.ValidatorDefinition that defines the external validator
}

// Name returns the name of the validator: external.<name> where name is the name
// of the external validator defined in the config
func (v externalValidator) Name() string {
	return "external." + v.Definition.Name
}

// Validate validates an item by serializing it and sending it to the external
// process then returning an error according to the status code of the external
// program
func (v externalValidator) Validate(item validation.PipelineItem) error {
	jsonString, err := json.Marshal(item)
	outString := &strings.Builder{}
	if err != nil {
		return errors.CannotSerializeItemError.Wrap(err, "cannot serialize item: %s", err)
	}
	cmd := exec.Command(v.Definition.Path, v.Definition.Args...)
	log.DaemonLogger.Infof("Running external validator %s with args %v", v.Definition.Path, v.Definition.Args)
	if v.Definition.ReadFromStdin {
		cmd.Stdin = strings.NewReader(string(jsonString))
		cmd.Stdout = outString
	}
	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			for _, exitCode := range v.Definition.ExitCodes {
				if exitCode.Code == int64(exitError.ExitCode()) {
					if exitCode.Ok {
						return nil
					}
					return errors.ValidationError.New("validation failed for endpoint %s: %s: %s", item.Url, exitCode.Description, outString)
				}
			}
			return errors.ValidationError.New("validation failed for endpoint %s: %s: %s", item.Url, err, outString)
		}
		return errors.ValidationError.New("validation failed for endpoint %s: %s: %s", item.Url, err, outString)
	} else {
		if len(v.Definition.ExitCodes) > 0 {
			for _, exitCode := range v.Definition.ExitCodes {
				if exitCode.Code == 0 {
					if exitCode.Ok {
						return nil
					}
					return errors.ValidationError.New("validation failed for endpoint %s: %s: %s", item.Url, exitCode.Description, outString)
				}
			}
		}
	}
	return nil
}

func (v externalValidator) Fatal() bool {
	return v.Definition.Fatal
}
