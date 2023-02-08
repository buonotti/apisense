package variables

import (
	"os"
	"time"
)

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
