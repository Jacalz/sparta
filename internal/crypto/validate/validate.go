package validate

import (
	"errors"
	"regexp"

	"github.com/Jacalz/sparta/internal/crypto/argon2"
	passwordvalidator "github.com/lane-c-wagner/go-password-validator"

	"fyne.io/fyne"
)

// usermatch holds a precompiled regex rule for verifying usernames
var usermatch = regexp.MustCompile(`^\w+$`)

const passwordEntropy = 60

// Username validates the received username
func Username(username string) error {
	if !usermatch.MatchString(username) {
		return errors.New("username contains invalid characters")
	}

	return nil
}

// Password validates the received password
func Password(password string) error {
	if err := passwordvalidator.Validate(password, passwordEntropy); err != nil {
		return err
	}

	return nil
}

// Credentials returns if the username and password are correct or not.
func Credentials(username, password string, a fyne.App, w fyne.Window) (key []byte, err error) {
	// Grab the password hash of the user.
	hash := a.Preferences().String("Username:" + username)

	key, err = argon2.CompareHashAndPasswordAES256(hash, []byte(password))
	if err != nil {
		return nil, err
	}

	return key, nil
}

// NewRegexp returns a new regexp validator
func NewRegexp(regex *regexp.Regexp, reason string) fyne.StringValidator {
	return func(input string) error {
		if regex != nil && !regex.MatchString(input) {
			return errors.New(reason)
		}

		return nil // Nothing to validate with, same as having no validator.
	}
}
