package errors

import (
	"github.com/joomcode/errorx"
)

var (
	TuiErrors    = errorx.NewNamespace("tui")
	WatcherError = TuiErrors.NewType("watcher_error", fatalTrait)
	NotifyError  = TuiErrors.NewType("notify_error", fatalTrait)
	UnknownError = TuiErrors.NewType("unknown_error", fatalTrait)
	ModelError   = TuiErrors.NewType("model_error", fatalTrait)
)
