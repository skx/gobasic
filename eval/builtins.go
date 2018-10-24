// builtins.go - Implementation of several built-in functions.
//
// This is where we implement functions/statements that are outside
// of our core.
//

package eval

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/skx/gobasic/object"
)

// init ensures that we've initialized our random-number state
func init() {
	rand.Seed(time.Now().UnixNano())
}

// DUMP just displays the only argument it received.
func DUMP(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() == object.NUMBER {
		i := args[0].(*object.NumberObject).Value
		fmt.Printf("NUMBER: %f\n", i)
	}
	if args[0].Type() == object.STRING {
		s := args[0].(*object.StringObject).Value
		fmt.Printf("STRING: %s\n", s)
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: 0}, nil
}

// ABS implements ABS
func ABS(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// If less than zero make it positive.
	if i < 0 {
		return &object.NumberObject{Value: -1 * i}, nil
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: i}, nil
}

// BIN converts a number from binary.
func BIN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// TODO!
	return &object.NumberObject{Value: float64(i)}, nil

}

// CHR returns the character specified by the given ASCII code.
func CHR(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Now
	r := rune(i)

	return &object.StringObject{Value: string(r)}, nil
}

// CODE returns the integer value of the specified character.
func CODE(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.StringObject).Value

	if len(i) > 0 {
		s := i[0]
		return &object.NumberObject{Value: float64(rune(s))}, nil
	}
	return &object.NumberObject{Value: float64(0)}, nil

}

// INT implements INT
func INT(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Truncate.
	return &object.NumberObject{Value: float64(int(i))}, nil
}

// LEFT returns the N left-most characters of the string.
func LEFT(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	n := args[1].(*object.NumberObject).Value

	if int(n) > len(in) {
		n = float64(len(in))
	}

	left := in[0:int(n)]

	return &object.StringObject{Value: left}, nil
}

// LEN returns the length of the given string
func LEN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	return &object.NumberObject{Value: float64(len(in))}, nil
}

// MID returns the N characters from the given offset
func MID(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	offset := args[1].(*object.NumberObject).Value

	// Get the (float) argument.
	if args[2].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	count := args[2].(*object.NumberObject).Value

	// too far
	if int(offset) > len(in) {
		return &object.StringObject{Value: ""}, nil
	}

	// get the string from the position
	out := in[int(offset):]

	// now cut, by length
	if int(count) > len(out) {
		count = float64(len(out))
	}
	out = out[:int(count)]
	return &object.StringObject{Value: out}, nil
}

// RIGHT returns the N right-most characters of the string.
func RIGHT(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	n := args[1].(*object.NumberObject).Value

	if int(n) > len(in) {
		n = float64(len(in))
	}
	right := in[len(in)-int(n):]

	return &object.StringObject{Value: right}, nil
}

// RND implements RND
func RND(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Return the random number
	return &object.NumberObject{Value: float64(rand.Intn(int(i)))}, nil
}

// SGN is the sign function (sometimes called signum).
func SGN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	if i < 0 {
		return &object.NumberObject{Value: -1}, nil
	}
	if i == 0 {
		return &object.NumberObject{Value: 0}, nil
	}
	return &object.NumberObject{Value: 1}, nil

}

// SQR implements square root.
func SQR(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Sqrt(i)}, nil
}

// TL returns a string, minus the first character.
func TL(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	if len(in) > 1 {
		rest := in[1:]

		return &object.StringObject{Value: rest}, nil
	}
	return &object.StringObject{Value: ""}, nil
}

// PI returns the value of PI
func PI(env Interpreter, args []object.Object) (object.Object, error) {
	return &object.NumberObject{Value: math.Pi}, nil
}

// COS implements the COS function..
func COS(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Cos(i)}, nil
}

// SIN operats the sin function.
func SIN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Sin(i)}, nil
}

// TAN implements the tan function.
func TAN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Tan(i)}, nil
}

// ASN (arcsine)
func ASN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Asin(i)}, nil
}

// ACS (arccosine)
func ACS(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Acos(i)}, nil
}

// ATN (arctan)
func ATN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Atan(i)}, nil
}

// EXP x=e^x EXP
func EXP(env Interpreter, args []object.Object) (object.Object, error) {
	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Exp(i)}, nil
}

// LN calculates logarithms to the base e - LN
func LN(env Interpreter, args []object.Object) (object.Object, error) {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type"), fmt.Errorf("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Log(i)}, nil
}
