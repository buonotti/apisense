package errors

import (
	"github.com/joomcode/errorx"
)

var (
	CliErrors                        = errorx.NewNamespace("cli")
	CannotRegisterCompletionFunction = CliErrors.NewType("cannot_register_comp_func")
	CannotGetFlagValueError          = CliErrors.NewType("cannot_get_flag_value", fatalTrait)
	UnknownReportError               = CliErrors.NewType("unknown_report", fatalTrait)
	UnknownExportFormatError         = CliErrors.NewType("unknown_report", fatalTrait)
	CannotReadInConfigError          = CliErrors.NewType("cannot_read_in_config", fatalTrait)
)
