// Package object contains code to store values passed to/from BASIC.
//
// Go allows a rich number of types, but when interpreting BASIC programs
// we only support numbers & strings, as well as two-dimensional arrays
// containing those values.
//
// Note that numbers are stored as `float64`, to allow holding both
// integers and floating-point numbers.
//
package object

import (
	"fmt"
)

// Type describes the type of an object.
type Type string

// These are our object-types.
const (
	ERROR  = "ERROR"
	NUMBER = "NUMBER"
	STRING = "STRING"
	ARRAY  = "ARRAY"
)

// Object is the interface that our types must implement.
type Object interface {

	// Type returns the type of the object.
	Type() Type

	// String converts the object to a printable version for debugging.
	String() string
}

// ArrayObject holds an array.
//
// We allow only two-dimensional arrays, and the size is set at the time
// the array is constructed.
type ArrayObject struct {

	// We store objects in our array
	Contents []Object

	// X is the X-size of the array, fixed at creation-time
	X int

	// Y is the Y-size of the array, fixed at creation-time.
	Y int
}

// Type returns the type of this object.
func (a *ArrayObject) Type() Type {
	return ARRAY
}

// Array creates a new array of the given dimensions
func Array(x int, y int) *ArrayObject {

	// Our semantics ensure that we allow "0-N".
	if x != 0 {
		x++
	}
	if y != 0 {
		y++
	}

	// setup the sizes
	a := &ArrayObject{X: x, Y: y}

	// for each entry ensure we store a value.
	var c int
	if x == 0 {
		c = y
	} else {
		c = x * y
	}

	// we default to "0"
	for c >= 0 {
		a.Contents = append(a.Contents, Number(0))
		c--
	}

	return a
}

// Get the value at the given X,Y coordinate
func (a *ArrayObject) Get(x int, y int) Object {
	offset := int(x*a.X + y)

	if a.X == 0 && offset >= a.Y {
		return &ErrorObject{Value: "Get-Array access out of bounds (Y)"}
	}
	if (a.X != 0) && (offset > a.X*a.Y) {
		return &ErrorObject{Value: "Get-Array access out of bounds (X,Y)"}
	}
	if offset < 0 {
		return &ErrorObject{Value: "Get-Array access out of bounds (negative index)"}
	}
	if offset > len(a.Contents) {
		return &ErrorObject{Value: "Get-Array access out of bounds (LEN)"}
	}
	return (a.Contents[offset])
}

// Set the value at the given X,Y coordinate
func (a *ArrayObject) Set(x int, y int, obj Object) Object {
	offset := int(x*a.X + y)

	if a.X == 0 && offset >= a.Y {
		return &ErrorObject{Value: "Set-Array access out of bounds (Y)"}
	}
	if (a.X != 0) && (offset > a.X*a.Y) {
		return &ErrorObject{Value: "Set-Array access out of bounds (X,Y)"}
	}
	if offset < 0 {
		return &ErrorObject{Value: "Set-Array access out of bounds (negative index)"}
	}
	if offset > len(a.Contents) {
		return &ErrorObject{Value: "Set-Array access out of bounds (LEN)"}
	}

	a.Contents[offset] = obj
	return obj
}

// String returns the string-contents of the string
func (a *ArrayObject) String() string {

	out := fmt.Sprintf("Array{X:%d, Y:%d, <%v>}",
		a.X, a.Y, a.Contents)
	return (out)
}

// StringObject holds a string.
type StringObject struct {

	// Value is the value our object wraps.
	Value string
}

// Type returns the type of this object.
func (s *StringObject) Type() Type {
	return STRING
}

// String returns a string representation of this object.
func (s *StringObject) String() string {
	return (fmt.Sprintf("Object{Type:string, Value:%s}", s.Value))
}

// String is a helper for creating a new string-object with the given value.
func String(val string) *StringObject {
	return &StringObject{Value: val}
}

// NumberObject holds a number.
type NumberObject struct {

	// Value is the value our object wraps.
	Value float64
}

// Type returns the type of this object.
func (s *NumberObject) Type() Type {
	return NUMBER
}

// String returns a string representation of this object.
func (s *NumberObject) String() string {
	return (fmt.Sprintf("Object{Type:number, Value:%f}", s.Value))
}

// Number is a helper for creating a new number-object with the given value.
func Number(val float64) *NumberObject {
	return &NumberObject{Value: val}
}

// ErrorObject holds a string, which describes an error
type ErrorObject struct {

	// Value is the message our object wraps.
	Value string
}

// Error is a helper for creating a new error-object with the given message.
func Error(format string, args ...interface{}) *ErrorObject {
	msg := fmt.Sprintf(format, args...)
	return &ErrorObject{Value: msg}
}

// Type returns the type of this object.
func (s *ErrorObject) Type() Type {
	return ERROR
}

// String returns a string representation of this object.
func (s *ErrorObject) String() string {
	return (fmt.Sprintf("Object{Type:error, Value:%s}", s.Value))
}
