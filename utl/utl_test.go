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
	if DeDuplicate("cool") != "col" {
		t.Errorf("DeDuplicate fails to deduplicate the word 'cool'.")
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

func TestIntersect(t *testing.T) {
	characterList1 := []string{"c", "o", "n", "t", "e", "n", "t", "c", "o", "n",
		"t", "e", "n", "t"}
	characterList2 := []string{"c", "o", "n", "t", "e", "x", "t"}
	result := Intersect(characterList1, characterList2)
	if stringsNotEqual(result, []string{"c", "o", "n", "t", "e", "t"}) {
		t.Errorf("Intersect unable to calculate intersection correctly.")
	}
}

func ExampleIntersect() {
	characterList1 := []string{"c", "o", "n", "t", "e", "x", "t"}
	characterList2 := []string{"c", "o", "n", "t", "e", "n", "t", "c", "o", "n",
		"t", "e", "n", "t"}
	result := Intersect(characterList1, characterList2)
	fmt.Print(result)
	// Output: [c o n t e t]
}

func TestTokenize(t *testing.T) {
	token := []string{"a", "b", "c", "d"}
	if stringsNotEqual(Tokenize([]string{"a", "b", "c", "d"}, 1), token) {
		t.Errorf("Failed to tokenize string by ngram 1.")
	}
	token = []string{"ab", "bc", "cd"}
	if stringsNotEqual(Tokenize([]string{"a", "b", "c", "d"}, 2), token) {
		t.Errorf("Failed to tokenize string by ngram 2.")
	}
	token = []string{"abc", "bcd", "cde"}
	if stringsNotEqual(Tokenize([]string{"a", "b", "c", "d", "e"}, 3), token) {
		t.Errorf("Failed to tokenize string by ngram 3.")
	}
	token = make([]string, 0)
	if stringsNotEqual(Tokenize(make([]string, 0), 3), token) {
		t.Errorf("Failed to tokenize string by ngram 3.")
	}
}
func ExampleTokenize() {
	result := Tokenize([]string{"a", "b", "c", "d"}, 2)
	fmt.Print(result)
	// Output: [ab bc cd]
}

func TestContains(t *testing.T) {
	if Contains([]rune("abc"), 'z') {
		t.Errorf("Failed to not find 'z' in 'abc'.")
	}
}
func ExampleContains() {
	success := Contains([]rune("abc"), 'b')
	fmt.Print(success)
	// Output: true
}

func TestMaxInt(t *testing.T) {
	if MaxInt([]int{100, 2, -3}) != 100 {
		t.Errorf("MaxInt failed to detect that the largest int was 100.")
	}
	if MaxInt([]int{-100, 2, -3}) != 2 {
		t.Errorf("MaxInt failed to detect that the largest int was 2.")
	}
}
func ExampleMaxInt() {
	maxInt := MaxInt([]int{1, 2, 3})
	fmt.Print(maxInt)
	// Output: 3
}

func TestMinInt(t *testing.T) {
	if MinInt([]int{100, 2, -3}) != -3 {
		t.Errorf("MinInt failed to detect that the smallest int was -3.")
	}
	if MinInt([]int{-100, 2, -3}) != -100 {
		t.Errorf("MinInt failed to detect that the smallest int was -100.")
	}
}
func ExampleMinInt() {
	minInt := MinInt([]int{1, 2, 3})
	fmt.Print(minInt)
	// Output: 1
}

/* Helper functions */

func runesNotEqual(a []rune, b []rune) bool {
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

// Compares string slices.
func stringsNotEqual(a []string, b []string) bool {
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
