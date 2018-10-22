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
	"github.com/skx/gobasic/token"
)

// init ensures that we've initialized our random-number state
func init() {
	rand.Seed(time.Now().UnixNano())
}

// TokenToFloat is a helper for getting the value of a token as a floating
// point number.
//
// If we're given a literal we return it.  Otherwise we look it up from
// our variable-store and validate the type is correct.
func TokenToFloat(env Interpreter, tok token.Token) (float64, error) {
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

// TokenToString is a helper for getting the value of a token as a string.
//
// If we're given a literal we return it.  Otherwise we look it up from
// our variable-store and validate the type is correct.
func TokenToString(env Interpreter, tok token.Token) (string, error) {

	// We were given a literal string as an argument return it.
	if tok.Type == token.STRING {
		return tok.Literal, nil
	}

	// We were given a variable as an argument.
	if tok.Type == token.IDENT {

		// Get the variable
		value := env.GetVariable(tok.Literal)

		// Ensure it is a string
		if value.Type() != object.STRING {
			return "", fmt.Errorf("Wrong type for variable %s - received %s", tok.Literal, value.Type())
		}

		return value.(*object.StringObject).Value, nil
	}

	// Can't happen?
	return "", nil
}

// ABS implements ABS
func ABS(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	// If less than zero make it positive.
	if i < 0 {
		return &object.NumberObject{Value: -1 * i}, nil
	}

	// Otherwise return as-is.
	return &object.NumberObject{Value: i}, nil
}

// BIN converts a number from binary.
func BIN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	// hack
	s := fmt.Sprintf("%d", int(i))

	b, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}
	return &object.NumberObject{Value: float64(b)}, nil

}

// CHR returns the character specified by the given ASCII code.
func CHR(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	// Now
	r := rune(i)

	return &object.StringObject{Value: string(r)}, nil
}

// CODE returns the integer value of the specified character.
func CODE(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (string) argument.
	i, err := TokenToString(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	if len(i) > 0 {
		s := i[0]
		return &object.NumberObject{Value: float64(rune(s))}, nil
	}
	return &object.NumberObject{Value: float64(0)}, nil

}

// INT implements INT
func INT(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	// Truncate.
	return &object.NumberObject{Value: float64(int(i))}, nil
}

// LEFT returns the N left-most characters of the string.
func LEFT(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the string
	in, err := TokenToString(env, args[0])
	if err != nil {
		return &object.StringObject{Value: ""}, err
	}

	// args[1] == COMMA

	// Get the number of characters to return
	n, err := TokenToFloat(env, args[2])
	if err != nil {
		return nil, err
	}

	if int(n) > len(in) {
		n = float64(len(in))
	}

	left := in[0:int(n)]

	return &object.StringObject{Value: left}, nil
}

// LEN returns the length of the given string
func LEN(env Interpreter, args []token.Token) (object.Object, error) {

	in, err := TokenToString(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, nil
	}
	return &object.NumberObject{Value: float64(len(in))}, nil
}

// MID returns the N characters from the given offset
func MID(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the string
	in, err := TokenToString(env, args[0])
	if err != nil {
		return &object.StringObject{Value: ""}, err
	}

	// args[1] == COMMA

	// Get the number of characters to return
	offset, err := TokenToFloat(env, args[2])
	if err != nil {
		return nil, err
	}

	// args[3] == COMMA

	count, err := TokenToFloat(env, args[4])
	if err != nil {
		return nil, err
	}

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
func RIGHT(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the string
	in, err := TokenToString(env, args[0])
	if err != nil {
		return &object.StringObject{Value: ""}, err
	}

	// args[1] == COMMA

	// Get the number of characters to return
	n, err := TokenToFloat(env, args[2])
	if err != nil {
		return nil, err
	}

	if int(n) > len(in) {
		n = float64(len(in))
	}
	right := in[len(in)-int(n):]

	return &object.StringObject{Value: right}, nil
}

// RND implements RND
func RND(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	max, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	// Return the random number
	return &object.NumberObject{Value: float64(rand.Intn(int(max)))}, nil
}

// SGN is the sign function (sometimes called signum).
func SGN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	if i < 0 {
		return &object.NumberObject{Value: -1}, nil
	}
	if i == 0 {
		return &object.NumberObject{Value: 0}, nil
	}
	return &object.NumberObject{Value: 1}, nil

}

// SQR implements square root.
func SQR(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Sqrt(i)}, nil
}

// TL returns a string, minus the first character.
func TL(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the string
	in, err := TokenToString(env, args[0])
	if err != nil {
		return &object.StringObject{Value: ""}, err
	}

	if len(in) > 1 {
		rest := in[1:]

		return &object.StringObject{Value: rest}, nil
	}
	return &object.StringObject{Value: ""}, nil
}

// PI returns the value of PI
func PI(env Interpreter, args []token.Token) (object.Object, error) {
	return &object.NumberObject{Value: math.Pi}, nil
}

// COS implements the COS function..
func COS(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Cos(i)}, nil
}

// SIN operats the sin function.
func SIN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Sin(i)}, nil
}

// TAN implements the tan function.
func TAN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Tan(i)}, nil
}

// ASN (arcsine)
func ASN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Asin(i)}, nil
}

// ACS (arccosine)
func ACS(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Acos(i)}, nil
}

// ATN (arctan)
func ATN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Atan(i)}, nil
}

// EXP x=e^x EXP
func EXP(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Exp(i)}, nil
}

// LN calculates logarithms to the base e - LN
func LN(env Interpreter, args []token.Token) (object.Object, error) {

	// Get the (float) argument.
	i, err := TokenToFloat(env, args[0])
	if err != nil {
		return &object.NumberObject{Value: 0}, err
	}

	return &object.NumberObject{Value: math.Log(i)}, nil
}
