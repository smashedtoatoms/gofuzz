package gofuzz

import (
	"testing"
)

func TestJaro(t *testing.T) {
	if Jaro("", "") != 0.0 {
		t.Errorf("Jaro Test Failed.")
	}
}
