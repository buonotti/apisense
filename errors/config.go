package errors

import "github.com/joomcode/errorx"

var (
	ConfigErrors = errorx.NewNamespace("config")
)

var (
	InvalidConfigFileError = ConfigErrors.NewType("invalid_config")
)
