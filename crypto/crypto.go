package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"

	"fyne.io/fyne"
)

// Hash uses returns the sha512/256 hash of the username and password.
func Hash(username, password string) [32]byte {
	return sha512.Sum512_256([]byte(username + password))
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
