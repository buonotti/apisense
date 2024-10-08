package errors

import (
	"github.com/joomcode/errorx"
)

var (
	FileSystemErrors               = errorx.NewNamespace("filesystem")
	FileNotFoundError              = FileSystemErrors.NewType("file_not_found")
	CannotParseDefinitionFileError = FileSystemErrors.NewType("cannot_parse_definition_file")
	CannotWriteFileError           = FileSystemErrors.NewType("cannot_write_file")
	CannotReadFileError            = FileSystemErrors.NewType("cannot_read_file")
	CannotReadLockFileError        = FileSystemErrors.NewType("cannot_read_lock_file")
	CannotLockFileError            = FileSystemErrors.NewType("cannot_lock_file")
	CannotUnlockFileError          = FileSystemErrors.NewType("cannot_unlock_file")
	CannotCreateDirectoryError     = FileSystemErrors.NewType("cannot_create_directory")
	CannotFindReportFile           = FileSystemErrors.NewType("cannot_find_report_file")
	CannotWriteConfigError         = FileSystemErrors.NewType("cannot_write_config")
	CannotCreateFileError          = FileSystemErrors.NewType("cannot_create_file")
	CannotReadDirectoryError       = FileSystemErrors.NewType("cannot_read_directory")
	CannotRemoveFileError          = FileSystemErrors.NewType("cannot_remove_file")
	CannotCloseFileError           = FileSystemErrors.NewType("cannot_close_file")
)
