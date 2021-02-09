package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"fyne.io/fyne/v2"
)

// commonCipther holds common cipther code between encryption and decryption.
func commonCipher(key []byte) cipher.AEAD {
	c, err := aes.NewCipher(key)
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
func Encrypt(key *[]byte, content []byte) []byte {
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
func Decrypt(key *[]byte, encrypted []byte) ([]byte, error) {
	// Make use of common code in encryption and decryption.
	gcm := commonCipher(*key)

	// Make sure that the nonceSize is bigger than the content.
	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		return nil, errors.New("the length of the encrypted content is longer than the nonceSize")
	}

	// Unencrypt the content into plaintext.
	plaintext, err := gcm.Open(nil, encrypted[:nonceSize], encrypted[nonceSize:], nil)
	if err != nil {
		fyne.LogError("Error on decrypting content", err)
		return nil, err
	}

	return plaintext, nil
}
