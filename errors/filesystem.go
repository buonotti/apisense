package errors

import (
	"github.com/joomcode/errorx"
)

// FileSystemErrors is the namespace holding all file system related errors
var FileSystemErrors = errorx.NewNamespace("filesystem")

// FileNotFoundError is returned when a file is not found
var FileNotFoundError = FileSystemErrors.NewType("file_not_found")

// CannotParseDefinitionFileError is returned when the daemon fails to parse the definition file
var CannotParseDefinitionFileError = FileSystemErrors.NewType("cannot_parse_definition_file")

// CannotWriteFileError is returned when a functions fails to read a file
var CannotWriteFileError = FileSystemErrors.NewType("cannot_write_file")

// CannotReadStatusError is returned when the file containing the daemon status cannot be read
var CannotReadStatusError = FileSystemErrors.NewType("cannot_read_status")

// CannotReadPidError is returned when the file containing the daemon pid cannot be read
var CannotReadPidError = FileSystemErrors.NewType("cannot_read_pid")

// CannotReadLockFileError is returned when the lock file cannot be read or does not exist
var CannotReadLockFileError = FileSystemErrors.NewType("cannot_read_lock_file")

// CannotLockFileError is returned when the lock file cannot be locked by the current process. It generally means that another daemon is already running.
var CannotLockFileError = FileSystemErrors.NewType("cannot_lock_file")

// CannotUnlockFileError is returned when the already locked lock file cannot be unlocked
var CannotUnlockFileError = FileSystemErrors.NewType("cannot_unlock_file")

// CannotCreateDirectoryError is returned when any of the setup functions fails to create the required directories
var CannotCreateDirectoryError = FileSystemErrors.NewType("cannot_create_directory")
