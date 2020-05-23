package standardpaths

import (
	"os"
	"path/filepath"
)

// AppConfig returns the path of the folder where to store the app config
func AppConfig() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, appName(), "config"), nil
}
