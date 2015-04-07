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

/* Helper functions */

func match(s1 string, s2 string) float32 {
	max := utl.Size(s2) / 2
	s1Size := utl.Size(s1)
	s2Size := utl.Size(s2)
	commons := 0
	transpositions := 0
	previous := -1
	for i := 0; i < s1Size; i = i + 1 {
		char1, _, _ := utl.GetLetter(s1, i)
		from := utl.MaxInt([]int{0, i - max})
		to := utl.MinInt([]int{s2Size, i + max})
	SecondWordLoop:
		for from < to {
			char2, _, _ := utl.GetLetter(s2, from)
			switch {
			case char1 == char2 && previous != -1 && from < previous:
				commons = commons + 1
				transpositions = transpositions + 1
				previous = from
				from = from + 1
				break SecondWordLoop
			case char1 == char2:
				previous = from
				commons = commons + 1
				from = from + 1
				break SecondWordLoop
			default:
				from = from + 1
			}
		}
	}
	switch {
	case commons == 0:
		return 0.0
	default:
		return ((float32(commons) / float32(s1Size)) +
			(float32(commons) / float32(s2Size)) +
			((float32(commons) - float32(transpositions)) / float32(commons))) / 3.0
	}
}
