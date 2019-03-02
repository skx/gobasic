// misc_test.go - Simple test-cases for misc. primitives.

package builtin

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/skx/gobasic/object"
)

type bufferEnv struct {
	writer *bufio.Writer
}

func (b *bufferEnv) StdInput() *bufio.Reader {
	return nil
}

func (b *bufferEnv) StdOutput() *bufio.Writer {
	return b.writer
}

func (b *bufferEnv) Data() interface{} {
	return nil
}

func TestDump(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	env := &bufferEnv{}
	env.writer = bufio.NewWriter(buf)

	//
	// Number
	//
	var in1 []object.Object
	in1 = append(in1, object.Number(1))
	out1 := DUMP(env, in1)
	if out1.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	str := buf.String()
	if str != "NUMBER: 1.000000\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
	buf.Reset()

	//
	// String
	//
	var in2 []object.Object
	in2 = append(in2, object.String("Stve"))
	out2 := DUMP(env, in2)
	if out2.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	str = buf.String()
	if str != "STRING: Stve\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
	buf.Reset()

	//
	// Error
	//
	var in3 []object.Object
	in3 = append(in3, object.Error("Stve"))
	out3 := DUMP(env, in3)
	if out3.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	str = buf.String()
	if str != "Error: Stve\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
}

func TestPrint(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	env := &bufferEnv{}
	env.writer = bufio.NewWriter(buf)

	//
	// Number
	//
	var in1 []object.Object
	in1 = append(in1, object.Number(1))
	out1 := PRINT(env, in1)
	if out1.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out1.(*object.NumberObject).Value != 1 {
		t.Errorf("We didn't print one item: %f",
			out1.(*object.NumberObject).Value)
	}
	str := buf.String()
	if str != "1\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
	buf.Reset()

	//
	// String
	//
	var in2 []object.Object
	in2 = append(in2, object.String("Stve"))
	out2 := PRINT(env, in2)
	if out2.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out2.(*object.NumberObject).Value != 1 {
		t.Errorf("We didn't print one item: %f",
			out2.(*object.NumberObject).Value)
	}
	str = buf.String()
	if str != "Stve\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
	buf.Reset()

	//
	// Error
	//
	var in3 []object.Object
	in3 = append(in3, object.Error("Stve"))
	out3 := PRINT(env, in3)
	if out3.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out3.(*object.NumberObject).Value != 1 {
		t.Errorf("We didn't print one item:%f",
			out3.(*object.NumberObject).Value)
	}
	str = buf.String()
	if str != "Stve\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
	buf.Reset()

	//
	// Now a bunch of things
	//
	var in4 []object.Object
	in4 = append(in4, object.Error("Stve"))
	in4 = append(in4, object.String("Stve"))
	in4 = append(in4, object.Number(3))
	in4 = append(in4, object.Number(4.3))
	out4 := PRINT(env, in4)
	if out4.Type() != object.NUMBER {
		t.Errorf("We didn't receive a number in response")
	}
	if out4.(*object.NumberObject).Value != 4 {
		t.Errorf("We didn't print the expected count of items:%f",
			out4.(*object.NumberObject).Value)
	}
	str = buf.String()
	if str != "StveStve34.300000\n" {
		t.Errorf("We didn't print the correct string: %s", str)
	}
}
