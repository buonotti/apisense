package log

import (
	"github.com/sirupsen/logrus"
)

var DaemonLogger = logrus.WithFields(logrus.Fields{"system": "daemon"})
var DefaultLogger = logrus.WithFields(logrus.Fields{"system": "default"})
var SSHLogger = logrus.WithFields(logrus.Fields{"system": "ssh"})
var ApiLogger = logrus.WithFields(logrus.Fields{"system": "api"})
