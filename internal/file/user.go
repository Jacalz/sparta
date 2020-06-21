package file

import (
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne"
)

// NoExistingUsers tells us if any users exist.
func NoExistingUsers() bool {
	f, err := os.Open(ConfigDir())
	if err != nil {
		fyne.LogError("Error on opening the config dir", err)
		return false
	}

	defer f.Close() // #nosec - We are not writing to the file.

	_, err = f.Readdirnames(1)

	return err == io.EOF
}

// CreateNewUser creates our new user.
func CreateNewUser(username string) error {
	if _, err := os.Stat(ConfigDir()); os.IsNotExist(err) {
		if err := os.Mkdir(ConfigDir(), os.ModePerm); err != nil {
			fyne.LogError("Error on creating config directory", err)
			return err
		}
	} else if err != nil {
		return err
	}

	_, err := os.Create(filepath.Join(ConfigDir(), username+"-exercises.json"))
	if err != nil {
		fyne.LogError("Error on creating user file", err)
		return err
	}

	return nil
}

// ExistingUser tells us if the user exists or not.
func ExistingUser(username string) bool {
	if _, err := os.Stat(filepath.Join(ConfigDir(), username+"-exercises.json")); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}

	return false
}
