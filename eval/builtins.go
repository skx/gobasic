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
func ABS(env Variables, args []token.Token) (float64, error) {

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
		val := env.Get(args[0].Literal)

		// Cast.
		iVal, ok := val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to int", args[0].Literal)
		}

		// <0 ?
		if iVal < 0 {
			return (-1 * iVal), nil
		}
		return iVal, nil
	}

	return 0, fmt.Errorf("Invalid type in input argument: %v\n", args[0])
}

// INT implements INT
func INT(env Variables, args []token.Token) (float64, error) {

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
		val := env.Get(args[0].Literal)

		// Cast.
		iVal, ok := val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

		return float64(int(iVal)), nil
	}

	return 0, fmt.Errorf("Invalid type in input argument: %v\n", args[0])
}

// RND implements RND
func RND(env Variables, args []token.Token) (float64, error) {
	return float64(rand.Intn(100)), nil
}

// SGN is the sign function (sometimes called signum). It is the first function you have seen that has nothing to do with strings, because both its argument and its result are numbers. The result is +1 if the argument is positive, 0 if the argument is zero, and -1 if the argument is negative.
// INT implements INT
func SGN(env Variables, args []token.Token) (float64, error) {

	var i int

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.Atoi(args[0].Literal)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(int)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to int", args[0].Literal)
		}

	}

	if i == 0 {
		return 0.0, nil
	}
	if i < 0 {
		return -1.0, nil
	}
	return 1.0, nil
}

// SQR: Square root
func SQR(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Sqrt(i), nil
}

// PI: Return PI
func PI(env Variables, args []token.Token) (float64, error) {

	return math.Pi, nil
}

//
// Functions I'm missing
//
// TODO: Need floats.
//

// COS
func COS(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Cos(i), nil
}

// SIN.
func SIN(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Sin(i), nil
}

// TAN.
func TAN(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Tan(i), nil
}

// ASN (arcsine)
func ASN(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Asin(i), nil
}

// ACS (arccosine)
func ACS(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Acos(i), nil
}

// ATN (arctan)
func ATN(env Variables, args []token.Token) (float64, error) {

	var i float64

	// We were given a literal int.
	if args[0].Type == token.INT {
		i, _ = strconv.ParseFloat(args[0].Literal, 64)
	}
	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		var ok bool
		i, ok = val.(float64)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to float64", args[0].Literal)
		}

	}

	return math.Atan(i), nil
}

// TODO: EXP x=e^x EXP
// TODO: LN which calculates logarithms to the base e - LN
