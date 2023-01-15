package pkg

import (
	cryptoRand "crypto/rand"
	"encoding/base64"
	"math"
	"math/big"
)

func RandIntInInterval(max int, min int) int {
	number, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0
	}

	return int(number.Int64()) + min
}

func RandMaxInt(max int) int {
	number, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0
	}

	return int(number.Int64())
}

func bigInt(max int64) int64 {
	nBig, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return nBig.Int64()
}

const step = 53

func RandMaxFloat64(max float64, precision int) float64 {
	randFloat64 := (float64(bigInt(1<<step)) / (1 << step)) * max

	return math.Round(randFloat64*10*float64(precision)) / 10 * float64(precision)
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cryptoRand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// CryptoRandString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue. This should be used
// when there are concerns about security and need something
// cryptographically secure.
func CryptoRandString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func CryptoRandInInterval(max int, min int) int {
	if max == 0 {
		return 0
	}

	if min == 0 {
		return RandMaxInt(max)
	}

	return RandMaxInt(max-min) + min
}

func CryptoRandSequence(max int, min int) []int {
	length := max - min

	res := make([]int, length)

	inserted := make(map[int]bool)

	for i := 0; ; {
		try := CryptoRandInInterval(max, min)

		_, ok := inserted[try]
		if !ok {
			inserted[try] = true
			res[i] = try
			i++

			if try == max {
				max--
			}

			if try == min {
				min++
			}
		}

		if i == length {
			return res
		}
	}
}
