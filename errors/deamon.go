package errors

import (
	"github.com/joomcode/errorx"
)

var DaemonErrors = errorx.NewNamespace("daemon")
var CannotReadStatusError = DaemonErrors.NewType("cannot_read_status")
var CannotReadPidError = DaemonErrors.NewType("cannot_read_pid")
var CannotReadLockFileError = DaemonErrors.NewType("cannot_read_lock_file", fatalTrait)
var CannotLockFileError = DaemonErrors.NewType("cannot_lock_file")
var CannotUnlockFileError = DaemonErrors.NewType("cannot_unlock_file")
var CannotCreateDirectoryError = DaemonErrors.NewType("cannot_create_directory", fatalTrait)
var CannotRequestDataError = DaemonErrors.NewType("cannot_request_data")
var CannotParseDataError = DaemonErrors.NewType("cannot_parse_data")
var UnknownFormatError = DaemonErrors.NewType("unknown_format")
var CannotApplyTemplateError = DaemonErrors.NewType("cannot_apply_template")
var CannotWriteStatusFileError = DaemonErrors.NewType("cannot_write_status_file")
var CannotReloadDaemonError = DaemonErrors.NewType("cannot_reload_daemon")
var CannotWritePidFileError = DaemonErrors.NewType("cannot_write_pid_file")
var CannotGetProcessInfoError = DaemonErrors.NewType("cannot_get_process_info")
var CannotSendSignalError = DaemonErrors.NewType("cannot_send_signal")
var DaemonNotRunningError = DaemonErrors.NewType("daemon_not_running")
