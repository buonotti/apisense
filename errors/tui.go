package errors

import (
	"github.com/joomcode/errorx"
)

// TODO chris

var TuiErrors = errorx.NewNamespace("tui")
var WatcherError = TuiErrors.NewType("watcher_error")
var NotifyError = TuiErrors.NewType("notify_error")
var UnknownError = TuiErrors.NewType("unknown_error")
