package validators

import (
	"github.com/buonotti/apisense/util"
	"github.com/buonotti/apisense/validation"
)

var allValidators = map[string]validation.Validator{
	"status": statusValidator{OkStatus: 200},
	"range":  rangeValidator{},
	"schema": schemaValidator{},
}

func All() []validation.Validator {
	return util.Values(allValidators)
}
