package security

import (
	"bytes"
	"encoding/base64"
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
	hashedPassword := argon2.IDKey(plainPassword, salt, pkg.ArgonTime, pkg.ArgonMemory, pkg.ArgonThreads, pkg.ArgonKeyLength)

	return append(salt, hashedPassword...)
}

// HashPassword generate password hash using Argon2
func HashPassword(plainPassword string) (string, error) {
	salt, err := genSalt()

	if err != nil {
		return "", err
	}

	hashedPassword := getHash(salt, []byte(plainPassword))

	return base64.StdEncoding.EncodeToString(hashedPassword), nil
}

// IsPasswordsEqual return true if passwords equal, false otherwise
func IsPasswordsEqual(hashedPassword, plainPassword string) bool {
	hashDecoded, _ := base64.StdEncoding.DecodeString(hashedPassword)

	salt := hashDecoded[0:pkg.SaltLength]

	userPasswordHash := getHash(salt, []byte(plainPassword))

	return bytes.Equal([]byte(hashDecoded), userPasswordHash)
}
