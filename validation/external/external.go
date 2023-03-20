package external

import (
	"github.com/spf13/viper"

	"github.com/buonotti/apisense/errors"
)

// ValidatorDefinition is the definition of an external validator
type ValidatorDefinition struct {
	Name          string   // Name is the name of the validator
	Path          string   // Path is the path to the executable
	Args          []string // Args are the arguments to pass to the executable
	ReadFromStdin bool     // ReadFromStdin controls whether the validator expects the item to validate on stdin
	Fatal         bool     // Fatal controls whether the validator is fatal or not that is if it fails the pipeline should stop
}

// ExitCode is the definition of an exit code
type ExitCode struct {
	Code        int64  // Code is the exit code of the validator
	Ok          bool   // Ok sets whether the exit code is considered as a pass or not
	Description string // Description is the description in the report that should be put alongside the error if the item fails to validate
}

// Parse parses the external validators in the config file and returns a slice containing all validators to later use in the pipeline
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

		args, err := parseArgs(obj["args"])
		if err != nil {
			return nil, err
		}

		validators[i] = ValidatorDefinition{
			Name:          obj["name"].(string),
			Path:          obj["path"].(string),
			Args:          args,
			ReadFromStdin: obj["read-from-stdin"].(bool),
			Fatal:         obj["fatal"].(bool),
		}
	}
	return validators, nil
}

func parseArgs(i interface{}) ([]string, error) {
	arr, isArray := i.([]interface{})
	if !isArray {
		return nil, errors.ExternalValidatorParseError.New("cannot parse external validator. expected []interface{}, got %T", i)
	}

	if len(arr) == 0 {
		return []string{}, nil
	}

	args := make([]string, len(arr))
	for _, elem := range arr {
		if _, isString := elem.(string); !isString {
			return nil, errors.ExternalValidatorParseError.New("cannot parse external validator. expected []string, got []%T", elem)
		}
		args = append(args, elem.(string))
	}

	return args, nil
}
