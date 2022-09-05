//go:build go1.18
// +build go1.18

package eval

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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

	//
	// Load each of our examples as a seed too.
	//
	files, err := filepath.Glob("../examples/*.bas")
	if err == nil {

		// For each example
		for _, name := range files {
			var data []byte

			// Read the contents
			data, err = os.ReadFile(name)

			if err == nil {
				// If no error then seed.
				fmt.Printf("Seeded with %s\n", name)
				f.Add(data)
			}
		}
	}

	f.Fuzz(func(t *testing.T, input []byte) {

		//
		// Expected errors, caused by bad syntax,
		// invalid types, etc.
		//
		// We hope that the fuzz-tester will result
		// in a panic, or error, but we know that
		// some programs that are malformed aren't
		// actually worth aborting for.
		//
		// For example this program:
		//
		//  10 PRINT "STEVE"
		//  20 GOTO 100
		//
		// Is invalid, as there is no line 100.  That's
		// not something the fuzz-tester should regard as
		// an interesting result.
		//
		// Similarly this is gonna cause an error:
		//
		//  10 GOTO 10
		//
		// Because it'll get caught by the timeout we've defined,
		// but that's not something we regard as interesting either.
		//
		expected := []string{
			"expect an integer",
			"got token",
			"access out of bounds",
			"argument count mis-match",
			"def fn: expected ",
			"dimension too large",
			"division by zero",
			"don't support operations",
			"mod 0 is an error",
			"object is not an array",
			"only handles string-prompts",
			"unclosed bracket around",
			"wrong type",
			"array indexes must be",
			"does not exist",
			"doesn't exist",
			"end of program processing",
			"expected ident after ",
			"expected assignment",
			"expected identifier",
			"index out of range",
			"input should be",
			"invalid prompt-type",
			"length of strings cannot exceed",
			"must be an integer",
			"must be >0",
			"next variable",
			"nil terminal",
			"not supported for strings",
			"only handles string-multiplication and integer-operations",
			"only integers are used for dimensions",
			"positive argument only",
			"read past the end of our data storage",
			"received a nil value",
			"return without gosub",
			"setarrayvariable",
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
		// Setup a timeout period to avoid infinite loops.
		//
		ctx, cancel := context.WithTimeout(
			context.Background(),
			500*time.Millisecond,
		)
		defer cancel()

		// Tokenize
		toker := tokenizer.New(string(input))

		// Prepare to run
		e, err := NewWithContext(ctx, toker)
		if err != nil {

			// Lower case the error
			fail := strings.ToLower(err.Error())

			// Is this failure a known one?  Then return
			for _, txt := range expected {
				if strings.Contains(fail, txt) {
					return
				}
			}

			// This is a panic caused by the fuzzer.
			// Report it.
			panic(fmt.Sprintf("input %s gave error %s", input, err))
		}

		// Run
		err = e.Run()

		if err != nil {

			// Lower case the error
			fail := strings.ToLower(err.Error())

			// Is this failure a known one?  Then return
			for _, txt := range expected {
				if strings.Contains(fail, txt) {
					return
				}
			}

			// This is a panic caused by the fuzzer.
			// Report it.
			panic(fmt.Sprintf("input %s gave error %s", input, err))
		}
	})
}
