package security

import (
	"bytes"

	stdErrors "github.com/pkg/errors"
	"golang.org/x/crypto/argon2"

	innerPKG "go-park-mail-ru/2022_2_BugOverload/internal/pkg"
	"go-park-mail-ru/2022_2_BugOverload/pkg"
)

// genSalt generate slice of random bytes with SaltLength(Params.go) length
func genSalt() ([]byte, error) {
	salt, err := pkg.GenerateRandomBytes(innerPKG.SaltLength)

	if err != nil {
		return []byte{}, err
	}

	return salt, nil
}

// getHash return salt + hash using Argon2
func getHash(salt, plainPassword []byte) []byte {
	hashedPassword := argon2.IDKey(plainPassword, salt, innerPKG.ArgonTime, innerPKG.ArgonMemory, innerPKG.ArgonThreads, innerPKG.ArgonKeyLength)

	return append(salt, hashedPassword...)
}

// HashPassword generate password hash using Argon2
func HashPassword(plainPassword string) (string, error) {
	salt, err := genSalt()

	if err != nil {
		return "", stdErrors.Wrap(err, "GenSalt falls")
	}

	hashedPassword := getHash(salt, []byte(plainPassword))

	return string(hashedPassword), nil
}

// IsPasswordsEqual return true if passwords equal, false otherwise
func IsPasswordsEqual(hashedPassword, plainPassword string) bool {
	salt := hashedPassword[0:innerPKG.SaltLength]

	userPasswordHash := getHash([]byte(salt), []byte(plainPassword))

	return bytes.Equal([]byte(hashedPassword), userPasswordHash)
}
