[![Travis CI](https://img.shields.io/travis/skx/gobasic/master.svg?style=flat-square)](https://travis-ci.org/skx/gobasic)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/gobasic)](https://goreportcard.com/report/github.com/skx/gobasic)
[![license](https://img.shields.io/github/license/skx/gobasic.svg)](https://github.com/skx/gobasic/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/gobasic.svg)](https://github.com/skx/gobasic/releases/latest)

# 10 PRINT "GOBASIC!"

This repository contains a naive implementation of BASIC, written in Golang.

The implementation is simple for three main reasons:

* There is no UI, which means any and all graphics-primitives are ruled out.
* I deliberately set a low bar for myself, originally this was going to be a [weekend project](https://blog.steve.fi/monkeying_around_with_intepreters.html).
  * This is _still_ a weekend-project, but happened over the course of a couple  of hours of evening/morning time instead.
* I didn't implement the full BASIC set of primitives.
  * Not even remotely.

Currently the following primitives work:

* `END`
  * Exit the program.
* `GOTO`
  * Jump to the given line.
* `GOSUB` / `RETURN`
  * Used to call subroutines, via line-indexes.
* `IF` / `THEN` / `ELSE`
  * Conditional execution.
* `INPUT`
  * Allow reading (numeric) input from the user.
* `LET`
  * Assign an integer value to a variable.
* `FOR` & `NEXT`
  * Looping constructs.
* `PRINT`
  * Print a string, an integer, or variable.
  * Multiple arguments may be separated by comma.
* `REM`
  * A single-line comment (BASIC has no notion of multi-line comments).

Most of the maths-related primitives I'm familiar with from my days
coding on a ZX Spectrum are present, for example SIN, COS, PI, ABS,
however __none__ of the string-related primitives are present.



## Limitations

This is a quick hack, so there are some (important) limitations:

* Only a single statement is allowed upon each line.
* Only a subset of the language is implemented.
  * I allow assignment, prints, loops, and control-flow primitives.
  * All string operations are missing.
  * Most of the math-related operations are present, but there may be omissions depending upon the BASIC dialect you're familiar with.
* Only floating-point and string values are parsed.
  * Strings can only be used literally, not stored in a variable.

The handling of the IF statement is perhaps a little unusual, since I'm
used to the BASIC provided by the ZX Spectrum which had no ELSE clause!

The general form of the IF statement is:

    IF $CONDITIONAL THEN $STATEMENT1 [ELSE $STATEMENT2]

Only a single statement is permitted between "THEN" and "ELSE", and again between "ELSE" and NEWLINE.  These are valid IF statements:

    IF 1 > 0 THEN PRINT "OK"
    IF 1 > 3 THEN PRINT "SOMETHING IS BROKEN": ELSE PRINT "Weird!"

In that second example you see that "`:`" was used to terminate the `PRINT` statement, which otherwise would have tried to consume all input until it hit a newline.


## Installation

Providing you have a working [go-installation](https://golang.org/) you should be able to install this software by running:

    go get -u github.com/skx/gobasic

If you don't have a golang environment setup you should be able to download a binary from the github release page:

* [Binary Release](https://github.com/skx/gobasic/releases)



## Usage

gobasic is very simple, and just requires the name of a BASIC-program to
execute.  Write your input in a file and invoke `gobasic` with the path.

For example the following program was useful to test my implementation of the `TOTO` primitive:

     10 GOTO 80
     20 GOTO 70
     30 GOTO 60
     40 PRINT "Hello, world!\n"
     50 END
     60 GOTO 40
     70 GOTO 30
     80 GOTO 20

Execute it like this:

    $ gobasic examples/10-goto.bas

**NOTE**: I feel nostalgic seeing keywords in upper-case, but `PRINT` and `print` are treated identically.  (As is "`PrInT`" for that matter!)



## Implementation

A traditional interpreter for a scripting/toy language would have a series of
well-defined steps:

* Split the input into a series of tokens ("lexing").
* Parse those tokens and build an abstract syntax tree (AST).
* Walk that tree, evaluating as you go.

As is common with early 8-bit home-computers this implementation is a little more basic:

* We parse the input into a series of tokens, defined in [token/token.go](token/token.go)
  * The parsing happens in [tokenizer/tokenizer.go](tokenizer/tokenizer.go)
* We then __directly__ execute those tokens.
  * The execution happens in [eval/eval.go](eval/eval.go) with a couple of small helpers:
    * [eval/for_loop.go](eval/for_loop.go) holds a simple data-structure for handling `FOR`/`NEXT` loops.
    * [eval/stack.go](eval/stack.go) holds a call-stack to handle `GOSUB`/`RETURN`
    * [eval/vars.go](eval/vars.go) holds all our variable references.
    * We have a facility to allow golang code to be made available to BASIC programs, and we use that facility to implement a bunch of our functions.
    * Specifically we use [eval/builtin-support.go](eval/builtin-support.go) to define a lot of functions in [eval/builtins.go](eval/builtins.go) which allow BASIC to call SIN, ABS, PI, etc.

As there is no AST step errors cannot be detected prior to the execution of programs - because we only hit them after we've started running.



## Sample Code

There are a small number of sample-programs located beneath [examples/](examples/).   These were written in an adhoc fashion to test various parts of the implementation.

Perhaps the best demonstration of the code is the "guessing game" program:

* [examples/99-game.bas](examples/99-game.bas)



## Embedding

The interpreter is designed to be easy to embed into your application(s)
if you're crazy enough to want to do that!

You can see an example in the file [embed/main.go](embed/main.go).

The example defines four new functions:

* `PEEK`
* `POKE`
* `DOT`
* `SAVE`

When the script runs it does some basic variable manipulation and it also
creates a PNG file - the `DOT` function allows your script to create a black
image, and draw red dots on it.  The `SAVE` function writes out the result.

Extending this example to draw circles, boxes, etc, is left as an exercise ;)

Hopefully this example shows that making your own functions available to
BASIC scripts is pretty simple.  (This is how SIN, COS, etc are implemented
in the standalone interpreter.)

A caveat of the current implementation of the embedded functions is that
they only work if called in an expression-context.  So this fails:

   10 POKE 3,33

But this succeeds:

    10 LET bogus = POKE 3,33

This might be fixed in the future.



## Bugs?

Probably.  Good luck!

The code _does_ contain a number of test-cases.  You can exercise them via:

    $ go test ./...

Test coverage should exceed 90%, as you can verify and view via:

    $ go test -coverprofile=c.out ./...
    $ go tool cover -html=c.out


Steve
--
