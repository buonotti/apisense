package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/buonotti/apisense/api/db"
	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var apiUserCreateCmd = &cobra.Command{
	Use:   "create [USERNAME]",
	Short: "Create a new user",
	Long:  `Create a new user with the given username and password. The user will be enabled by default.`,
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		username := args[0]
		fmt.Print("Password: ")
		bytes, err := term.ReadPassword(int(os.Stdin.Fd()))
		errors.CheckErr(err)

		fmt.Print("\nConfirm password: ")
		bytesRepeat, err := term.ReadPassword(int(os.Stdin.Fd()))
		errors.CheckErr(err)

		password := strings.TrimSpace(string(bytes))
		passwordRepeat := strings.TrimSpace(string(bytesRepeat))

		fmt.Println()
		os.Stdout.Sync()

		if password != passwordRepeat {
			log.CliLogger.Error("passwords do not match")
		} else if len(password) == 0 {
			log.CliLogger.Error("password cannot be empty")
		} else {
			_, err = db.RegisterUser(username, password)
			errors.CheckErr(err)

			log.CliLogger.WithField("username", username).Info("user created")
		}
	},
}

func init() {
	apiUserCmd.AddCommand(apiUserCreateCmd)
}
