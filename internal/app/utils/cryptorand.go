package utils

import (
	"crypto/rand"
	"math/big"
)

func Rand(max int) int {
	number, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}

	return int(number.Int64())
}
