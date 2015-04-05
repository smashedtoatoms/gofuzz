package gofuzz

import (
	"testing"
)

func TestJaro(t *testing.T) {
	error := "Unable to calculate Jaro against empty string."
	_, err1 := Jaro("", "")
	if err1.Error() != error {
		t.Errorf("Jaro Test Failed.")
	}
}
