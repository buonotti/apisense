package errors

import (
	"github.com/joomcode/errorx"
)

var (
	FileSystemErrors               = errorx.NewNamespace("filesystem")
	FileNotFoundError              = FileSystemErrors.NewType("file_not_found", fatalTrait)
	CannotParseDefinitionFileError = FileSystemErrors.NewType("cannot_parse_definition_file", fatalTrait)
	CannotWriteFileError           = FileSystemErrors.NewType("cannot_write_file", fatalTrait)
	CannotReadFileError            = FileSystemErrors.NewType("cannot_read_file", fatalTrait)
	CannotReadLockFileError        = FileSystemErrors.NewType("cannot_read_lock_file", fatalTrait)
	CannotLockFileError            = FileSystemErrors.NewType("cannot_lock_file", fatalTrait)
	CannotUnlockFileError          = FileSystemErrors.NewType("cannot_unlock_file", fatalTrait)
	CannotCreateDirectoryError     = FileSystemErrors.NewType("cannot_create_directory", fatalTrait)
	CannotFindReportFile           = FileSystemErrors.NewType("cannot_find_report_file", fatalTrait)
	CannotWriteConfigError         = FileSystemErrors.NewType("cannot_write_config", fatalTrait)
	CannotCreateFileError          = FileSystemErrors.NewType("cannot_create_file", fatalTrait)
	CannotReadDirectoryError       = FileSystemErrors.NewType("cannot_read_directory", fatalTrait)
	CannotRemoveFileError          = FileSystemErrors.NewType("cannot_remove_file", fatalTrait)
	CannotCloseFileError           = FileSystemErrors.NewType("cannot_close_file", fatalTrait)
)
