package utl

import (
	"testing"
)

func TestDeDuplicate(t *testing.T) {
	if DeDuplicate("a") != "a" {
		t.Errorf("DeDuplicate fails to deduplicate the single letter a.")
	}
	if DeDuplicate("c") != "c" {
		t.Errorf("DeDuplicate fails to deduplicate the single letter c.")
	}
	if DeDuplicate("buzz") != "buz" {
		t.Errorf("DeDuplicate fails to deduplicate the letter z from buzz.")
	}
	if DeDuplicate("accept") != "accept" {
		t.Errorf("DeDuplicate failed to not deduplicate the letter c from accept.")
	}
}

func TestIsAlphabetic(t *testing.T) {
	if IsAlphabetic("abc3") != false {
		t.Errorf("Fauled to determine that abc3 is not alphabetic.")
	}
	if IsAlphabetic("abc") != true {
		t.Errorf("Fauled to determine that abc is alphabetic.")
	}
}

func TestIntersect(t *testing.T) {
	if Intersect("context", "contentcontent") != "contet" {
		t.Errorf("Failed to find the intersect of context and contentcontent.")
	}
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

/* Examples */

func ExampleDeDuplicate(t *testing.T) {
	DeDuplicate("buzz")
	// Output: buz
}
func ExampleDeDuplicate_2(t *testing.T) {
	DeDuplicate("accept")
	// Output: accept
}
func ExampleIsAlphabetic(t *testing.T) {
	IsAlphabetic("abc")
	// Output: true
}
func ExampleIsAlphabetic_2(t *testing.T) {
	IsAlphabetic("abc1")
	// Output: false
}
func ExampleIntersect(t *testing.T) {
	Intersect("context", "contentcontent")
	// Output: "contet"
}
func ExampleTokenize(t *testing.T) {
	Tokenize("abcd", 2)
	// Output: []string{"ab", "bc", "cd"}]
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
