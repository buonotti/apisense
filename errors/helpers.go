package errors

import (
	"fmt"

	"github.com/joomcode/errorx"
)

// NewF is a helper function that allows to create a new *errorx.Error of the given type with a formatted message
func NewF(t *errorx.Type, format string, args ...interface{}) *errorx.Error {
	return t.New(fmt.Sprintf(format, args...))
}

// WrapF is a helper function that allows to wrap an existing error around the given type with a formatted message
func WrapF(t *errorx.Type, err error, format string, args ...interface{}) *errorx.Error {
	return t.Wrap(err, fmt.Sprintf(format, args...))
}

// SafeWrap is a helper function to wrap an existing error. The function returns nil if the error is nil or a new error of the given type wrapping the given error
func SafeWrap(t *errorx.Type, err error, message string, args ...interface{}) error {
	if err != nil {
		return t.Wrap(err, message, args...)
	}
	return nil
}

// SafeWrapF does the same thing as SafeWrap but accepts a format string with arguments as message
func SafeWrapF(t *errorx.Type, err error, format string, args ...interface{}) error {
	if err != nil {
		return WrapF(t, err, format, args...)
	}
	return nil
}
