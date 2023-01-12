package errors

import (
	"github.com/joomcode/errorx"
)

// TODO chris

var TuiErrors = errorx.NewNamespace("tui")
var WatcherError = TuiErrors.NewType("watcher_error")
var ModelError = TuiErrors.NewType("model_error")
var UnknownError = TuiErrors.NewType("unknown_error")
