/*
  Utl contains utility fuctions used by the GoFuzz library.
*/
package utl

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

// DeDuplicate removes duplicates from a string (except for the letter c).
func DeDuplicate(s string) string {
	if Size(s) <= 1 {
		return s
	}
	newVal := string(s[0])
	for _, char := range s[1:] {
		newValLastLetter, _ := utf8.DecodeLastRuneInString(newVal)
		c, _ := utf8.DecodeLastRuneInString("c")
		if char == c || char != newValLastLetter {
			newVal = newVal + string(char)
		}
	}
	return newVal
}

// Determines if a string is alphabetic or not.
func IsAlphabetic(s string) bool {
	isAlphabetic, _ := regexp.MatchString("[\\W0-9]", s)
	return !isAlphabetic
}

// Intersect finds the intersection of two strings.
func Intersect(s1 string, s2 string) string {
	short := s1
	long := s2
	if Size(s1) > Size(s2) {
		long = s1
		short = s2
	}
	intersect := ""
	for _, char1 := range short {
		for _, char2 := range long {
			lastRune, _ := utf8.DecodeLastRuneInString(intersect)
			if char1 == char2 && char1 != lastRune {
				intersect = intersect + string(char1)
			}
		}
	}
	return intersect
}

// Tokenize ngram tokenizes the provided string in groups of n.
func Tokenize(s string, n int) []string {
	count := Size(s)
	if n <= 0 || count < n {
		return make([]string, 0)
	}
	tokenized := make([]string, len(s)-(n-1))
	for i, j := 0, n; j <= count; i, j = i+1, j+1 {
		tokenized[i] = string(s[i : i+(n)])
	}
	return tokenized
}

// GetLetter gets the specified letter (via an index i) from a string, and
// returns it as well as the rest of the string following the specified letter.
func GetLetter(s string, i int) (string, string, error) {
	size := Size(s)
	if size <= i || i < 0 {
		return s, "", errors.New("String index out of range.")
	}
	return string([]rune(s)[i]), string([]rune(s)[i+1 : size]), nil
}

// FirstLetter gets the last letter from the specified string.
func FirstLetter(s string) string {
	size := Size(s)
	if size < 1 {
		return ""
	}
	return string([]rune(s)[0:1])
}

// LastLetter gets the last letter from the specified string.
func LastLetter(s string) string {
	size := Size(s)
	if size < 1 {
		return ""
	}
	return string([]rune(s)[size-1 : size])
}

// SplitAt splits a string at the specified index i into two strings.
func SplitAt(s string, i int) (string, string, error) {
	size := Size(s)
	if size <= i || i < 0 {
		return s, "", errors.New("String index out of range.")
	}
	return string([]rune(s)[0:i]), string([]rune(s)[i:size]), nil
}

// Contains checks an array of strings contains the specified string.
func Contains(vals []string, val string) bool {
	for _, s := range vals {
		if val == s {
			return true
		}
	}
	return false
}

// Size gets the current size of the string.
func Size(s string) int {
	return utf8.RuneCountInString(s)
}
