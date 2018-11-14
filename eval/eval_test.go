// eval_test.go - Simple test-cases for our evaluator.

package eval

import (
	"testing"

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

// TestTrace tests that getting/setting the tracing-flag works as expected.
func TestTrace(t *testing.T) {
	input := `10 PRINT "OK\n"`
	tokener := tokenizer.New(input)
	e, _ := New(tokener)

	// Tracing is off by default
	if e.GetTrace() != false {
		t.Errorf("tracing should not be enabled by default")
	}

	// Enable tracing
	e.SetTrace(true)

	// Tracing is now on
	if e.GetTrace() != true {
		t.Errorf("tracing should have been enabled, but was not")
	}
}

// TestVariables gets/sets some variables and ensures they work
func TestVariables(t *testing.T) {

	type Test struct {
		Name   string
		Object object.Object
	}

	var vars []Test

	//
	// Setup some test variables.
	//
	vars = append(vars, Test{Name: "number", Object: object.Number(33)})
	vars = append(vars, Test{Name: "string", Object: object.String("Steve")})
	vars = append(vars, Test{Name: "error", Object: object.Error("Blah")})

	//
	// Test getting/setting each variable.
	//
	for _, v := range vars {

		input := `10 PRINT "OK\n"`
		tokener := tokenizer.New(input)
		e, _ := New(tokener)

		//
		// By default the variable won't exist.
		//
		cur := e.GetVariable(v.Name)
		if cur.Type() != object.ERROR {
			t.Errorf("Unexpectedly managed to retrieve a missing variable")
		}

		//
		// Set it
		//
		e.SetVariable(v.Name, v.Object)

		//
		// Ensure it was set
		//
		cur = e.GetVariable(v.Name)
		if cur.Type() != v.Object.Type() {
			t.Errorf("Retrieved variable '%s' had the wrong type %s != %s", v.Name, cur.Type(), v.Object.Type())
		}
	}
}

// TestData tests that invalid data items cause the program to fail.
func TestData(t *testing.T) {
	type Test struct {
		Input string
		Valid bool
	}

	vars := []Test{{Input: `10 DATA 2,1,2`, Valid: true},
		{Input: `10 DATA "2","1","2"`, Valid: true},
		{Input: `10 DATA "2","steve",2
`, Valid: true},
		{Input: `10 DATA LET, b, c`, Valid: false},
	}

	//
	// Test reading each set of data.
	//
	for _, v := range vars {

		tokener := tokenizer.New(v.Input)
		_, err := New(tokener)

		if v.Valid {
			if err != nil {
				t.Errorf("Expected error, received one: %s!", err.Error())
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, received none for input %s", v.Input)
			}
		}
	}
}
