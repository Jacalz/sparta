package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"golang.org/x/crypto/bcrypt"
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

// CorrectCredentials returns if the username and  password is correct or not.
func CorrectCredentials(username, password string, a fyne.App, w fyne.Window) bool {
	if !ValidInput(username, password, w) {
		return false
	}

	// Grab the password hash of the user.
	hash := a.Preferences().String("Username:" + username)

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}

// GenerateEncryptionKey uses returns the sha512/256 hash of the password to ensure that the encryption key is 32bytes long.
func GenerateEncryptionKey(password string) [32]byte {
	return sha512.Sum512_256([]byte(password))
}

// GeneratePasswordHash returns the hash and an eventual error
func GeneratePasswordHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}

	return string(hash)
}

// commonCipther holds common cipther code between encryption and decryption.
func commonCipher(key [32]byte) cipher.AEAD {
	// Generate a cipher to use. [:] is used to convert [N]array to []slice.
	c, err := aes.NewCipher(key[:])
	if err != nil {
		fyne.LogError("Error on generating new AES cipher", err)
		return nil
	}

	// Create a new Galois Counter Mode Cipher.
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fyne.LogError("Error on creating new GCM cipher", err)
		return nil
	}

	return gcm
}

// Encrypt is the initial encryption function.
func Encrypt(key *[32]byte, content []byte) []byte {
	// Make use of common code in encryption and decryption.
	gcm := commonCipher(*key)

	// Create a byte slice with the size of the nonce.
	nonce := make([]byte, gcm.NonceSize())

	// We then populate nonce with a crypto secure random sequence.
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fyne.LogError("Error on applying the nonce", err)
		return nil
	}

	// Do the actuall encryption and return the []byte slice.
	return gcm.Seal(nonce, nonce, content, nil)
}

// Decrypt is the initial decryption function.
func Decrypt(key *[32]byte, encrypted []byte) ([]byte, error) {
	// Make use of common code in encryption and decryption.
	gcm := commonCipher(*key)

	// Make sure that the nonceSize is bigger than the content.
	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		err := errors.New("The length of the encrypted content is longer than the nonceSize")
		return nil, err
	}

	// Unencrypt the content in to plaintext.
	plaintext, err := gcm.Open(nil, encrypted[:nonceSize], encrypted[nonceSize:], nil)
	if err != nil {
		fyne.LogError("Error on decrypting content", err)
		return nil, err
	}

	return plaintext, nil
}
