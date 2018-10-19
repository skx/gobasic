// vars_test.go - Simple test-cases for our variable-store

package eval

import (
	"testing"
)

// TestInt: Test we can store/retrieve an int.
func TestInt(t *testing.T) {

	// Holder for variables
	v := NewVars()

	// Set "steve" -> int
	v.Set("steve", 42)

	// Get the value
	out := v.Get("steve")

	// Ensure it is an int
	value, ok := out.(int)
	if !ok {
		t.Errorf("Failed to cast to int!")
	}
	if value != 42 {
		t.Errorf("Value had the wrong content!")
	}
}

// TestString: Test we can store/retrieve a string.
func TestString(t *testing.T) {

	// Holder for variables
	v := NewVars()

	// Set "steve$" -> string
	v.Set("steve$", "kemp")

	// Get the value
	out := v.Get("steve$")

	// Ensure it is a string
	value, ok := out.(string)
	if !ok {
		t.Errorf("Failed to cast to string!")
	}
	if value != "kemp" {
		t.Errorf("Value had the wrong content!")
	}
}
