package main

import "errors"

// ErrNegativValue defines the error returned by a generator
// if a negative value is passed as param for a Generate() implementation.
var ErrNegativValue = errors.New("could not generate a string of negative size")

// Generator defines methods that should
// be implemented by any generator implementation.
type Generator interface {
	Generate(size int) (string, error)
}
