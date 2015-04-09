package gofuzz

import (
	"errors"
	"github.com/smashedtoatoms/gofuzz/utl"
)

// Overlap calculates the Overlap distance between two strings.
func Overlap(s1 string, s2 string, ngramSize int) (float32, error) {
	c1 := []rune(s1)
	c2 := []rune(s2)
	s1Size := len(c1)
	s2Size := len(c2)
	switch {
	case ngramSize <= 0 || s1Size < ngramSize || s2Size < ngramSize:
		return 0, errors.New("String is too small to calculate Overlap for " +
			"the given ngram size.")
	case s1 == s2:
		return 1.0, nil
	default:
		tokens1 := utl.Tokenize(runesToStrings(c1), ngramSize)
		tokens2 := utl.Tokenize(runesToStrings(c2), ngramSize)
		ms := float32(len(utl.Intersect(tokens1, tokens2)))
		return ms / float32(utl.MinInt([]int{len(tokens1), len(tokens2)})), nil
	}
}

/* Helper functions */

func runesToStrings(r []rune) []string {
	strings := make([]string, len(r))
	for i, val := range r {
		strings[i] = string(val)
	}
	return strings
}
