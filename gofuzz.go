/*
  gofuzz is a collection of metrics and phonetic (fuzzy string matching)
  algorithms for Go.

  All of the phonetic functions are defined within the 'phon' subdirectory, and
  all of the similarity functions are defined within the 'sim' subdirectory.
  They are also set up to be called from the gofuzz package, whichever is
  easier.
*/
package gofuzz

import (
	"github.com/smashedtoatoms/gofuzz/phon"
	"github.com/smashedtoatoms/gofuzz/sim"
)

/* Similarity */

// Jaro calculates the Jaro distance between two strings.
func Jaro(s1 string, s2 string) (float32, error) {
	return sim.Jaro(s1, s2)
}

// JaroWinkler calculates the Jaro-Winkler distance between two strings.
func JaroWinkler(s1 string, s2 string) (float32, error) {
	return sim.JaroWinkler(s1, s2)
}

// Overlap calculates the Overlap distance between two strings.
//
// Note: the ngramSize is the number of letters to clump together when doing
// the comparison.  When in doubt, set it to 1.
func Overlap(s1 string, s2 string, ngramSize int) (float32, error) {
	return sim.Overlap(s1, s2, ngramSize)
}

/* Phonetic */

// Metaphone returns the metaphone phonetic version of the provided string.
func Metaphone(val string) (string, error) {
	return phon.Metaphone(val)
}

// MetaphoneMetric compares two strings and returns a boolean of whether they
// match or not.
func MetaphoneMetric(val1 string, val2 string) (bool, error) {
	return phon.MetaphoneMetric(val1, val2)
}
