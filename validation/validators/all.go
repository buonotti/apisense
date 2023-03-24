package validators

import (
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
)

var allValidators = map[string]validation.Validator{
	"status": NewStatusValidatorC(200),
	"range":  NewRangeValidator(),
	"schema": NewSchemaValidator(),
}

func All() []validation.Validator {
	return util.Values(allValidators)
}

func Without(names ...string) []validation.Validator {
	var validators []validation.Validator
	for _, validator := range All() {
		if !util.Contains(names, validator.Name()) {
			validators = append(validators, validator)
		}
	}
	return validators
}
