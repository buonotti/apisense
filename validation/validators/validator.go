package validators

import (
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/definitions"
)

type ValidationItem struct {
	Response   validation.EndpointResponse
	Definition definitions.Endpoint
}

type Validator interface {
	Name() string
	Validate(item ValidationItem) error
	IsFatal() bool
}
