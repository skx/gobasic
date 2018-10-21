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

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/token"
)

// tokenToFloat is a helper for getting the value of a token as a floating
// point number.
//
// If we're given a literal we return it.  Otherwise we look it up from
// our variable-store and validate the type is correct.
func tokenToFloat(env Interpreter, tok token.Token) (float64, error) {
	var i float64
	var err error

	// We were given a literal integer as an argument - get it in i.
	if tok.Type == token.INT {

		// Convert from string.  (Yeah.)
		i, err = strconv.ParseFloat(tok.Literal, 64)
		if err != nil {
			return 0.0, err
		}
	}

	// We were given a variable as an argument.
	if tok.Type == token.IDENT {

		// Get the variable
		value := env.GetVariable(tok.Literal)

		// Ensure it is a number
		if value.Type() != object.NUMBER {
			return 0.0, fmt.Errorf("Wrong type for variable %s - received %s", tok.Literal, value.Type())
		}

		i = value.(*object.NumberObject).Value
	}

	return i, nil

}

// ABS implements ABS
func ABS(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	// If less than zero make it positive.
	if i < 0 {
		return &object.NumberObject{Value: -1 * i}, nil
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: i}, nil
}

// INT implements INT
func INT(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	// Truncate.
	return &object.NumberObject{Value: float64(int(i))}, nil
}

// RND implements RND
func RND(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	max, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	// Truncate.
	return &object.NumberObject{Value: float64(rand.Intn(int(max)))}, nil
}

// SGN is the sign function (sometimes called signum).
func SGN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	// If less than zero make it positive.
	if i < 0 {
		return &object.NumberObject{Value: -1 * i}, nil
	}

	if i == 0 {
		return &object.NumberObject{Value: 0}, nil
	}
	if i < 0 {
		return &object.NumberObject{Value: -1}, nil
	}
	return &object.NumberObject{Value: 1}, nil

}

// SQR implements square root.
func SQR(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Sqrt(i)}, nil
}

// PI returns the value of PI
func PI(env Interpreter, args []token.Token) (object.Object, error) {

	return &object.NumberObject{Value: math.Pi}, nil
}

// COS implements the COS function..
func COS(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Cos(i)}, nil
}

// SIN operats the sin function.
func SIN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Sin(i)}, nil
}

// TAN implements the tan function.
func TAN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Tan(i)}, nil
}

// ASN (arcsine)
func ASN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Asin(i)}, nil
}

// ACS (arccosine)
func ACS(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Acos(i)}, nil
}

// ATN (arctan)
func ATN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Atan(i)}, nil
}

// EXP x=e^x EXP
func EXP(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Exp(i)}, nil
}

// LN calculates logarithms to the base e - LN
func LN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := tokenToFloat(env, args[0])
	if err != nil {
		return nil, err
	}

	return &object.NumberObject{Value: math.Log(i)}, nil
}
