package gofuzz

import (
	"testing"
)

func TestJaroWinkler(t *testing.T) {
	if JaroWinkler("", "") != 0.0 {
		t.Errorf("JaroWinkler Test Failed.")
	}
}
