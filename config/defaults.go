package config

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func setupDefaults() {
	viper.SetDefault("APISENSE_EMAIL_USER", "")
	viper.SetDefault("APISENSE_EMAIL_PASS", "")

	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", "")
	viper.SetDefault("daemon.interval", "* * * * *")
	viper.SetDefault("daemon.run_on_startup", true)
	viper.SetDefault("daemon.ignore_prefix", "_")
	viper.SetDefault("daemon.discard.enabled", true)
	viper.SetDefault("daemon.discard.max_lifetime", "720h")
	viper.SetDefault("daemon.notification.enabled", false)
	viper.SetDefault("daemon.notification.only_on_error", true)
	viper.SetDefault("daemon.notification.receiver", "")
	viper.SetDefault("daemon.notification.smtp_server", "")
	viper.SetDefault("daemon.notification.smtp_port", 587)
	viper.SetDefault("ssh.host", "")
	viper.SetDefault("ssh.port", 23232)
	viper.SetDefault("tui.refresh", 10)
	viper.SetDefault("api.host", "")
	viper.SetDefault("api.port", 8080)
	viper.SetDefault("api.key", uuid.New().String())
	viper.SetDefault("api.auth", true)
	viper.SetDefault("validation.excluded_builtin_validators", []string{})
}
