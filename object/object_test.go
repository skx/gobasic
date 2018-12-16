package object

import (
	"math"
	"strings"
	"testing"
)

// Test we can create error/int/string and that they have the correct types
func TestTypes(t *testing.T) {

	v := StringObject{Value: "Steve"}
	if v.Type() != STRING {
		t.Errorf("Wrong type for String")
	}
	if !strings.Contains(v.String(), ":string") {
		t.Errorf("Unexpected value for stringified object")
	}

	n := NumberObject{Value: math.Pi}
	if n.Type() != NUMBER {
		t.Errorf("Wrong type for Number")
	}
	if !strings.Contains(n.String(), ":number") {
		t.Errorf("Unexpected value for stringified object")
	}

	e := ErrorObject{Value: "You fail!"}
	if e.Type() != ERROR {
		t.Errorf("Wrong type for Error")
	}
	if !strings.Contains(e.String(), ":error") {
		t.Errorf("Unexpected value for stringified object")
	}
}

func TestError(t *testing.T) {

	a := Error("Test")
	b := Error("Test %d", 3)
	c := Error("Test %s", "me")

	// Test types
	if a.Type() != ERROR {
		t.Errorf("Object has the wrong type!")
	}
	if b.Type() != ERROR {
		t.Errorf("Object has the wrong type!")
	}
	if c.Type() != ERROR {
		t.Errorf("Object has the wrong type!")
	}

	// Test values
	if a.Value != "Test" {
		t.Errorf("Wrong value for error-message")
	}
	if b.Value != "Test 3" {
		t.Errorf("Wrong value for error-message")
	}
	if c.Value != "Test me" {
		t.Errorf("Wrong value for error-message")
	}
}

func TestNumber(t *testing.T) {

	a := Number(33)

	// Test types
	if a.Type() != NUMBER {
		t.Errorf("Object has the wrong type!")
	}

	// Test values
	if a.Value != 33 {
		t.Errorf("Wrong value for number-object")
	}

}

func TestString(t *testing.T) {

	a := String("Test")

	// Test types
	if a.Type() != STRING {
		t.Errorf("Object has the wrong type!")
	}

	// Test values
	if a.Value != "Test" {
		t.Errorf("Wrong value for string-object")
	}
}

func Test1DArray(t *testing.T) {

	// Create an array of one dimension
	a := Array(0, 5)

	if a.Type() != ARRAY {
		t.Errorf("Object has the wrong type!")
	}

	// Ensure each dimension is settable
	for i := 0; i < 5; i++ {
		a.Set(0, i, Number(float64(i)))
	}

	// Ensure each dimension is gettable
	sum := 0
	for i := 0; i < 5; i++ {
		out := a.Get(0, i)

		if out.Type() == NUMBER {
			sum += int(out.(*NumberObject).Value)
		}
	}

	if sum != 10 {
		t.Errorf("Sum was %d not %d", sum, 10)
	}

	// Test that we have bound-checking on Get
	err1 := a.Get(0, 6)
	if err1.Type() != ERROR {
		t.Errorf("Expected a bounds-error, got none")
	}

	// Test that we have bound-checking on Set
	err2 := a.Set(0, 6, err1)
	if err2.Type() != ERROR {
		t.Errorf("Expected a bounds-error, got none")
	}
}

func Test2DArray(t *testing.T) {

	// Create an array
	a := Array(5, 5)

	if a.Type() != ARRAY {
		t.Errorf("Object has the wrong type!")
	}

	// Ensure each dimension is gettable
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a.Set(i, j, Number(float64(i*j)))
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			expected := float64(i * j)

			actual := a.Get(i, j)

			if actual.Type() != NUMBER {
				t.Errorf("Failed to get the correct type")
			}
			if expected != actual.(*NumberObject).Value {
				t.Errorf("Failed to get the right value for %d,%d!", i, j)
			}
		}
	}

	// Convert the object to a string
	out := a.String()
	if !strings.Contains(out, "Value:16.00") {
		t.Errorf("Failed to find a decent value in our string-representation")
	}

	// Get an out of bounds entry
	err := a.Get(6, 4)
	if err.Type() != ERROR {
		t.Errorf("Expected error - got none!")
	}

	// Set an out of bounds entry
	e := a.Set(6, 4, Number(3))
	if e == nil {
		t.Errorf("Expected error - got none!")
	}
	e = a.Set(3, 2, Number(3))
	if e.Type() == ERROR {
		t.Errorf("We didn't expect an error, but we found one: %v", e)
	}
}
