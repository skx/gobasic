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
	"strconv"
	"time"

	"github.com/skx/gobasic/object"
)

// init ensures that we've initialized our random-number state
func init() {
	rand.Seed(time.Now().UnixNano())
}

// DUMP just displays the only argument it received.
func DUMP(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() == object.NUMBER {
		i := args[0].(*object.NumberObject).Value
		fmt.Printf("NUMBER: %f\n", i)
	}
	if args[0].Type() == object.STRING {
		s := args[0].(*object.StringObject).Value
		fmt.Printf("STRING: %s\n", s)
	}
	if args[0].Type() == object.ERROR {
		s := args[0].(*object.ErrorObject).Value
		fmt.Printf("Error: %s\n", s)
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: 0}
}

// PRINT handles displaying strings, integers, and errors.
func PRINT(env Interpreter, args []object.Object) object.Object {

	for _, ent := range args {
		switch ent.Type() {
		case object.NUMBER:
			n := ent.(*object.NumberObject).Value
			if n == float64(int(n)) {
				fmt.Printf("%d", int(n))
			} else {
				fmt.Printf("%f", n)
			}
		case object.STRING:
			fmt.Printf("%s", ent.(*object.StringObject).Value)
		case object.ERROR:
			fmt.Printf("%s", ent.(*object.ErrorObject).Value)
		}
	}

	// Return the count of values we printed.
	return &object.NumberObject{Value: float64(len(args))}
}

// ABS implements ABS
func ABS(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// If less than zero make it positive.
	if i < 0 {
		return &object.NumberObject{Value: -1 * i}
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: i}
}

// BIN converts a number from binary.
func BIN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	s := fmt.Sprintf("%d", int(i))

	b, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return object.Error("BIN:%s", err.Error())
	}

	return &object.NumberObject{Value: float64(b)}

}

// CHR returns the character specified by the given ASCII code.
func CHR(env Interpreter, args []object.Object) object.Object {

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
func CODE(env Interpreter, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.StringObject).Value

	if len(i) > 0 {
		s := i[0]
		return &object.NumberObject{Value: float64(rune(s))}
	}
	return &object.NumberObject{Value: float64(0)}

}

// INT implements INT
func INT(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Truncate.
	return &object.NumberObject{Value: float64(int(i))}
}

// LEFT returns the N left-most characters of the string.
func LEFT(env Interpreter, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	n := args[1].(*object.NumberObject).Value

	if int(n) > len(in) {
		n = float64(len(in))
	}

	left := in[0:int(n)]

	return &object.StringObject{Value: left}
}

// LEN returns the length of the given string
func LEN(env Interpreter, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	return &object.NumberObject{Value: float64(len(in))}
}

// MID returns the N characters from the given offset
func MID(env Interpreter, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

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
	return &object.StringObject{Value: out}
}

// RIGHT returns the N right-most characters of the string.
func RIGHT(env Interpreter, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	// Get the (float) argument.
	if args[1].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	n := args[1].(*object.NumberObject).Value

	if int(n) > len(in) {
		n = float64(len(in))
	}
	right := in[len(in)-int(n):]

	return &object.StringObject{Value: right}
}

// RND implements RND
func RND(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Ensure it is valid.
	if i < 1 {
		return object.Error("Argument to RND must be >0")
	}

	// Return the random number
	return &object.NumberObject{Value: float64(rand.Intn(int(i)))}
}

// SGN is the sign function (sometimes called signum).
func SGN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	if i < 0 {
		return &object.NumberObject{Value: -1}
	}
	if i == 0 {
		return &object.NumberObject{Value: 0}
	}
	return &object.NumberObject{Value: 1}

}

// SQR implements square root.
func SQR(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Sqrt(i)}
}

// TL returns a string, minus the first character.
func TL(env Interpreter, args []object.Object) object.Object {

	// Get the (string) argument.
	if args[0].Type() != object.STRING {
		return object.Error("Wrong type")
	}
	in := args[0].(*object.StringObject).Value

	if len(in) > 1 {
		rest := in[1:]

		return &object.StringObject{Value: rest}
	}
	return &object.StringObject{Value: ""}
}

// PI returns the value of PI
func PI(env Interpreter, args []object.Object) object.Object {
	return &object.NumberObject{Value: math.Pi}
}

// COS implements the COS function..
func COS(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Cos(i)}
}

// SIN operats the sin function.
func SIN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Sin(i)}
}

// TAN implements the tan function.
func TAN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Tan(i)}
}

// ASN (arcsine)
func ASN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Asin(i)}
}

// ACS (arccosine)
func ACS(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Acos(i)}
}

// ATN (arctan)
func ATN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Atan(i)}
}

// EXP x=e^x EXP
func EXP(env Interpreter, args []object.Object) object.Object {
	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Exp(i)}
}

// LN calculates logarithms to the base e - LN
func LN(env Interpreter, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Log(i)}
}

// VAL converts a string to a number
func VAL(env Interpreter, args []object.Object) object.Object {

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

// STR converts a number to a string
func STR(env Interpreter, args []object.Object) object.Object {

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
