// eval_test.go - Simple test-cases for our evaluator.

//
// TODO:
//  DEF FN
//  CALL FN
//  Builtin
//  IF
//  FOR
//

package eval

import (
	"strings"
	"testing"

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

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
		{Input: `170 IF 1=1 XOR 33=2 THEN let J=358 ELSE LET J=131`, Var: "J", Val: 131},
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

		tokener := tokenizer.New(v.Input + "\n")
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
		{Input: `10 DATA LET, b, c`, Valid: false},
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
		"10 IF 3 ",
		"10 IF \"steve\" ",
		"10 IF  ",
		"10 FOR I = 1 TO 3 STEP",
		"10 FOR I = 1 TO ",
		"10 FOR I = 1 ",
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

		// multi-line tests:
		`10 DATA 3,4,5
20 READ`,
		`10 DEF FN double(x) = x * x
20 LET a = 3 + FN double`,
		`10 DEF FN double(x) = x * x
20 LET a = 3 + FN `,
		`10 DEF FN double(x) = x * x
20 LET a = 3 + FN( 3 `,
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

		// multi-line tests
		`10 LET start="steve"
20 FOR I = start TO 20`,
		`10 LET end="steve"
20 FOR I = 1 TO end`,
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
	}

	for _, test := range tests {

		// TODO: Tests fail without the trailing newline - BUG
		tokener := tokenizer.New(test.Input + "\n")
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
	if !strings.Contains(err.Error(), "Expected THEN after IF") {
		t.Errorf("The error we found was not what we expected: %s", err.Error())
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
	// Read a string
	//
	// NOTE: This requires (hacked) support in eval.go
	//
	ok1 := `
10 INPUT "give me a string", a$
`
	e, err := FromString(ok1)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok1, err.Error())
	}
	err = e.Run()
	if err != nil {
		t.Errorf("Unexpected error, reading input %s", err.Error())
	}
	//
	// Now a$ should be a string
	//
	cur := e.GetVariable("a$")
	if cur.Type() != object.STRING {
		t.Errorf("Variable a$ had wrong type: %s", cur.String())
	}
	out := cur.(*object.StringObject).Value
	if out != "steve" {
		t.Errorf("Reading INPUT returned the wrong string: %s", out)
	}

	//
	// Read a number
	//
	// NOTE: This requires (hacked) support in eval.go
	//
	ok2 := `
10 LET p="Give me a number"
20 INPUT p,b
`
	e, err = FromString(ok2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", ok2, err.Error())
	}
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
	if out2 != 3.21 {
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
	if !strings.Contains(err.Error(), "Expected IDENT after NEXT in FOR loop") {
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
	if !strings.Contains(err.Error(), "Expected identifier") {
		t.Errorf("Our error-message wasn't what we expected")
	}

	//
	// This will fail because we READ too far.
	//
	fail2 := `
10 DATA "a", "b", "c"
20 READ a, b, c, d
`
	e, err = FromString(fail2)
	if err != nil {
		t.Errorf("Error parsing %s - %s", fail2, err.Error())
	}
	err = e.Run()
	if err == nil {
		t.Errorf("Expected to see an error, but didn't.")
	}
	if !strings.Contains(err.Error(), "Read past the end of our DATA storage") {
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
	if !strings.Contains(err.Error(), "Unclosed FOR loop") {
		t.Errorf("Our error-message wasn't what we expected")
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
