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
	if out.Type() != object.NUMBER {
		t.Errorf("Object %s was not a number", name)
	}
	return (out.(*object.NumberObject).Value)
}

// getString is a helper for retrieving the value of a string-object
func getString(t *testing.T, e *Interpreter, name string) string {
	out := e.GetVariable(name)
	if out.Type() != object.STRING {
		t.Errorf("Object %s was not a string", name)
	}
	return (out.(*object.StringObject).Value)
}

// TestGetSet ensures a value is set.  It is naive.
func TestGetSet(t *testing.T) {
	input := "10 LET a = a + 1\n"

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
`
	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "a") != 11 {
		t.Errorf("LEN 1 Failed!")
	}
	if getFloat(t, obj, "b") != 10 {
		t.Errorf("LEN 2 Failed!")
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
}

// TestFor runs a single simple FOR loop
func TestFor(t *testing.T) {
	input := `
05 LET SUM = 0
10 FOR I = 1 TO 10 STEP 1
20 LET SUM = SUM + I
30 NEXT I
`

	obj := Compile(input)
	obj.Run()

	if getFloat(t, obj, "SUM") != 55 {
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
		//		"10 FOR I=1 TO N\n",
		"10 FOR I=1 TO 10 STEP STEP\n",
		"10 FOR I=1 TO 20\n20NEXT 3\n",
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
`

	obj := Compile(input)
	obj.Run()

	// Nothing useful to test here
	// unless we use an I/O writer..?
	//
	// TODO: Reconsider
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
100 LET x=1
`

	obj := Compile(input)
	obj.Run()

	//
	// Get our variables - they should all be equal to one
	//
	vars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "x"}

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
`

	obj := Compile(input)
	obj.Run()

	if getString(t, obj, "B$") != "SATURDAY" {
		t.Errorf("LEFT$ failed! - got '%s'", getString(t, obj, "B$"))
	}
	if getString(t, obj, "C$") != "MORNING" {
		t.Errorf("RIGHT$ failed! - got '%s'", getString(t, obj, "C$"))
	}
}
