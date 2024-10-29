package validators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/buonotti/apisense/errors"
	"github.com/buonotti/apisense/filesystem/locations/directories"
	"github.com/buonotti/apisense/log"
	"github.com/buonotti/apisense/util"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
)

type FindExecFunc func(string) (string, error)

var knownProjectMap map[string]FindExecFunc = map[string]FindExecFunc{
	"go.mod": func(s string) (string, error) {
		goModFile := filepath.FromSlash(directories.ValidatorCustomDirectory() + "/" + s + "/go.mod")
		contents, err := os.ReadFile(goModFile)
		if err != nil {
			return "", errors.CannotReadFileError.WrapWithNoMessage(err)
		}
		spl := strings.Split(string(contents), "\n")
		if len(spl) == 0 {
			return "", errors.GoModFileEmptyError.New("go.mod file is empty: %s", goModFile)
		}
		modSpl := strings.Split(spl[0], " ")
		if len(modSpl) != 2 {
			return "", errors.ModuleLineMalformedError.New("line 1: expected two tokens separated by space got something else")
		}
		slashSpl := strings.Split(modSpl[1], "/")
		return slashSpl[len(slashSpl)-1], nil
	},
	"Cargo.toml": func(s string) (string, error) {
		cargoTomlFile := filepath.FromSlash(directories.ValidatorCustomDirectory() + "/" + s + "/Cargo.toml")
		contents, err := os.ReadFile(cargoTomlFile)
		if err != nil {
			return "", errors.CannotReadFileError.WrapWithNoMessage(err)
		}
		var cargoToml map[string]any
		err = toml.Unmarshal(contents, &cargoToml)
		if err != nil {
			return "", errors.CannotUnmarshalCargoTomlError.WrapWithNoMessage(err)
		}

		return fmt.Sprintf("target/release/%s", cargoToml["package"].(map[string]any)["name"]), nil
	},
}

func AutoDiscoverExternal(override bool) error {
	dirEntries, err := os.ReadDir(filepath.FromSlash(directories.ValidatorCustomDirectory()))
	if err != nil {
		return errors.CannotReadDirectoryError.WrapWithNoMessage(err)
	}

	discovered := make([]ValidatorDefinition, 0)

	for _, dirEntry := range dirEntries {
		log.DefaultLogger().Debug("Try to get project type for folder", "folder", dirEntry.Name())
		projEntries, err := os.ReadDir(filepath.FromSlash(directories.ValidatorCustomDirectory() + "/" + dirEntry.Name()))
		if err != nil {
			return errors.CannotReadDirectoryError.WrapWithNoMessage(err)
		}
		for _, projFile := range projEntries {
			if f, ok := knownProjectMap[projFile.Name()]; ok {
				path, err := f(dirEntry.Name())
				if err != nil {
					return errors.CannotDetermineValidatorPathError.Wrap(err, "cannot find exec name for validator %s", dirEntry.Name())
				}

				discovered = append(discovered, ValidatorDefinition{
					Name:  dirEntry.Name(),
					Path:  filepath.FromSlash(directories.ValidatorCustomDirectory() + "/" + path),
					Args:  make([]string, 0),
					Fatal: false,
				})
				log.DefaultLogger().Debug("Found project type", "folder", dirEntry.Name(), "file", projFile.Name())
			}
		}

	}

	externalValidators, err := parse()
	if err != nil {
		return err
	}

	for _, val := range discovered {
		exists := util.FindFirst(externalValidators, func(def ValidatorDefinition) bool {
			return def.Name == val.Name
		})
		if exists == nil || override {
			externalValidators = append(externalValidators, val)
		} else {
			log.DefaultLogger().Warn("Validator with the same name already exists. Not overriding", "name", val.Name)
		}
	}

	viper.Set("validation.external_validators", externalValidators)
	err = viper.WriteConfig()
	if err != nil {
		return errors.CannotWriteConfigError.WrapWithNoMessage(err)
	}

	return nil
}
