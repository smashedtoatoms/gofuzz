package gofuzz

import (
	"fmt"
	"testing"
)

func TestOverlap(t *testing.T) {
	// Error conditions
	error := "String is too small to calculate Overlap for the given ngram size."
	_, err1 := Overlap("n", "naght", 2)
	if err1.Error() != error {
		t.Errorf("Overlap ngram too large, large string second..")
	}
	_, err2 := Overlap("night", "n", 2)
	if err2.Error() != error {
		t.Errorf("Overlap ngram too large, large string first.")
	}
	_, err3 := Overlap("ni", "naght", 3)
	if err3.Error() != error {
		t.Errorf("Overlap ngram too large, large string second.")
	}
	_, err4 := Overlap("naght", "na", 3)
	if err4.Error() != error {
		t.Errorf("Overlap ngram too large, large string first.")
	}
	_, err5 := Overlap("", "", 1)
	if err5.Error() != error {
		t.Errorf("Overlap two empty strings failed")
	}
	_, err6 := Overlap("abc", "", 3)
	if err6.Error() != error {
		t.Errorf("Overlap second string is empty failed.")
	}
	_, err7 := Overlap("", "abc", 3)
	if err7.Error() != error {
		t.Errorf("Overlap first string is empty failed")
	}

	// Total Equality
	q1, _ := Overlap("abc", "abc", 1)
	if q1 != 1.0 {
		t.Errorf("Equal letters, ngram 1, expected 1.0, failed with: %f.", q1)
	}
	q2, _ := Overlap("abc", "abc", 2)
	if q2 != 1.0 {
		t.Errorf("Equal letters, ngram 2, expected 1.0, failed with: %f.", q2)
	}
	q3, _ := Overlap("abc", "abc", 3)
	if q3 != 1.0 {
		t.Errorf("Equal letters, ngram 3, expected 1.0, failed with: %f.", q3)
	}

	// Total Inequality
	r1, _ := Overlap("abc", "xyz", 1)
	if r1 != 0.0 {
		t.Errorf("Unequal letters, ngram 1, expected 0.0, failed with: %f.", r1)
	}
	r2, _ := Overlap("abc", "xyz", 2)
	if r2 != 0.0 {
		t.Errorf("Unequal letters, ngram 2, expected 0.0, failed with: %f.", r2)
	}
	r3, _ := Overlap("abc", "xyz", 3)
	if r3 != 0.0 {
		t.Errorf("Unequal letters, ngram 3, expected 0.0, failed with: %f.", r3)
	}

	// Valid tests
	s1, _ := Overlap("bob", "bobman", 1)
	if s1 != 1.0 {
		t.Errorf("'bob' and 'bobman', ngram 1 should equal 1.0, they equaled:"+
			" %f.", s1)
	}
	s2, _ := Overlap("bob", "manbobman", 1)
	if s2 != 1.0 {
		t.Errorf("'bob' and 'manbobman', ngram 1 should equal 1, they equaled:"+
			" %f.", s2)
	}
	s3, _ := Overlap("night", "nacht", 1)
	if s3 != 0.6 {
		t.Errorf("'night' and 'nacht', ngram 1 should equal 0.6, they equaled:"+
			" %f.", s3)
	}
	s4, _ := Overlap("night", "naght", 1)
	if s4 != 0.8 {
		t.Errorf("'night' and 'naght', ngram 1 should equal 0.8, they equaled:"+
			" %f.", s4)
	}
	s5, _ := Overlap("context", "contact", 1)
	if s5 != 0.7142857142857143 {
		t.Errorf("'context' and 'contact', ngram 1 should equal 0.7142857142857"+
			"143, they equaled: %f.", s5)
	}
	s6, _ := Overlap("night", "nacht", 2)
	if s6 != 0.25 {
		t.Errorf("'night' and 'nacht', ngram 2 should equal 0.25, they equaled:"+
			" %f.", s6)
	}
	s7, _ := Overlap("night", "naght", 2)
	if s7 != 0.5 {
		t.Errorf("'night' and 'nacht', ngram 2 should equal 0.5, they equaled:"+
			"%f.", s7)
	}
	s8, _ := Overlap("context", "contact", 2)
	if s8 != 0.5 {
		t.Errorf("'context' and 'contact', ngram 2 should equal 0.5, they "+
			"equaled: %f.", s8)
	}
	s9, _ := Overlap("contextcontext", "contact", 2)
	if s9 != 0.5 {
		t.Errorf("'contextcontext' and 'contact', ngram 2 should equal 0.5, they"+
			" equaled: %f.", s9)
	}
	s10, _ := Overlap("contact", "contextcontext", 2)
	if s10 != 0.5 {
		t.Errorf("'contact' and 'contextcontext', ngram 2 should equal 0.5, they"+
			" equaled: %f.", s10)
	}
	s11, _ := Overlap("ht", "nacht", 2)
	if s11 != 1.0 {
		t.Errorf("'ht' and 'nacht', ngram2 should equal 1.0 they equaled:"+
			" %f.", s11)
	}
	s12, _ := Overlap("xp", "nacht", 2)
	if s12 != 0.0 {
		t.Errorf("'xp' and 'nacht', ngram2  should equal 0.0, they equaled:"+
			" %f.", s12)
	}
	s13, _ := Overlap("ht", "hththt", 2)
	if s13 != 1.0 {
		t.Errorf("'ht' and 'hththt', ngram2  should equal 1.0, they equaled:"+
			" %f.", s13)
	}
	s14, _ := Overlap("night", "nacht", 3)
	if s14 != 0.0 {
		t.Errorf("'night' and 'nacht', ngram 3 should equal 0.0, they equaled:"+
			" %f.", s14)
	}
	s15, _ := Overlap("night", "naght", 3)
	if s15 > 0.3333332 && s15 > 0.3333334 {
		t.Errorf("'night' and 'naght', ngram 3 should equal 0.333333, they "+
			"equaled: %f.", s15)
	}
	s16, _ := Overlap("context", "contact", 3)
	if s16 != 0.4 {
		t.Errorf("'context' and 'contact', ngram3 should equal 0.4, they "+
			"equaled: %f.", s16)
	}
}
func ExampleOverlap() {
	result, err := Overlap("Chipotle", "Chipotle Mexican Grill", 2)
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: 1
}
