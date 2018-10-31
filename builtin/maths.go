// The builtin package provides the ability to register our built-in functions.
//
// maths.go implements our math-related primitives

package builtin

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

// ABS implements ABS
func ABS(env interface{}, args []object.Object) object.Object {

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

// ACS (arccosine)
func ACS(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Acos(i)}
}

// ASN (arcsine)
func ASN(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Asin(i)}
}

// ATN (arctan)
func ATN(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Atan(i)}
}

// BIN converts a number from binary.
func BIN(env interface{}, args []object.Object) object.Object {

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

// COS implements the COS function..
func COS(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Cos(i)}
}

// EXP x=e^x EXP
func EXP(env interface{}, args []object.Object) object.Object {
	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Exp(i)}
}

// INT implements INT
func INT(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Truncate.
	return &object.NumberObject{Value: float64(int(i))}
}

// LN calculates logarithms to the base e - LN
func LN(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Log(i)}
}

// PI returns the value of PI
func PI(env interface{}, args []object.Object) object.Object {
	return &object.NumberObject{Value: math.Pi}
}

// RND implements RND
func RND(env interface{}, args []object.Object) object.Object {

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
func SGN(env interface{}, args []object.Object) object.Object {

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

// SIN operats the sin function.
func SIN(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Sin(i)}
}

// SQR implements square root.
func SQR(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	// Ensure it is valid.
	if i < 1 {
		return object.Error("Argument to SQR must be >0")
	}
	return &object.NumberObject{Value: math.Sqrt(i)}
}

// TAN implements the tan function.
func TAN(env interface{}, args []object.Object) object.Object {

	// Get the (float) argument.
	if args[0].Type() != object.NUMBER {
		return object.Error("Wrong type")
	}
	i := args[0].(*object.NumberObject).Value

	return &object.NumberObject{Value: math.Tan(i)}
}
