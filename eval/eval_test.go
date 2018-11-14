// eval_test.go - Simple test-cases for our evaluator.

package eval

import (
	"testing"

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/tokenizer"
)

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
		// TODO: if 3 THEN
		// TODO: if "steve" THEN
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
