package util

import (
	"errors"
	"math/big"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func permutation(n, k int) *big.Int {
	numerator := factorial(n)
	denominator := factorial(n - k)
	return new(big.Int).Div(numerator, denominator)
}

func RandomCode(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateCodeSet(length, setSize int) ([]string, error) {
	permutations := permutation(len(letterRunes), length)

	if setSize > int(permutations.Int64()) {
		return nil, errors.New("setSize is greater than the number of possible permutations")
	}

	set := make(map[string]bool)
	for len(set) < setSize {
		code := RandomCode(length)
		if !set[code] {
			set[code] = true
		}
	}

	uniqueCodes := make([]string, setSize)
	i := 0
	for code := range set {
		uniqueCodes[i] = code
		i++
	}

	return uniqueCodes, nil
}
