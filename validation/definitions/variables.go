package definitions

import (
	"os"
	"time"
)

// Variable describes a variable that should be interpolated in the base url and the query parameters
type Variable struct {
	Name       string   `yaml:"name" json:"name" validate:"required"`         // Name is the name of the variable
	IsConstant bool     `yaml:"constant" json:"constant" validate:"required"` // IsConstant is true if the value of the variable is constant or else false
	Values     []string `yaml:"values" json:"values" validate:"required"`     // Values are all the possible values of the variable (only 1 in case of a constant)
}

// Value returns the value of the variable according to the index of the test case
func (v Variable) Value(index int) any {
	if v.IsConstant {
		return v.Values[0]
	}
	return v.Values[index]
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
