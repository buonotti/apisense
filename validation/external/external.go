package external

import (
	"github.com/spf13/viper"

	"github.com/buonotti/odh-data-monitor/errors"
)

type ValidatorDefinition struct {
	Name          string
	Path          string
	Args          []string
	ReadFromStdin bool
	ExitCodes     []ExitCode
}

type ExitCode struct {
	Code        int
	Ok          bool
	Description string
}

func Parse() ([]ValidatorDefinition, error) {
	object := viper.Get("validation.external")
	if object == nil {
		return []ValidatorDefinition{}, nil
	}
	arr, isArray := object.([]interface{})
	if !isArray {
		return nil, errors.ExternalValidatorParseError.New("cannot parse external validators. Expected []any, got %T", object)
	}
	validators := make([]ValidatorDefinition, len(arr))
	for i, arrayEntry := range arr {
		obj, isStringMap := arrayEntry.(map[string]interface{})
		if !isStringMap {
			return nil, errors.ExternalValidatorParseError.New("cannot parse external validators. Expected map[string]any, got %T", arrayEntry)
		}
		exitCodes, err := parseExitCodes(obj["exitCodes"])
		if err != nil {
			return nil, err
		}
		validators[i] = ValidatorDefinition{
			Name:          obj["name"].(string),
			Path:          obj["path"].(string),
			Args:          obj["args"].([]string),
			ReadFromStdin: obj["read-from-stdin"].(bool),
			ExitCodes:     exitCodes,
		}
	}
	return validators, nil
}

func parseExitCodes(object interface{}) ([]ExitCode, error) {
	arr, isArray := object.([]interface{})
	if !isArray {
		return nil, errors.ExternalValidatorParseError.New("cannot parse external validators. Expected []any, got %T", object)
	}
	exitCodes := make([]ExitCode, len(arr))
	for i, arrayEntry := range arr {
		obj, isStringMap := arrayEntry.(map[string]interface{})
		if !isStringMap {
			return nil, errors.ExternalValidatorParseError.New("cannot parse external validators. Expected map[string]any, got %T", arrayEntry)
		}
		exitCodes[i] = ExitCode{
			Code:        obj["code"].(int),
			Ok:          obj["ok"].(bool),
			Description: obj["description"].(string),
		}
	}
	return exitCodes, nil
}
