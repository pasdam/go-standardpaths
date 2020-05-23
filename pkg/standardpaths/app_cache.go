package standardpaths

import (
	"os"
	"path/filepath"
)

// AppCache returns the path of the folder where to store the app cache
func AppCache() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, appName(), "cache"), nil
}
