// vars_test.go - Simple test-cases for our variable-store

package eval

import (
	"testing"
)

// TestFuzz is just a simple wrapper that pretends we cover the fuzzer.
func TestFuzz(t *testing.T) {

	data := []byte(`
10 REM
`)

	out := Fuzz(data)
	if out != 1 {
		t.Errorf("We found an unexpected result in Fuzz!")
	}

}
