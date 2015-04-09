package sim

import (
	"fmt"
	"testing"
)

func TestJaroWinkler(t *testing.T) {
	// Error conditions
	error := "Unable to calculate Jaro-Winkler against empty string."
	_, err1 := JaroWinkler("", "")
	if err1.Error() != error {
		t.Errorf("JaroWinkler two empty strings failed, expected error.")
	}
	_, err2 := JaroWinkler("abc", "")
	if err2.Error() != error {
		t.Errorf("JaroWinkler second string empty failed, expected error.")
	}
	_, err3 := JaroWinkler("", "xyz")
	if err3.Error() != error {
		t.Errorf("JaroWinkler first string empty failed, expected error.")
	}

	// Total Equality
	q1, _ := JaroWinkler("a", "a")
	if q1 != 1.0 {
		t.Errorf("Equal letter expected 1.0, failed with: %f.", q1)
	}
	q2, _ := JaroWinkler("abc", "abc")
	if q2 != 1.0 {
		t.Errorf("Equal letters expected 1.0, failed with: %f.", q2)
	}
	q3, _ := JaroWinkler("123", "123")
	if q3 != 1.0 {
		t.Errorf("Equal numbers expected 1.0, failed with: %f.", q3)
	}

	// Total Inequality
	r1, _ := JaroWinkler("123", "456")
	if r1 != 0.0 {
		t.Errorf("Unequal numeric arguments expected 0.0, failed with: %f.", r1)
	}
	r2, _ := JaroWinkler("abc", "xyz")
	if r2 != 0.0 {
		t.Errorf("Unequal alphabetic arguments expected 0.0, failed with: %f.", r2)
	}

	// Main tests
	s1, _ := JaroWinkler("aa", "a")
	if s1 > 0.8500001 || s1 < 0.8499999 {
		t.Errorf("'aa' to 'a' failed with: %f.", s1)
	}
	s2, _ := JaroWinkler("a", "aa")
	if s2 > 0.8500001 || s2 < 0.8499999 {
		t.Errorf("'a' to 'aa' failed with: %f.", s2)
	}
	s3, _ := JaroWinkler("veryveryverylong", "v")
	if s3 != 0.71875 {
		t.Errorf("'veryveryverylong' to 'v' expected 0.6875, failed with: %f.", s3)
	}
	s4, _ := JaroWinkler("v", "veryveryverylong")
	if s4 != 0.71875 {
		t.Errorf("'v' to 'veryveryverylong' expected 0.6875, failed with: %f.", s4)
	}
	s5, _ := JaroWinkler("martha", "marhta")
	if s5 > 0.961112 || s5 < 0.961110 {
		t.Errorf("'martha' to 'marhta' expected 0.961111, failed with: %f.", s5)
	}
	s6, _ := JaroWinkler("dwayne", "duane")
	if s6 > 0.840001 || s6 < 0.839999 {
		t.Errorf("'dwayne' to 'duane' expected 0.8400001, failed with: %f.", s6)
	}
	s7, _ := JaroWinkler("dixon", "dicksonx")
	if s7 > 0.813334 || s7 < 0.813332 {
		t.Errorf("'dixon' to 'dicksonx' expected 0.8133333, failed with: %f.", s7)
	}
	s8, _ := JaroWinkler("abcvwxyz", "cabvwxyz")
	if s8 > 0.958334 || s8 < 0.958332 {
		t.Errorf("'abcvwxyz' to 'cabvwxyz' expected 0.958333, failed with: %f.", s8)
	}
	s9, _ := JaroWinkler("jones", "johnson")
	if s9 > 0.832381 || s9 < 0.832379 {
		t.Errorf("'jones' to 'johnson' expected 0.832380, failed with: %f.", s9)
	}
	s10, _ := JaroWinkler("henka", "henkan")
	if s10 > 0.96667 || s10 < 0.96665 {
		t.Errorf("'henka' to 'henkan' expected 0.9666666, failed with: %f.", s10)
	}
	s11, _ := JaroWinkler("fvie", "ten")
	if s11 != 0.0 {
		t.Errorf("'fvie' to 'ten' expected 0.0, failed with: %f.", s11)
	}

	// Pointless comparisons.
	t1, _ := JaroWinkler("zac ephron", "zac efron")
	t2, _ := JaroWinkler("zac ephron", "kai ephron")
	if t1 <= t2 {
		t.Errorf("failed bool comparison 1")
	}
	t3, _ := JaroWinkler("brittney spears", "britney spears")
	t4, _ := JaroWinkler("brittney spears", "brittney startzman")
	if t3 <= t4 {
		t.Errorf("failed bool comparison 2")
	}
}
func ExampleJaroWinkler() {
	result, err := JaroWinkler("bob", "bobby")
	if err != nil {
		fmt.Print("an unexpected error was encountered")
	}
	fmt.Print(result)
	// Output: 0.8288889
}
