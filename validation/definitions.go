package validation

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"

	"github.com/buonotti/odh-data-monitor/errors"
)

func DefinitionsLocation() string {
	return filepath.FromSlash(viper.GetString("daemon.definitions-dir"))
}

type EndpointDefinition struct {
	Name            string
	BaseUrl         string           `toml:"base-url"`
	QueryParameters []QueryParameter `toml:"query"`
	Method          string
	Format          string
	Variables       []VariableSchema `toml:"variable"`
	ResultSchema    ResultSchema     `toml:"result"`
}

type QueryParameter struct {
	Name  string
	Value string
}

type ResultSchema struct {
	Entries []SchemaEntry `toml:"entry"`
}

type VariableSchema struct {
	Name       string
	IsConstant bool `toml:"constant"`
	Values     []string
}

type SchemaEntry struct {
	Name         string
	Type         string
	Minimum      interface{}   `toml:"min"`
	Maximum      interface{}   `toml:"max"`
	IsRequired   bool          `toml:"required"`
	ChildEntries []SchemaEntry `toml:"fields"`
}

func parseDefinition(filename string) (EndpointDefinition, error) {
	definitionContent, err := os.ReadFile(DefinitionsLocation() + string(filepath.Separator) + filename)
	if err != nil {
		return EndpointDefinition{}, errors.FileNotFound.Wrap(err, "Cannot read definition file")
	}
	var definition EndpointDefinition
	err = toml.Unmarshal(definitionContent, &definition)
	if err != nil {
		return EndpointDefinition{}, errors.CannotParseDefinitionFile.Wrap(err, "Cannot parse definition file")
	}
	return definition, nil
}

func endpointDefinitions() ([]EndpointDefinition, error) {
	definitionsFiles, err := os.ReadDir(DefinitionsLocation())
	if err != nil {
		return []EndpointDefinition{}, errors.FileNotFound.Wrap(err, "Cannot read definitions directory")
	}
	var definitions []EndpointDefinition
	for _, definitionFile := range definitionsFiles {
		definition, err := parseDefinition(definitionFile.Name())
		if err != nil {
			return []EndpointDefinition{}, err
		}
		definitions = append(definitions, definition)
	}
	return definitions, nil
}
