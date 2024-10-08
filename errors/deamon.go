package errors

import (
	"github.com/joomcode/errorx"
)

var (
	DaemonErrors                              = errorx.NewNamespace("daemon")
	CannotRequestDataError                    = DaemonErrors.NewType("cannot_request_data")
	CannotParseDataError                      = DaemonErrors.NewType("cannot_parse_data", fatalTrait)
	UnknownFormatError                        = DaemonErrors.NewType("unknown_format", fatalTrait)
	CannotApplyTemplateError                  = DaemonErrors.NewType("cannot_apply_template", fatalTrait)
	CannotReloadDaemonError                   = DaemonErrors.NewType("cannot_reload_daemon")
	CannotGetProcessInfoError                 = DaemonErrors.NewType("cannot_get_process_info", fatalTrait)
	CannotSendSignalError                     = DaemonErrors.NewType("cannot_send_signal")
	DaemonNotRunningError                     = DaemonErrors.NewType("daemon_not_running")
	DuplicateDefinitionError                  = DaemonErrors.NewType("duplicate_definition", fatalTrait)
	CannotAddWorkFunctionToCronSchedulerError = DaemonErrors.NewType("cannot_add_work_function_to_cron_scheduler", fatalTrait)
)

