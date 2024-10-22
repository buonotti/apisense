package config

import (
	"github.com/spf13/viper"
)

func setupDefaults() {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", "")

	viper.SetDefault("daemon.interval", "* * * * *")
	viper.SetDefault("daemon.run_on_startup", true)
	viper.SetDefault("daemon.ignore_prefix", "_")
	viper.SetDefault("daemon.discard.enabled", true)
	viper.SetDefault("daemon.discard.max_lifetime", "720h")
	viper.SetDefault("daemon.notification.enabled", false)
	viper.SetDefault("daemon.notification.only_on_error", true)
	viper.SetDefault("daemon.notification.sender", "")
	viper.SetDefault("daemon.notification.receiver", "")
	viper.SetDefault("daemon.notification.smtp_server", "")
	viper.SetDefault("daemon.notification.username", "")
	viper.SetDefault("daemon.notification.password", "")
	viper.SetDefault("daemon.notification.smtp_port", 587)
	viper.SetDefault("daemon.rpc", true)

	viper.SetDefault("ssh.host", "")
	viper.SetDefault("ssh.port", 23232)

	viper.SetDefault("tui.refresh", 10)

	viper.SetDefault("api.host", "")
	viper.SetDefault("api.port", 8080)
	viper.SetDefault("api.auth", true)
	viper.SetDefault("api.swagger", true)
	viper.SetDefault("api.signing_key", "")

	viper.SetDefault("validation.external_validators", []any{})
}
