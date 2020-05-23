package standardpaths

import (
	"os"
	"path/filepath"
)

// AppData returns the path of the folder where to store the app data
func AppData() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, appName(), "data"), nil
}
