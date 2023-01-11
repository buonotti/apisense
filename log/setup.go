package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Setup reads the config and configures log level and log output of all loggers
func Setup() error {
	// read configs
	doPrettyLog := viper.GetBool("daemon.log.pretty")
	forceColorLog := viper.GetBool("daemon.log.force-color")
	logLevel := viper.GetString("daemon.log.level")
	logFileName := viper.GetString("daemon.log.file")

	// set logger type
	if doPrettyLog {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   forceColorLog,
			DisableColors: logFileName != "",
			PadLevelText:  true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	// parse the level from the config into a logrus.Level
	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)

	// set the log output
	if logFileName != "" {
		logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}

		logrus.SetOutput(logFile)
	}

	return nil
}
