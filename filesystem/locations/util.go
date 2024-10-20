package locations

import (
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
)

// Returns the full path of a definition given its filename without extension
func Definition(name string) string {
	return filepath.FromSlash(directories.DefinitionsDirectory() + "/" + name + "apisensedef.yml")
}

// Returns the full path of a definition given its filename with extension
func DefinitionExt(nameExt string) string {
	return filepath.FromSlash(directories.DefinitionsDirectory() + "/" + nameExt)
}
