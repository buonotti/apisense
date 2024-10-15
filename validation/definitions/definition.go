package definitions

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/viper"
)

// JwtLoginOptions is the definition of optional authentication options for jwt
type JwtLoginOptions struct {
	Url          string         `yaml:"url" json:"url" validate:"required"`   // Url is the url to the login endpoint
	LoginPayload map[string]any `yaml:"login_payload" json:"login_payload"`   // LoginPayload is the json or yml payload to send
	TokenKeyName string         `yaml:"token_key_name" json:"token_key_name"` // TokenKeyName is the name of the key in the response which contains the token
}

// QueryDefinition is a query parameter that should be added to the call
type QueryDefinition struct {
	Name  string `yaml:"name" json:"name" validate:"required"`   // Name is the name of the query parameter
	Value string `yaml:"value" json:"value" validate:"required"` // Value is the value of the query parameter
}

// Endpoint is the definition of an endpoint to test with all its query
// parameters, variables and its result schema
type Endpoint struct {
	FileName           string            `yaml:"-" json:"-"`                                                                    // FileName is the name of the file that contains the definition
	FullPath           string            `yaml:"-" json:"-"`                                                                    // FullPath is the full path of the file that contains the definition
	Secrets            map[string]string `yaml:"-" json:"-"`                                                                    // Secrets are the secrets for this definition loaded from the secrets file
	Name               string            `yaml:"name,omitempty" json:"name,omitempty" validate:"required"`                      // Name is the name of the endpoint
	IsEnabled          bool              `yaml:"enabled,omitempty" json:"enabled,omitempty"`                                    // IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)
	BaseUrl            string            `yaml:"base_url,omitempty" json:"baseUrl,omitempty" validate:"required"`               // BaseUrl is the base path of the endpoint
	Method             string            `yaml:"method,omitempty" json:"method,omitempty"`                                      // Method is the name of the http-method to use for the request
	Payload            map[string]any    `yaml:"payload,omitempty" json:"payload,omitempty"`                                    // Payload is the payload to use in case of a POST or PUT request
	Authorization      string            `yaml:"authorization,omitempty" json:"authorization,omitempty"`                        // Authorization is the value to set for the authorization header
	JwtLogin           JwtLoginOptions   `yaml:"jwt_login,omitempty" json:"jwt_login,omitempty"`                                // JwtLogin are options to auto-get a login token for a request.
	Headers            map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`                                    // Headers are additional headers to set for the request
	ExcludedValidators []string          `yaml:"excluded_validators,omitempty" json:"excludedValidators,omitempty"`             // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []QueryDefinition `yaml:"query_parameters,omitempty" json:"queryParameters,omitempty"`                   // QueryParameters are all the query parameters that should be added to the call
	Format             string            `yaml:"format,omitempty" json:"format,omitempty" validate:"required"`                  // Format is the response format of the
	Variables          []Variable        `yaml:"variables,omitempty" json:"variables,omitempty"`                                // Variables are all the variables that should be interpolated in the base url and the query parameters
	OkCode             int               `yaml:"ok_code,omitempty" json:"ok_code,omitempty"`                                    // The expected status code
	ResponseSchema     map[string]any    `yaml:"response_schema,omitempty" json:"responseSchema,omitempty" validate:"required"` // ResponseSchema describes how the response should look like
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

// validateDefinitions validates the definitions to check if they are correct
func validateDefinitions(definitions []Endpoint, definition Endpoint) bool {
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

	compiler := jsonschema.NewCompiler()
	err := compiler.AddResource("schema.json", definition.ResponseSchema)
	_, err = compiler.Compile("schema.json")
	if err != nil {
		log.DaemonLogger().Error("Definition has an invalid schema: "+err.Error(), "name", definition.Name)
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
			if !validateDefinitions(definitions, definition) {
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
