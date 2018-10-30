// string_test.go - Simple test-cases for string-related primitives.

package builtin

import (
	"testing"

	"github.com/skx/gobasic/object"
)

func TestChr(t *testing.T) {

	//
	// Call with a non-number argument.
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := CHR(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now do it properly
	//
	var validArgs []object.Object
	validArgs = append(validArgs, &object.NumberObject{Value: 42})
	out2 := CHR(nil, validArgs)
	if out2.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", out2.String())
	}
	if out2.(*object.StringObject).Value != "*" {
		t.Errorf("Function returned a surprising result")
	}

}

func TestCode(t *testing.T) {

	//
	// Call with a non-string argument.
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := CODE(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now do it properly, we expect * -> 42.
	//
	var validArgs []object.Object
	validArgs = append(validArgs, &object.StringObject{Value: "*"})
	out2 := CODE(nil, validArgs)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a number return, but didn't get one: %s", out2.String())
	}
	if out2.(*object.NumberObject).Value != 42 {
		t.Errorf("Function returned a surprising result")
	}

	//
	// With an empty string we receive zero as a result
	//
	var emptyArgs []object.Object
	emptyArgs = append(emptyArgs, &object.StringObject{Value: ""})
	out3 := CODE(nil, emptyArgs)
	if out3.Type() != object.NUMBER {
		t.Errorf("We expected a number return, but didn't get one: %s", out3.String())
	}
	if out3.(*object.NumberObject).Value != 0 {
		t.Errorf("Function returned a surprising result")
	}
}

func TestLeft(t *testing.T) {

	//
	// Call with an initial argument which is a non-string.
	//
	var fail1 []object.Object
	fail1 = append(fail1, object.Error("Bogus type"))
	out1 := LEFT(nil, fail1)
	if out1.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now call with a string for the first argument, but a non-number
	// for the second.
	//
	var fail2 []object.Object
	fail2 = append(fail2, &object.StringObject{Value: "Valid type"})
	fail2 = append(fail2, object.Error("Invalid"))
	out2 := LEFT(nil, fail2)
	if out2.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Setup a structure for testing.
	//
	type LeftTest struct {
		Input  string
		Count  float64
		Output string
	}

	// Define some tests
	tests := []LeftTest{{Input: "Steve", Count: 2, Output: "St"},
		{Input: "Steve", Count: 200, Output: "Steve"},
		// TODO: Fix me		{Input: "ウェブの国際化", Count: 3, Output: "ウェブ"},
	}

	for _, test := range tests {

		var args []object.Object
		args = append(args, &object.StringObject{Value: test.Input})
		args = append(args, &object.NumberObject{Value: test.Count})
		output := LEFT(nil, args)
		if output.Type() != object.STRING {
			t.Errorf("We expected a string-result, but got something else")
		}
		if output.(*object.StringObject).Value != test.Output {
			t.Errorf("LEFT %s,%f gave '%s' not '%s'",
				test.Input, test.Count, output.(*object.StringObject).Value, test.Output)
		}
	}
}

func TestLen(t *testing.T) {

	//
	// Call with a non-string argument.
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := LEN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now test the function with a number of random-strings.
	//
	// UTF-8 included, obviously.
	//
	type LenTest struct {
		Input  string
		Length float64
	}

	// Define some tests - Ensure we have a UTF-8 string.
	tests := []LenTest{
		{Input: "Steve", Length: 5},
		{Input: "", Length: 0},
		{Input: "ウェブの国際化", Length: 7}}

	for _, test := range tests {

		var validArgs []object.Object
		validArgs = append(validArgs, &object.StringObject{Value: test.Input})
		out2 := LEN(nil, validArgs)
		if out2.Type() != object.NUMBER {
			t.Errorf("We expected a number return, but didn't get one: %s", out2.String())
		}
		if out2.(*object.NumberObject).Value != float64(test.Length) {
			t.Errorf("Function returned a surprising result")
		}
	}

}

func TestMid(t *testing.T) {
}

func TestRight(t *testing.T) {
}

func TestStr(t *testing.T) {
}

func TestTl(t *testing.T) {
}

func TestVal(t *testing.T) {
}
