//go:build go1.18
// +build go1.18

package eval

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/skx/gobasic/tokenizer"
)

func FuzzEval(f *testing.F) {
	f.Add([]byte(""))

	// Simple
	f.Add([]byte("10 REM OK"))
	f.Add([]byte("10 PRINT \"foo\"\r"))
	f.Add([]byte("10 LET a = 3 + 4 * 5\r\n"))

	// Broken
	f.Add([]byte("20 PRINT \"incomplete\n"))
	f.Add([]byte("10 GOTO 100\n"))
	f.Add([]byte("10 GOTO 10\xbc\n"))

	// Bigger
	f.Add([]byte(`
	00 REM This program tests GOTO-handling.
	10 GOTO 80
	20 GOTO 70
	30 GOTO 60
	40 PRINT "Hello-GOTO!\n"
	50 END
	60 GOTO 40
	70 GOTO 30
	80 GOTO 20
	`))

	f.Add([]byte(`
 50 DEF FN double(x) = x + x
 60 DEF FN square(x) = x * x
 70 DEF FN cube(x)   = x * x * x
 80 DEF FN quad(x)   = x * x * x * x
 90 PRINT "N\tDoubled\tSquared\tCubed\tQuadded (?)\n"
100 FOR I = 1 TO 10
110   PRINT I, "\t", FN double(I), "\t", FN square(I), "\t", FN cube(I), "\t", FN quad(I), "\n"
120 NEXT I
`))

	f.Fuzz(func(t *testing.T, input []byte) {

		// Expected errors
		expected := []string{
			"expect an integer",
			"got token",
			"argument count mis-match",
			"def fn: expected ",
			"dimension too large",
			"division by zero",
			"mod 0 is an error",
			"object is not an array",
			"unclosed bracket around",
			"wrong type",
			"array indexes must be",
			"does not exist",
			"doesn't exist",
			"end of program processing",
			"expected ident after ",
			"expected assignment",
			"length of strings cannot exceed",
			"must be an integer",
			"not supported for strings",
			"only handles string-multiplication and integer-operations",
			"only integers are used for dimensions",
			"positive argument only",
			"should be followed by an integer",
			"strconv.parse",
			"the variable",
			"timeout during execution",
			"type mismatch between",
			"unclosed for loop",
			"unexpected token",
			"unexpected value found when looking for index",
			"unhandled token",
			"while searching for argument",
			"without opening for",
		}

		//
		// Setup a timeout period to avoid infinite loops
		//
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
		defer cancel()

		// Tokenize
		toker := tokenizer.New(string(input))

		// Prepare to run
		e, err := NewWithContext(ctx, toker)
		if err != nil {

			ignore := false

			// Lower case the error
			fail := strings.ToLower(err.Error())

			for _, txt := range expected {
				if strings.Contains(fail, txt) {
					ignore = true
				}
			}

			if !ignore {
				panic(fmt.Sprintf("input %s gave error %s", input, err))
			}
			return
		}

		// Run
		err = e.Run()

		if err != nil {
			ignore := false

			// Lower case the error
			fail := strings.ToLower(err.Error())

			for _, txt := range expected {
				if strings.Contains(fail, txt) {
					ignore = true
				}
			}

			if !ignore {
				panic(fmt.Sprintf("input %s gave error %s", input, err))
			}
		}
	})

}
