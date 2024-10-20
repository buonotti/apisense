package locations

import (
	"path/filepath"

	"github.com/buonotti/apisense/filesystem/locations/directories"
)

func Definition(name string) string {
	return filepath.FromSlash(directories.DefinitionsDirectory() + name + "apisensedef.yml")
}
