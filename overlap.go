package gofuzz

import (
	"errors"
	"github.com/smashedtoatoms/gofuzz/utl"
)

// Overlap calculates the Overlap distance between two strings.
func Overlap(s1 string, s2 string, ngramSize int) (float32, error) {
	chars1 := []rune(s1)
	chars2 := []rune(s2)
	s1Size := len(chars1)
	s2Size := len(chars2)
	switch {
	case ngramSize <= 0 || s1Size < ngramSize || s2Size < ngramSize:
		return 0, errors.New("String is too small to calculate Overlap for " +
			"the given ngram size.")
	case s1 == s2:
		return 1.0, nil
	default:
		strings1 := make([]string, s1Size)
		strings2 := make([]string, s2Size)
		for i, val := range chars1 {
			strings1[i] = string(val)
		}
		for i, val := range chars2 {
			strings2[i] = string(val)
		}
		tokens1 := utl.Tokenize(strings1, ngramSize)
		tokens2 := utl.Tokenize(strings2, ngramSize)
		ms := float32(len(utl.Intersect(tokens1, tokens2)))
		return ms / float32(utl.MinInt([]int{len(tokens1), len(tokens2)})), nil
	}
}
