package file

import (
	"os"
	"path/filepath"

	"fyne.io/fyne"
)

// ConfigDir returns the config directory where files are being stored.
func ConfigDir() string {
	// Get the config directory of the user.
	dir, err := os.UserConfigDir()
	if err != nil {
		fyne.LogError("Error on reading config directory", err)
	}

	return filepath.Join(dir, "com.github.jacalz.sparta")
}
