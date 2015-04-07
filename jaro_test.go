package gofuzz

import (
	"fmt"
	"testing"
)

func TestJaro(t *testing.T) {
	// Error conditions
	error := "Unable to calculate Jaro against empty string."
	_, err1 := Jaro("", "")
	if err1.Error() != error {
		t.Errorf("Jara two empty strings failed, expected error.")
	}
	_, err2 := Jaro("abc", "")
	if err2.Error() != error {
		t.Errorf("Jaro second string empty failed, expected error.")
	}
	_, err3 := Jaro("", "xyz")
	if err3.Error() != error {
		t.Errorf("Jaro first string empty failed, expected error.")
	}

	// Total Equality
	q1, _ := Jaro("a", "a")
	if q1 != 1.0 {
		t.Errorf("Equal letter expected 1.0, failed with: %f.", q1)
	}
	q2, _ := Jaro("abc", "abc")
	if q2 != 1.0 {
		t.Errorf("Equal letters expected 1.0, failed with: %f.", q2)
	}
	q3, _ := Jaro("123", "123")
	if q3 != 1.0 {
		t.Errorf("Equal numbers expected 1.0, failed with: %f.", q3)
	}

	// Total Inequality
	r1, _ := Jaro("123", "456")
	if r1 != 0.0 {
		t.Errorf("Unequal numeric arguments expected 0.0, failed with: %f.", r1)
	}
	r2, _ := Jaro("abc", "xyz")
	if r2 != 0.0 {
		t.Errorf("Unequal alphabetic arguments expected 0.0, failed with: %f.", r2)
	}

	// Main tests
	s1, _ := Jaro("aa", "a")
	if s1 > 0.8333334 || s1 < 0.8333332 {
		t.Errorf("'aa' to 'a' failed with: %f.", s1)
	}
	s2, _ := Jaro("a", "aa")
	if s1 > 0.8333334 || s1 < 0.8333332 {
		t.Errorf("'a' to 'aa' failed with: %f.", s2)
	}
	s3, _ := Jaro("veryveryverylong", "v")
	if s3 != 0.6875 {
		t.Errorf("'veryveryverylong' to 'v' expected 0.6875, failed with: %f.", s3)
	}
	s4, _ := Jaro("v", "veryveryverylong")
	if s4 != 0.6875 {
		t.Errorf("'v' to 'veryveryverylong' expected 0.6875, failed with: %f.", s4)
	}
	s5, _ := Jaro("martha", "marhta")
	if s5 > 0.94445 || s5 < 0.94443 {
		t.Errorf("'martha' to 'marhta' expected 0.94444, failed with: %f.", s5)
	}
	s6, _ := Jaro("dwayne", "duane")
	if s6 > 0.822223 || s6 < 0.822221 {
		t.Errorf("'dwayne' to 'duane' expected 0.82222, failed with: %f.", s6)
	}
	s7, _ := Jaro("dixon", "dicksonx")
	if s7 > 0.766668 || s7 < 0.766666 {
		t.Errorf("'dixon' to 'dicksonx' expected 0.766667, failed with: %f.", s7)
	}
	s8, _ := Jaro("abcvwxyz", "cabvwxyz")
	if s8 > 0.958334 || s8 < 0.958332 {
		t.Errorf("'abcvwxyz' to 'cabvwxyz' expected 0.95833, failed with: %f.", s8)
	}
	s9, _ := Jaro("jones", "johnson")
	if s9 > 0.790477 || s9 < 0.790475 {
		t.Errorf("'jones' to 'johnson' expected 0.790476, failed with: %f.", s9)
	}
	s10, _ := Jaro("henka", "henkan")
	if s10 > 0.944445 || s10 < 0.944444 {
		t.Errorf("'henka' to 'henkan' expected 0.944444, failed with: %f.", s10)
	}
	s11, _ := Jaro("fvie", "ten")
	if s11 != 0.0 {
		t.Errorf("'fvie' to 'ten' expected 0.0, failed with: %f.", s11)
	}

	// Pointless comparisons.
	t1, _ := Jaro("zac ephron", "zac efron")
	t2, _ := Jaro("zac ephron", "kai ephron")
	if t1 <= t2 {
		t.Errorf("failed bool comparison 1")
	}
	t3, _ := Jaro("brittney spears", "britney spears")
	t4, _ := Jaro("brittney spears", "brittney startzman")
	if t3 <= t4 {
		t.Errorf("failed bool comparison 2")
	}
}
func ExampleJaro() {
	result, err := Jaro("bob", "bobby")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: 0.75555557
}
