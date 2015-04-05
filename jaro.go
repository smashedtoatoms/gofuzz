package gofuzz

import (
	"errors"
	"github.com/smashedtoatoms/gofuzz/utl"
)

// Jaro calculates the Jaro distance between two strings.
func Jaro(s1 string, s2 string) (float32, error) {
	s1Size := utl.Size(s1)
	s2Size := utl.Size(s2)
	switch {
	case s1Size == 0 || s2Size == 0:
		return 0, errors.New("Unable to calculate Jaro against empty string.")
	case s1 == s2:
		return 1.0, nil
	case s1Size > s2Size:
		return match(s2, s1), nil
	default:
		return match(s1, s2), nil
	}
	return 0.0, nil
}

func match(s1 string, s2 string) float32 {
	return 0.0
}
