package validation

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/validation/query"
	"github.com/buonotti/apisense/validation/response"
	"github.com/buonotti/apisense/validation/variables"
)

// DefinitionsLocation returns the location of the definitions directory
func DefinitionsLocation() string {
	return os.Getenv("HOME") + "/apisense/definitions"
}

// EndpointDefinition is the definition of an endpoint to test with all its query
// parameters, variables and its result schema
type EndpointDefinition struct {
	FileName           string
	Name               string                 `toml:"name"`                // Name is the name of the endpoint
	IsEnabled          bool                   `toml:"enabled"`             // IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)
	BaseUrl            string                 `toml:"base-url"`            // BaseUrl is the base path of the endpoint
	ExcludedValidators []string               `toml:"excluded-validators"` // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []query.Definition     `toml:"query"`               // QueryParameters are all the query parameters that should be added to the call
	Format             string                 `toml:"format"`              // Format is the response format of the
	Variables          []variables.Definition `toml:"variable"`            // Variables are all the variables that should be interpolated in the base url and the query parameters
	ResultSchema       response.Schema        `toml:"result"`              // ResultSchema describes how the response should look like
}

// parseDefinition reads a given file and returns and EndpointDefinition.
// If the file could not be parsed the function returns an *errors.FileNotFoundError
func parseDefinition(filename string) (EndpointDefinition, error) {
	definitionContent, err := os.ReadFile(DefinitionsLocation() + string(filepath.Separator) + filename)
	if err != nil {
		return EndpointDefinition{}, errors.FileNotFoundError.Wrap(err, "cannot read definition file")
	}

	var definition EndpointDefinition
	err = toml.Unmarshal(definitionContent, &definition)
	if err != nil {
		return EndpointDefinition{}, errors.CannotParseDefinitionFileError.Wrap(err, "cannot parse definition file")
	}
	definition.FileName = filename
	definition.IsEnabled = !strings.HasPrefix(filename, viper.GetString("daemon.ignore-prefix"))

	return definition, nil
}

// EndpointDefinitions uses parseDefinition to parse all the definitions found in
// the definitions/ directory. Directories and Files that start with the
// ignorePrefix are ignored.
func EndpointDefinitions() ([]EndpointDefinition, error) {
	definitionsFiles, err := os.ReadDir(DefinitionsLocation())
	if err != nil {
		return []EndpointDefinition{}, errors.FileNotFoundError.Wrap(err, "cannot read definitions directory")
	}
	var definitions []EndpointDefinition
	for _, definitionFile := range definitionsFiles {
		if definitionFile.IsDir() {
			continue
		}

		definition, err := parseDefinition(definitionFile.Name())
		if err != nil {
			return []EndpointDefinition{}, err
		}

		definitions = append(definitions, definition)
	}
	return definitions, nil
}
