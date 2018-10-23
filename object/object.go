// Package object contains code to store values passed to/from BASIC.
//
// Go allows a rich number of types, but when interpreting BASIC only
// two types are supported: Numbers and Strings.
//
// Numbers are stored as `float64`, to allow holding both integers and
// floating-point numbers.  When it comes to output our interpreter will
// round values that are int-like to avoid showing "3.0000" when "3"
// would be sufficient.
package object

import "fmt"

// Type describes the type of an object.
type Type string

// These are our object-types.
const (
	ERROR  = "ERROR"
	NUMBER = "NUMBER"
	STRING = "STRING"
)

// Object is the interface that our types must implement.
type Object interface {

	// Type returns the type of the object.
	Type() Type
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

// NumberObject holds a number.
type NumberObject struct {

	// Value is the value our object wraps.
	Value float64
}

// Type returns the type of this object.
func (s *NumberObject) Type() Type {
	return NUMBER
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
