package errors

import (
	"github.com/joomcode/errorx"
)

var (
	TuiErrors    = errorx.NewNamespace("tui")
	WatcherError = TuiErrors.NewType("watcher_error")
	NotifyError  = TuiErrors.NewType("notify_error")
	UnknownError = TuiErrors.NewType("unknown_error")
	ModelError   = TuiErrors.NewType("model_error")
)
