package eval

import "github.com/skx/gobasic/tokenizer"

func Fuzz(data []byte) int {

	tokener := tokenizer.New(string(data))
	e := New(tokener)
	e.Run()
	return 1

}
