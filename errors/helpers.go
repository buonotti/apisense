package errors

import (
	"fmt"

	"github.com/joomcode/errorx"
)

func Newf(t *errorx.Type, format string, args ...interface{}) *errorx.Error {
	return t.New(fmt.Sprintf(format, args...))
}

func Wrapf(t *errorx.Type, err error, format string, args ...interface{}) *errorx.Error {
	return t.Wrap(err, fmt.Sprintf(format, args...))
}
