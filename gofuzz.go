/*
  gofuzz is a collection of metrics and phonetic (fuzzy string matching)
  algorithms for Go.
*/
package gofuzz

/* Similarity Algorithms */

// Jaro calculates the Jaro distance between two strings.
func Jaro(val1 string, val2 string) float32 {
	return 0.0
}

// JaroWinkler calculates the Jaro-Winkler distance between two strings.
func JaroWinkler(val1 string, val2 string) float32 {
	return 0.0
}

// Overlap calculates the Overlap distance between two strings.
func Overlap(val1 string, val2 string, ngramSize int) float32 {
	return 0.0
}

/* Phonetic Algorithms */

// Metaphone returns the metaphone phonetic version of the provided string.
func Metaphone(val string) string {
	return ""
}

// MetaphoneMetric compares two strings and returns a boolean of whether they
// match or not.
func MetaphoneMetric(val1 string, val2 string) bool {
	return false
}
