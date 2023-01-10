package validators

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/buonotti/odh-data-monitor/validation/external"
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
	if err != nil {
		return errors.CannotSerializeItemError.Wrap(err, "cannot serialize item: %s", err)
	}
	cmd := exec.Command(v.Definition.Path, v.Definition.Args...)
	if v.Definition.ReadFromStdin {
		cmd.Stdin = strings.NewReader(string(jsonString))
	}
	err = cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			for _, exitCode := range v.Definition.ExitCodes {
				if exitCode.Code == exitError.ExitCode() {
					if exitCode.Ok {
						return nil
					}
					return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Url, exitCode.Description)
				}
			}
			return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Url, err)
		}
		return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Url, err)
	} else {
		if len(v.Definition.ExitCodes) > 0 {
			for _, exitCode := range v.Definition.ExitCodes {
				if exitCode.Code == 0 {
					if exitCode.Ok {
						return nil
					}
					return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Url, exitCode.Description)
				}
			}
		}
	}
	return nil
}

func (v externalValidator) Fatal() bool {
	return v.Definition.Fatal
}
