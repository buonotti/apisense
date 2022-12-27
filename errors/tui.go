package errors

import (
	"github.com/joomcode/errorx"
)

var TuiErrors = errorx.NewNamespace("tui")
var UnknownError = TuiErrors.NewType("unknown error")
