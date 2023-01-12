package errors

import (
	"github.com/joomcode/errorx"
)

var CliErrors = errorx.NewNamespace("cli")
var CannotRegisterCompletionFunction = CliErrors.NewType("cannot_register_comp_func")
var CannotGetFlagValueError = CliErrors.NewType("cannot_get_flag_value", fatalTrait)
var UnknownReportError = CliErrors.NewType("unknown_report", fatalTrait)
var UnknownExportFormatError = CliErrors.NewType("unknown_report", fatalTrait)
var CannotReadInConfigError = CliErrors.NewType("cannot_read_in_config", fatalTrait)
