// eval_test.go - Simple test-cases for our evaluator

package eval

import (
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

// TestMaths does some simple sums
func TestMaths(t *testing.T) {
	input := `
10 LET A = 1 + 2
20 LET B = 6 - 1
30 LET C = 4 * 5
40 LET D = 100 / 5
50 LET E = 5 % 2
60 LET F = ( ( 3 + 1 ) / 2 ) + ( 3 * 33 ) - 1
70 LET G = ABS(3)
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

// TestPrint runs some prints.
func TestPrint(t *testing.T) {
	input := `
10 LET a=3
20 PRINT a
30 PRINT "Test\n"
40 PRINT ( 3 * ( 3 + 4  ) ) "\n"
`

	obj := Compile(input)
	obj.Run()

	// Nothing useful to test here
	// unless we use an I/O writer..?
	//
	// TODO: Reconsider
}
