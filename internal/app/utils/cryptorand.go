package utils

import (
	crand "crypto/rand"
	"encoding/base64"
	"math/big"
)

func Rand(max int) int {
	number, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}

	return int(number.Int64())
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// CryptoString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue. This should be used
// when there are concerns about security and need something
// cryptographically secure.
func CryptoString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
