[![Travis CI](https://img.shields.io/travis/skx/gobasic/master.svg?style=flat-square)](https://travis-ci.org/skx/gobasic)
[![Coverage Status](https://coveralls.io/repos/github/skx/gobasic/badge.svg?branch=master)](https://coveralls.io/github/skx/gobasic?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/skx/gobasic)](https://goreportcard.com/report/github.com/skx/gobasic)
[![license](https://img.shields.io/github/license/skx/gobasic.svg)](https://github.com/skx/gobasic/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/skx/gobasic.svg)](https://github.com/skx/gobasic/releases/latest)



# 05 PRINT "Index"

* [10 PRINT "GOBASIC!"](#10-print-gobasic)
* [20 PRINT "Limitations"](#20-print-limitations)
  * [Arrays](#arrays)
  * [Line Numbers](#line-numbers)
  * [IF Statement](#if-statement)
  * [DATA / READ Statements](#data--read-statements)
  * [Builtin Functions](#builtin-functions)
  * [Types](#types)
* [30 PRINT "Installation"](#30-print-installation)
  * [Build without Go Modules (Go before 1.11)](#build-without-go-modules-go-before-111)
  * [Build with Go Modules (Go 1.11 or higher)](#build-with-go-modules-go-111-or-higher)
* [40 PRINT "Usage"](#40-print-usage)
* [50 PRINT "Implementation"](#50-print-implementation)
* [60 PRINT "Sample Code"](#60-print-sample-code)
* [70 PRINT "Embedding"](#70-print-embedding)
* [80 PRINT "Visual BASIC!"](#80-print-visual-basic)
* [90 PRINT "Bugs?"](#90-print-bugs)
* [100 PRINT "Project Goals / Links"](#100-print-project-goals--links)

<br />
<br />
<br />
<br />

# 10 PRINT "GOBASIC!"

This repository contains a naive implementation of BASIC, written in Golang.

> If you'd prefer to see a more "real" interpreted language, implemented in Go, you might prefer [monkey](https://github.com/skx/monkey/).

The implementation is simple for two main reasons:

* There is no UI, which means any and all graphics-primitives are ruled out.
  * However the embedded sample, described later in this file, demonstrates using BASIC to create a PNG image.
  * There is also a HTTP-based BASIC server, also described later, which allows you to create images "interactively".
* I didn't implement the full BASIC set of primitives.
  * Although most of the commands available to the ZX Spectrum are implemented. I only excluded things relating to tape, PEEK, POKE, etc.
  * If you want to add new BASIC keywords this is easy, and the samples mentioned above do that.

The following obvious primitives work as you'd expect:

* `DIM`
  * Create an array.  Note that only one and two-dimensional arrays are supported.
  * See [examples/95-arrays.bas](examples/95-arrays) and [examples/100-array-sort.bas](examples/100-array-sort.bas) for quick samples.
* `END`
  * Exit the program.
* `GOTO`
  * Jump to the given line.
* `GOSUB` / `RETURN`
  * Used to call the subroutines at the specified line.
* `IF` / `THEN` / `ELSE`
  * Conditional execution.
* `INPUT`
  * Allow reading a string, or number (see later note about types).
* `LET`
  * Assign a value to a variable, creating it if necessary.
* `FOR` & `NEXT`
  * Looping constructs.
* `PRINT`
  * Print a string, an integer, or variable.
  * Multiple arguments may be separated by commas.
* `REM`
  * A single-line comment (BASIC has no notion of multi-line comments).
* `READ` & `DATA`
  * Allow reading from stored data within the program.
  * See [examples/35-read-data.bas](examples/35-read-data.bas) for a demonstration, along with [examples/100-array-sort.bas](examples/100-array-sort.bas).
* `SWAP`
  * Allow swapping the contents of two variables.
  * Useful for sorting arrays, as shown in [examples/100-array-sort.bas](examples/100-array-sort.bas).
* `DEF FN` & `FN`
  * Allow user-defined functions to be defined or invoked.
  * See [examples/25-def-fn.bas](examples/25-def-fn.bas) for an example.

Most of the maths-related primitives I'm familiar with are also present, for example SIN, COS, PI, ABS, along with the similar string-related primitives:

* `LEN "STEVE"`
  * Returns the length of a string "STEVE" (5).
* `LEFT$ "STEVE", 2`
  * Returns the left-most 2 characters of "STEVE" ("ST").
* `RIGHT$ "STEVE", 2`
  * Returns the right-most 2 characters of "STEVE" ("VE").
* `CHR$ 42`
  * Converts the integer 42 to a character (`*`).  (i.e. ASCII value.)
* `CODE " "`
  * Converts the given character to the integer value (32).

<br />
<br />
<br />
<br />

## 20 PRINT "Limitations"

This project was started as [a weekend-project](https://blog.steve.fi/so_i_wrote_a_basic_basic.html), although it has subsequently been improved and extended.

The code has near-total test-coverage, and has been hammered with multiple days of fuzz-testing (i.e. Feeding random programs into the interpreter to see if it will die - see [FUZZING.md](FUZZING.md) for more details on that.)

That said there are some (obvious) limitations:

* Only a single statement is allowed upon each line.
* Only a subset of the language is implemented.
  * If there are specific primitives you miss, then please [report a bug](https://github.com/skx/gobasic/issues/).
    * The project is open to suggestions, but do bear in mind the [project goals]((#100-print-project-goals--links)) listed later on.
* When it comes to types only floating-point and string values are permitted.
  * There is support for arrays but only one or two dimensional ones.

### Arrays

Arrays are used just like normal variables, but they need to be declared using the `DIM` statement.   Individual elements are accessed using the offsets in brackets after the variable name:

    10 DIM a(10,10)
    20 LET a[1,1]=10
    30 PRINT a[1,1], "\n"

Arrays are indexed from 0-N, so with an array size of ten you can access eleven
elements:

     10 DIM a(10)
     20 a[0] = 0
     30 a[1] = 1
     40 ..
     90 a[9] = 9
    100 a[10] = 10

ZX Spectrum BASIC indexed arrays from 1, denying the ability to use the zeroth element, which I've long considered a mistake.


### Line Numbers

Line numbers are _mostly_ optional, for example the following program is valid and correct:

     10 READ a
     20 IF a = 999 THEN GOTO 100
     30 PRINT a, "\n"
     40 GOTO 10
    100 END
        DATA 1, 2, 3, 4, 999

The main reason you need line-numbers is for the `GOTO` and `GOSUB` functions,
if you prefer to avoid them then you're welcome to do so.

### `IF` Statement

The handling of the IF statement is perhaps a little unusual, since I'm
used to the BASIC provided by the ZX Spectrum which had no ELSE clause.
The general form of the IF statement I've implemented is:

    IF $CONDITIONAL THEN $STATEMENT1 [ELSE $STATEMENT2]

Only a single statement is permitted between "THEN" and "ELSE", and again between "ELSE" and NEWLINE.  These are valid IF statements:

    IF 1 > 0 THEN PRINT "OK"
    IF 1 > 3 THEN PRINT "SOMETHING IS BROKEN": ELSE PRINT "Weird!"

In that second example you'll see that "`:`" was used to terminate the `PRINT` statement, which otherwise would have tried to consume all input until it hit a newline.

The set of comparison functions _probably_ includes everything you need:

* `IF a < b THEN ..`
* `IF a > b THEN ..`
* `IF a <= b THEN ..`
* `IF a >= b THEN ..`
* `IF a = b THEN ..`
* `IF a <> b THEN ..`
* `IF a THEN ..`
  * This passes if `a` is a number which is not zero.
  * This passes if `a` is a string which is non-empty.

If you're missing something from that list please let me know by filing an issue.


### `DATA` / `READ` Statements

The `READ` statement allows you to read the next value from the data stored
in the program, via `DATA`.  There is no support for the `RESTORE` function,
so once your data is read it cannot be re-read.


### Builtin Functions

You'll also notice that the primitives which are present all suffer from the flaw that they don't allow brackets around their arguments.  So this is valid:

    10 PRINT RND 100

But this is not:

    10 PRINT RND(100)

This particular problem could be fixed, but I've not considered it significant.


### Types

There are no type restrictions on variable names vs. their contents, so these statements are each valid:

* `LET a = "steve"`
* `LET a = 3.2`
* `LET a% = ""`
* `LET a$ = "steve"`
* `LET a$ = 17 + 3`
* `LET a% = "string"`

The __sole__ exception relates to the `INPUT` statement.  The `INPUT` statement prompts a user for input, and returns it as a value - it doesn't know whether to return a "string" or a "number".  So it returns a string if it sees a `$` in the variable name.

This means this reads a string:

    10 INPUT "Enter a string", a$

But this prompts for a number:

    10 INPUT "Enter a number", a

This seemed better than trying to return a string, unless the input looked like a number (i.e. the input matched `/^([0-9\.]+)$/` we could store a number, otherwise a string).


<br />
<br />
<br />
<br />


## 30 PRINT "Installation"

### Build without Go Modules (Go before 1.11)

Providing you have a working [go-installation](https://golang.org/) you should be able to install this software by running:

    go get -u github.com/skx/gobasic

**NOTE** This will only install the command-line driver, rather than the HTTP-server, or the embedded example code.

### Build with Go Modules (Go 1.11 or higher)

    git clone https://github.com/skx/gobasic ;# make sure to clone outside of GOPATH
    cd gobasic
    go install

If you don't have a golang environment setup you should be able to download various binaries from the github release page:

* [Binary Releases](https://github.com/skx/gobasic/releases)


<br />
<br />
<br />
<br />


## 40 PRINT "Usage"

gobasic is very simple, and just requires the name of a BASIC-program to
execute.  Write your input in a file and invoke `gobasic` with the path.

For example the following program was useful to test my implementation of the `GOTO` primitive:

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

**NOTE**: I feel nostalgic seeing keywords in upper-case, but `PRINT` and `print` are treated identically.


<br />
<br />
<br />
<br />


## 50 PRINT "Implementation"

A traditional interpreter for a scripting/toy language would have a series of
well-defined steps:

* Split the input into a series of tokens ("lexing").
* Parse those tokens and build an abstract syntax tree (AST).
* Walk that tree, evaluating as you go.

As is common with early 8-bit home-computers this implementation is a little more BASIC:

* We parse the input into a series of tokens, defined in [token/token.go](token/token.go)
  * The parsing happens in [tokenizer/tokenizer.go](tokenizer/tokenizer.go)
* We then __directly__ execute those tokens.
  * The execution happens in [eval/eval.go](eval/eval.go) with a couple of small helpers:
    * [eval/for_loop.go](eval/for_loop.go) holds a simple data-structure for handling `FOR`/`NEXT` loops.
    * [eval/stack.go](eval/stack.go) holds a call-stack to handle `GOSUB`/`RETURN`
    * [eval/vars.go](eval/vars.go) holds all our variable references.
    * We have a facility to allow golang code to be made available to BASIC programs, and we use that facility to implement a bunch of our functions as "builtins".
      * Our builtin-functions are implemented beneath [builtin/](builtin/).
* Because we support both strings and ints/floats in our BASIC scripts we use a wrapper to hold them on the golang-side.  This can be found in [object/object.go](object/object.go).

As there is no AST step errors cannot be detected prior to the execution of programs - because we only hit them after we've started running.


<br />
<br />
<br />
<br />


## 60 PRINT "Sample Code"

There are a small number of sample-programs located beneath [examples/](examples/).   These were written in an adhoc fashion to test various parts of the implementation.

Perhaps the best demonstration of the code are the following two samples:

* [examples/90-stars.bas](examples/90-stars.bas)
  * Prompt the user for their name and the number of stars to print.
  * Then print them.  Riveting!  Fascinating!  A program for the whole family!
* [examples/99-game.bas](examples/99-game.bas)
  * A classic game where you guess the random number the computer has thought of.

<br />
<br />
<br />
<br />


## 70 PRINT "Embedding"

The interpreter is designed to be easy to embed into your application(s)
if you're crazy enough to want to do that!

You can see an example in the file [embed/main.go](embed/main.go).

The example defines several new functions which can be called by BASIC:

* `PEEK`
* `POKE`
* `PLOT`
* `SAVE`
* `CIRCLE`

When the script runs it does some BASIC variable manipulation and it also
creates a PNG file - the `PLOT` function allows your script to set a pixel
and the `CIRCLE` primitive draws an outline of a circle.  Finally the
`SAVE` function writes out the result.

Extending this example to draw filled circles, boxes, etc, is left as an
exercise ;)

Hopefully this example shows that making your own functions available to
BASIC scripts is pretty simple.  (This is how SIN, COS, etc are implemented
in the standalone interpreter.)


<br />
<br />
<br />
<br />


## 80 PRINT "Visual BASIC!"

Building upon the code in the embedded-example I've also implemented a simple
HTTP-server which will accept BASIC code, and render images!

To run this:

    cd goserver ; go build . ; ./goserver

Once running open your browser at the URL:

* [http://localhost:8080](http://localhost:8080)

The view will have an area of entering code, and once you run it the result will
be shown in the bottom of the screen.  Something like this:

![alt text](https://github.com/skx/gobasic/raw/master/goserver/screenshot.png "Sample view")

There are several included examples which you can load/launch by clicking upon them.


<br />
<br />
<br />
<br />


## 90 PRINT "Bugs?"

It is probable that bugs exist in this interpreter, but I've tried to do
as much testing as I can.  If you spot anything that
seems wrong please do [report an issue](https://github.com/skx/gobasic/issues).

* If the interpreter segfaults that is a bug.
  * Even if the program is invalid, bogus, or malformed the interpreter should cope with it.
* If a valid program produces the wrong output then that is also a bug.

The project contain a number of test-cases, which you can execute like so:

    $ go test ./...

Finally __if our test-coverage drops beneath 95%__ that is __a bug__.  The
test coverage of most of our packages is 100%, unfortunately the main `eval/`
package is not yet completely covered.

You can see the __global__ coverage via:

    $ ./test-coverage
    97.9%

In addition to the test-cases which have been manually written the interpreter
has also been fuzz-tested, which has resulted in some significant improvements.

See [FUZZING.md](FUZZING.md) for details of how to run the fuzz-tests.


<br />
<br />
<br />
<br />


## 100 PRINT "Project Goals / Links"

It is never the intention of this project to support _all_ things that are
possible in the various dialects of BASIC.

There are facilities which will make porting programs useful, such as
the ability to use `WHILE`/`END` loops, functions with named-parameters,
and primitives such as SLEEP, BEEP, & etc.

Above all else this project is supposed to be fun, for me.  Which means
if there are two ways of implementing something I'll pick the way I remember
back when I was 12 and computers were .. fun!

If there are feature-requests which seem less fun, and less immediately
useful to me - with my biased memories of coding on a ZX Spectrum - I will
tag them "wontfix".  If you contribute a pull-request to support them I will
accept them, but I'm probably not likely to work upon them directly.

That said there are cases where I can be persuaded, and there are a lot
of other BASIC intepreters out there, so I won't feel bad if this particular
project doesn't suit your needs.

One project, slightly related to this, which might be worth checking up
on is this one:

* https://github.com/udhos/basgo


Steve
--
