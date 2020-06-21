package validate

import (
	"crypto/subtle"
	"errors"
	"regexp"

	"github.com/Jacalz/sparta/internal/crypto/argon2"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
)

// usermatch holds a precompiled regex rule for verifying usernames
var usermatch = regexp.MustCompile(`^\w+$`)

// Input checks if the inputed username and passwords are valid and creates a dialog message inside the window if they are not.
func Input(username, password string, w fyne.Window) bool {
	if subtle.ConstantTimeCompare([]byte(username), []byte(password)) == 1 {
		dialog.ShowInformation("Identical username and password", "The username and password can't be identical.", w)
	} else if subtle.ConstantTimeCompare([]byte(password), nil) == 1 || username == "" {
		dialog.ShowInformation("Missing username/password", "Please provide both a username and a password.", w)
	} else if len(password) < 8 {
		dialog.ShowInformation("Too short password", "The password should be eight characters or longer.", w)
	} else if !usermatch.MatchString(username) {
		dialog.ShowInformation("Invalid username", "The username needs to be a single word with valid word characters.", w)
	} else {
		return true
	}

	return false
}

// CorrectCredentials returns if the username and password is correct or not.
func CorrectCredentials(username, password string, a fyne.App, w fyne.Window) (key []byte, err error) {
	if !Input(username, password, w) {
		return nil, errors.New("username or password is invalid")
	}

	// Grab the password hash of the user.
	hash := a.Preferences().String("Username:" + username)

	key, err = argon2.CompareHashAndPasswordAES256(hash, []byte(password))
	if err != nil {
		dialog.ShowInformation("Wrong username and/or password", "The login credentials are incorrect, please try again.", w)
		return nil, err
	}

	return key, nil
}
