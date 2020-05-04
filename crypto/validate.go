package crypto

import (
	"sparta/crypto/argon2"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
)

// ValidInput checks if the inputed username and passwords are valid and creates a message if they are not.
func ValidInput(username, password string, w fyne.Window) (valid bool) {
	if username == "" || password == "" {
		dialog.ShowInformation("Missing username/password", "Please provide both username and password.", w)
	} else if username == password {
		dialog.ShowInformation("Identical username and password", "Use separate usernames and passwords.", w)
	} else if len(password) < 8 {
		dialog.ShowInformation("Too short password", "The password should be eight characters or longer.", w)
	} else {
		valid = true
	}

	return valid
}

// CorrectCredentials returns if the username and password is correct or not.
func CorrectCredentials(username, password string, a fyne.App, w fyne.Window) (key []byte, err error) {
	if !ValidInput(username, password, w) {
		return nil, err
	}

	// Grab the password hash of the user.
	hash := a.Preferences().String("Username:" + username)

	key, err = argon2.CompareHashAndPasswordAES256(hash, []byte(password))
	if err != nil {
		fyne.LogError("Error in comparison of hashes", err)
		return nil, err
	}

	return key, nil
}
