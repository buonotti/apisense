package errors

import (
	"github.com/joomcode/errorx"
)

var SSHErrors = errorx.NewNamespace("ssh")
var CannotStartSSHServer = SSHErrors.NewType("cannot_start_ssh_server")
var CannotStopSSHServer = SSHErrors.NewType("cannot_stop_ssh_server")
var CannotCreateSSHServer = SSHErrors.NewType("cannot_create_ssh_server")
