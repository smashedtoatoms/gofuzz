/*
  Phon contains all of the functions for doing phonetic comparisons.
*/
package phon

import (
	"errors"
	"github.com/smashedtoatoms/gofuzz/utl"
	"strings"
)

const (
	none = rune(0)
	ks   = rune(1)
)

// Metaphone returns the metaphone phonetic version of the provided string.
func Metaphone(val string) (string, error) {
	if len([]rune(val)) == 0 || !utl.IsAlphabetic(val) {
		return "", errors.New("String is empty or non-alphabetic.")
	}
	return transcode(
		transcodeFirstCharacter(
			utl.DeDuplicate(
				strings.ToLower(val)))), nil
}

// MetaphoneMetric compares two strings and returns a boolean of whether they
// match or not.
func MetaphoneMetric(val1 string, val2 string) (bool, error) {
	s1Size := len([]rune(val1))
	s2Size := len([]rune(val2))
	s1IsAlphabetic := utl.IsAlphabetic(val1)
	s2IsAlphabetic := utl.IsAlphabetic(val2)
	switch s1Size == 0 || !s1IsAlphabetic || s2Size == 0 || !s2IsAlphabetic {
	case false:
		phonetic1, _ := Metaphone(val1)
		phonetic2, _ := Metaphone(val2)
		return phonetic1 == phonetic2, nil
	default:
		return false, errors.New("Unable to Metaphone compare the two values.")
	}
}

/* Helper functions */

// transcode transcodes the string using the Metaphone algorithm.
func transcode(characters []rune) string {
	size := len(characters) // number of chars total
	iey := []rune{'i', 'e', 'y'}
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	pCount := 0 // number of chars processed
	output := make([]rune, size*2)
	oCount := 0 // number of chars in output

	// builds the new slice of characters and increments indexes.
	shift := func(count int, character rune) {
		switch character {
		case none:
			pCount = pCount + count
		case ks:
			output[oCount] = 'k'
			output[oCount+1] = 's'
			oCount = oCount + 2
			pCount = pCount + count
		default:
			output[oCount] = character
			oCount = oCount + 1
			pCount = pCount + count
		}
	}

	// returns the remaining number of characters to process.
	rSize := func() int {
		return size - (pCount + 1)
	}

	// returns the nth char off the beginning of the remaining characters
	// that need to be processed.
	rChar := func(n int) rune {
		i := size - rSize() + n
		return characters[i]
	}

	// returns the nth char off of the end of the characters that have been
	// processed thus far.
	pChar := func(n int) rune {
		if pCount < 1+n {
			return rune(0)
		}
		return characters[pCount-(1+n)]
	}

	for pCount < size {
		switch characters[pCount] {
		case 'a', 'e', 'i', 'o', 'u':
			switch {
			case pCount == 0:
				shift(1, characters[pCount])
			default:
				shift(1, none)
			}
		case 'f', 'j', 'l', 'm', 'n', 'r':
			shift(1, characters[pCount])
		case 'b':
			switch {
			case pCount >= 1 && pChar(0) == 'm' && rSize() == 0:
				shift(1, none)
			default:
				shift(1, 'b')
			}
		case 'c':
			switch {
			case rSize() >= 1 && rChar(0) == 'h' && pCount >= 1 &&
				pChar(0) == 's':
				shift(1, 'k')
			case rSize() >= 2 && rChar(0) == 'i' && rChar(1) == 'a':
				shift(3, 'x')
			case rSize() >= 1 && rChar(0) == 'h':
				shift(2, 'x')
			case pCount >= 1 && rSize() >= 1 && pChar(0) == 's' &&
				utl.Contains(iey, rChar(0)):
				shift(1, none)
			case rSize() >= 1 && utl.Contains(iey, rChar(0)):
				shift(1, 's')
			default:
				shift(1, 'k')
			}
		case 'd':
			switch {
			case rSize() >= 2 && rChar(0) == 'g' && utl.Contains(iey, rChar(1)):
				shift(1, 'j')
			default:
				shift(1, 't')
			}
		case 'g':
			switch {
			case (rSize() > 1 && rChar(0) == 'h') ||
				(rSize() == 1 && rChar(0) == 'n') ||
				(rSize() == 3 && rChar(0) == 'n' && rChar(2) == 'd'):
				shift(1, none)
			case rSize() >= 1 && utl.Contains(iey, rChar(0)):
				shift(2, 'j')
			default:
				shift(1, 'k')
			}
		case 'h':
			switch {
			case (pCount >= 1 && utl.Contains(vowels, pChar(0)) &&
				(rSize() == 0 || utl.Contains(vowels, rChar(0)))) ||
				pCount >= 2 && pChar(0) == 'h' && pChar(1) == 't' ||
				pChar(1) == 'g':
				shift(1, none)
			default:
				shift(1, 'h')
			}
		case 'k':
			switch {
			case pCount >= 1 && pChar(0) == 'c':
				shift(1, none)
			default:
				shift(1, 'k')
			}
		case 'p':
			switch {
			case rSize() >= 1 && rChar(0) == 'h':
				shift(3, 'f')
			default:
				shift(1, 'p')
			}
		case 'q':
			shift(1, 'k')
		case 's':
			switch {
			case rSize() >= 2 && rChar(0) == 'i' &&
				utl.Contains([]rune{'a', 'o'}, rChar(1)):
				shift(3, 'x')
			case rSize() >= 1 && rChar(0) == 'h':
				shift(2, 'x')
			default:
				shift(1, 's')
			}
		case 't':
			switch {
			case rSize() >= 2 && rChar(0) == 'i' &&
				utl.Contains([]rune{'a', 'o'}, rChar(1)):
				shift(3, 'x')
			case rSize() >= 1 && rChar(0) == 'h':
				shift(2, '0')
			case rSize() >= 2 && rChar(0) == 'c' && rChar(1) == 'h':
				shift(1, none)
			default:
				shift(1, 't')
			}
		case 'v':
			shift(1, 'f')
		case 'w', 'y':
			switch {
			case rSize() == 0 || !utl.Contains(vowels, rChar(0)):
				shift(1, none)
			default:
				shift(1, characters[pCount])
			}
		case 'x':
			shift(1, ks)
		case 'z':
			shift(1, 's')
		default:
			shift(1, none)
		}
	}
	return string(output[0:oCount])
}

// transcodeFirstCharacter transcodes the first character.
func transcodeFirstCharacter(s string) []rune {
	characters := []rune(s)
	switch len(characters) {
	case 0:
		return characters
	case 1:
		firstLetter := characters[0]
		switch firstLetter {
		case 'x':
			return []rune{'s'}
		default:
			return characters
		}
	default:
		letter1 := characters[0]
		letter2 := characters[1]
		restOfWord1 := characters[1:]
		restOfWord2 := characters[2:]
		switch letter1 {
		case 'a':
			switch letter2 {
			case 'e':
				return restOfWord1
			default:
				return characters
			}
		case 'g', 'k', 'p':
			switch letter2 {
			case 'n':
				return restOfWord1
			default:
				return characters
			}
		case 'w':
			switch letter2 {
			case 'r':
				return restOfWord1
			case 'h':
				return append([]rune{'w'}, restOfWord2...)
			default:
				return characters
			}
		case 'x':
			return append([]rune{'s'}, restOfWord1...)
		default:
			return characters
		}
	}
}
