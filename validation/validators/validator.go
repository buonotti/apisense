package validators

import (
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/definitions"
)

type ValidationItem interface {
	Response() validation.EndpointResponse
	Definition() definitions.Endpoint
}

type ExtendedValidationItem struct {
	Response   validation.EndpointResponse
	Definition definitions.Endpoint
}

type SlimValidationItem struct {
	Response validation.EndpointResponse
}

type Validator interface {
	Name() string
	Validate(item ValidationItem) error
	IsFatal() bool
	IsSlim() bool
}
