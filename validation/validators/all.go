package validators

import (
	"github.com/buonotti/apisense/util"
)

func All() []Validator {
	return []Validator{NewSchemaValidator()}
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
