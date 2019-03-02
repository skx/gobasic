// vars_test.go - Simple test-cases for our variable-store

package eval

import (
	"testing"

	"github.com/andydotxyz/gobasic/object"
)

// TestInt: Test we can store/retrieve an int.
func TestInt(t *testing.T) {

	// Holder for variables
	v := NewVars()

	// Set "steve" -> int
	v.Set("steve", &object.NumberObject{Value: 42})

	// Get the value
	out := v.Get("steve")

	// Ensure it is an int
	if out.Type() != object.NUMBER {
		t.Errorf("The value was the wrong type!")
	}

	// And check the value is correct.
	if out.(*object.NumberObject).Value != 42 {
		t.Errorf("Our value was lost!")
	}
}

// TestString: Test we can store/retrieve a string.
func TestString(t *testing.T) {

	// Holder for variables
	v := NewVars()

	// Set "steve" -> int
	v.Set("steve$", &object.StringObject{Value: "Kemp"})

	// Get the value
	out := v.Get("steve$")

	// Ensure it is a string
	if out.Type() != object.STRING {
		t.Errorf("The value was the wrong type!")
	}

	// And check the value is correct.
	if out.(*object.StringObject).Value != "Kemp" {
		t.Errorf("Our value was lost!")
	}
}
