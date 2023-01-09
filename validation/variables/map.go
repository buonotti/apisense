package variables

import (
	"os"
	"time"
)

type VariableMap map[string]any

func (m VariableMap) Env(key string) string {
	return os.Getenv(key)
}

func (m VariableMap) Now(format string) string {
	return time.Now().Format(format)
}
