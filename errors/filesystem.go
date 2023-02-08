package errors

import (
	"github.com/joomcode/errorx"
)

var FileSystemErrors = errorx.NewNamespace("filesystem")
var FileNotFoundError = FileSystemErrors.NewType("file_not_found", fatalTrait)
var CannotParseDefinitionFileError = FileSystemErrors.NewType("cannot_parse_definition_file", fatalTrait)
var CannotWriteFileError = FileSystemErrors.NewType("cannot_write_file", fatalTrait)
var CannotReadFileError = FileSystemErrors.NewType("cannot_read_file", fatalTrait)
var CannotReadLockFileError = FileSystemErrors.NewType("cannot_read_lock_file", fatalTrait)
var CannotLockFileError = FileSystemErrors.NewType("cannot_lock_file", fatalTrait)
var CannotUnlockFileError = FileSystemErrors.NewType("cannot_unlock_file", fatalTrait)
var CannotCreateDirectoryError = FileSystemErrors.NewType("cannot_create_directory", fatalTrait)
var CannotFindReportFile = FileSystemErrors.NewType("cannot_find_report_file", fatalTrait)
var CannotWriteConfigError = FileSystemErrors.NewType("cannot_write_config", fatalTrait)
var CannotCreateFileError = FileSystemErrors.NewType("cannot_create_file", fatalTrait)
var CannotReadDirectoryError = FileSystemErrors.NewType("cannot_read_directory", fatalTrait)
var CannotUnmarshalThemeError = FileSystemErrors.NewType("cannot_unmarshal_theme", fatalTrait)
