package log

import (
	"os"

	"github.com/buonotti/apisense/v2/util"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func hasLogFile() bool {
	return logFile != nil
}

var logFile *os.File = nil

// CloseLogFile closes the log file if one is in use
func CloseLogFile() error {
	if logFile != nil {
		err := logFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// Setup reads the config and configures log level and log output of all loggers
func Setup() error {
	logFilePath := viper.GetString("log.file")
	if logFilePath != "" {
		osFile, err := os.OpenFile(viper.GetString("log.file"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
		if err != nil {
			return err
		}
		logFile = osFile
	}

	lvl, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		return err
	}

	log.SetLevel(lvl)
	log.SetTimeFormat(util.ApisenseTimeFormat)
	log.SetReportCaller(log.GetLevel() == log.DebugLevel)
	log.SetOutput(getWriter())

	return nil
}
