// eval_test.go - Simple test-cases for our evaluator.

package eval

import (
	"math"
	"strings"
	"testing"

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

// Create a new interpreter from the given string of text.
func Compile(input string) *Interpreter {
	tokener := tokenizer.New(input)
	e := New(tokener)
	return e
}

// getFloat is a helper for retrieving the value of a number-object
func getFloat(t *testing.T, e *Interpreter, name string) float64 {
	out := e.GetVariable(name)
	if out == nil {
		t.Errorf("Loading variable '%s' failed\n", name)
	}

	if out.Type() != object.NUMBER {
		t.Errorf("Object %s was not a number", name)
	}
	return (out.(*object.NumberObject).Value)
}

// getError is a helper for getting an error.
func getError(t *testing.T, e *Interpreter, name string) string {
	out := e.GetVariable(name)
	if out == nil {
		t.Errorf("Loading variable '%s' failed\n", name)
	}
	if out.Type() != object.ERROR {
		t.Errorf("Object %s was not an error", name)
	}
	return (out.(*object.ErrorObject).Value)
}

// getString is a helper for retrieving the value of a string-object
func getString(t *testing.T, e *Interpreter, name string) string {
	out := e.GetVariable(name)
	if out == nil {
		t.Errorf("Loading variable '%s' failed\n", name)
	}
	if out.Type() != object.STRING {
		t.Errorf("Object %s was not a string", name)
	}
	return (out.(*object.StringObject).Value)
}

// TestGetSet ensures a value is set.  It is naive.
func TestGetSet(t *testing.T) {
	input := "10 LET a = ( a + 1 ) \n"

	obj := Compile(input)
	obj.SetVariable("a", &object.NumberObject{Value: 17})
	obj.Run()

	out := getFloat(t, obj, "a")
	if out != 18 {
		t.Errorf("Value not expected!")
	}
}

// TestLen tests our LEN implementation.
func TestLen(t *testing.T) {
	input := `
10 LET I="Hello World"
20 LET a=LEN I
30 LET b=LEN "Steve Kemp"
40 LET t = "Steve"
50 LET t= t + " "
60 LET t= t + "Kemp"
70 LET c= LEN t
80 LET d = LEN 0
`
	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != 11 {
		t.Errorf("LEN 1 Failed!")
	}
	if getFloat(t, obj, "b") != 10 {
		t.Errorf("LEN 2 Failed!")
	}
	if getFloat(t, obj, "c") != 10 {
		t.Errorf("LEN 3 Failed!")
	}
	if !strings.Contains(getError(t, obj, "d"), "doesn't exist") {
		t.Errorf("LEN 4 Failed!")
	}
}

// TestLeftRight tests our LEFT$/RIGHT$ implementation.
func TestLeftRight(t *testing.T) {
	input := `
10 LET I="Hello World"
20 LET a=LEFT$ I, 4
30 LET b=LEFT$ "Steve", 2
40 LET c=RIGHT$ I, 5
50 LET d=RIGHT$ I, 2

60 LET e=RIGHT$ I,100
70 LET f=LEFT$ I,100
`
	obj := Compile(input)
	obj.Run()

	if getString(t, obj, "a") != "Hell" {
		t.Errorf("LEFT$ 1 Failed! Got %v", getString(t, obj, "a"))
	}
	if getString(t, obj, "b") != "St" {
		t.Errorf("LEFT$ 2 Failed!")
	}
	if getString(t, obj, "c") != "World" {
		t.Errorf("RIGHT$ 1 Failed!")
	}
	if getString(t, obj, "d") != "ld" {
		t.Errorf("RIGHT$ 2 Failed!")
	}
	if getString(t, obj, "e") != "Hello World" {
		t.Errorf("RIGHT$ 3 Failed!")
	}
	if getString(t, obj, "f") != "Hello World" {
		t.Errorf("LEFT$ 3 Failed!")
	}

}

// TestLet ensures a value is set.  It is naive.
func TestLet(t *testing.T) {
	input := "10 LET a = 33\n"

	obj := Compile(input)
	obj.Run()

	out := getFloat(t, obj, "a")
	if out != 33 {
		t.Errorf("Value not expected!")
	}
}

// TestPI ensures that PI and INT works
func TestPI(t *testing.T) {
	input := `10 LET a = PI
20 LET b = INT a
30 LET c = INT PI
`

	obj := Compile(input)
	obj.Run()

	out := getFloat(t, obj, "a")
	if out != math.Pi {
		t.Errorf("Value not expected!")
	}

	// PI -> int == 3
	if getFloat(t, obj, "b") != 3 {
		t.Errorf("INT didn't work as expected!")
	}
	if getFloat(t, obj, "c") != 3 {
		t.Errorf("INT didn't work as expected!")
	}
}

// TestSGN ensures that the sign-extend function works
func TestSGN(t *testing.T) {
	input := `10 REM
20 LET a = SGN -10
30 LET b = SGN 13
40 LET c = SGN 0

50 LET X = -10
60 LET Y = 0
64 LET Z = 300

70 LET d = SGN X
80 LET e = SGN Y
90 LET f = SGN Z
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != -1 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "b") != 1 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "c") != 0 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "d") != -1 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "e") != 0 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "f") != 1 {
		t.Errorf("Value not expected!")
	}
}

// TestSQR ensures that the square-root code is sane.
func TestSQR(t *testing.T) {
	input := `10 REM
20 LET a = SQR 9
30 LET X = 100
40 LET b = SQR X
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != 3 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "b") != 10 {
		t.Errorf("Value not expected!")
	}
}

// TestBinOp tests that binary AND + binary OR work
func TestBinOp(t *testing.T) {
	input := `
10 LET a = ( BIN 00001111 ) OR ( BIN 01110000 )
20 LET b = 129 AND 128
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != 255-128 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "b") != 128 {
		t.Errorf("Value not expected!")
	}

}

// TestBogusLet tests error-handling in LET
func TestBogusLet(t *testing.T) {

	txt := []string{"10 LET 3\n",
		"10 LET a _ 3\n"}

	for _, prg := range txt {

		obj := Compile(prg)
		err := obj.Run()

		if err == nil {
			t.Errorf("Expected to receive an error in the program - but didn't")
		}
		if !strings.Contains(err.Error(), "LET") {
			t.Errorf("Received error, but the wrong thing?")
		}
	}
}

// TestGoSub ensures a value is set.  It is naive.
func TestGoSub(t *testing.T) {
	input := `
10 LET A=33
20 GOSUB 300
30 END
300 LET A = 1002
310 RETURN
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "A") != 1002 {
		t.Errorf("Value not expected!")
	}
}

// TestBogusGoSub ensures that bogus-gosubs are found
func TestBogusGoSub(t *testing.T) {

	txt := []string{"10 GOSUB A\n",
		"10 GOSUB 1000\n",
		"10 RETURN\n"}

	for _, prg := range txt {

		obj := Compile(prg)
		err := obj.Run()

		if err == nil {
			t.Errorf("Expected to receive an error in the program - but didn't")
		}
		if !strings.Contains(err.Error(), "GOSUB") {
			t.Errorf("Received error, but the wrong thing?")
		}
	}
}

// TestGoTo ensures a value is set.  It is naive.
func TestGoTo(t *testing.T) {
	input := `
 10 GOTO 100
 20 GOTO 90
 30 GOTO 80
 40 GOTO 70
 50 LET a=333333
 60 END
 70 GOTO 50
 80 GOTO 40
 90 GOTO 30
100 GOTO 20
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != 333333 {
		t.Errorf("Value not expected!")
	}
}

// TestBogusGoTO ensures that bogus-gotos are found
func TestBogusGoTO(t *testing.T) {

	txt := []string{"10 GOTO A\n",
		"10 GOTO 1000\n"}

	for _, prg := range txt {

		obj := Compile(prg)
		err := obj.Run()

		if err == nil {
			t.Errorf("Expected to receive an error in the program - but didn't")
		}
		if !strings.Contains(err.Error(), "GOTO") {
			t.Errorf("Received error, but the wrong thing?")
		}
	}
}

// TestMaths does some simple sums
func TestMaths(t *testing.T) {
	input := `
10 LET A = 1 + 2
20 LET B = 6 - 1
30 LET C = 4 * 5
40 LET D = 100 / 5
50 LET E = 5 % 2
60 LET F = ( ( 3 + 1 ) / 2 ) + ( 3 * 33 ) - 1
70 LET G = ABS -3
80 LET H = 10 - 20
90 LET H = ABS H
95 LET H = 33
99 LET H = ABS H
110 LET R = RND 100
120 LET KEY = "STEVE"
130 LET RT = RND KEY
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "A") != 3 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "B") != 5 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "C") != 20 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "D") != 20 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "E") != 1 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "F") != 100 {
		t.Errorf("Value not expected!")
	}
	if getFloat(t, obj, "H") != 33 {
		t.Errorf("Value not expected!")
	}
	if !strings.Contains(getError(t, obj, "RT"), "doesn't exist") {
		t.Errorf("Value not expected!")
	}
	if getString(t, obj, "KEY") != "STEVE" {
		t.Errorf("Value not expected!")
	}
}

// TestFor runs a single simple FOR loop
func TestFor(t *testing.T) {
	input := `
10 LET SUM = 0
20 LET N=10
30 FOR I = 1 TO N STEP 1
40 LET SUM = SUM + I
50 NEXT I
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "SUM") != 55 {
		t.Errorf("Value not expected!")
	}
}

// TestForTerm ensures that a for-loop with start/end the same runs only once.
func TestForTerm(t *testing.T) {
	input := `
10 LET COUNT=0
20 FOR I = 1 TO 1
30  LET COUNT = COUNT + 1
40 NEXT I
50 PRINT COUNT
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "COUNT") != 1 {
		t.Errorf("Value not expected!")
	}
}

// TestBogusFor ensures that bogus-FOR-loops are found
func TestBogusFor(t *testing.T) {

	txt := []string{"10 FOR \n",
		"10 FOR I\n",
		"10 FOR I=\n",
		"10 FOR I=1\n",
		"10 FOR I=1 TO\n",
		"10 FOR I=1 TO 10 STEP STEP\n",
		"10 FOR I=1 TO 20\n20NEXT 3\n",
		`10 LET TERM="steve"
20 FOR I = 1 TO TERM`,
	}

	for _, prg := range txt {

		obj := Compile(prg)
		err := obj.Run()

		if err == nil {
			t.Errorf("Expected to receive an error in the program '%s' - but didn't", prg)
		}
		if !strings.Contains(err.Error(), "FOR") {
			t.Errorf("Received error, but the wrong thing?")
		}
	}
}

// TestPrint runs some prints.
func TestPrint(t *testing.T) {
	input := `
10 LET a=3
20 PRINT a
30 PRINT "Test\n"
40 PRINT ( 3 * ( 3 + 4  ) ) "\n"
50 PRINT "OK","OK"
60 LET a = PI
70 PRINT a
80 LET a = "STEVE"
90 PRINT a
95 PRINT PI + 2
97 PRINT "steve" + " kemp"
99 PRINT 3 + 5 "\n"
100 PRINT LEN "STEVE"
110 PRINT LEFT$ "Steve" 2
120 PRINT DUMP 3, "Steve", 22, 32-1, "\n"
`

	obj := Compile(input)
	err := obj.Run()

	if err != nil {
		t.Errorf("Unexpected error in runPRINT")
	}
}

// TestREM ensures that REM is handled.
func TestREM(t *testing.T) {

	txt := []string{"10 REM\n",
		"10 REM"}

	for _, prg := range txt {

		obj := Compile(prg)
		err := obj.Run()

		if err != nil {
			t.Errorf("Error parsing program '%s'", prg)
		}
	}
}

// TestIF tests an IF statement
func TestIf(t *testing.T) {
	input := `
10 IF 1 < 10 THEN let a=1 ELSE PRINT "FAIL1\n"
20 IF 1 <= 10 THEN let b=1 ELSE PRINT "FAIL2\n"
30 IF 11 > 7 THEN let c=1 ELSE PRINT "FAIL3\n"
40 IF 11 >= 7 THEN let d=1 ELSE PRINT "FAIL4\n"
50 IF a = b THEN let e=1 ELSE PRINT "FAIL5\n"
60 IF a = 8 THEN let e=0 ELSE PRINT "FAIL6\n"
70 IF a <> b THEN PRINT "FAIL7\n": ELSE let f=1
80 IF a <> 100 THEN let g=1 ELSE LET g=1
90 IF a = 1 THEN let h=1
100 IF "steve" = "steve" THEN LET i=1
110 IF "steve" <> "kemp" THEN LET j=1
120 IF "steve" > "STEVE" THEN LET k=1
130 IF "STEVE" >= "STEVE" THEN LET l=1
140 IF "STEVE" < "steve" then let m=1
150 IF "steve" <= "steve" then let n=1
160 IF "steve" = "fsteve" then PRINT "NOP"
170 LET x=1
`

	obj := Compile(input)
	obj.Run()

	//
	// Get our variables - they should all be equal to one
	//
	vars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "x", "k", "l", "m", "n"}

	for _, nm := range vars {
		out := getFloat(t, obj, nm)
		if out != 1 {
			t.Errorf("Value not expected - got %f for %s", out, nm)
		}

	}
}

// TestSubstr tests LEFT$ and RIGHT$
func TestSubstr(t *testing.T) {
	input := `
10 LET A$="SATURDAY MORNING"
20 LET B$=LEFT$ A$ , 8
30 LET C$=RIGHT$ A$ , 7
40 LET blah=0
50 LET D$=LEFT$ blah, 3
`

	obj := Compile(input)
	obj.Run()

	if getString(t, obj, "B$") != "SATURDAY" {
		t.Errorf("LEFT$ failed! - got '%s'", getString(t, obj, "B$"))
	}
	if getString(t, obj, "C$") != "MORNING" {
		t.Errorf("RIGHT$ failed! - got '%s'", getString(t, obj, "C$"))
	}

	if !strings.Contains(getError(t, obj, "D$"), "doesn't exist") {
		t.Errorf("RIGHT$ failed!")
	}
}

// TestMath is a lie - it just invokes sin/cos/tan/etc.  It does
// no results-testing.
func TestMath(t *testing.T) {
	input := `
10 PRINT SIN 3, "\n"
20 PRINT COS 3, "\n"
30 PRINT TAN 3, "\n"
40 PRINT ATN 3, "\n"
50 PRINT ASN 3, "\n"
60 PRINT ACS 3, "\n"
70 PRINT EXP 3, "\n"
80 PRINT LN 3, "\n"
`

	obj := Compile(input)
	obj.Run()
}

func TestTL(t *testing.T) {
	input := `
10 LET a = TL$ "Hello World"
20 LET b = TL$ "S"
30 LET c = TL$ ""
40 LET s = "Steve"
50 LET d = TL$ s
`
	obj := Compile(input)
	obj.Run()

	if getString(t, obj, "a") != "ello World" {
		t.Errorf("TL 1 Failed!")
	}
	if getString(t, obj, "b") != "" {
		t.Errorf("TL 2 Failed!")
	}
	if getString(t, obj, "c") != "" {
		t.Errorf("TL 3 Failed!")
	}
	if getString(t, obj, "d") != "teve" {
		t.Errorf("TL 4 Failed!")
	}
}

// TestMID tests our MID$ function.
func TestMID(t *testing.T) {
	input := `
10 LET IN = "Hello World"
20 LET a = MID$ IN, 1000, 1
30 LET b = MID$ IN, 6, 4
40 LET c = MID$ IN, 6, 400
`
	obj := Compile(input)
	obj.Run()

	if getString(t, obj, "a") != "" {
		t.Errorf("MID$ 1 Failed!")
	}
	if getString(t, obj, "b") != "Worl" {
		t.Errorf("CODE 2 Failed!")
	}
	if getString(t, obj, "c") != "World" {
		t.Errorf("CODE 3 Failed!")
	}
}

// TestCode tests our CODE function
func TestCode(t *testing.T) {
	input := `
10 LET a = CODE "Hello World"
20 LET b = CODE "S"
30 LET c = CODE ""
`
	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != 72 {
		t.Errorf("CODE 1 Failed!")
	}
	if getFloat(t, obj, "b") != 83 {
		t.Errorf("CODE 2 Failed!")
	}
	if getFloat(t, obj, "c") != 0 {
		t.Errorf("CODE 3 Failed!")
	}
}

// TestCHR tests our CHR$ function
func TestCHR(t *testing.T) {
	input := `
10 LET a = CHR$ 42
20 LET b = CHR$ 32
`
	obj := Compile(input)
	obj.Run()

	if getString(t, obj, "a") != "*" {
		t.Errorf("CHR$ 1 Failed!")
	}
	if getString(t, obj, "b") != " " {
		t.Errorf("CHR$ 2 Failed!")
	}
}

// TestBIN tests our BIN function
func TestBIN(t *testing.T) {
	input := `
10 LET a = BIN 11111111
20 LET b = BIN 00000010
20 LET c = BIN 00002010
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Found no error, but expected one")
	}
	if getFloat(t, obj, "a") != 255 {
		t.Errorf("BIN 1!")
	}
	if getFloat(t, obj, "b") != 2 {
		t.Errorf("BIN 2!")
	}
	if !strings.Contains(getError(t, obj, "d"), "doesn't exist") {
		t.Errorf("BIN 3!")
	}
}

// TestIfBIN tests our IF function
func TestIFBIN(t *testing.T) {
	input := `
10 LET A = 1
20 LET B = 0
30 LET C = 0
40 LET D = 0
50 IF A = 1 OR B = 1 THEN LET C = 1
60 IF A = 1 AND B = 0 THEN LET D = 1
`
	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "C") != 1 {
		t.Errorf("1 Failed!")
	}
	if getFloat(t, obj, "D") != 1 {
		t.Errorf("2 Failed!")
	}
}

// TestBogusVariable tests that retrieving a bogus variable fails
func TestBogusVariable(t *testing.T) {
	input := `10 PRINT A`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "The variable") {
		t.Errorf("Our error-message wasn't what we expected")
	}

}

// TestBogusBuiltIn tests that retrieving a bogus builtin fails
func TestBogusBuiltIn(t *testing.T) {
	input := `10 AARDVARK`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Token not handled") {
		t.Errorf("Our error-message wasn't what we expected: %s", err.Error())
	}
}

// TestMismatchedTypes tests that expr() errors on mismatched types.
func TestMismatchedTypes(t *testing.T) {
	input := `10 LET a=3
20 LET b="steve"
30 LET c = a + b
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "type mismatch") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestMismatchedTypesTerm tests that term() errors on mismatched types.
func TestMismatchedTypesTerm(t *testing.T) {
	input := `10 LET a="steve"
20 LET b = ( a * 2 ) + ( a * 33 )
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "handles integers") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestStringFail tests that expr() errors on bogus string operations.
func TestStringFail(t *testing.T) {
	input := `10 LET a="steve"
20 LET b="steve"
30 LET c = a - b
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "not supported for strings") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestExprTerm tests that expr() errors on unclosed brackets.
func TestExprTerm(t *testing.T) {
	input := `10 LET a = ( 3 + 3 * 33
20 PRINT a "\n"
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Unclosed bracket") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestIfFail tests that IF returns an error.
func TestIfFail(t *testing.T) {
	tst := []string{"10 IF \"*\" * 10 = 42 THEN PRINT \"OK\"",
		"10 IF 42 = \"*\" / 44 THEN PRINT \"OK\"",
	}

	for _, input := range tst {
		obj := Compile(input)
		err := obj.Run()

		if err == nil {
			t.Errorf("Expected to see an error, but didn't.")
		}
		if !strings.Contains(err.Error(), "only handles integers") {
			t.Errorf("Our error-message wasn't what we expected:%s", err.Error())
		}
	}
}

// TestBogusInput ensures that bogus INPUTs are handled.
func TestBogusInput(t *testing.T) {

	txt := []string{"10 INPUT \n20 REM\n",
		"10 INPUT ID ID\n20 REM\n",
		"10 INPUT \"steve\", 33\n20 REM\n",
	}

	for _, prg := range txt {

		obj := Compile(prg)
		err := obj.Run()

		if err == nil {
			t.Errorf("Expected to receive an error in the program '%s' - but didn't", prg)
		}
		if !strings.Contains(err.Error(), "INPUT") {
			t.Errorf("Received error for '%s' but the wrong thing? %s", prg, err.Error())
		}
	}
}

// TestDump calls DUMP
func TestDump(t *testing.T) {

	obj := Compile("10 DUMP 3\n")
	err := obj.Run()

	if err != nil {
		t.Errorf("Found error calling DUMP\n")
	}
}

// TestBuiltinError tests that a builtin-error is handled.
func TestBuiltinError(t *testing.T) {

	obj := Compile("10 ABS \"steve\"\n")
	err := obj.Run()

	if err == nil {
		t.Errorf("Didn't find a type error, and we should have done.")
	}
	if !strings.Contains(err.Error(), "Wrong type") {
		t.Errorf("Received error but the wrong thing? %s", err.Error())
	}
}

// Test the start/end condition of a loop can be variables
func TestIfStartEnd(t *testing.T) {
	type IfTest struct {
		Input  string
		Output float64
	}

	tsts := []IfTest{{Input: `10 LET OUT = 33
20 LET STOP=10
30 FOR I = 5 TO STOP
40  LET OUT = OUT * I
50 NEXT I
`, Output: 4989600},
		{
			Input: `10 LET START=10
20 LET OUT = 22
30 FOR I = START TO 100
40  LET OUT = OUT + I
50 NEXT I
`,
			Output: 5027},
		{
			Input: `10 LET START=1
15 LET OUT = 1
20 LET STOP=10
30 FOR I = START TO STOP
40  LET OUT = OUT * I
50 NEXT I
`,
			Output: 3628800},
	}
	for _, prg := range tsts {

		obj := Compile(prg.Input)
		err := obj.Run()

		if err != nil {
			t.Errorf("Found error running '%s' - %s", prg.Input, err.Error())
		}

		if getFloat(t, obj, "OUT") != prg.Output {
			t.Errorf("Output of program was %f not %f\n",
				getFloat(t, obj, "OUT"), prg.Output)
		}

	}
}

// Test VAL
func TestVAL(t *testing.T) {

	input := `
10 LET A = VAL 33
20 LET B = VAL "3.44"
`
	obj := Compile(input)
	err := obj.Run()

	if err != nil {
		t.Errorf("Found error running '%s' - %s", input, err.Error())
	}

	if getFloat(t, obj, "A") != 33 {
		t.Errorf("Wrong value for VAL output")
	}
	if getFloat(t, obj, "B") != 3.44 {
		t.Errorf("Wrong value for VAL output")
	}
}

// Test STR$
func TestSTR(t *testing.T) {

	input := `
10 LET A = STR$ 33
20 LET B = STR$ 19.22
30 LET B = LEFT$ B 5
40 LET C = STR$ "steve"
`
	obj := Compile(input)
	err := obj.Run()

	if err != nil {
		t.Errorf("Found error running '%s' - %s", input, err.Error())
	}

	if getString(t, obj, "A") != "33" {
		t.Errorf("Wrong value for STR")
	}
	if getString(t, obj, "B") != "19.22" {
		t.Errorf("Wrong value for STR: %v", getString(t, obj, "B"))
	}
	if getString(t, obj, "C") != "steve" {
		t.Errorf("Wrong value for STR")
	}
}

// TestMismatchedNext ensures NEXT is paired with a FOR.
func TestMisMatchedNext(t *testing.T) {
	input := `
10 NEXT I
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("We didn't find an error, and should have done")
	}

	if !strings.Contains(err.Error(), "without opening FOR") {
		t.Errorf("Wrong error-message was found")
	}
}

// TestNextType ensures NEXT works on ints.
func TestNextType(t *testing.T) {
	input := `
10 FOR I = 1 TO 20
20   LET I = "steve"
30 NEXT I
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("We didn't find an error, and should have done")
	}

	if !strings.Contains(err.Error(), "is not a number!") {
		t.Errorf("Wrong error-message was found")
	}
}

// TestIssue32 is the test case for https://github.com/skx/gobasic/issues/32
func TestIssue32(t *testing.T) {
	input := `
10 LET A = LEFT$ STR$ 49.31321, 5
`
	obj := Compile(input)
	err := obj.Run()

	if err != nil {
		t.Errorf("Found error running '%s' - %s", input, err.Error())
	}

	if getString(t, obj, "A") != "49.31" {
		t.Errorf("Wrong value for LEFT$ STR$, got '%s'",
			getString(t, obj, "A"))
	}
}

// TestIssue34 is the test case for https://github.com/skx/gobasic/issues/34
func TestIssue34(t *testing.T) {

	// Broken examples
	inputs := []string{
		`10 PRINT RND 0`,
		`10 LET A = 3 - 5
10 PRINT RND A`,
	}

	for _, txt := range inputs {

		obj := Compile(txt)
		err := obj.Run()

		if err == nil {
			t.Errorf("We expected to find an error, but didn't")
		}
		if !strings.Contains(err.Error(), "Argument to RND must be >") {
			t.Errorf("The error we found was not what we expected: %s", err.Error())
		}

	}

	// Valid example
	obj := Compile("10 LET A = RND 55")
	err := obj.Run()

	if err != nil {
		t.Errorf("We didn't expect an error here")
	}

}

// Test that too-few arguments are caught to builtins.
func TestBuiltinArguments(t *testing.T) {

	inputs := []string{
		"10 PRINT STR$\n",
		"10 PRINT STR$"}

	for _, txt := range inputs {

		obj := Compile(txt)
		err := obj.Run()

		if err == nil {
			t.Errorf("We expected to find an error, but didn't")
		}
		if !strings.Contains(err.Error(), "while searching for argument") {
			t.Errorf("The error we found was not what we expected: %s", err.Error())
		}

	}

}

// TestIssue37 is the test case for https://github.com/skx/gobasic/issues/37
func TestIssue37(t *testing.T) {

	input := `10 FOR I = 1 TO 10
20 PRINT I
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("We expected to find an error, but didn't")
	}
	if !strings.Contains(err.Error(), "Unclosed FOR loop") {
		t.Errorf("The error we found was not what we expected: %s", err.Error())
	}

}

// TestIssue43 is the test case for https://github.com/skx/gobasic/issues/43
func TestIssue43(t *testing.T) {
	input := `
10 LET A = 3 % 0
`
	obj := Compile(input)
	err := obj.Run()

	if err == nil {
		t.Errorf("We expected an error and found none")
	}

	if !strings.Contains(err.Error(), "MOD 0") {
		t.Errorf("Wrong error found for MOD 0 : %s\n", err.Error())
	}
}

// TestIssue42 is the test case for https://github.com/skx/gobasic/issues/42
func TestIssue42(t *testing.T) {
	input := `
10 IF 1 THEN LET a=1 ELSE LET a = 3
10 IF 3 + 3 THEN LET b=1 ELSE LET b = 3
30 IF -3 + 3 THEN LET c = 3 ELSE LET c = 1
40 LET t=1
50 LET f=0
60 IF t THEN LET d=1 ELSE let d=3
60 IF f THEN LET e=3 ELSE let e=1
`
	obj := Compile(input)
	err := obj.Run()

	if err != nil {
		t.Errorf("We expected no error, but we got one!")
	}

	vars := []string{"a", "b", "c", "d", "e"}

	for _, nm := range vars {
		out := getFloat(t, obj, nm)
		if out != 1 {
			t.Errorf("Value not expected - got %f for %s", out, nm)
		}
	}
}
