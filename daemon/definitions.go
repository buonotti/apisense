package daemon

import (
	"os"

	"github.com/buonotti/odh-data-monitor/config"
	"github.com/buonotti/odh-data-monitor/errors"
	"github.com/pelletier/go-toml/v2"
)

type EndpointDefinition struct {
	// Name of the endpoint
	Name string
	// URL of the endpoint
	Endpoint string
	// Method of the endpoint
	Method string
	// Parameters of the endpoint
	Parameters []ParameterDefinition
}

type ParameterDefinition struct {
	// Name of the parameter
	Name string
	// Values of the parameter
	Values []interface{}
}

func loadDefinition(filename string) (EndpointDefinition, error) {
	definition := EndpointDefinition{}
	definitionContent, err := os.ReadFile(config.Directory + "/definitions.d/"+filename)
	if err != nil {
		return EndpointDefinition{}, errors.FileNotFound.Wrap(err, "Cannot read definition file")
	}
	err = toml.Unmarshal(definitionContent, &definition)
	if err != nil {
		return EndpointDefinition{}, errors.CannotParseDefinitionFile.Wrap(err, "Cannot parse definition file")
	}
	return definition, nil
}

func Definitions() ([]EndpointDefinition, error) {
	definitionsFiles, err := os.ReadDir(config.Directory + "/definitions.d")
	if err != nil {
		return []EndpointDefinition{}, errors.FileNotFound.Wrap(err, "Cannot read definitions directory")
	}
	definitions := []EndpointDefinition{}
	for _, definitionFile := range definitionsFiles {
		definition, err := loadDefinition(definitionFile.Name())
		if err != nil {
			return []EndpointDefinition{}, err
		}
		definitions = append(definitions, definition)
	}
	return definitions, nil
}