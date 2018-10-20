// eval_test.go - Simple test-cases for our evaluator

package eval

import (
	"strings"
	"testing"

	"github.com/skx/gobasic/tokenizer"
)

// Create a new interpreter from the given string of text.
func Compile(input string) *Interpreter {
	tokener := tokenizer.New(input)
	e := New(tokener)
	return e
}

// TestLet ensures a value is set.  It is naive.
func TestLet(t *testing.T) {
	input := "10 LET a = 33\n"

	obj := Compile(input)
	obj.Run()

	out := obj.GetVariable("a")
	if out != 33 {
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

	out := obj.GetVariable("A")
	if out != 1002 {
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

	out := obj.GetVariable("a")
	if out != 333333 {
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
70 LET G = ABS(-3)
80 LET H = 10 - 20
90 LET H = ABS(H)
95 LET H = 33
99 LET H = ABS(H)
100 LET I = ABS(0 - 2)
110 LET R = RND()
`

	obj := Compile(input)
	obj.Run()

	out := obj.GetVariable("A")
	if out != 3 {
		t.Errorf("Value not expected!")
	}
	out = obj.GetVariable("B")
	if out != 5 {
		t.Errorf("Value not expected!")
	}
	out = obj.GetVariable("C")
	if out != 20 {
		t.Errorf("Value not expected!")
	}
	out = obj.GetVariable("D")
	if out != 20 {
		t.Errorf("Value not expected!")
	}
	out = obj.GetVariable("E")
	if out != 1 {
		t.Errorf("Value not expected!")
	}
	out = obj.GetVariable("F")
	if out != 100 {
		t.Errorf("Value not expected!")
	}
	out = obj.GetVariable("H")
	if out != 33 {
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

	out := obj.GetVariable("SUM")
	if out != 55 {
		t.Errorf("Value not expected - got %d", out)
	}
}

// TestBogusFor ensures that bogus-FOR-loops are found
func TestBogusFor(t *testing.T) {

	txt := []string{"10 FOR \n",
		"10 FOR I\n",
		"10 FOR I=\n",
		"10 FOR I=1\n",
		"10 FOR I=1 TO\n",
		"10 FOR I=1 TO N\n",
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
