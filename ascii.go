package main

import (
	"crypto/rand"
	"math/big"
)

// asciiStart define the first printable
// index in the ascii table.
const asciiStart = 33

// asciiEnd define the last printable
// character in the ascii table.
const asciiEnd = 93

// ASCII is the ascii generator.
// Could be used to generate string containing
// only ascii printable characters.
type ASCII struct{}

// NewASCII returns a new ASCII generator.
func NewASCII() Generator {
	return ASCII{}
}

func generateASCIIChar() (byte, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(asciiEnd))
	if err != nil {
		return 0, err
	}

	c := asciiStart + r.Int64()

	return byte(c), nil
}

// Generate generates a string of size length using only
// printable ascii characters.
func (a ASCII) Generate(length int) (string, error) {

	if length < 0 {
		return "", ErrNegativValue
	}

	res := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		r, err := generateASCIIChar()
		if err != nil {
			return "", err
		}

		res = append(res, r)
	}

	return string(res), nil
}
