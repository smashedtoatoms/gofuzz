package gofuzz

import (
	"testing"
)

func TestJaro(t *testing.T) {
	if Jaro("", "") != 0.0 {
		t.Errorf("Jaro Test Failed.")
	}
}

func TestJaroWinkler(t *testing.T) {
	if JaroWinkler("", "") != 0.0 {
		t.Errorf("Jaro Test Failed.")
	}
}

func TestOverlap(t *testing.T) {
	if Overlap("", "", 2) != 0.0 {
		t.Errorf("Jaro Test Failed.")
	}
}

func TestMetaphone(t *testing.T) {
	if Metaphone("") != "" {
		t.Errorf("Jaro Test Failed.")
	}
}

func TestMetaphoneMetric(t *testing.T) {
	if MetaphoneMetric("", "") != false {
		t.Errorf("Jaro Test Failed.")
	}
}
