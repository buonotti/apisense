package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/spf13/cobra"
)

var definitionBoilerplate = `
# The name of the endpoint (used in the report)
name: "%s"

# The format of the response (json or xml)
# format: "json"

# The endpoint to call. Variables can be interpolated with {{ .ParamName }} (Go template syntax)
# There are also two builtin functions: Now and Env
# Now is used with the following syntax: {{ .Now "<format>" }} you can find the format here: https://pkg.go.dev/time#Time.Format
# Env is used with the following syntax: {{ .Env "<env_var_name>" }}
# base_url: 'https://exampleapi.com/v{{ .ApiVersion }}'

# List of names of validators that should not be run. Keep in mind that external validators are named with external.<name>
# to better symbolize that they are external
# excluded_validators:
    # - "range"

# The query parameters to pass to the endpoint. Variables can be interpolated with {{ .ParamName }} (Go template syntax)
# query_parameters:
  # - name: "limit"
    # value: "-1"

# Variable definitions
# For all non-constant variables, the number of values must be the same and will be used in order to call the endpoint
# variables:
  # - name: "ApiVersion"
    # constant: true
    # values:
      # - "2"

# The expected result
# response_schema:
  # The name of the field that contains the data
  # - name:
    # The type of the field (integer, string, array, object)
    # type: "number"
    # The minimum value of the field (only for integer and float)
    # min: "none"
    # The maximum value of the field (only for integer and float)
    # max: "none"
    # Whether the field is required or not (in case of an array, it means that the array must not be empty)
    # required: true
    # The fields of the object (only for object and array)
    # fields: []
`

var definitionCreateCmd = &cobra.Command{
	Use:   "create [NAME]",
	Short: "Creates a new definition file",
	Long:  `Creates a new definition file with the needed boilerplate and the given name`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fileName := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + name + ".apisensedef.yml")
		err := os.WriteFile(fileName, []byte(fmt.Sprintf(definitionBoilerplate, name)), os.ModePerm)
		errors.CheckErr(err)
		log.ApiLogger.Infof("definition file %s successfully created", fileName)
	},
}

func init() {
	definitionCmd.AddCommand(definitionCreateCmd)
}
