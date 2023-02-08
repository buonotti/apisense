package log

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/util"
)

// Setup reads the config and configures log level and log output of all loggers
func Setup() error {
	doPrettyLog := viper.GetBool("daemon.log.pretty")
	forceColorLog := viper.GetBool("daemon.log.force-color")
	logLevel := viper.GetString("daemon.log.level")
	logFileName := viper.GetString("daemon.log.file")

	hasLogFile := logFileName != ""

	if doPrettyLog {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   forceColorLog,
			DisableColors: hasLogFile,
			PadLevelText:  true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	lvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(lvl)

	if hasLogFile {
		path := logFileName
		util.ExpandHome(&path)
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		cobra.CheckErr(err) // TODO find better fix

		logrus.SetOutput(logFile)
	}

	return nil
}
