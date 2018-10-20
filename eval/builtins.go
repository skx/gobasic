// builtins.go - Implementation of several built-in functions.
//
// This is where we'll add the missing math-functions, etc.
//

package eval

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"github.com/skx/gobasic/token"
)

// ABS implements ABS
func ABS(env Interpreter, args []token.Token) (float64, error) {

	// We were given a literal integer as an argument
	if args[0].Type == token.INT {

		// Conver from string.  (Yeah.)
		i, err := strconv.ParseFloat(args[0].Literal, 64)
		if err != nil {
			return 0, err
		}

		if i < 0 {
			return (-1.0 * i), nil
		}
		return i, nil
	}

	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		iVal := env.GetVariable(args[0].Literal)

		// <0 ?
		if iVal < 0 {
			return (-1 * iVal), nil
		}
		return iVal, nil
	}

	return 0, fmt.Errorf("Invalid type in input argument: %v", args[0])
}

// INT implements INT
func INT(env Interpreter, args []token.Token) (float64, error) {

	// Truncate the given float to an int.
	if args[0].Type == token.INT {
		i, err := strconv.ParseFloat(args[0].Literal, 64)
		if err != nil {
			return 0, err
		}

		return float64(int(i)), nil
	}

	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.GetVariable(args[0].Literal)
		return float64(int(val)), nil
	}

	return 0, fmt.Errorf("Invalid type in input argument: %v", args[0])
}

// RND implements RND
func RND(env Interpreter, args []token.Token) (float64, error) {
	return float64(rand.Intn(100)), nil
}

// SGN is the sign function (sometimes called signum).
func SGN(env Interpreter, args []token.Token) (float64, error) {

	var i int

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.Atoi(args[0].Literal)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.GetVariable(args[0].Literal)
		i = int(val)

	}

	if i == 0 {
		return 0.0, nil
	}
	if i < 0 {
		return -1.0, nil
	}
	return 1.0, nil
}

// SQR implements square root.
func SQR(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)
	}

	return math.Sqrt(i), nil
}

// PI returns the value of PI
func PI(env Interpreter, args []token.Token) (float64, error) {

	return math.Pi, nil
}

// COS implements the COS function..
func COS(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {
		// Get.
		i = env.GetVariable(args[0].Literal)

	}

	return math.Cos(i), nil
}

// SIN operats the sin function.
func SIN(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)

	}

	return math.Sin(i), nil
}

// TAN implements the tan function.
func TAN(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {
		// Get.
		i = env.GetVariable(args[0].Literal)
	}

	return math.Tan(i), nil
}

// ASN (arcsine)
func ASN(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)
	}

	return math.Asin(i), nil
}

// ACS (arccosine)
func ACS(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)

	}

	return math.Acos(i), nil
}

// ATN (arctan)
func ATN(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)
	}

	return math.Atan(i), nil
}

// EXP x=e^x EXP
func EXP(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)
	}

	return math.Exp(i), nil
}

// LN calculates logarithms to the base e - LN
func LN(env Interpreter, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		i = env.GetVariable(args[0].Literal)

	}

	return math.Log(i), nil
}
