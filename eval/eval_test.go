// eval_test.go - Simple test-cases for our evaluator.

package eval

import (
	"bufio"
	"strings"
	"testing"

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

// TestBuiltin covers some of our builtins, however it doesn't test
// the implementation of them - they are covered in their own package - just
// that we can call them.
func TestBuiltin(t *testing.T) {

	//
	// Three different ways argument-counting can fail:
	//
	//  EOF :  TODO - this get caught by expr()
	//  :
	//  Newline
	//
	tests := []string{
		"20 PRINT RND :",
		"20 PRINT RND:",
		"30 PRINT RND\n"}

	for _, test := range tests {

		tokener := tokenizer.New(test)
		e, err := New(tokener)
		if err != nil {
			t.Errorf("Error parsing %s - %s", test, err.Error())
		}

		err = e.Run()
		if err == nil {
			t.Errorf("Expected an error - found none")
		}
		if !strings.Contains(err.Error(), "while searching for argument") {
			t.Errorf("Got an error, but it was the wrong one: %s", err.Error())
		}

	}

}

// TestCompare tests our comparison operation, via IF
func TestCompare(t *testing.T) {
	type Test struct {
		Input string
		Var   string
		Val   float64
	}

	tests := []Test{
		{Input: `10 IF 1 < 10 THEN LET a=1 ELSE LET a=0`, Var: "a", Val: 1},
		{Input: `20 IF 1 <= 10 THEN LET b=1 ELSE LET b=2`, Var: "b", Val: 1},
		{Input: `10 IF 11 > 7 THEN let c=1 ELSE LET c=0`, Var: "c", Val: 1},
		{Input: `40 IF 11 >= 7 THEN let d=1 ELSE LET d=3`, Var: "d", Val: 1},
		{Input: `50 IF 1 = 1 THEN let e=1 ELSE LET e=3`, Var: "e", Val: 1},
		{Input: `60 IF 1 <> 3 THEN let f=13 ELSE LET f=3`, Var: "f", Val: 13},
		{Input: `70 IF 1 <> 1 THEN let g=3 ELSE LET g=33`, Var: "g", Val: 33},
		{Input: `80 IF "a" < "b" THEN LET A=1 ELSE LET A=0`, Var: "A", Val: 1},
		{Input: `90 IF "a" <= "a" THEN LET B=1 ELSE LET B=2`, Var: "B", Val: 1},
		{Input: `100 IF "b" > "a" THEN let C=1 ELSE LET C=0`, Var: "C", Val: 1},
		{Input: `110 IF "c" >= "a" THEN let D=1 ELSE LET D=3`, Var: "D", Val: 1},
		{Input: `120 IF "moi" = "moi" THEN let E=1 ELSE LET E=3`, Var: "E", Val: 1},
		{Input: `130 IF "steve" <> "kemp" THEN let F=13 ELSE LET F=3`, Var: "F", Val: 13},
		{Input: `140 IF "a" <> "a" THEN let G=3 ELSE LET G=33`, Var: "G", Val: 33},
		{Input: `150 IF 1=1 AND 2=2 THEN let H=11 ELSE LET H=12`, Var: "H", Val: 11},
		{Input: `160 IF 1=1 OR 33=2 THEN let I=211 ELSE LET I=20`, Var: "I", Val: 211},
		{Input: `170 IF 1=1 XOR 33=2 THEN let J=358 ELSE LET J=131`, Var: "J", Val: 358},
		{Input: `10 LET a=1
20 IF a THEN LET t=11 ELSE let t=10
`, Var: "t", Val: 11},
		{Input: `10 LET a=0
20 IF a THEN LET t=11 ELSE let t=10
`, Var: "t", Val: 10},
		{Input: `10 LET a="steve"
20 IF a THEN LET tt=11 ELSE let tt=10
`, Var: "tt", Val: 11},
		{Input: `10 LET a=""
20 IF a THEN LET tt=11 ELSE let tt=10
`, Var: "tt", Val: 10},
	}

	//
	// Test each comparison
	//
	for _, v := range tests {

		tokener := tokenizer.New(v.Input)
		e, err := New(tokener)
		if err != nil {
			t.Errorf("Error parsing %s - %s", v.Input, err.Error())
		}

		e.Run()

		//
		// By default the variable won't exist.
		//
		cur := e.GetVariable(v.Var)
		if cur.Type() == object.ERROR {
			t.Errorf("Variable %s does not exist", v.Var)
		}
		if cur.Type() != object.NUMBER {
			t.Errorf("Variable %s had wrong type: %s", v.Var, cur.String())
		}
		out := cur.(*object.NumberObject).Value
		if out != v.Val {
			t.Errorf("Expected %s to be %f, got %f", v.Var, v.Val, out)
		}
	}
}

// TestData tests that invalid data items cause the program to fail.
func TestData(t *testing.T) {
	type Test struct {
		Input string
		Valid bool
	}

	vars := []Test{{Input: `10 DATA 2,1,2`, Valid: true},
		{Input: `10 DATA "2","1","2"`, Valid: true},
		{Input: `10 DATA "2","steve",2
`, Valid: true},
		{Input: `10 DATA LET, b, c, + , -`, Valid: false},
	}

	//
	// Test reading each set of data.
	//
	for _, v := range vars {

		tokener := tokenizer.New(v.Input)
		_, err := New(tokener)

		if v.Valid {
			if err != nil {
				t.Errorf("Expected error, received one: %s!", err.Error())
			}
		} else {
			if err == nil {
				t.Errorf("Expected error, received none for input %s", v.Input)
			}
		}
	}
}

// TestDefFN tests that parsing user-defined functions works, somewhat.
func TestDefFn(t *testing.T) {
	tests := []string{
		"10 DEF FN square(x,y) = ",
		"10 DEF FN square(x) 33 ",
		"10 DEF FN square(x) ",
		"10 DEF FN square(x ",
		"10 DEF FN square(x 32",
		"10 DEF FN square",
		"10 DEF FN square 3 \"steve\"",
		"10 DEF FN 3",
		"10 DEF FN ",
		"10 DEF 3 ",
	}

	for _, test := range tests {

		_, err := FromString(test)
		if err == nil {
			t.Errorf("Expected error parsing '%s' - saw none", test)
		} else {
			if !strings.Contains(err.Error(), "DEF") {
				t.Errorf("The error didn't seem to match a DEF FN failure: %s", err.Error())
			}
		}
	}

}

// TestDim tests basic DIM functionality for single and 2dimensional
// arrays
func TestDim(t *testing.T) {

	//
	// Valid input
	//
	input := `10 DIM a(3)
20 DIM b(3,3)
24 LET i=2
30 LET a[i]="Steve"
40 LET b=a[2]
`
	tokener := tokenizer.New(input)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}

	err = e.Run()

	if err != nil {
		t.Errorf("Found an unexpected error parsing valid DIM statements:%s", err.Error())
	}

	//
	// Ensure we got the value we expected
	//
	out := e.GetVariable("b")
	if out.Type() != object.STRING {
		t.Errorf("Variable 'b' had wrong type: %s", out.String())
	}
	res := out.(*object.StringObject).Value
	if res != "Steve" {
		t.Errorf("Expected to get 'Steve' from array, got %v", res)
	}

	//
	// Now we have a series of invalid DIM statements
	// which have the wrong types
	//
	invalid := []string{"10 DIM 3",
		"10 DIM 3,",
		"10 DIM a(\"steve\"",
		"10 DIM a(3,,,,",
		"10 DIM a(3 (",
		"10 DIM a(3,4,",
		"10 DIM a(3, \"steve\")",
		"10 DIM a(3, 4 ]",
		"10 DIM a [",
	}

	for _, test := range invalid {

		tokener = tokenizer.New(test)
		e, _ = New(tokener)
		err = e.Run()

		if err == nil {
			t.Errorf("Expected an error parsing '%s' - Got none", test)
		} else {

			if !strings.Contains(err.Error(), "DIM") {
				t.Errorf("Error '%s' didn't contain DIM", err.Error())
			}
		}
	}

	//
	// Now we'll test get/set of an array value
	//
	getSet := `
10 DIM a(2,3)
20 LET a[2,0]=33
30 LET a[2,1]=2
40 LET a[2,2]=-2
50 LET a[2,3]=0.5
60 LET c = a[2,0] + a[2,1] + a[2,2] + a[2,3]
`

	// Run the script
	tokener = tokenizer.New(getSet)
	e, err = New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Found an unexpected error: %s", err.Error())
	}

	// Ensure our result is expected
	out = e.GetVariable("c")
	if out.Type() != object.NUMBER {
		t.Errorf("Variable 'c' had wrong type: %s", out.String())
	}
	result := out.(*object.NumberObject).Value
	if result != 33.5 {
		t.Errorf("Expected sum to be 33.5, got %f", result)
	}

	//
	// Now we test that operations will fail, as expected, upon
	// mis-matched types.
	//
	failTypes := `
10 DIM a(3)
20 DIM b(3)
30 LET c=a + b
`
	// Run the script
	tokener = tokenizer.New(failTypes)
	e, err = New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", failTypes, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("We expected an error, but found none!")
	}
	if !strings.Contains(err.Error(), "non-number/non-string") {
		t.Errorf("We found the wrong kind of error!")
	}

	//
	// Final test is that we handle array dimensions that are
	// too large.
	//
	dims := []string{
		"10 DIM a(3,100300)",
		"10 DIM b(100030,3)",
		"10 DIM c(100030000)"}

	for _, test := range dims {

		tokener := tokenizer.New(test)
		e, err := New(tokener)

		if err != nil {
			t.Errorf("Error parsing %s - %s", test, err.Error())
		}
		err = e.Run()
		if err == nil {
			t.Errorf("Expected error running '%s', got none", test)
		}
		if !strings.Contains(err.Error(), "dimension too large") {
			t.Errorf("Error '%s' wasn't the expected error!", err.Error())
		}
	}

}

// TestEOF ensures that our bounds-checking of program works.
//
// The bounds-checks we've added have largely been as a result of
// fuzz-testing, as described in `FUZZING.md`.  Very useful additions but also
// fiddly and annoying.
func TestEOF(t *testing.T) {

	tests := []string{
		"10 LET a = 3 *",
		"20 LET a = ( 3 * 3 ) + ",
		"30 LET a = ( 3",
		"40 LET a = (",
		"50 LET a = 3 * 3 / 3 +",
		"60 GOSUB",
		"70 GOTO",
		"80 INPUT",
		"90 INPUT \"test\"",
		"100 INPUT \"test\", ",
		"100 LET",
		"110 LET x",
		"120 LET x=",
		"130 NEXT",
		"140 LET x=3 +",
		"140 LET x=3 + 3 - 1 + 22",
		"150 LET x=3 * 3 * 93 / 3",
		"160 IF 3 < ",
		"170 READ ",
		"10 PRINT 3 +",
		"10 PRINT 3 /",
		"10 PRINT 3 *",
		"10 PRINT ,",
		"10 IF 3 ",
		"10 IF \"steve\" ",
		"10 IF  ",
		"10 FOR I = 1 TO 3 STEP",
		"10 FOR I = 1 TO ",
		"10 FOR I = 1 ",
		"10 DIM",
		"10 DIM x",
		"10 DIM x(",
		"10 DIM x(3",
		"10 DIM x(3,",
		"10 DIM x(3,3",
		"10 DEF FN x() = ",
		"10 DEF FN x()  ",
		"10 DEF FN x( ",
		"10 DEF FN x",
		"10 DEF FN",
		"10 DEF ",
		"10 LET a =RND",
		"10 LEFT$ \"steve\"",
		"10 FOR I=1 TO 10 STEP",
		"10 FOR I=1 TO 10",
		"10 FOR I=1 TO",
		"10 FOR I=1",
		"10 FOR I=",
		"10 FOR I",
		"10 FOR ",
		"10 LET a[0",
		"10 SWAP",
		"10 SWAP a",
		"10 SWAP a,",

		// multi-line tests:
		`10 DATA 3,4,5
20 READ`,
		`10 DEF FN double(x) = x * x
20 LET a = 3 + FN double`,
		`10 DEF FN double(x) = x * x
20 LET a = 3 + FN `,
		`10 DEF FN double(x) = x * x
20 LET a = 3 + FN double( 3 `,
	}
	for _, test := range tests {

		tokener := tokenizer.New(test)
		e, err := New(tokener)

		//
		// We handle two cases
		//
		//  1.  Error parsing.
		//
		//  2.  Error running.
		//
		// In both cases we're looking for a bounds-check, the reason
		// we need to care about parse-failures is because DEF FN
		// is parsed at load-time, not run-time.
		//
		if err != nil {

			//
			if !strings.Contains(err.Error(), "end of program") {
				t.Errorf("Error '%s' wasn't an end-of-program error!", err.Error())
			}
		} else {

			err = e.Run()
			if err == nil {
				t.Errorf("Expected error running '%s', got none", test)
			} else {
				if !strings.Contains(err.Error(), "end of program") {
					t.Errorf("Error '%s' wasn't an end-of-program error!", err.Error())
				}
			}
		}
	}
}

// TestExprTerm tests that expr() errors on unclosed brackets.
func TestExprTerm(t *testing.T) {
	input := `10 LET a = ( 3 + 3 * 33
20 PRINT a "\n"
`
	tokener := tokenizer.New(input)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}

	err = e.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Unclosed bracket") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestFN tests calling user-defined functions
func TestFN(t *testing.T) {

	//
	// call a function that doesn't exist.
	//
	fail1 := `
 10 LET t = FN foo("steve")
`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail1, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "User-defined function foo doesn't exist") {
		t.Errorf("Our error-message wasn't what we expected:%s", err.Error())
	}

	//
	// call a function with the wrong number of arguments.
	//
	fail2 := `
 10 DEF FN square(x) = x * x
 20 LET t = FN square(1,2)
`

	e, err = FromString(fail2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail2, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Argument count mis-match") {
		t.Errorf("Our error-message wasn't what we expected:%s", err.Error())
	}

	//
	// call a function with an argument that is an error.
	//
	fail3 := `
 10 DEF FN square(x) = x * x
 20 LET t = FN square( "steve" + 3 )
`

	e, err = FromString(fail3)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail3, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "type mismatch") {
		t.Errorf("Our error-message wasn't what we expected:%s", err.Error())
	}

	//
	// call a working function.
	//
	ok1 := `
 10 DEF FN square(x) = x * x
 20 DEF FN hello(x) = PRINT "Hello" + x
 30 LET t = FN square( 3 )
 40 FN hello( "Steve" )
`
	e, err = FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Expected to see no error, but got one: %s", err.Error())
	}
	cur := e.GetVariable("t")
	if cur.Type() != object.NUMBER {
		t.Errorf("Variable 't' had wrong type: %s", cur.String())
	}
	out := cur.(*object.NumberObject).Value
	if out != 9 {
		t.Errorf("Expected user-defined function to give 9, got %f", out)
	}
}

// TestFor performs testing of our looping primitive
func TestFor(t *testing.T) {

	//
	// These will each fail.
	//
	fails := []string{`10 FOR I=1 TO`,
		`10 FOR I=1 3`,
		`10 FOR I 3`,
		`10 FOR I`,
		`10 FOR 3`,
		`10 FOR I="steve" TO 10`,
		`10 FOR I=1 TO "kemp"`,

		// multi-line tests
		`10 LET start="steve"
20 FOR I = start TO 20`,
		`10 LET e="steve"
20 FOR I = 1 TO e`,
		`10 LET start = 1
20 LET en = 10
30 FOR I = start TO en STEP "steve" + "steve"
40 NEXT I
`,
	}

	for _, test := range fails {

		e, err := FromString(test)
		if err != nil {
			t.Errorf("Error parsing %s - %s", test, err.Error())
		}
		err = e.Run()
		if err == nil {
			t.Errorf("Expected to see an error, but didn't.")
		}
	}

	//
	// Now a valid loop
	//

	//
	// This will work.
	//
	ok1 := `
 10 LET s = 0
 20 FOR I = 1 TO 10 STEP 1
 20  LET s = s + I
 30 NEXT I
`
	e, err := FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("We found an unexpected error: %s", err.Error())
	}

	cur := e.GetVariable("s")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable s does not exist!")
	}
	if cur.Type() != object.NUMBER {
		t.Errorf("Variable s had wrong type: %s", cur.String())
	}
	out := cur.(*object.NumberObject).Value
	if out != 55 {
		t.Errorf("Expected s to be %d, got %f", 55, out)
	}

	//
	// This will also work, backwards.
	//
	ok2 := `
 10 LET s = 0
 20 FOR I = 10 TO 1 STEP -1
 20  LET s = s - I
 30 NEXT I
`
	e, err = FromString(ok2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok2, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("We found an unexpected error: %s", err.Error())
	}

	cur = e.GetVariable("s")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable s does not exist!")
	}
	if cur.Type() != object.NUMBER {
		t.Errorf("Variable s had wrong type: %s", cur.String())
	}
	out = cur.(*object.NumberObject).Value
	if out != -55 {
		t.Errorf("Expected s to be %d, got %f", -55, out)
	}

}

// TestGosub checks that GOSUB behaviour is reasonable.
func TestGoSub(t *testing.T) {

	//
	// This will fail because the target should be a literal.
	//
	fail1 := `
 10 LET t = 200
 20 GOSUB t
200 END
`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail1, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "GOSUB should be followed by an integer") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// Now we test that GOSUB to a missing line fails
	//
	fail2 := `10 GOSUB 1000
20 END`
	e, err = FromString(fail2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail2, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Line 1000 does not exist") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// This will work.
	//
	ok1 := `
 10 LET a="Kissa"
 20 GOSUB 100
 30 END
100 LET a= "Cat"
120 RETURN
`
	e, err = FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("We found an unexpected error: %s", err.Error())
	}

	cur := e.GetVariable("a")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable a does not exist!")
	}
	if cur.Type() != object.STRING {
		t.Errorf("Variable a had wrong type: %s", cur.String())
	}
	out := cur.(*object.StringObject).Value
	if out != "Cat" {
		t.Errorf("Expected x to be %s, got %s", "Cat", out)
	}

}

// TestGoto checks that GOTO behaviour is reasonable.
func TestGoto(t *testing.T) {

	//
	// This will fail because the target should be a literal.
	//
	fail1 := `
 10 LET t = 200
 20 GOTO t
200 END
`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail1, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "GOTO should be followed by an integer") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// Now we test that GOTO of a missing line fails
	//
	fail2 := `10 GOTO 1000
20 END`
	e, err = FromString(fail2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail2, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Line 1000 does not exist") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// This will work.
	//
	ok1 := `
 10 GOTO 40
 20 LET a="Steve"
 30 END
 40 GOTO 20
`

	e, err = FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("We found an unexpected error: %s", err.Error())
	}

	cur := e.GetVariable("a")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable a does not exist!")
	}
	if cur.Type() != object.STRING {
		t.Errorf("Variable a had wrong type: %s", cur.String())
	}
	out := cur.(*object.StringObject).Value
	if out != "Steve" {
		t.Errorf("Expected x to be %s, got %s", "Steve", out)
	}

}

// TestIF performs testing of our IF implementation.
func TestIF(t *testing.T) {
	type Test struct {
		Input  string
		Result float64
	}

	tests := []Test{
		{Input: "10 IF 1 < 3 THEN LET res=3", Result: 3},
		{Input: "20 IF 1 > 3 THEN LET res=1 ELSE let res=33", Result: 33},
		{Input: "30 IF 1 THEN LET res=21 ELSE PRINT \"OK\n\":", Result: 21},
		{Input: "30 IF 1 > 3 THEN LET res=21\n", Result: -1},
		{Input: "30 IF 1 < 3 THEN LET res=21\n", Result: -1},
		{Input: "30 IF 1 <> \"steve\" + 3 THEN LET res=21\n", Result: -1},
		{Input: "30 IF 1=0 XOR  3 <> \"steve\" + 3 THEN LET foo=3 ELSE LET res=21\n", Result: -1},
	}

	for _, test := range tests {

		tokener := tokenizer.New(test.Input)
		e, err := New(tokener)
		if err != nil {
			t.Errorf("Error parsing %s - %s", test.Input, err.Error())
		}

		e.Run()

		if test.Result > 0 {
			cur := e.GetVariable("res")
			if cur.Type() == object.ERROR {
				t.Errorf("Variable 'res' does not exist for %s", test.Input)
			}
			if cur.Type() != object.NUMBER {
				t.Errorf("Variable 'res' had wrong type: %s", cur.String())
			}
			out := cur.(*object.NumberObject).Value
			if out != test.Result {
				t.Errorf("Expected 'res' to be %f, got %f", test.Result, out)
			}
		}
	}

	//
	// Failure to parse
	//
	fail1 := `
10 IF 3 <> 3 3
`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Failed to parse program")
	}

	err = e.Run()
	if err == nil {
		t.Errorf("Expected runtime-error, received none")
	}
	if !strings.Contains(err.Error(), "expected THEN after IF") {
		t.Errorf("The error we found was not what we expected: %s", err.Error())
	}

	//
	// Now some simple tests of "GOTO" in IF
	//
	// The first with a `GOTO` token, the second without.
	//
	test1 := `
10 IF 2 < 10 THEN GOTO 40
20 LET a = 3
30 END
40 LET a = 33
50 END
`
	test2 := `
10 IF 2 < 10 THEN 40 ELSE 99
20 LET a = 2
30 END
40 LET a = 313
50 END
`
	//
	// test1
	//
	e, err = FromString(test1)
	if err != nil {
		t.Errorf("Failed to parse program")
	}

	err = e.Run()
	if err != nil {
		t.Errorf("Didn't expect a runtime-error, received one %s", err.Error())
	}
	cur := e.GetVariable("a")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable 'a' does not exist for %s", test1)
	}
	if cur.Type() != object.NUMBER {
		t.Errorf("Variable 'a' had wrong type: %s", cur.String())
	}
	out := cur.(*object.NumberObject).Value
	if out != 33 {
		t.Errorf("Expected 'a' to be %d, got %f", 33, out)
	}

	//
	// test2
	//
	e, err = FromString(test2)
	if err != nil {
		t.Errorf("Failed to parse program")
	}

	err = e.Run()
	if err != nil {
		t.Errorf("Didn't expect a runtime-error, received one %s", err.Error())
	}
	cur = e.GetVariable("a")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable 'a' does not exist for %s", test2)
	}
	if cur.Type() != object.NUMBER {
		t.Errorf("Variable 'a' had wrong type: %s", cur.String())
	}
	out = cur.(*object.NumberObject).Value
	if out != 313 {
		t.Errorf("Expected 'a' to be %d, got %f", 313, out)
	}
}

// TestINPUT performs testing of our INPUT implementation.
func TestINPUT(t *testing.T) {

	//
	// These will each fail.
	//
	fails := []string{`10 INPUT "" :`,
		`10 INPUT "", `,
		`10 INPUT "", ""`,
		`10 INPUT 33, b`,
		`10 LET a = 33
20 INPUT a, b`,
	}

	for _, test := range fails {

		e, err := FromString(test)
		if err != nil {
			t.Errorf("Error parsing %s - %s", test, err.Error())
		}
		err = e.Run()
		if err == nil {
			t.Errorf("Expected to see an error, but didn't.")
		}
		if !strings.Contains(err.Error(), "INPUT") {
			t.Errorf("Our error-message wasn't what we expected")
		}
	}

	//
	// Fake buffer for reading a string from.
	//
	strBuf := strings.NewReader("STEVE\n")

	//
	// Fake buffer for reading a number from.
	//
	numBuf := strings.NewReader("3.13\n")

	//
	// Read a string
	//
	ok1 := `
10 INPUT "give me a string", a$
`
	e, err := FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}

	// Fake input
	e.STDIN = bufio.NewReader(strBuf)
	err = e.Run()
	if err != nil {
		t.Errorf("Unexpected error, reading input %s", err.Error())
	}

	//
	//
	//
	// Now a$ should be a string
	//
	cur := e.GetVariable("a$")
	if cur.Type() != object.STRING {
		t.Errorf("Variable a$ had wrong type: %s", cur.String())
	}
	out := cur.(*object.StringObject).Value
	if out != "STEVE" {
		t.Errorf("Reading INPUT returned the wrong string: %s", out)
	}

	//
	// Read a number
	//
	ok2 := `
10 LET p="Give me a number"
20 INPUT p,b
`
	e, err = FromString(ok2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok2, err.Error())
	}
	// Fake input
	e.STDIN = bufio.NewReader(numBuf)
	err = e.Run()
	if err != nil {
		t.Errorf("Unexpected error, reading input %s", err.Error())
	}
	//
	// Now b should be a number
	//
	cur = e.GetVariable("b")
	if cur.Type() != object.NUMBER {
		t.Errorf("Variable b had wrong type: %s", cur.String())
	}
	out2 := cur.(*object.NumberObject).Value
	if out2 != 3.130000 {
		t.Errorf("Reading INPUT returned the wrong number: %f", out2)
	}
}

// TestLet performs sanity-checking on our LET implementation.
func TestLet(t *testing.T) {

	//
	// Failure cases.
	//
	fails := []string{"10 LET 3\n",
		"10 LET a _ 3\n"}

	for _, fail := range fails {

		e, err := FromString(fail)
		if err != nil {
			t.Errorf("Error parsing %s - %s", fail, err.Error())
		}
		err = e.Run()
		if err == nil {
			t.Errorf("Expected to see an error, but didn't.")
		}
		if !strings.Contains(err.Error(), "LET") {
			t.Errorf("Our error-message wasn't what we expected")
		}

	}

	//
	// Now a working example.
	//
	ok1 := "10 LET a = \"Steve\"\n"
	e, err := FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("We found an unexpected error: %s", err.Error())
	}

	cur := e.GetVariable("a")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable 'a' does not exist!")
	}
	if cur.Type() != object.STRING {
		t.Errorf("Variable 'a' had wrong type: %s", cur.String())
	}
	out := cur.(*object.StringObject).Value
	if out != "Steve" {
		t.Errorf("Expected 'a' to be %s, got %s", "Steve", out)
	}

	//
	// Now a working example, without a LET statement.
	//
	ok2 := "10 s = \"Steve Kemp\"\n"
	e, err = FromString(ok2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok2, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("We found an unexpected error: %s", err.Error())
	}

	cur = e.GetVariable("s")
	if cur.Type() == object.ERROR {
		t.Errorf("Variable 's' does not exist!")
	}
	if cur.Type() != object.STRING {
		t.Errorf("Variable 's' had wrong type: %s", cur.String())
	}
	out = cur.(*object.StringObject).Value
	if out != "Steve Kemp" {
		t.Errorf("Expected 'a' to be %s, got %s", "Steve Kemp", out)
	}
}

// TestMaths tests addition, subtraction, multiplication, division, etc.
func TestMaths(t *testing.T) {
	type Test struct {
		Input  string
		Result float64
	}

	tests := []Test{
		{Input: "3 + 3", Result: 6},
		{Input: "3 - 1", Result: 2},
		{Input: "6 / 2", Result: 3},
		{Input: "6 * 5", Result: 30},
		{Input: "2 ^ 3", Result: 8},
		{Input: "4 % 2", Result: 0},
		{Input: " ( BIN 00001111 ) OR ( BIN 01110000 )", Result: 255 - 128},
		{Input: "129 AND 128", Result: 128},
		{Input: "128 XOR 1", Result: 129},
	}

	for _, test := range tests {

		tokener := tokenizer.New("LET x =" + test.Input + "\n")
		e, err := New(tokener)
		if err != nil {
			t.Errorf("Error parsing %s - %s", test.Input, err.Error())
		}

		e.Run()

		cur := e.GetVariable("x")
		if cur.Type() == object.ERROR {
			t.Errorf("Variable x does not exist!")
		}
		if cur.Type() != object.NUMBER {
			t.Errorf("Variable x had wrong type: %s", cur.String())
		}
		out := cur.(*object.NumberObject).Value
		if out != test.Result {
			t.Errorf("Expected x to be %f, got %f", test.Result, out)
		}
	}
}

// TestMismatchedTypes tests that expr() errors on mismatched types.
func TestMismatchedTypes(t *testing.T) {
	input := `10 LET a=3
20 LET b="steve"
30 LET c = a + b
`
	tokener := tokenizer.New(input)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}

	err = e.Run()

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
	tokener := tokenizer.New(input)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}

	err = e.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "handles integers") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestNext ensures that the NEXT statement is sane.
func TestNext(t *testing.T) {

	//
	// This will fail because the "I" variable is a string.
	//
	fail1 := `
10 FOR I = 1 TO 20
20   LET I = "steve"
30 NEXT I
`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail1, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "NEXT variable I is not a number") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// This will fail because NEXT requires a variable name.
	//
	fail2 := `
05 LET SUM = 0
10 FOR I = 1 TO 20
20   LET SUM = SUM + 1
30 NEXT 3
`
	e, err = FromString(fail2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail2, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "expected IDENT after NEXT in FOR loop") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// This will fail because the NEXT variable is unknown.
	//
	fail3 := `
05 LET SUM = 0
10 FOR I = 1 TO 20
20   LET SUM = SUM + 1
30 NEXT J
`
	e, err = FromString(fail3)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail3, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "NEXT J found - without opening FOR") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// Now a working example.
	//
	ok1 := `
05 LET SUM = 0
10 FOR I = 1 TO 20
20   LET SUM = SUM + 1
30 NEXT I
`
	e, err = FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Expected no error, but found one: %s", err.Error())
	}

	//
	// And another, with only a single iteration.
	//
	ok2 := `
05 LET SUM = 0
10 FOR I = 1 TO 1
20   LET SUM = SUM + 1
30 NEXT I
`
	e, err = FromString(ok2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok2, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Expected no error, but found one: %s", err.Error())
	}
}

// TestRead ensures that the READ statement is sane.
func TestRead(t *testing.T) {

	//
	// This will fail because READ requires an ident.
	//
	fail1 := `
10 DATA "foo", "bar", "baz"
20 READ 3
`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail1, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "expected identifier") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// This will fail because we READ too far.
	//
	fail2 := `
10 DATA "a", "b", "c"
20 READ a, b, c, d, e, f
`
	e, err = FromString(fail2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail2, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "read past the end of our DATA storage") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// Now a working example.
	//
	ok1 := `
10 DATA "Cat", "Kissa"
20 READ a
`
	e, err = FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Expected no error, but found one: %s", err.Error())
	}

	//
	// Now we should be able to validate our read succeeded.
	//
	out := e.GetVariable("a")
	if out.Type() != object.STRING {
		t.Errorf("Variable %s had wrong type: %s", "a", out.String())
	}
	val := out.(*object.StringObject).Value
	if val != "Cat" {
		t.Errorf("Expected %s to be %s, got %s", "a", "Cat", val)
	}

	//
	// Now a "working" example.
	//
	ok2 := `
10 DATA "Cat", "Kissa"
20 READ ,,,,,,,,,,,,,,,,,,,,,,
`
	e, err = FromString(ok2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok2, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Expected no error, but found one: %s", err.Error())
	}

}

// TestRem ensures we get some coverage of swallowLine
func TestRem(t *testing.T) {

	tests := []string{
		"10 REM ",
		"20 REM\n",
		"30 REM REM REM",
	}
	for _, test := range tests {

		tokener := tokenizer.New(test)
		e, err := New(tokener)
		if err != nil {
			t.Errorf("unexpected error parsing '%s' - %s", test, err.Error())
		}

		err = e.Run()
		if err != nil {
			t.Errorf("Unexpected error running '%s' - %s", test, err.Error())
		}
	}
}

// TestReturn tests that RETURN works as expected
func TestReturn(t *testing.T) {

	//
	// This will fail because there has been no GOSUB.
	//
	fail1 := `10 RETURN`

	e, err := FromString(fail1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail1, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "RETURN without GOSUB") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// Now a valid subroutine call
	//
	ok1 := `
10 GOSUB 30
20 END
30 RETURN
`
	e, err = FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Expected no error, but found one: %s", err.Error())
	}
}

// TestRun tests the Run method of our interpreter - just looking for
// any unopen for-loops.
func TestRun(t *testing.T) {
	input := `10 LET SUM = 0
20 FOR I = 1 TO 20
30 LET SUM = SUM + I
`
	tokener := tokenizer.New(input)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}

	err = e.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "unclosed FOR loop") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestSwap ensures that the SWAP statement is sane.
func TestSwap(t *testing.T) {

	//
	// Some simple failures
	//
	fails := []string{
		"10 SWAP 3",
		"10 SWAP a",
		"10 SWAP a,3",
		"10 SWAP \"steve\", 3",
		"10 SWAP \"steve\" 3",
		"10 SWAP a a a a",
	}

	for _, test := range fails {

		tokener := tokenizer.New(test)
		e, err := New(tokener)
		if err != nil {
			t.Errorf("Error parsing %s - %s", test, err.Error())
		}

		err = e.Run()
		if err == nil {
			t.Errorf("Expected an error - found none")
		}
		if !strings.Contains(err.Error(), "SWAP") {
			t.Errorf("Got an error, but it was the wrong one: %s", err.Error())
		}
	}

	//
	// Simple Swap
	//
	test1 := `10 REM simple swap
20 a ="Kemp"
30 LET b = 44
40 SWAP a,b`

	tokener := tokenizer.New(test1)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", test1, err.Error())
	}
	err = e.Run()

	if err != nil {
		t.Errorf("Found an unexpected error:%s", err.Error())
	}

	A := e.GetVariable("a")
	if A.Type() != object.NUMBER {
		t.Errorf("Unexpectedly managed to retrieve a missing variable")
	}
	B := e.GetVariable("b")
	if B.Type() != object.STRING {
		t.Errorf("Unexpectedly managed to retrieve a missing variable")
	}

	test2 := `10 REM indexed-swap
	20 DIM A(3)
	30 A[2] = "Steve"
	40 A[1] = "Kemp"
	50 SWAP A[1], A[2]
	`
	tokener = tokenizer.New(test2)
	e, err = New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", test2, err.Error())
	}
	err = e.Run()

	if err != nil {
		t.Errorf("Found an unexpected error:%s", err.Error())
	}

	var a []int
	a = append(a, 1)
	A = e.GetArrayVariable("A", a)
	if A.Type() != object.STRING {
		t.Errorf("Array variable has the wrong type")
	}
	if A.(*object.StringObject).Value != "Steve" {
		t.Errorf("Failed to swap array")
	}

	b := append(a, 2)
	B = e.GetArrayVariable("A", b)
	if B.Type() != object.STRING {
		t.Errorf("Array variable has the wrong type")
	}
	if B.(*object.StringObject).Value != "Kemp" {
		t.Errorf("Failed to swap array")
	}

}

// TestStringFail tests that expr() errors on bogus string operations.
func TestStringFail(t *testing.T) {
	input := `10 LET a="steve"
20 LET b="steve"
30 LET c = a - b
`
	tokener := tokenizer.New(input)
	e, err := New(tokener)
	if err != nil {
		t.Errorf("Error parsing %s - %s", input, err.Error())
	}

	err = e.Run()

	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "not supported for strings") {
		t.Errorf("Our error-message wasn't what we expected")
	}
}

// TestTrace tests that getting/setting the tracing-flag works as expected.
func TestTrace(t *testing.T) {
	input := `10 PRINT "OK\n"`
	tokener := tokenizer.New(input)
	e, _ := New(tokener)

	// Tracing is off by default
	if e.GetTrace() != false {
		t.Errorf("tracing should not be enabled by default")
	}

	// Enable tracing
	e.SetTrace(true)

	// Tracing is now on
	if e.GetTrace() != true {
		t.Errorf("tracing should have been enabled, but was not")
	}

	//
	// Run a script which calls multiple functions which
	// output trace-text.  This is a bit ad-hoc.
	//
	ok1 := `
 10 DEF FN square(x) = x * x
 20 LET t = FN square( 3 )
 30 LET r = RND 3
`
	e, err := FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	e.SetTrace(true)
	err = e.Run()

	if err != nil {
		t.Errorf("Expected to see no error, but got one: %s", err.Error())
	}

}

// TestVariables gets/sets some variables and ensures they work
func TestVariables(t *testing.T) {

	type Test struct {
		Name   string
		Object object.Object
	}

	var vars []Test

	//
	// Setup some test variables.
	//
	vars = append(vars, Test{Name: "number", Object: object.Number(33)})
	vars = append(vars, Test{Name: "string", Object: object.String("Steve")})
	vars = append(vars, Test{Name: "error", Object: object.Error("Blah")})

	//
	// Test getting/setting each variable.
	//
	for _, v := range vars {

		input := `10 PRINT "OK\n"`
		tokener := tokenizer.New(input)
		e, _ := New(tokener)

		//
		// By default the variable won't exist.
		//
		cur := e.GetVariable(v.Name)
		if cur.Type() != object.ERROR {
			t.Errorf("Unexpectedly managed to retrieve a missing variable")
		}

		//
		// Set it
		//
		e.SetVariable(v.Name, v.Object)

		//
		// Ensure it was set
		//
		cur = e.GetVariable(v.Name)
		if cur.Type() != v.Object.Type() {
			t.Errorf("Retrieved variable '%s' had the wrong type %s != %s", v.Name, cur.Type(), v.Object.Type())
		}
	}
}

// TestZero ensures that division/modulo by zero errors.
func TestZero(t *testing.T) {

	divTests := []string{
		`10 LET a = 3 / 0
`,
		`10 LET a = 3
20 LET b = 0
30 LET c = a / b
`}

	for _, div := range divTests {
		e, err := FromString(div)
		if err != nil {
			t.Errorf("Error parsing %s - %s", div, err.Error())
		}
		err = e.Run()
		if err == nil {
			t.Errorf("Expected to see an error, but didn't.")
		}
		if !strings.Contains(err.Error(), "Division by zero") {
			t.Errorf("Our error-message wasn't what we expected:%s", err.Error())
		}
	}

	modTests := []string{
		`10 LET a = 3 % 0
`,
		`10 LET a = 3
20 LET b = 0
30 LET c = a % b
`}

	for _, mod := range modTests {
		e, err := FromString(mod)
		if err != nil {
			t.Errorf("Error parsing %s - %s", mod, err.Error())
		}
		err = e.Run()
		if err == nil {
			t.Errorf("Expected to see an error, but didn't.")
		}
		if !strings.Contains(err.Error(), "MOD 0 is an error") {
			t.Errorf("Our error-message wasn't what we expected:%s", err.Error())
		}
	}

}
