package gofuzz

import (
	"errors"
	"github.com/smashedtoatoms/gofuzz/utl"
	"log"
	"strconv"
	"strings"
)

// Metaphone returns the metaphone phonetic version of the provided string.
func Metaphone(s string) (string, error) {
	if utl.Size(s) == 0 || !utl.IsAlphabetic(s) {
		return "", errors.New("String is empty or non-alphabetic.")
	}
	return transcode(
		transcodeFirstCharacter(
			utl.DeDuplicate(
				strings.ToLower(s)))), nil
}

// MetaphoneMetric compares two strings and returns a boolean of whether they
// match or not.
func MetaphoneMetric(s1 string, s2 string) bool {
	return false
}

/* Helper functions */

// transcode transcodes the string using the Metaphone algorithm.
func transcode(s string) string {
	iey := []string{"i", "e", "y"}
	vowels := []string{"a", "e", "i", "o", "u"}
	var processed string
	var output string

	shift := func(count int, character string) {
		head, tail, _ := utl.SplitAt(s, count)
		output = output + character
		processed = processed + head
		s = tail
	}

	for utl.Size(s) > 0 {
		current, remainder, err := utl.GetLetter(s, 0)
		if err != nil {
			return output
		}
		switch current {
		case "a", "e", "i", "o", "u":
			pSize := utl.Size(processed)
			if pSize == 0 {
				shift(1, current)
			} else {
				shift(1, "")
			}
		case "f", "j", "l", "m", "n", "r":
			shift(1, current)
		case "b":
			pSize := utl.Size(processed)
			pLast := utl.LastLetter(processed)
			rSize := utl.Size(remainder)
			if pSize >= 1 && pLast == "m" && rSize == 0 {
				shift(1, "")
			} else {
				shift(1, "b")
			}
		case "c":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			pSize := utl.Size(processed)
			pLast := utl.LastLetter(processed)
			rSecond, _, _ := utl.GetLetter(remainder, 1)
			rHasIEY := utl.Contains(iey, rFirst)
			if rSize >= 1 && rFirst == "h" && pSize >= 1 && pLast == "s" {
				shift(1, "k")
			} else if rSize >= 2 && rFirst == "i" && rSecond == "a" {
				shift(3, "x")
			} else if rSize >= 1 && rFirst == "h" {
				shift(2, "x")
			} else if pSize >= 1 && rSize >= 1 && pLast == "s" && rFirst == "h" {
				shift(2, "x")
			} else if pSize >= 1 && rSize >= 1 && pLast == "s" && rHasIEY {
				shift(1, "")
			} else if rSize >= 1 && rHasIEY {
				shift(1, "s")
			} else {
				shift(1, "k")
			}
		case "d":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			rSecond, _, _ := utl.GetLetter(remainder, 1)
			rSecondHasIEY := utl.Contains(iey, rSecond)
			if rSize >= 2 && rFirst == "g" && rSecondHasIEY {
				shift(1, "j")
			} else {
				shift(1, "t")
			}
		case "g":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			rThird, _, _ := utl.GetLetter(remainder, 2)
			rFirstHasIEY := utl.Contains(iey, rFirst)
			if (rSize > 1 && rFirst == "h") || (rSize == 1 && rFirst == "n") ||
				(rSize == 3 && rFirst == "n" && rThird == "d") {
				shift(1, "")
			} else if rSize >= 1 && rFirstHasIEY {
				shift(2, "j")
			} else {
				shift(1, "k")
			}
		case "h":
			pSize := utl.Size(processed)
			pLast := utl.LastLetter(processed)
			pLastInVowels := utl.Contains(vowels, pLast)
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			rFirstInVowels := utl.Contains(vowels, rFirst)
			pSecondToLast, _, _ := utl.GetLetter(processed, utl.Size(processed)-2)
			if (pSize >= 1 && pLastInVowels && rSize == 0 || rFirstInVowels) ||
				pSize >= 2 && pLast == "h" && pSecondToLast == "t" ||
				pSecondToLast == "g" {
				shift(1, "")
			} else {
				shift(1, "h")
			}
		case "k":
			pSize := utl.Size(processed)
			pLast := utl.LastLetter(processed)
			if pSize >= 1 && pLast == "c" {
				shift(1, "")
			} else {
				shift(1, "k")
			}
		case "p":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			if rSize >= 1 && rFirst == "h" {
				shift(3, "f")
			} else {
				shift(1, "p")
			}
		case "q":
			shift(1, "k")
		case "s":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			rSecond, _, _ := utl.GetLetter(remainder, 1)
			rSecondInAo := utl.Contains([]string{"a", "o"}, rSecond)
			if rSize >= 2 && rFirst == "i" && rSecondInAo {
				shift(3, "x")
			} else if rSize >= 1 && rFirst == "h" {
				shift(2, "x")
			} else {
				shift(1, "s")
			}
		case "t":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			rSecond, _, _ := utl.GetLetter(remainder, 1)
			rSecondInAo := utl.Contains([]string{"a", "o"}, rSecond)
			if rSize >= 2 && rFirst == "i" && rSecondInAo {
				shift(3, "x")
			} else if rSize >= 1 && rFirst == "h" {
				shift(2, "0")
			} else if rSize >= 2 && rFirst == "c" && rSecond == "h" {
				shift(1, "")
			} else {
				shift(1, "t")
			}
		case "v":
			shift(1, "f")
		case "w", "y":
			rSize := utl.Size(remainder)
			rFirst := utl.FirstLetter(remainder)
			rFirstInVowels := utl.Contains(vowels, rFirst)
			if rSize == 0 || !rFirstInVowels {
				shift(1, "")
			} else {
				shift(1, current)
			}
		case "x":
			shift(1, "ks")
		case "z":
			shift(1, "s")
		default:
			shift(1, "")
		}
	}
	return output
}

// transcodeFirstCharacter transcodes the first character.
func transcodeFirstCharacter(s string) string {
	switch utl.Size(s) {
	case 0:
		return s
	case 1:
		firstLetter, _, _ := utl.GetLetter(s, 0)
		switch firstLetter {
		case "x":
			return "s"
		default:
			return s
		}
	default:
		letter1, restOfWord1, _ := utl.GetLetter(s, 0)
		letter2, restOfWord2, _ := utl.GetLetter(s, 1)
		switch letter1 {
		case "a":
			switch letter2 {
			case "e":
				return restOfWord1
			default:
				return s
			}
		case "g", "k", "p":
			switch letter2 {
			case "n":
				return restOfWord1
			default:
				return s
			}
		case "w":
			switch letter2 {
			case "r":
				return restOfWord1
			case "h":
				return "w" + restOfWord2
			default:
				return s
			}
		case "x":
			return "s" + restOfWord1
		default:
			return s
		}
	}
}
