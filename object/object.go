// Package object contains code to store values passed to/from BASIC.
//
// Go allows a rich number of types, but when interpreting BASIC programs
// only two types are supported: Numbers and Strings.
//
// Numbers are stored as `float64`, to allow holding both integers and
// floating-point numbers.
//
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

	// String converts the object to a printable version for
	// debugging.
	String() string
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
