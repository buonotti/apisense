package errors

import (
	"github.com/joomcode/errorx"
)

var (
	SSHErrors                  = errorx.NewNamespace("ssh")
	CannotCreateSSHServerError = SSHErrors.NewType("cannot_create_ssh_server")
	CannotStopSSHServerError   = SSHErrors.NewType("cannot_stop_ssh_server")
)
