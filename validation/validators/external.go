package validators

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/buonotti/apisense/errors"
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
func (v externalValidator) Validate(item validation.PipelineTestCase) error {
	jsonString, err := json.Marshal(item)
	outString := &strings.Builder{}
	if err != nil {
		return errors.CannotSerializeItemError.Wrap(err, "cannot serialize item: %s", err)
	}
	cmd := exec.Command(v.Definition.Path, v.Definition.Args...)
	if v.Definition.ReadFromStdin {
		cmd.Stdin = strings.NewReader(string(jsonString))
		cmd.Stdout = outString
	}

	validatorOut := strings.Builder{}
	validatorErr := strings.Builder{}
	cmd.Stdout = &validatorOut
	cmd.Stderr = &validatorErr

	err = cmd.Run()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				return errors.NewF(errors.ValidationError, "validation failed for endpoint %s: %s", item.EndpointName, validatorErr.String())
			}
		}
	}
	return nil
}

func (v externalValidator) IsFatal() bool {
	return v.Definition.Fatal
}
