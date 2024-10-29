package validators

import (
	"github.com/buonotti/apisense/validation"
	"github.com/buonotti/apisense/validation/definitions"
)

func NewValidationItem(response validation.EndpointResponse, definition definitions.Endpoint) ValidationItem {
	return ExtendedValidationItem{
		response,
		definition,
	}
}

type ValidationItem interface {
	Response() validation.EndpointResponse
	Definition() definitions.Endpoint
}

type ExtendedValidationItem struct {
	response   validation.EndpointResponse
	definition definitions.Endpoint
}

func (i ExtendedValidationItem) Response() validation.EndpointResponse {
	return i.response
}

func (i ExtendedValidationItem) Definition() definitions.Endpoint {
	return i.definition
}

type SlimValidationItem struct {
	response validation.EndpointResponse
}

func (i SlimValidationItem) Response() validation.EndpointResponse {
	return i.response
}

func (i SlimValidationItem) Definition() definitions.Endpoint {
	return definitions.Endpoint{}
}

type Validator interface {
	Name() string
	Validate(item ValidationItem) error
	IsFatal() bool
	IsSlim() bool
}
