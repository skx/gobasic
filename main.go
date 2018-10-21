package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/skx/gobasic/eval"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

// This version-string will be updated via travis for generated binaries.
var version = "master/unreleased"

func main() {

	//
	// Ensure we seed our random-number source
	//
	// This is required such that RND() returns suitable values.
	//
	rand.Seed(time.Now().UnixNano())

	//
	// Setup some command-line flags
	//
	lex := flag.Bool("lex", false, "Show the output of the lexer.")
	vers := flag.Bool("version", false, "Show our version and exit.")

	//
	// Parse the flags
	//
	flag.Parse()

	//
	// Showing the version?
	//
	if *vers {
		fmt.Printf("gobasic %s\n", version)
		os.Exit(1)
	}

	//
	// Test we have a file to interpret
	//
	if len(flag.Args()) != 1 {
		fmt.Printf("Usage: gobasic /path/to/input/script.bas\n")
		os.Exit(2)
	}

	//
	// Load the file.
	//
	data, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Printf("Error reading %s - %s\n", flag.Args()[0], err.Error())
		os.Exit(3)
	}

	//
	// Tokenize
	//
	t := tokenizer.New(string(data))

	//
	// Are we dumping tokens?
	//
	if *lex {
		for {
			tok := t.NextToken()
			if tok.Type == token.EOF {
				break
			}
			fmt.Printf("%v\n", tok)
		}
		os.Exit(0)
	}

	//
	// OK we're running the code.  Do it.
	//
	e := eval.New(t)
	err = e.Run()
	if err != nil {
		fmt.Printf("Error running program: %s\n", err.Error())
	}
}
