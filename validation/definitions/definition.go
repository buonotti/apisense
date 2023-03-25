package definitions

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/validation/query"
)

// SchemaEntry is a field definition of the response
type SchemaEntry struct {
	Name       string        `yaml:"name"`     // Name is the name of the field
	Type       string        `yaml:"type"`     // Type is the type of the field
	Minimum    interface{}   `yaml:"min"`      // Minimum is the minimum allowed value of the field
	Maximum    interface{}   `yaml:"max"`      // Maximum is the maximum allowed value of the field
	IsRequired bool          `yaml:"required"` // Required is true if the field is required (not null or not empty in case of an array)
	Fields     []SchemaEntry `yaml:"fields"`   // Fields describe the children of this field if the field is an object or array
}

// Endpoint is the definition of an endpoint to test with all its query
// parameters, variables and its result schema
type Endpoint struct {
	FileName           string             `yaml:"-"`                   // FileName is the name of the file that contains the definition
	FullPath           string             `yaml:"-"`                   // FullPath is the full path of the file that contains the definition
	Name               string             `yaml:"name"`                // Name is the name of the endpoint
	IsEnabled          bool               `yaml:"enabled"`             // IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)
	BaseUrl            string             `yaml:"base_url"`            // BaseUrl is the base path of the endpoint
	ExcludedValidators []string           `yaml:"excluded_validators"` // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []query.Definition `yaml:"query_parameters"`    // QueryParameters are all the query parameters that should be added to the call
	Format             string             `yaml:"format"`              // Format is the response format of the
	Variables          []Variable         `yaml:"variables"`           // Variables are all the variables that should be interpolated in the base url and the query parameters
	ResultSchema       []SchemaEntry      `yaml:"result"`              // ResultSchema describes how the response should look like
}

// parseDefinition reads a given file and returns and EndpointDefinition.
// If the file could not be parsed the function returns an *errors.FileNotFoundError
func parseDefinition(filename string) (Endpoint, error) {
	definitionContent, err := os.ReadFile(filepath.FromSlash(directories.DefinitionsDirectory() + "/" + filename))
	if err != nil {
		return Endpoint{}, errors.FileNotFoundError.Wrap(err, "cannot read definition file")
	}

	if !strings.HasSuffix(filename, ".apisensedef.yaml") && !strings.HasSuffix(filename, ".apisensedef.yml") {
		return Endpoint{}, errors.CannotParseDefinitionFileError.Wrap(err, "found file that is not a definition file: "+filename)
	}

	var definition Endpoint
	err = yaml.Unmarshal(definitionContent, &definition)
	if err != nil {
		return Endpoint{}, errors.CannotParseDefinitionFileError.Wrap(err, "cannot parse definition file")
	}
	definition.FileName = filename
	definition.FullPath = filepath.FromSlash(directories.DefinitionsDirectory() + "/" + filename)
	definition.IsEnabled = !strings.HasPrefix(filename, viper.GetString("daemon.ignore_prefix"))

	return definition, nil
}

func validateDefinition(definitions []Endpoint, definition Endpoint) bool {
	for _, def := range definitions {
		if def.Name == definition.Name {
			log.DaemonLogger.Warnf("duplicate definition found: %s (%s)\n", definition.Name, definition.FileName)
			return false
		}
	}
	if definition.BaseUrl == "" {
		log.DaemonLogger.Errorf("definition %s (%s) has no base url\n", definition.Name, definition.FileName)
		return false
	}
	if definition.Format == "" {
		log.DaemonLogger.Errorf("definition %s (%s) has no format\n", definition.Name, definition.FileName)
		return false
	} else if definition.Format != "json" && definition.Format != "xml" {
		log.DaemonLogger.Errorf("definition %s (%s) has an invalid format: %s. Found %s expected either 'json' or 'xml'\n", definition.Name, definition.FileName, definition.Format, definition.Format)
		return false
	}
	if len(definition.ResultSchema) == 0 {
		log.DaemonLogger.Errorf("schema has no entries\n")
		return false
	}
	return true
}

// Endpoints uses parseDefinition to parse all the definitions found in
// the definitions/ directory. Directories and Files that start with the
// ignorePrefix are ignored.
func Endpoints() ([]Endpoint, error) {
	definitionsFiles, err := os.ReadDir(filepath.FromSlash(directories.DefinitionsDirectory()))
	if err != nil {
		return []Endpoint{}, errors.FileNotFoundError.Wrap(err, "cannot read definitions directory")
	}
	var definitions []Endpoint
	for _, definitionFile := range definitionsFiles {
		if definitionFile.IsDir() {
			continue
		}

		definition, err := parseDefinition(definitionFile.Name())
		if err != nil {
			return []Endpoint{}, err
		}
		if !validateDefinition(definitions, definition) {
			log.DaemonLogger.Errorf("validation failed for definition %s (%s). skipping", definition.Name, definition.FileName)
			continue
		}
		definitions = append(definitions, definition)
	}
	return definitions, nil
}
