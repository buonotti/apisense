package validators

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/buonotti/odh-data-monitor/validation"
	"github.com/buonotti/odh-data-monitor/validation/external"
)

func NewExternalValidator(definition external.ValidatorDefinition) validation.Validator {
	return externalValidator{
		Definition: definition,
	}
}

type externalValidator struct {
	Definition external.ValidatorDefinition
}

func (v externalValidator) Name() string {
	return "external." + v.Definition.Name
}

func (v externalValidator) Validate(item validation.Item) error {
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
					return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Endpoint, exitCode.Description)
				}
			}
			return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Endpoint, err)
		}
		return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Endpoint, err)
	} else {
		if len(v.Definition.ExitCodes) > 0 {
			for _, exitCode := range v.Definition.ExitCodes {
				if exitCode.Code == 0 {
					if exitCode.Ok {
						return nil
					}
					return errors.ValidationError.New("validation failed for endpoint %s: %s", item.Endpoint, exitCode.Description)
				}
			}
		}
	}
	return nil
}
