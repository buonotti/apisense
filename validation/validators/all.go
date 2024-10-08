package validators

import (
	"github.com/buonotti/apisense/util"
)

var allValidators = map[string]Validator{
	"range":  NewRangeValidator(),
	"schema": NewSchemaValidator(),
}

func All() []Validator {
	return util.Values(allValidators)
}

func Without(names ...string) []Validator {
	var validators []Validator
	for _, validator := range All() {
		if !util.Contains(names, validator.Name()) {
			validators = append(validators, validator)
		}
	}
	return validators
}
