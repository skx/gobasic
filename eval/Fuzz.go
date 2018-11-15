package eval

import "github.com/skx/gobasic/tokenizer"

// Fuzz is the function that our fuzzer-application uses.
// See `FUZZING.md` in our distribution for how to invoke it.
func Fuzz(data []byte) int {

	tokener := tokenizer.New(string(data))
	e, err := New(tokener)
	if err == nil {
		e.Run()
	}
	return 1

}
