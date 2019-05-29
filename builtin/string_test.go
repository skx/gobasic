// string_test.go - Simple test-cases for string-related primitives.

package builtin

import (
	"fmt"
	"strings"
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
	// Call with a negative argument.
	//
	var failArgs2 []object.Object
	failArgs2 = append(failArgs2, object.Number(-33))
	out2 := CHR(nil, failArgs2)
	if out2.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}
	if !strings.Contains(out2.String(), "Positive") {
		t.Errorf("Our error message wasn't what we expected")
	}

	//
	// Now do it properly
	//
	var validArgs []object.Object
	validArgs = append(validArgs, object.Number(42))
	out3 := CHR(nil, validArgs)
	if out3.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", out2.String())
	}
	if out3.(*object.StringObject).Value != "*" {
		t.Errorf("Function returned a surprising result")
	}

	//
	// A large number should also return the appropriate UTF-character.
	//
	var utf []object.Object
	utf = append(utf, object.Number(12454))
	out4 := CHR(nil, utf)
	if out4.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", out3.String())
	}
	if out4.(*object.StringObject).Value != "ウ" {
		t.Errorf("Function returned a surprising result: %s",
			out4.(*object.StringObject).Value)
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
	validArgs = append(validArgs, object.String("*"))
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
	emptyArgs = append(emptyArgs, object.String(""))
	out3 := CODE(nil, emptyArgs)
	if out3.Type() != object.NUMBER {
		t.Errorf("We expected a number return, but didn't get one: %s", out3.String())
	}
	if out3.(*object.NumberObject).Value != 0 {
		t.Errorf("Function returned a surprising result")
	}

	//
	// A UTF-character should also return a number
	//
	var utf []object.Object
	utf = append(utf, object.String("ウ"))
	out4 := CODE(nil, utf)
	if out3.Type() != object.NUMBER {
		t.Errorf("We expected a number return, but didn't get one: %s", out4.String())
	}
	if out4.(*object.NumberObject).Value != 12454 {
		t.Errorf("Function returned a surprising result: %f",
			out4.(*object.NumberObject).Value)
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
	fail2 = append(fail2, object.String("Valid type"))
	fail2 = append(fail2, object.Error("Invalid"))
	out2 := LEFT(nil, fail2)
	if out2.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// The final bogus-test would be to give a negative number as an argument.
	//
	var fail3 []object.Object
	fail3 = append(fail3, object.String("Valid type"))
	fail3 = append(fail3, object.Number(-33))
	out3 := LEFT(nil, fail3)
	if out3.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}
	if !strings.Contains(out3.String(), "Positive") {
		t.Errorf("Our error message wasn't what we expected")
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
	tests := []LeftTest{
		{Input: "Steve", Count: 2, Output: "St"},
		{Input: "Steve", Count: 200, Output: "Steve"},
		{Input: "Test", Count: 2, Output: "Te"},
		{Input: "ウェブの国際化", Count: 3, Output: "ウェブ"},
	}

	for _, test := range tests {

		var args []object.Object
		args = append(args, object.String(test.Input))
		args = append(args, object.Number(test.Count))
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
		validArgs = append(validArgs, object.String(test.Input))
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

	//
	// Valid arguments are string, int, int
	//

	//
	// 1.  Call with error, error, error
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	failArgs = append(failArgs, object.Error("Bogus type"))
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := MID(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// 2.  Call with number, string, string
	//
	var failArgs2 []object.Object
	failArgs2 = append(failArgs2, object.Number(1))
	failArgs2 = append(failArgs2, object.String("Blah"))
	failArgs2 = append(failArgs2, object.String("Blah"))
	out2 := MID(nil, failArgs2)
	if out2.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// 3.  Call with string, string, string
	//
	var failArgs3 []object.Object
	failArgs3 = append(failArgs3, object.String("ok"))
	failArgs3 = append(failArgs3, object.String("Blah"))
	failArgs3 = append(failArgs3, object.String("Blah"))
	out3 := MID(nil, failArgs3)
	if out3.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// 4.  Call with string, number, string
	//
	var failArgs4 []object.Object
	failArgs4 = append(failArgs4, object.String("ok"))
	failArgs4 = append(failArgs4, object.Number(3))
	failArgs4 = append(failArgs4, object.String("Blah"))
	out4 := MID(nil, failArgs4)
	if out4.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// 5.  Call with string, number, -number
	//
	var failArgs5 []object.Object
	failArgs5 = append(failArgs5, object.String("ok"))
	failArgs5 = append(failArgs5, object.Number(3))
	failArgs5 = append(failArgs5, object.Number(-54))
	out5 := MID(nil, failArgs5)
	if out5.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}
	if !strings.Contains(out5.String(), "Positive") {
		t.Errorf("Our error message wasn't what we expected")
	}

	//
	// 6.  Call with string, -number, number
	//
	var failArgs6 []object.Object
	failArgs6 = append(failArgs6, object.String("ok"))
	failArgs6 = append(failArgs6, object.Number(-3))
	failArgs6 = append(failArgs6, object.Number(54))
	out6 := MID(nil, failArgs6)
	if out6.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}
	if !strings.Contains(out6.String(), "Positive") {
		t.Errorf("Our error message wasn't what we expected")
	}

	//
	// Now test things that work properly.
	//
	//
	// Setup a structure for testing.
	//
	type MIDTest struct {
		Input  string
		Offset float64
		Count  float64
		Output string
	}

	// Define some tests
	tests := []MIDTest{{Input: "Steve", Offset: 1, Count: 2, Output: "te"},
		{Input: "Steve", Offset: 4, Count: 100, Output: "e"},
		{Input: "Steve", Offset: 100, Count: 100, Output: ""},
		{Input: "ウェブの国際化", Offset: 1, Count: 2, Output: "ェブ"},
	}

	for _, test := range tests {

		var args []object.Object
		args = append(args, object.String(test.Input))
		args = append(args, object.Number(test.Offset))
		args = append(args, object.Number(test.Count))
		output := MID(nil, args)
		if output.Type() != object.STRING {
			t.Errorf("We expected a string-result, but got something else")
		}
		if output.(*object.StringObject).Value != test.Output {
			t.Errorf("LEFT %s,%f,%f gave '%s' not '%s'",
				test.Input, test.Offset, test.Count, output.(*object.StringObject).Value, test.Output)
		}
	}
}

func TestRight(t *testing.T) {

	//
	// Call with an initial argument which is a non-string.
	//
	var fail1 []object.Object
	fail1 = append(fail1, object.Error("Bogus type"))
	out1 := RIGHT(nil, fail1)
	if out1.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now call with a string for the first argument, but a non-number
	// for the second.
	//
	var fail2 []object.Object
	fail2 = append(fail2, object.String("Valid type"))
	fail2 = append(fail2, object.Error("Invalid"))
	out2 := RIGHT(nil, fail2)
	if out2.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// The final bogus-test would be to give a negative number as an argument.
	//
	var fail3 []object.Object
	fail3 = append(fail3, object.String("Valid type"))
	fail3 = append(fail3, object.Number(-33))
	out3 := RIGHT(nil, fail3)
	if out3.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}
	if !strings.Contains(out3.String(), "Positive") {
		t.Errorf("Our error message wasn't what we expected")
	}

	//
	// Setup a structure for testing.
	//
	type RightTest struct {
		Input  string
		Count  float64
		Output string
	}

	// Define some tests
	tests := []RightTest{{Input: "Steve", Count: 3, Output: "eve"},
		{Input: "Steve", Count: 200, Output: "Steve"},
		{Input: "ウェブの国際化", Count: 1, Output: "化"},
	}

	for _, test := range tests {

		var args []object.Object
		args = append(args, object.String(test.Input))
		args = append(args, object.Number(test.Count))
		output := RIGHT(nil, args)
		if output.Type() != object.STRING {
			t.Errorf("We expected a string-result, but got something else")
		}
		if output.(*object.StringObject).Value != test.Output {
			t.Errorf("RIGHT %s,%f gave '%s' not '%s'",
				test.Input, test.Count, output.(*object.StringObject).Value, test.Output)
		}
	}

}

func TestSpc(t *testing.T) {

	//
	// Call with a non-number argument.
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := SPC(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now a string which will also fail.
	//
	var nArgs []object.Object
	nArgs = append(nArgs, object.String("steve"))
	out = SPC(nil, nArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call with a negative argument.
	//
	var failArgs2 []object.Object
	failArgs2 = append(failArgs2, object.Number(-33))
	out2 := SPC(nil, failArgs2)
	if out2.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}
	if !strings.Contains(out2.String(), "Positive") {
		t.Errorf("Our error message wasn't what we expected")
	}

	//
	// Now do it properly.
	//
	var fArgs []object.Object
	fArgs = append(fArgs, object.Number(17))
	fOut := SPC(nil, fArgs)
	if fOut.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", fOut.String())
	}
	if len(fOut.(*object.StringObject).Value) != 17 {
		t.Errorf("Function returned the wrong length: '%s' - %d", fOut.String(), len(fOut.String()))
	}

	//
	// Now do it properly - int
	//
	var iArgs []object.Object
	iArgs = append(iArgs, object.Number(3))
	iOut := SPC(nil, iArgs)
	if iOut.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", iOut.String())
	}
	if len(iOut.(*object.StringObject).Value) != 3 {
		t.Errorf("Function returned the wrong length: '%s'", fOut.String())
	}

	if (iOut.(*object.StringObject).Value) != "   " {
		t.Errorf("Function returned a surprising result '%s'", iOut.String())
	}

}

func TestStr(t *testing.T) {

	//
	// Call with a non-number argument.
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := STR(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now a string
	//
	var nArgs []object.Object
	nArgs = append(nArgs, object.String("steve"))
	out = STR(nil, nArgs)
	if out.Type() != object.STRING {
		t.Errorf("We expected a string, but didn't receive one")
	}

	//
	// Now do it properly - float
	//
	var fArgs []object.Object
	fArgs = append(fArgs, object.Number(17.8))
	fOut := STR(nil, fArgs)
	if fOut.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", fOut.String())
	}
	if !strings.HasPrefix(fOut.(*object.StringObject).Value, "17.8") {
		t.Errorf("Function returned a surprising result for float: %s", fOut.String())
	}

	//
	// Now do it properly - int
	//
	var iArgs []object.Object
	iArgs = append(iArgs, object.Number(99))
	iOut := STR(nil, iArgs)
	if iOut.Type() != object.STRING {
		t.Errorf("We expected a string return, but didn't get one: %s", iOut.String())
	}
	if !strings.HasPrefix(iOut.(*object.StringObject).Value, "99") {
		t.Errorf("Function returned a surprising result for int: %s", iOut.String())
	}

}

func TestTl(t *testing.T) {

	// Call with a non-string argument.
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := TL(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	// Now test with some valid strings.
	//
	// Setup a structure for testing.
	//
	type TLTest struct {
		Input  string
		Output string
	}

	// Define some tests
	tests := []TLTest{{Input: "Steve", Output: "teve"},
		{Input: "", Output: ""},
		{Input: "ウェブの国際化", Output: "ェブの国際化"},
	}

	for _, test := range tests {

		var args []object.Object
		args = append(args, object.String(test.Input))
		output := TL(nil, args)
		if output.Type() != object.STRING {
			t.Errorf("We expected a string-result, but got something else")
		}
		if output.(*object.StringObject).Value != test.Output {
			t.Errorf("TL %s gave '%s' not '%s'",
				test.Input, output.(*object.StringObject).Value, test.Output)
		}
	}
}

func TestVal(t *testing.T) {

	// Inputs
	num := object.Number(3.2)
	err := object.Error("Error")
	str := object.String("3.11")

	// err
	var eArr []object.Object
	eArr = append(eArr, err)
	eOut := VAL(nil, eArr)
	if eOut.Type() != object.ERROR {
		fmt.Printf("Failed to find error")
	}

	var nArr []object.Object
	nArr = append(nArr, num)
	nOut := VAL(nil, nArr)
	if nOut.Type() != object.NUMBER {
		fmt.Printf("Failed to find number")
	}

	// str - this should become a number
	var sArr []object.Object
	sArr = append(sArr, str)
	sOut := VAL(nil, sArr)
	if sOut.Type() != object.NUMBER {
		fmt.Printf("Failed to convert string to number")
	}

	// invalid input - this should become an error
	var fArr []object.Object
	fArr = append(fArr, object.String("Not a number!"))
	fOut := VAL(nil, fArr)
	if fOut.Type() != object.ERROR {
		fmt.Printf("Error-handling failed")
	}
}
