package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

func peekFunction(env eval.Variables, args []token.Token) (float64, error) {
	fmt.Printf("PEEK called with %v\n", args[0])
	return 0, nil
}
func pokeFunction(env eval.Variables, args []token.Token) (float64, error) {
	fmt.Printf("POKE called.\n")
	for i, e := range args {
		fmt.Printf(" Arg %d -> %v\n", i, e)
	}
	return 0, nil
}

func main() {

	//
	// Ensure we seed a random-number source
	//
	// This is required such that RND() returns suitable
	// values that change.
	//
	rand.Seed(time.Now().UnixNano())

	//
	// This is the program we're going to execute
	//
	prog := `
10 PRINT "HELLO, I AM EMBEDDED BASIC\n"
20 LET S = S + PI
30 LET R = POKE 23659 , 0
40 LET n = PEEK 30
50 PRINT "I'M DONE!\n"
`

	//
	// Load the program
	//
	t := tokenizer.New(prog)

	//
	// Create an interpreter
	//
	e := eval.New(t)

	//
	// Register a pair of functions.
	//
	e.RegisterBuiltin("PEEK", 1, peekFunction)
	e.RegisterBuiltin("POKE", 3, pokeFunction)

	//
	// Set an initial value to the variable "S".
	//
	e.SetVariable("S", 3)

	//
	// Run the code.
	//
	err := e.Run()
	if err != nil {
		fmt.Printf("Error running program: %s\n", err.Error())
	}

	//
	// The value of the variable is now different
	//
	fmt.Printf("Output value is %v\n", e.GetVariable("S"))
}
