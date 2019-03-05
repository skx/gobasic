// maths_test.go - Simple test-cases for maths-related primitives.

package builtin

import (
	"math"
	"testing"

	"github.com/skx/gobasic/object"
)

func TestABS(t *testing.T) {

	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := ABS(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now test that postivie and negative values work.
	//
	// Setup a structure for testing.
	//
	type ABSTest struct {
		In  float64
		Out float64
	}

	// Define some tests
	tests := []ABSTest{
		{In: 3, Out: 3},
		{In: -3, Out: 3},
	}

	for _, test := range tests {

		args := []object.Object{object.Number(test.In)}
		output := ABS(nil, args)
		if output.Type() != object.NUMBER {
			t.Errorf("We expected a number-result, but got something else")
		}
		if output.(*object.NumberObject).Value != test.Out {
			t.Errorf("Abs %f gave '%f' not '%f'",
				test.In, output.(*object.NumberObject).Value, test.Out)
		}
	}
}

func TestACS(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := ACS(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := ACS(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}

}

func TestASN(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := ASN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := ASN(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}
}

func TestATN(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := ATN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := ATN(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}

}

func TestBIN(t *testing.T) {

	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := BIN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Test an invalid input
	//
	var bogus []object.Object
	bogus = append(bogus, object.Number(2001))
	out = BIN(nil, bogus)
	if out.Type() != object.ERROR {
		t.Errorf("We expected an error, but didn't receive one")
	}

	//
	// Test a valid input
	//
	var valid []object.Object
	valid = append(valid, object.Number(1001))
	out = BIN(nil, valid)
	if out.Type() != object.NUMBER {
		t.Errorf("We expected no error, but got one")
	}
	if out.(*object.NumberObject).Value != 9 {
		t.Errorf("Wrong result for binary conversion of 1001")
	}

}

func TestCOS(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := COS(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := COS(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}
}

func TestEXP(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := EXP(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := EXP(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}
}

func TestINT(t *testing.T) {

	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := INT(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now test that truncation works
	//
	// Setup a structure for testing.
	//
	type INTTest struct {
		In  float64
		Out int
	}

	// Define some tests
	tests := []INTTest{
		{In: 3.1, Out: 3},
		{In: -3.2, Out: -3},
	}

	for _, test := range tests {

		args := []object.Object{object.Number(test.In)}
		output := INT(nil, args)
		if output.Type() != object.NUMBER {
			t.Errorf("We expected a number-result, but got something else")
		}
		if int(output.(*object.NumberObject).Value) != test.Out {
			t.Errorf("INT %f gave '%d' not '%d'",
				test.In, int(output.(*object.NumberObject).Value), test.Out)
		}
	}
}

func TestLN(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := LN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := LN(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}
}

func TestPI(t *testing.T) {

	out := PI(nil, nil)
	if out.Type() != object.NUMBER {
		t.Errorf("Invalid type in return value")
	}
	if out.(*object.NumberObject).Value != math.Pi {
		t.Errorf("Invalid return value")
	}
}

func TestRND(t *testing.T) {

	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := RND(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Requires a POSITIVE argument
	//
	var negs []object.Object
	negs = append(negs, object.Number(-32))
	out = RND(nil, negs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected an error, but didn't receive one")
	}

	//
	// Now call normally.
	//
	var ok []object.Object
	ok = append(ok, object.Number(32))
	out2 := RND(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a number, but didn't receive one")
	}

}

func TestSGN(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := SGN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Now test that sign extension works.
	//
	// Setup a structure for testing.
	//
	type SGNTest struct {
		In  float64
		Out float64
	}

	// Define some tests
	tests := []SGNTest{
		{In: 3.1, Out: 1},
		{In: 0, Out: 0},
		{In: -3.2, Out: -1},
	}

	for _, test := range tests {

		args := []object.Object{object.Number(test.In)}
		output := SGN(nil, args)
		if output.Type() != object.NUMBER {
			t.Errorf("We expected a number-result, but got something else")
		}
		if output.(*object.NumberObject).Value != test.Out {
			t.Errorf("INT %f gave '%f' not '%f'",
				test.In, output.(*object.NumberObject).Value, test.Out)
		}
	}
}

func TestSIN(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := SIN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := SIN(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}
}

func TestSQR(t *testing.T) {

	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := SQR(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Requires a POSITIVE argument
	//
	var negs []object.Object
	negs = append(negs, object.Number(-32))
	out = SQR(nil, negs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected an error, but didn't receive one")
	}

	//
	// Now call normally.
	//
	var ok []object.Object
	ok = append(ok, object.Number(16))
	out2 := SQR(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a number, but didn't receive one")
	}
	if out2.(*object.NumberObject).Value != 4 {
		t.Errorf("Square-root result was wrong")
	}

}

func TestTAN(t *testing.T) {
	//
	// Requires a number argument
	//
	var failArgs []object.Object
	failArgs = append(failArgs, object.Error("Bogus type"))
	out := TAN(nil, failArgs)
	if out.Type() != object.ERROR {
		t.Errorf("We expected a type-error, but didn't receive one")
	}

	//
	// Call normally - ignore the result.  Sorry.
	//
	var ok []object.Object
	ok = append(ok, object.Number(21))
	out2 := TAN(nil, ok)
	if out2.Type() != object.NUMBER {
		t.Errorf("We expected a numeric-result, but didn't receive one.")
	}
}
