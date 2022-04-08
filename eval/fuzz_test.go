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
			"end of program processing",
			"DEF FN: expected ",
			"unexpected token",
            "without opening FOR",
			" expect an integer ",
			"unclosed FOR loop",
			" got Token",
			"must be an integer",
			"Argument count mis-match",
			"expected IDENT after ",
			"does not exist",
			"only handles string-multiplication and integer-operations",
			"the variable ",
			"doesn't exist",
			"factor() - unhandled token",
			"expected assignment",
			"Unclosed bracket around",
			"should be followed by an integer",
			"array indexes must be",
			"type mismatch between",
			"unexpected value found when looking for index",
			"MOD 0 is an error",
			"expr() operation '-' not supported for strings",
			"Wrong type",
			"Division by zero",
			"timeout during execution",
			"while searching for argument",
			"Object is not an array",
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

			for _, txt := range expected {
				if strings.Contains(err.Error(), txt) {
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

			for _, txt := range expected {
				if strings.Contains(err.Error(), txt) {
					ignore = true
				}
			}

			if !ignore {
				panic(fmt.Sprintf("input %s gave error %s", input, err))
			}
		}
	})

}
