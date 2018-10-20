// builtins.go - Implementation of several built-in functions.
//
// This is where we'll add the missing math-functions, etc.
//

package eval

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/skx/gobasic/token"
)

// ABS implements ABS
func ABS(env Variables, args []token.Token) (int, error) {

	// We were given a literal integer as an argument
	if args[0].Type == token.INT {

		// Conver from string.  (Yeah.)
		i, err := strconv.Atoi(args[0].Literal)
		if err != nil {
			return 0, err
		}

		if i < 0 {
			return (-1 * i), nil
		}
		return i, nil
	}

	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		iVal, ok := val.(int)
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
func INT(env Variables, args []token.Token) (int, error) {

	// Truncate the given float to an int.
	if args[0].Type == token.INT {
		i, err := strconv.Atoi(args[0].Literal)
		if err != nil {
			return 0, err
		}

		return int(i), nil
	}

	// We were given a variable as an argument.
	if args[0].Type == token.IDENT {

		// Get.
		val := env.Get(args[0].Literal)

		// Cast.
		iVal, ok := val.(int)
		if !ok {
			return 0, fmt.Errorf("Error casting variable '%s' to int", args[0].Literal)
		}

		return iVal, nil
	}

	return 0, fmt.Errorf("Invalid type in input argument: %v\n", args[0])
}

// RND implements RND
func RND(env Variables, args []token.Token) (int, error) {
	return rand.Intn(100), nil
}

// SGN is the sign function (sometimes called signum). It is the first function you have seen that has nothing to do with strings, because both its argument and its result are numbers. The result is +1 if the argument is positive, 0 if the argument is zero, and -1 if the argument is negative.
// INT implements INT
func SGN(env Variables, args []token.Token) (int, error) {

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
		return 0, nil
	}
	if i < 0 {
		return -1, nil
	}
	return 1, nil
}

func SQR(env Variables, args []token.Token) (int, error) {

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

	i++
	// TODO: Need floats
	//	return math.Sqrt(i), nil
	return 0, nil
}

func PI(env Variables, args []token.Token) (int, error) {

	// TODO: Need floats
	//	return math.Pi, nil
	return 0, nil
}

//
// Functions I'm missing
//
// TODO: Need floats.
//

// TODO: SQN: Square root
// TODO: PI: PI
// TODO: EXP x=e^x EXP
// TODO: LN which calculates logarithms to the base e - LN

// TODO: SIN
// TODO: COS
// TODO: TAN

// TODO: ASN (arcsine)
// TODO: ACS (arccosine )
// TODO: ATN (arctan)
