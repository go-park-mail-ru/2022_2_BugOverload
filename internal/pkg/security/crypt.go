package security

import (
	"bytes"

	"golang.org/x/crypto/argon2"

	"go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	commonPkg "go-park-mail-ru/2022_2_BugOverload/pkg"
)

// genSalt generate slice of random bytes with SaltLength(Params.go) length
func genSalt() ([]byte, error) {
	salt, err := commonPkg.GenerateRandomBytes(pkg.SaltLength)

	if err != nil {
		return []byte{}, err
	}

	return salt, nil
}

// getHash return salt + hash using Argon2
func getHash(salt, plainPassword []byte) []byte {
	hashedPassword := argon2.IDKey(plainPassword, salt, 1, 32*1024, 4, 32)

	return append(salt, hashedPassword...)
}

// HashPassword generate password hash using Argon2
func HashPassword(plainPassword string) (string, error) {
	salt, err := genSalt()

	if err != nil {
		return "", err
	}

	hashedPassword := getHash(salt, []byte(plainPassword))

	return string(hashedPassword), nil
}

// IsPasswordsEqual return true if passwords equal, false otherwise
func IsPasswordsEqual(hashedPassword, plainPassword string) bool {
	if len(hashedPassword) < pkg.HashLength+pkg.SaltLength {
		return false
	}

	salt := hashedPassword[0:pkg.SaltLength]

	userPasswordHash := getHash([]byte(salt), []byte(plainPassword))

	return bytes.Equal([]byte(hashedPassword), userPasswordHash)
}
