package validation

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"

	"github.com/buonotti/odh-data-monitor/errors"
)

func DefinitionsLocation() string {
	path := viper.GetString("daemon.definitions-dir")
	if strings.Contains(path, "~") {
		path = strings.Replace(path, "~", os.Getenv("HOME"), 1)
	}
	return filepath.FromSlash(path)
}

type endpointDefinition struct {
	Name            string
	BaseUrl         string           `toml:"base-url"`
	QueryParameters []queryParameter `toml:"query"`
	Method          string
	Format          string
	Variables       []variableSchema `toml:"variable"`
	ResultSchema    resultSchema     `toml:"result"`
}

type queryParameter struct {
	Name  string
	Value string
}

type resultSchema struct {
	Entries []SchemaEntry `toml:"entry"`
}

type variableSchema struct {
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

func parseDefinition(filename string) (endpointDefinition, error) {
	definitionContent, err := os.ReadFile(DefinitionsLocation() + string(filepath.Separator) + filename)
	if err != nil {
		return endpointDefinition{}, errors.FileNotFound.Wrap(err, "Cannot read definition file")
	}
	var definition endpointDefinition
	err = toml.Unmarshal(definitionContent, &definition)
	if err != nil {
		return endpointDefinition{}, errors.CannotParseDefinitionFile.Wrap(err, "Cannot parse definition file")
	}
	return definition, nil
}

func endpointDefinitions() ([]endpointDefinition, error) {
	definitionsFiles, err := os.ReadDir(DefinitionsLocation())
	if err != nil {
		return []endpointDefinition{}, errors.FileNotFound.Wrap(err, "Cannot read definitions directory")
	}
	var definitions []endpointDefinition
	for _, definitionFile := range definitionsFiles {
		definition, err := parseDefinition(definitionFile.Name())
		if err != nil {
			return []endpointDefinition{}, err
		}
		definitions = append(definitions, definition)
	}
	return definitions, nil
}
