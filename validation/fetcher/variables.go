package fetcher

import (
	"os"
	"time"
)

// EndpointParameter is an interface that defines a parameter that will be interpolated as variable in an endpoint request
type EndpointParameter interface {
	Value(index int) any // Value returns the value of the parameter at the given index
}

// NewVariableEndpointParameter returns a new EndpointParameter with the given values
func NewVariableEndpointParameter(values []string) EndpointParameter {
	return VariableEndpointParameter{values: values}
}

// VariableEndpointParameter is a parameter that returns a different value based on the index in the given collection
type VariableEndpointParameter struct {
	values []string // values is the collection of values that will be returned
}

// Value returns the value in the initial collection at the given index
func (p VariableEndpointParameter) Value(index int) any {
	return p.values[index]
}

// ConstantEndpointParameter is a parameter that always returns the same value
type ConstantEndpointParameter struct {
	value string // value is the value that is always returned
}

// NewConstantEndpointParameter returns a new EndpointParameter with the given value
func NewConstantEndpointParameter(value string) EndpointParameter {
	return ConstantEndpointParameter{value: value}
}

// Value always returns the same value
func (p ConstantEndpointParameter) Value(int) any {
	return p.value
}

// VariableMap is a map of variables that will be generated from a collection of
// EndpointParameter and will be used when executing the go template
type VariableMap map[string]any

// Env is a function that returns the value of a given environment variable
func (m VariableMap) Env(key string) string {
	return os.Getenv(key)
}

// Now is a function that returns the current time in the given format (see Time.Format)
func (m VariableMap) Now(format string) string {
	return time.Now().Format(format)
}
