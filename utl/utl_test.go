package utl

import (
	"fmt"
	"testing"
)

func TestDeDuplicate(t *testing.T) {
	if DeDuplicate("a") != "a" {
		t.Errorf("DeDuplicate fails to deduplicate the single letter a.")
	}
	if DeDuplicate("c") != "c" {
		t.Errorf("DeDuplicate fails to deduplicate the single letter c.")
	}
}
func ExampleDeDuplicate() {
	result := DeDuplicate("buzz")
	fmt.Print(result)
	// Output: buz
}
func ExampleDeDuplicate_2() {
	result := DeDuplicate("accept")
	fmt.Print(result)
	// Output: accept
}

func TestIsAlphabetic(t *testing.T) {
	if IsAlphabetic("abc3") != false {
		t.Errorf("Fauled to determine that abc3 is not alphabetic.")
	}
	if IsAlphabetic("abc") != true {
		t.Errorf("Fauled to determine that abc is alphabetic.")
	}
}
func ExampleIsAlphabetic() {
	result := IsAlphabetic("abc")
	fmt.Print(result)
	// Output: true
}
func ExampleIsAlphabetic_2() {
	result := IsAlphabetic("abc1")
	fmt.Print(result)
	// Output: false
}

func ExampleIntersect() {
	result := Intersect("context", "contentcontent")
	fmt.Print(result)
	// Output: contet
}

func TestTokenize(t *testing.T) {
	token := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	if stringArrayNotEqual(Tokenize("abcdefghijklmnopqrstuvwxyz", 1), token) {
		t.Errorf("Failed to tokenize string by ngram 1.")
	}
	token = []string{"ab", "bc", "cd", "de", "ef", "fg", "gh", "hi", "ij",
		"jk", "kl", "lm", "mn", "no", "op", "pq", "qr", "rs", "st", "tu", "uv",
		"vw", "wx", "xy", "yz"}
	if stringArrayNotEqual(Tokenize("abcdefghijklmnopqrstuvwxyz", 2), token) {
		t.Errorf("Failed to tokenize string by ngram 2.")
	}
	token = []string{"abc", "bcd", "cde", "def", "efg", "fgh", "ghi", "hij",
		"ijk", "jkl", "klm", "lmn", "mno", "nop", "opq", "pqr", "qrs", "rst",
		"stu", "tuv", "uvw", "vwx", "wxy", "xyz"}
	if stringArrayNotEqual(Tokenize("abcdefghijklmnopqrstuvwxyz", 3), token) {
		t.Errorf("Failed to tokenize string by ngram 3.")
	}
}
func ExampleTokenize() {
	result := Tokenize("abcd", 2)
	fmt.Print(result)
	// Output: [ab bc cd]
}

func TestGetLetter(t *testing.T) {
	e := "String index out of range."
	l1, w1, err := GetLetter("bob世jerry", 3)
	if l1 != "世" || w1 != "jerry" || err != nil {
		t.Errorf("Failed to get letter and word correctly.")
	}
	l2, w2, err := GetLetter("a", 1)
	if l2 != "a" || w2 != "" || err.Error() != e {
		t.Errorf("Failed to return correct letters when range was too large.")
	}
	l3, w3, err := GetLetter("", 0)
	if l3 != "" || w3 != "" || err.Error() != e {
		t.Errorf("Failed to return correct letters when string was empty.")
	}
}
func ExampleTestGetLetter() {
	firstLetter, restOfWord, err := GetLetter("bob", 0)
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(firstLetter + " : " + restOfWord)
	// Output: b : ob
}

func ExampleContains() {
	success := Contains([]string{"a", "b", "jerry"}, "jerry")
	fmt.Print(success)
	// Output: true
}

func TestSplitAt(t *testing.T) {
	s1, s2, err := SplitAt("toolong", 7)
	e := "String index out of range."
	if s1 != "toolong" || s2 != "" || err.Error() != e {
		t.Errorf("Failed to return correct split strings when range was too large.")
	}
}
func ExampleSplitAt() {
	s1, s2, err := SplitAt("splitme", 5)
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(s1 + " : " + s2)
	// Output: split : me
}

func TestLastLetter(t *testing.T) {
	if LastLetter("jason") != "n" {
		t.Errorf("failed to get last letter from the word 'Jason'.")
	}
	if LastLetter("j") != "j" {
		t.Errorf("failed to get last letter from the word 'j'.")
	}
	if LastLetter("") != "" {
		t.Errorf("failed to get last letter from empty string.")
	}
}
func ExampleLastLetter() {
	lastLetter := LastLetter("these three words")
	fmt.Print(lastLetter)
	// Output: s
}

func TestFirstLetter(t *testing.T) {
	if FirstLetter("jason") != "j" {
		t.Errorf("failed to get last letter from the word 'Jason'.")
	}
	if FirstLetter("j") != "j" {
		t.Errorf("failed to get last letter from the word 'j'.")
	}
	if FirstLetter("") != "" {
		t.Errorf("failed to get last letter from empty string.")
	}
}
func ExampleFirstLetter() {
	firstLetter := FirstLetter("these three words")
	fmt.Print(firstLetter)
	// Output: t
}

func TestSize(t *testing.T) {
	if Size("") != 0 {
		t.Errorf("failed to determine size of empty string.")
	}
}
func ExampleSize() {
	stringSize := Size("this string")
	fmt.Print(stringSize)
	// Output: 11
}

/* Helper Functions */

// Compares string slices.
func stringArrayNotEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
}
