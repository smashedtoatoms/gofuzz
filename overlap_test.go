package gofuzz

import (
	"testing"
)

func TestOverlap(t *testing.T) {
	if Overlap("", "", 2) != 0.0 {
		t.Errorf("Overlap Test Failed.")
	}
}
