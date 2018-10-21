package object

import (
	"math"
	"testing"
)

// Test we can create string/ints and that they have the correct types
func TestTypes(t *testing.T) {

	v := StringObject{Value: "Steve"}
	if v.Type() != STRING {
		t.Errorf("Wrong type for String")
	}

	n := NumberObject{Value: math.Pi}
	if n.Type() != NUMBER {
		t.Errorf("Wrong type for Number")
	}
}
