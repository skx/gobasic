// misc_test.go - Simple test-cases for misc. primitives.

package builtin

import (
	"testing"

	"github.com/skx/gobasic/object"
)

func TestDump(t *testing.T) {

	//
	// Number
	//
	var in1 []object.Object
	in1 = append(in1, object.Number(1))
	out1 := DUMP(nil, in1)
	if out1.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}

	//
	// String
	//
	var in2 []object.Object
	in2 = append(in2, object.String("Stve"))
	out2 := DUMP(nil, in2)
	if out2.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}

	//
	// Error
	//
	var in3 []object.Object
	in3 = append(in3, object.Error("Stve"))
	out3 := DUMP(nil, in3)
	if out3.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
}

func TestPrint(t *testing.T) {

	//
	// Number
	//
	var in1 []object.Object
	in1 = append(in1, object.Number(1))
	out1 := PRINT(nil, in1)
	if out1.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out1.(*object.NumberObject).Value != 1 {
		t.Errorf("We didn't print one item: %f",
			out1.(*object.NumberObject).Value)
	}

	//
	// String
	//
	var in2 []object.Object
	in2 = append(in2, object.String("Stve"))
	out2 := PRINT(nil, in2)
	if out2.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out2.(*object.NumberObject).Value != 1 {
		t.Errorf("We didn't print one item: %f",
			out2.(*object.NumberObject).Value)
	}

	//
	// Error
	//
	var in3 []object.Object
	in3 = append(in3, object.Error("Stve"))
	out3 := PRINT(nil, in3)
	if out3.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out3.(*object.NumberObject).Value != 1 {
		t.Errorf("We didn't print one item:%f",
			out3.(*object.NumberObject).Value)
	}

	//
	// Now a bunch of things
	//
	var in4 []object.Object
	in4 = append(in4, object.Error("Stve"))
	in4 = append(in4, object.String("Stve"))
	in4 = append(in4, object.Number(3))
	in4 = append(in4, object.Number(4.3))
	out4 := PRINT(nil, in4)
	if out4.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out4.(*object.NumberObject).Value != 4 {
		t.Errorf("We didn't print the expected count of items:%f",
			out4.(*object.NumberObject).Value)
	}

}
