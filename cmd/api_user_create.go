package cmd

import (
	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var apiUserCreateCmd = &cobra.Command{
	Use:   "create [USERNAME]",
	Short: "Create a new user",
	Long:  `Create a new user with the given username and password. The user will be enabled by default.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		username := args[0]
		var password string
		err := huh.
			NewInput().
			EchoMode(huh.EchoModePassword).
			Prompt("Input password: ").
			Value(&password).
			WithTheme(huh.ThemeCatppuccin()).
			Run()
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotRunPromptError.WrapWithNoMessage(err))
		}

		var passwordRepeat string
		err = huh.
			NewInput().
			EchoMode(huh.EchoModePassword).
			Prompt("Repeat password: ").
			Value(&passwordRepeat).
			WithTheme(huh.ThemeCatppuccin()).
			Run()
		if err != nil {
			log.DefaultLogger().Fatal(errors.CannotRunPromptError.WrapWithNoMessage(err))
		}

		if password != passwordRepeat {
			log.DefaultLogger().Fatal("Passwords do not match")
		} else if len(password) == 0 {
			log.DefaultLogger().Fatal("Password cannot be empty")
		} else {
			_, err = db.RegisterUser(username, password)
			if err != nil {
				log.DefaultLogger().Fatal(err)
			}

			log.DefaultLogger().Info("User created", "username", username)
		}
	},
}

func init() {
	apiUserCmd.AddCommand(apiUserCreateCmd)
}
