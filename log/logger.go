package log

import (
	"github.com/sirupsen/logrus"
)

// DaemonLogger is the logger used to log all daemon related messages
var DaemonLogger = logrus.WithFields(logrus.Fields{"system": "daemon"})

// DefaultLogger is used to log general purpose messages
var DefaultLogger = logrus.WithFields(logrus.Fields{"system": "default"})

// SSHLogger is used by the ssh server to log messages and is also used in log.WishMiddleware
var SSHLogger = logrus.WithFields(logrus.Fields{"system": "ssh"})

// ApiLogger is used to log all api related messages
var ApiLogger = logrus.WithFields(logrus.Fields{"system": "api"})
