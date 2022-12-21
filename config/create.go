package config

import (
	"os"
)

const (
	TEMPLATE = `
# This is the example configuration file for the ODH Data Monitor.

[log]
# Set the log level. Valid values are: debug, info, warn, error, fatal, panic
level = "info"

# Set the log file. Leave empty to log to stdout.
file = ""

# Enable pretty log output. This will colorize the log output and print it in a readable format.
# If set to false, the log output will be in JSON format.
pretty = true

# Set to true to force color logging. Only has an effect if pretty is set to true.
force-color = true
`
)

func create() error {
	err := os.Mkdir(configDir, 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, []byte(TEMPLATE), os.ModePerm)
}
