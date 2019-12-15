package encrypt

import (
	"crypto/sha256"
)

// EncryptionKey uses returns the sha256 hash of the username and password.
func EncryptionKey(username, password string) [32]byte {
	return sha256.Sum256([]byte(username + password))
}
