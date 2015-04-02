/*
  Utl contains utility fuctions used by the GoFuzz library.
*/
package utl

import (
	"regexp"
	"unicode/utf8"
)

// DeDuplicate removes duplicates from a string (except for the letter c).
func DeDuplicate(s string) string {
	if utf8.RuneCountInString(s) <= 1 {
		return s
	}
	newVal := string(s[0])
	for _, char := range s {
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
	if utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2) {
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
	count := utf8.RuneCountInString(s)
	if n <= 0 || count < n {
		return make([]string, 0)
	}
	tokenized := make([]string, len(s)-(n-1))
	for i, j := 0, n; j <= count; i, j = i+1, j+1 {
		tokenized[i] = string(s[i : i+(n)])
	}
	return tokenized
}
