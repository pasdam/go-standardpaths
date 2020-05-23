package standardpaths

import (
	"os"
	"path/filepath"
)

func appName() string {
	return filepath.Base(os.Args[0])
}
