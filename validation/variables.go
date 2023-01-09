package validation

import (
	"os"
	"time"
)

type EndpointParameter interface {
	Value(index int) any
}

type VariableEndpointParameter struct {
	values []string
}

func (p VariableEndpointParameter) Value(index int) any {
	return p.values[index]
}

type ConstantEndpointParameter struct {
	value string
}

func (p ConstantEndpointParameter) Value(int) any {
	return p.value
}

type VariableMap map[string]any

func (m VariableMap) Env(key string) string {
	return os.Getenv(key)
}

func (m VariableMap) Now(format string) string {
	return time.Now().Format(format)
}
