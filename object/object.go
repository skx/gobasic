// Package object contains code to store values passed too/from BASIC.
//
// Our values are either numbers (float64) or strings.
package object

// Type describes the type of an object.
type Type string

// pre-defined constant ObjectType
const (
	NUMBER = "NUMBER"
	STRING = "STRING"
)

// Object is the interface that all of our various objects must implmenet.
type Object interface {

	// Type returns the type of our objects contents.
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

// NumberObject holds a number (float64).
type NumberObject struct {

	// Value is the value our object wraps.
	Value float64
}

// Type returns the type of this object.
func (s *NumberObject) Type() Type {
	return NUMBER
}
