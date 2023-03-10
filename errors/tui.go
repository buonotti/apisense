package errors

import (
	"github.com/joomcode/errorx"
)

var TuiErrors = errorx.NewNamespace("tui")
var WatcherError = TuiErrors.NewType("watcher_error", fatalTrait)
var NotifyError = TuiErrors.NewType("notify_error", fatalTrait)
var UnknownError = TuiErrors.NewType("unknown_error", fatalTrait)
var ModelError = TuiErrors.NewType("model_error", fatalTrait)
