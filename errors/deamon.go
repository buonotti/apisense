package errors

import (
	"github.com/joomcode/errorx"
)

var DaemonErrors = errorx.NewNamespace("daemon")
var CannotRequestDataError = DaemonErrors.NewType("cannot_request_data")
var CannotParseDataError = DaemonErrors.NewType("cannot_parse_data", fatalTrait)
var UnknownFormatError = DaemonErrors.NewType("unknown_format", fatalTrait)
var CannotApplyTemplateError = DaemonErrors.NewType("cannot_apply_template", fatalTrait)
var CannotReloadDaemonError = DaemonErrors.NewType("cannot_reload_daemon")
var CannotGetProcessInfoError = DaemonErrors.NewType("cannot_get_process_info", fatalTrait)
var CannotSendSignalError = DaemonErrors.NewType("cannot_send_signal")
var DaemonNotRunningError = DaemonErrors.NewType("daemon_not_running")
var DuplicateDefinitionError = DaemonErrors.NewType("duplicate_definition", fatalTrait)
