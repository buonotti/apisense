package errors

import (
	"github.com/joomcode/errorx"
)

var (
	TuiErrors            = errorx.NewNamespace("tui")
	WatcherError         = TuiErrors.NewType("watcher")
	NotifyError          = TuiErrors.NewType("notify")
	UnknownError         = TuiErrors.NewType("unknown")
	ModelError           = TuiErrors.NewType("model")
	RenderError          = TuiErrors.NewType("render")
	CannotRunPromptError = TuiErrors.NewType("cannotRunPrompt")
)
