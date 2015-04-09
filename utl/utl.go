/*
  Utl contains utility fuctions used by the GoFuzz library.
*/
package utl

import (
	"regexp"
)

// DeDuplicate removes duplicates from a string (except for the letter c).
func DeDuplicate(s string) string {
	runes := []rune(s)
	if len(runes) <= 1 {
		return s
	}
	newVal := make([]rune, len(runes))
	newVal[0] = runes[0]
	lastCharIndex := 0
	for _, char := range runes[1:] {
		lastChar := newVal[lastCharIndex]
		// 99 below is rune number representing the lowercase letter c
		if char == 99 || char != lastChar {
			lastCharIndex = lastCharIndex + 1
			newVal[lastCharIndex] = char
		}
	}
	return string(newVal[0 : lastCharIndex+1])
}

// Determines if a string is alphabetic or not.
func IsAlphabetic(s string) bool {
	isAlphabetic, err := regexp.MatchString("[\\W0-9]", s)
	if err != nil {
		return false
	}
	return !isAlphabetic
}

// Intersect finds the intersection of two char slices.
func Intersect(s1 []string, s2 []string) []string {
	if len(s1) > len(s2) {
		temp := s1
		s1 = s2
		s2 = temp
	}
	intersect := make([]string, len(s1))
	intersectIndex := 0
	for _, char1 := range s1 {
		for _, char2 := range s2 {
			switch {
			case char1 == char2 && intersectIndex == 0:
				intersect[0] = char1
				intersectIndex = intersectIndex + 1
			case char1 == char2 && char1 != intersect[intersectIndex-1]:
				intersect[intersectIndex] = char1
				intersectIndex = intersectIndex + 1
			}
		}
	}
	return intersect[0:intersectIndex]
}

// Tokenize ngram tokenizes the provided string in groups of n.
func Tokenize(s1 []string, n int) []string {
	count := len(s1)
	if n <= 0 || count < n {
		return make([]string, 0)
	}
	tokenized := make([]string, len(s1)-(n-1))
	for i, j := 0, n; j <= count; i, j = i+1, j+1 {
		newString := ""
		strings := s1[i : i+(n)]
		for _, s := range strings {
			newString = newString + s
		}
		tokenized[i] = newString
	}
	return tokenized
}

// Contains checks if a slice of characters contains the specified character.
func Contains(vals []rune, val rune) bool {
	for _, s := range vals {
		if val == s {
			return true
		}
	}
	return false
}

// MaxInt finds the largest integer in a slice of ints and returns it.
func MaxInt(ints []int) int {
	biggest := ints[0]
	for _, i := range ints {
		if i > biggest {
			biggest = i
		}
	}
	return biggest
}

// MinInt finds the smallest integer in a slice of ints and returns it.
func MinInt(ints []int) int {
	smallest := ints[0]
	for _, i := range ints {
		if i < smallest {
			smallest = i
		}
	}
	return smallest
}
