package errors

import (
	"github.com/joomcode/errorx"
)

// DaemonErrors is the namespace holding all daemon related errors
var DaemonErrors = errorx.NewNamespace("daemon")

// CannotRequestDataError is returned when the daemon fails to request data from the ODH APIs
var CannotRequestDataError = DaemonErrors.NewType("cannot_request_data")

// CannotParseDataError is returned when the daemon fails to deserialize the data received from the ODH APIs to a map
var CannotParseDataError = DaemonErrors.NewType("cannot_parse_data")

// UnknownFormatError is returned when the config declares a format that is not supported by the daemon
var UnknownFormatError = DaemonErrors.NewType("unknown_format")

// CannotApplyTemplateError is returned when the daemon fails to apply the
// variables using the go template engine (see package text/template)
var CannotApplyTemplateError = DaemonErrors.NewType("cannot_apply_template")

// CannotReloadDaemonError is returned when the daemon.ReloadDaemon function
// fails to send a SIGHUP to the daemon or get the daemons pid
var CannotReloadDaemonError = DaemonErrors.NewType("cannot_reload_daemon")

// CannotGetProcessInfoError is returned when the daemon.Stop function cannot get the process info of the daemon by its pid
var CannotGetProcessInfoError = DaemonErrors.NewType("cannot_get_process_info")

// CannotSendSignalError is returned when the cli fails to send a signal to the daemon
var CannotSendSignalError = DaemonErrors.NewType("cannot_send_signal")

// DaemonNotRunningError is returned when the daemon is not running and the cli tries to stop it
var DaemonNotRunningError = DaemonErrors.NewType("daemon_not_running")
