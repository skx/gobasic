// The builtin package provides the ability to register our built-in functions.
//
// string.go implements our string-related primitives

package builtin

import (
	"fmt"
	"strconv"
	"unicode/utf8"

	"github.com/skx/gobasic/object"
)

// CHR returns the character specified by the given ASCII code.
func CHR(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Now
	r := rune(i)

	return &object.StringObject{Value: string(r)}
}

// CODE returns the integer value of the specified character.
func CODE(env interface{}, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}

	// We convert this to an array of runes because we
	// want to handle unicode strings.
	i := []rune(args[0].(*object.StringObject).Value)

	if len(i) > 0 {
		s := rune(i[0])
		return &object.NumberObject{Value: float64(rune(s))}
	}
	return &object.NumberObject{Value: float64(0)}

}

// LEFT returns the N left-most characters of the string.
func LEFT(env interface{}, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}

	// We convert this to an array of runes because we
	// want to handle unicode strings.
	in := []rune(args[0].(*object.StringObject).Value)

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	n := args[1].(*object.NumberObject).Value

	if int(n) > len(in) {
		n = float64(len(in))
	}

	left := in[0:int(n)]

	return &object.StringObject{Value: string(left)}
}

// LEN returns the length of the given string
func LEN(env interface{}, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	// We need to count in UTF-8 characters.
	len := utf8.RuneCountInString(in)

	return &object.NumberObject{Value: float64(len)}
}

// MID returns the N characters from the given offset
func MID(env interface{}, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}

	// We convert this to an array of runes because we
	// want to handle unicode strings.
	in := []rune(args[0].(*object.StringObject).Value)

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	offset := args[1].(*object.NumberObject).Value

	// Get the (float) argument.
	if args[2].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	count := args[2].(*object.NumberObject).Value

	// too far
	if int(offset) > len(in) {
		return &object.StringObject{Value: ""}
	}

	// get the string from the position
	out := in[int(offset):]

	// now cut, by length
	if int(count) > len(out) {
		count = float64(len(out))
	}
	out = out[:int(count)]
	return &object.StringObject{Value: string(out)}
}

// RIGHT returns the N right-most characters of the string.
func RIGHT(env interface{}, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}

	// We convert this to an array of runes because we
	// want to handle unicode strings.
	in := []rune(args[0].(*object.StringObject).Value)

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	n := args[1].(*object.NumberObject).Value

	if int(n) > len(in) {
		n = float64(len(in))
	}
	right := in[len(in)-int(n):]

	return &object.StringObject{Value: string(right)}
}

// STR converts a number to a string
func STR(env interface{}, args []object.Object) object.Object {

	// Error?
	if args[0].Type() == object.ERROR {
		return args[0]
	}

	// Already a string?
	if args[0].Type() == object.STRING {
		return args[0]
	}

	// Get the value
	var i float64
	i = args[0].(*object.NumberObject).Value
	s := ""

	if i == float64(int(i)) {
		s = fmt.Sprintf("%d", int(i))
	} else {
		s = fmt.Sprintf("%f", i)
	}
	return &object.StringObject{Value: s}
}

// TL returns a string, minus the first character.
func TL(env interface{}, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}

	// We convert this to an array of runes because we
	// want to handle unicode strings.
	in := []rune(args[0].(*object.StringObject).Value)

	if len(in) > 1 {
		rest := in[1:]

		return &object.StringObject{Value: string(rest)}
	}
	return &object.StringObject{Value: ""}
}

// VAL converts a string to a number
func VAL(env interface{}, args []object.Object) object.Object {

	// Error?
	if args[0].Type() == object.ERROR {
		return args[0]
	}

	// Already a number?
	if args[0].Type() == object.NUMBER {
		return args[0]
	}

	// Get the value
	s := args[0].(*object.StringObject).Value
	b, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return object.Error("VAL: %s", err.Error())
	}

	return &object.NumberObject{Value: float64(b)}
}
