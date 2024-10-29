package definitions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"github.com/spf13/viper"
)

const SpecVersion int = 1

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
	Version            int               `json:"version" yaml:"version"`                                            // Version is the version of the definition
	FileName           string            `yaml:"-" json:"-"`                                                        // FileName is the name of the file that contains the definition
	FullPath           string            `yaml:"-" json:"-"`                                                        // FullPath is the full path of the file that contains the definition
	Secrets            map[string]any    `yaml:"-" json:"-"`                                                        // Secrets are the secrets for this definition loaded from the secrets file
	IsEnabled          bool              `yaml:"-" json:"-"`                                                        // IsEnabled is a boolean that indicates if the endpoint is enabled (not contained in the definition)
	Name               string            `yaml:"name,omitempty" json:"name,omitempty" validate:"required"`          // Name is the name of the endpoint
	BaseUrl            string            `yaml:"base_url,omitempty" json:"baseUrl,omitempty" validate:"required"`   // BaseUrl is the base path of the endpoint
	Method             string            `yaml:"method,omitempty" json:"method,omitempty"`                          // Method is the name of the http-method to use for the request
	Payload            map[string]any    `yaml:"payload,omitempty" json:"payload,omitempty"`                        // Payload is the payload to use in case of a POST or PUT request
	Authorization      string            `yaml:"authorization,omitempty" json:"authorization,omitempty"`            // Authorization is the value to set for the authorization header
	JwtLogin           JwtLoginOptions   `yaml:"jwt_login,omitempty" json:"jwt_login,omitempty"`                    // JwtLogin are options to auto-get a login token for a request.
	Headers            map[string]string `yaml:"headers,omitempty" json:"headers,omitempty"`                        // Headers are additional headers to set for the request
	ExcludedValidators []string          `yaml:"excluded_validators,omitempty" json:"excludedValidators,omitempty"` // ExcludedValidators is a list of validators that should not be used for this endpoint
	QueryParameters    []QueryDefinition `yaml:"query_parameters,omitempty" json:"queryParameters,omitempty"`       // QueryParameters are all the query parameters that should be added to the call
	Format             string            `yaml:"format,omitempty" json:"format,omitempty" validate:"required"`      // Format is the response format of the
	Variables          []Variable        `yaml:"variables,omitempty" json:"variables,omitempty"`                    // Variables are all the variables that should be interpolated in the base url and the query parameters
	TestCaseNames      []string          `yaml:"test_case_names,omitempty" json:"test_case_names,omitempty"`
	OkCode             int               `yaml:"ok_code,omitempty" json:"ok_code,omitempty"`                                    // The expected status code
	ResponseSchema     any               `yaml:"response_schema,omitempty" json:"responseSchema,omitempty" validate:"required"` // ResponseSchema describes how the response should look like
}

// parseDefinition reads a given file and returns and EndpointDefinition.
// If the file could not be parsed the function returns an *errors.FileNotFoundError
func parseDefinition(filename string) (Endpoint, error) {
	definitionFile := locations.DefinitionExt(filename)
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
	definition.FullPath = definitionFile
	definition.IsEnabled = !strings.HasPrefix(filename, viper.GetString("daemon.ignore_prefix"))
	if definition.Authorization != "" || definition.JwtLogin.Url != "" {
		secretsFileName, _ := strings.CutSuffix(filename, ".apisensedef.yml")
		secretsFilePath := filepath.FromSlash(directories.DefinitionsDirectory() + "/" + secretsFileName + ".secrets.yml")
		secretsContent, err := os.ReadFile(secretsFilePath)
		if err != nil {
			return Endpoint{}, errors.FileNotFoundError.Wrap(err, "cannot read secrets file")
		}

		var secrets map[string]any
		err = yaml.Unmarshal(secretsContent, &secrets)
		if err != nil {
			return Endpoint{}, errors.CannotParseDefinitionFileError.Wrap(err, "cannot parse secrets file")
		}
		definition.Secrets = secrets
	}

	return definition, nil
}

// ValidateDefinition validates the definition and adds sensible defaults if needed
func ValidateDefinition(definition *Endpoint) error {
	if definition.Version != SpecVersion {
		return errors.InvalidVersionError.New("definition file: %s has an invalid version (%d). This apisense version only supports %d", definition.FileName, definition.Version, SpecVersion)
	}

	if strings.Contains(definition.Name, ".") || strings.Contains(definition.Name, "/") {
		return errors.InvalidCharacterError.New("definition name: %s contains invalid characters ('.', '/')", definition.Name)
	}

	if definition.BaseUrl == "" {
		return errors.NoBaseUrlError.New("definition file: %s has no base url", definition.FileName)
	}

	if definition.Format == "" {
		definition.Format = "json"
	} else if definition.Format != "json" && definition.Format != "xml" {
		return errors.InvalidFormatError.New("definition file: %s has invalid format. Expected either 'json' or 'xml', got %s", definition.FileName, definition.Format)
	}

	if len(definition.Variables) > 0 {
		firstVariableVar := util.FindFirst(definition.Variables, func(param Variable) bool { return !param.IsConstant })
		valueCount := 1
		if firstVariableVar != nil {
			valueCount = len(firstVariableVar.Values)
		}
		// check if any of the non-constant variables has a different length of values than the first non-constant variable
		for _, param := range definition.Variables {
			if !param.IsConstant && len(param.Values) != valueCount {
				return errors.VariableValueLengthMismatchError.New("variable %s has %d values, but %d are expected", param.Name, len(param.Values), valueCount)
			}
		}
		if definition.TestCaseNames == nil || len(definition.TestCaseNames) == 0 {
			definition.TestCaseNames = make([]string, 0)
			for i := range valueCount {
				definition.TestCaseNames = append(definition.TestCaseNames, fmt.Sprintf("TestCase%d", i+1))
			}
		}
		if len(definition.TestCaseNames) != valueCount {
			return errors.TestCaseNamesLengthMismatchError.New("test_case_names length (%d) does not match with the amount of test cases generated (%d)", len(definition.TestCaseNames), valueCount)
		}
	} else {
		if len(definition.TestCaseNames) == 0 {
			definition.TestCaseNames = []string{"TestCase1"}
		} else if len(definition.TestCaseNames) > 1 {
			return errors.TestCaseNamesLengthMismatchError.New("no variables set so only one test case will be generated")
		}
	}

	for k, v := range definition.Secrets {
		if _, ok := v.(string); !ok {
			return errors.InvalidSecretsError.New("secrets %s is not of type string", k)
		}
	}

	compiler := jsonschema.NewCompiler()
	err := compiler.AddResource("schema.json", definition.ResponseSchema)
	if err != nil {
		return errors.InvalidSchemaError.New("could not add resource to schema: %s", err.Error())
	}
	_, err = compiler.Compile("schema.json")
	if err != nil {
		return errors.InvalidSchemaError.New("schema is invalid: %s", err.Error())
	}

	return nil
}

// checkDuplicate checks whether there is a definition in definitions with the same name as definition
func checkDuplicate(definition Endpoint, definitions []Endpoint) error {
	for _, def := range definitions {
		if def.Name == definition.Name {
			return errors.DuplicateDefinitionError.New("duplicate definition name found: %s in file: %s", def.Name, definition.FileName)
		}
	}
	return nil
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
			err = checkDuplicate(definition, definitions)
			if err != nil {
				log.DaemonLogger().Error("Validation failed for definition. Skipping", "name", definition.Name, "reason", err.Error())
				continue
			}
			err = ValidateDefinition(&definition)
			if err != nil {
				log.DaemonLogger().Error("Validation failed for definition. Skipping", "name", definition.Name, "reason", err.Error())
				continue
			}
			definitions = append(definitions, definition)
		} else if !strings.HasSuffix(definitionFile.Name(), ".secrets.yml") {
			log.DaemonLogger().Error("Found file that is not a definition or a secrets file", "filename", definitionFile.Name())
		}
	}
	return definitions, nil
}
