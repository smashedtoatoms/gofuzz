package sim

import (
	"errors"
	"github.com/smashedtoatoms/gofuzz/utl"
)

// JaroWinkler calculates the Jaro-Winkler distance between two strings.
func JaroWinkler(s1 string, s2 string) (float32, error) {
	chars1 := []rune(s1)
	chars2 := []rune(s2)
	s1Size := len(chars1)
	s2Size := len(chars2)
	switch {
	case s1Size == 0 || s2Size == 0:
		return 0.0, errors.New("Unable to calculate Jaro-Winkler against empty " +
			"string.")
	case s1 == s2:
		return 1.0, nil
	case s1Size > s2Size:
		jaroScore, _ := Jaro(s2, s1)
		modifiedPrefix := float32(modifyPrefix(chars2, chars1))
		return jaroScore + ((modifiedPrefix * (1.0 - jaroScore)) / 10.0), nil
	default:
		jaroScore, _ := Jaro(s1, s2)
		modifiedPrefix := float32(modifyPrefix(chars1, chars2))
		return jaroScore + ((modifiedPrefix * (1.0 - jaroScore)) / 10.0), nil
	}
}

/* Helper functions */

func modifyPrefix(chars1 []rune, chars2 []rune) int {
	prefixLength := 0
	lastChar := utl.MinInt([]int{4, len(chars1)})
	for prefixLength < lastChar && chars1[prefixLength] == chars2[prefixLength] {
		prefixLength = prefixLength + 1
	}
	return prefixLength
}
