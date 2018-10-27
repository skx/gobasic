package eval

import "github.com/skx/gobasic/tokenizer"

func Fuzz(data []byte) int {

	tokener := tokenizer.New(string(data))
	e := New(tokener)
	err := e.Run()
	if err != nil {
		return 0
	}
	return 1

}
