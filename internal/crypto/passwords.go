package crypto

import (
	"runtime"

	"github.com/Jacalz/sparta/internal/crypto/argon2"

	"fyne.io/fyne/v2"
)

// SpartaDefaults holds the default argon2 values. Based on official recommendations and suggestions.
var SpartaDefaults = argon2.Params{
	Time:    1,
	Memory:  64 * 1024,
	Threads: uint8(runtime.NumCPU()),
	KeyLen:  64,
}

// GeneratePasswordHash returns the key, storage hash and an eventual error
func GeneratePasswordHash(password string) ([]byte, string, error) {
	key, verific, err := argon2.GenerateFromPasswordAES256([]byte(password), SpartaDefaults)
	if err != nil {
		return []byte(""), "", err
	}

	return key, verific, err
}

// SaveNewPasswordHash saves the hash of the password.
func SaveNewPasswordHash(p, u string, a fyne.App) ([]byte, error) {
	// Get the encryption key and password verification string for the user.
	key, verific, err := GeneratePasswordHash(p)
	if err != nil {
		return nil, err
	}

	// Add the password hash for the user.
	a.Preferences().SetString("Username:"+u, verific)

	return key, nil
}
