package locations

import (
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
)

// Returns the full path of a definition given it's filename without extension
func Definition(name string) string {
	return filepath.FromSlash(directories.DefinitionsDirectory() + "/" + name + "apisensedef.yml")
}
