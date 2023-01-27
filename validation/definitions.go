package validation

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

// DefinitionsLocation returns the location of the definitions directory
func DefinitionsLocation() string {
	return os.Getenv("HOME") + "/apisense/definitions"
}

// EndpointDefinition is the definition of an endpoint to test with all its query
// parameters, variables and its result schema
type EndpointDefinition struct {
	FileName           string
	Name               string           `toml:"name"`                // Name is the name of the endpoint
	BaseUrl            string           `toml:"base-url"`            // BaseUrl is the base path of the endpoint
	ExcludedValidators []string         `toml:"excluded-validators"` // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []queryParameter `toml:"query"`               // QueryParameters are all the query parameters that should be added to the call
	Format             string           `toml:"format"`              // Format is the response format of the
	Variables          []variableSchema `toml:"variable"`            // Variables are all the variables that should be interpolated in the base url and the query parameters
	ResultSchema       resultSchema     `toml:"result"`              // ResultSchema describes how the response should look like
}

// queryParameter is a query parameter that should be added to the call
type queryParameter struct {
	Name  string // Name is the name of the query parameter
	Value string // Value is the value of the query parameter
}

// resultSchema describes how the response should look like
type resultSchema struct {
	Entries []SchemaEntry `toml:"entry"` // Entries are all the field definitions of the response
}

// SchemaEntry is a field definition of the response
type SchemaEntry struct {
	Name         string        `toml:"name"`     // Name is the name of the field
	Type         string        `toml:"type"`     // Type is the type of the field
	Minimum      interface{}   `toml:"min"`      // Minimum is the minimum allowed value of the field
	Maximum      interface{}   `toml:"max"`      // Maximum is the maximum allowed value of the field
	IsRequired   bool          `toml:"required"` // Required is true if the field is required (not null or not empty in case of an array)
	ChildEntries []SchemaEntry `toml:"fields"`   // ChildEntries describe the children of this field if the field is an object or array
}

// variableSchema describes a variable that should be interpolated in the base url and the query parameters
type variableSchema struct {
	Name       string   `toml:"name"`     // Name is the name of the variable
	IsConstant bool     `toml:"constant"` // IsConstant is true if the value of the variable is constant or else false
	Values     []string `toml:"values"`   // Values are all the possible values of the variable (only 1 in case of a constant)
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
		if definitionFile.IsDir() || strings.HasPrefix(definitionFile.Name(), viper.GetString("daemon.ignore-prefix")) {
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
