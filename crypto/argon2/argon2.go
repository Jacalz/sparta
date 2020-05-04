package argon2

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne"
	"golang.org/x/crypto/argon2"
)

// Params controls structural values for the argon2 implementation.
type Params struct {
	Salt    []byte
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

// randSalt generates a random salt value using crypto/rand.
func randSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		fyne.LogError("Error on generating random salt", err)
		return nil, err
	}

	return salt, nil
}

// encode takes the hash and parameters to encode a string with it all (not in accordance to the argon2 specs).
func encode(hash []byte, p Params) string {
	return fmt.Sprintf("$%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Time, p.Threads,
		base64.RawStdEncoding.EncodeToString(p.Salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
}

// decode takes the encoded hash string and extracts hash and parameters from it.
func decode(encoded string) (hashed []byte, p Params, err error) {
	data := strings.Split(encoded, "$")
	if len(data) != 5 {
		return nil, p, errors.New("Invalid hash")
	}

	if version, err := strconv.ParseInt(data[1], 10, 0); err != nil {
		return nil, p, err
	} else if version != argon2.Version {
		return nil, p, errors.New("Invalid argon2 version")
	}

	_, err = fmt.Sscanf(data[2], "m=%d,t=%d,p=%d", &p.Memory, &p.Time, &p.Threads)
	if err != nil {
		fyne.LogError("Error on scanning parameters", err)
		return nil, p, err
	}

	p.Salt, err = base64.RawStdEncoding.DecodeString(data[3])
	if err != nil {
		fyne.LogError("Error on decoding base64 salt", err)
		return nil, p, err
	}

	hashed, err = base64.RawStdEncoding.DecodeString(data[4])
	if err != nil {
		fyne.LogError("Error on decoding base64 hash", err)
		return nil, p, err
	}

	p.KeyLen = 2 * uint32(len(hashed))

	return hashed, p, nil
}

// GenerateFromPasswordAES256 generates a hash and returns the first part for use as an AES-256 key, the second part for storage and lastly an error.
func GenerateFromPasswordAES256(password []byte, p Params) (key []byte, verific string, err error) {
	p.Salt, err = randSalt()
	if err != nil {
		return nil, "", err
	}

	hash := argon2.IDKey(password, p.Salt, p.Time, p.Memory, p.Threads, p.KeyLen)

	return hash[:p.KeyLen/2], encode(hash[p.KeyLen/2:], p), nil
}

// CompareHashAndPasswordAES256 generates a hash from the password with the same parameters as the hash and then compares the second part of it.
func CompareHashAndPasswordAES256(encoded string, password []byte) (key []byte, err error) {
	stored, p, err := decode(encoded)
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey(password, p.Salt, p.Time, p.Memory, p.Threads, p.KeyLen)

	if subtle.ConstantTimeCompare(stored, hash[p.KeyLen/2:]) != 1 {
		return nil, errors.New("The hashes are not equal")
	}

	return hash[:p.KeyLen/2], nil
}
