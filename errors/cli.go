package errors

import (
	"github.com/joomcode/errorx"
)

var (
	CliErrors                        = errorx.NewNamespace("cli")
	CannotRegisterCompletionFunction = CliErrors.NewType("cannot_register_comp_func")
	CannotGetFlagValueError          = CliErrors.NewType("cannot_get_flag_value")
	CannotMarkFlagRequiredError      = CliErrors.NewType("cannot_mark_required")
	UnknownReportError               = CliErrors.NewType("unknown_report")
	UnknownExportFormatError         = CliErrors.NewType("unknown_report")
	CannotReadInConfigError          = CliErrors.NewType("cannot_read_in_config")
)
