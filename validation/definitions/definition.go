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
)

// SchemaEntry is a field definition of the response
type SchemaEntry struct {
	Name       string        `yaml:"name" json:"name" validate:"required"`         // Name is the name of the field
	Type       string        `yaml:"type" json:"type" validate:"required"`         // Type is the type of the field
	Minimum    interface{}   `yaml:"min" json:"min"`                               // Minimum is the minimum allowed value of the field
	Maximum    interface{}   `yaml:"max" json:"max"`                               // Maximum is the maximum allowed value of the field
	IsRequired bool          `yaml:"required" json:"required" validate:"required"` // Required is true if the field is required (not null or not empty in case of an array)
	Fields     []SchemaEntry `yaml:"fields" json:"fields" validate:"required"`     // Fields describe the children of this field if the field is an object or array
}

type JwtLoginOptions struct {
	Url          string         `yaml:"url" json:"url" validate:"required"`   // Url is the url to the login endpoint
	LoginPayload map[string]any `yaml:"login_payload" json:"login_payload"`   // LoginPayload is the json or yml payload to send
	TokenKeyName string         `yaml:"token_key_name" json:"token_key_name"` // TokenKeyName is the name of the key in the response which contains the token
}

// Definition is a query parameter that should be added to the call
type QueryDefinition struct {
	Name  string `yaml:"name" json:"name" validate:"required"`   // Name is the name of the query parameter
	Value string `yaml:"value" json:"value" validate:"required"` // Value is the value of the query parameter
}

// Endpoint is the definition of an endpoint to test with all its query
// parameters, variables and its result schema
type Endpoint struct {
	FileName           string            `yaml:"-" json:"-"`                                                // FileName is the name of the file that contains the definition
	FullPath           string            `yaml:"-" json:"-"`                                                // FullPath is the full path of the file that contains the definition
	Secrets            map[string]string `yaml:"-" json:"-"`                                                // Secrets are the secrets for this definition loaded from the secrets file
	Name               string            `yaml:"name" json:"name" validate:"required"`                      // Name is the name of the endpoint
	IsEnabled          bool              `yaml:"enabled" json:"enabled"`                                    // IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)
	BaseUrl            string            `yaml:"base_url" json:"baseUrl" validate:"required"`               // BaseUrl is the base path of the endpoint
	Method             string            `yaml:"method" json:"method"`                                      // Method is the name of the http-method to use for the request
	Payload            map[string]any    `yaml:"payload" json:"payload"`                                    // Payload is the payload to use in case of a POST or PUT request
	Authorization      string            `yaml:"authorization" json:"authorization"`                        // Authorization is the value to set for the authorization header
	JwtLogin           JwtLoginOptions   `yaml:"jwt_login" json:"jwt_login"`                                // JwtLogin are options to auto-get a login token for a request.
	Headers            map[string]string `yaml:"headers" json:"headers"`                                    // Headers are additional headers to set for the request
	ExcludedValidators []string          `yaml:"excluded_validators" json:"excludedValidators"`             // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []QueryDefinition `yaml:"query_parameters" json:"queryParameters"`                   // QueryParameters are all the query parameters that should be added to the call
	Format             string            `yaml:"format" json:"format" validate:"required"`                  // Format is the response format of the
	Variables          []Variable        `yaml:"variables" json:"variables"`                                // Variables are all the variables that should be interpolated in the base url and the query parameters
	OkCode             int               `yaml:"ok_code" json:"ok_code"`                                    // The expected status code
	ResponseSchema     []SchemaEntry     `yaml:"response_schema" json:"responseSchema" validate:"required"` // ResponseSchema describes how the response should look like
}

// parseDefinition reads a given file and returns and EndpointDefinition.
// If the file could not be parsed the function returns an *errors.FileNotFoundError
func parseDefinition(filename string) (Endpoint, error) {
	definitionFile := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + filename)
	definitionContent, err := os.ReadFile(definitionFile)
	if err != nil {
		return Endpoint{}, errors.FileNotFoundError.Wrap(err, "cannot read definition file")
	}

	var definition Endpoint
	err = yaml.Unmarshal(definitionContent, &definition)
	if err != nil {
		return Endpoint{}, errors.CannotParseDefinitionFileError.Wrap(err, "cannot parse definition file")
	}
	definition.FileName = filename
	definition.FullPath = filepath.FromSlash(directories.DefinitionsDirectory() + "/" + filename)
	definition.IsEnabled = !strings.HasPrefix(filename, viper.GetString("daemon.ignore_prefix"))
	if definition.Authorization != "" || definition.JwtLogin.Url != "" {
		secretsFileName, _ := strings.CutSuffix(filename, ".apisensedef.yml")
		secretsFilePath := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + secretsFileName + ".secrets.yml")
		secretsContent, err := os.ReadFile(secretsFilePath)
		if err != nil {
			return Endpoint{}, errors.FileNotFoundError.Wrap(err, "cannot read secrets file")
		}

		var secrets map[string]string
		err = yaml.Unmarshal(secretsContent, &secrets)
		if err != nil {
			return Endpoint{}, errors.CannotParseDefinitionFileError.Wrap(err, "cannot parse secrets file")
		}
		definition.Secrets = secrets
	}

	return definition, nil
}

func validateDefinition(definitions []Endpoint, definition Endpoint) bool {
	for _, def := range definitions {
		if def.Name == definition.Name {
			log.DaemonLogger().Error("Duplicate definition found", "name", definition.Name, "filename", definition.FileName)
			return false
		}
	}
	if definition.BaseUrl == "" {
		log.DaemonLogger().Error("Definition has no base url", "name", definition.Name, "filename", definition.FileName)
		return false
	}
	if definition.Format == "" {
		log.DaemonLogger().Error("Definition has no format", "name", definition.Name, "filename", definition.FileName)
		return false
	} else if definition.Format != "json" && definition.Format != "xml" {
		log.DaemonLogger().Error("Definition has an invalid format. Expected either 'json' or 'xml'", "name", definition.Name, "filename", definition.FileName)
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

		if strings.HasSuffix(definitionFile.Name(), ".apisensedef.yml") {
			definition, err := parseDefinition(definitionFile.Name())
			if err != nil {
				return []Endpoint{}, err
			}
			if !validateDefinition(definitions, definition) {
				log.DaemonLogger().Error("Validation failed for definition. skipping", "name", definition.Name, "filename", definition.FileName)
				continue
			}
			definitions = append(definitions, definition)
		} else if !strings.HasSuffix(definitionFile.Name(), ".secrets.yml") {
			log.DaemonLogger().Error("Found file that is not a definition or a secrets file", "filename", definitionFile.Name())
		}
	}
	return definitions, nil
}
