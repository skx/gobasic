// Package eval contains our evaluator
//
// This is pretty simple:
//
//  * The program is an array of tokens.
//
//  * We have one statement per line.
//
//  * We handle the different types of statements in their own functions.
//
package eval

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/skx/gobasic/object"
	"github.com/skx/gobasic/token"
	"github.com/skx/gobasic/tokenizer"
)

// Interpreter holds our state.
type Interpreter struct {

	// The program we execute is nothing more than an array of tokens.
	program []token.Token

	// Should we finish execution?
	// This is set by the `END` statement.
	finished bool

	// We execute from the given offset.
	//
	// Sequential execution just means bumping this up by one each
	// time we execute an instruction, or pick off the arguments to
	// one.
	//
	// But set it to 17, or some other random value, and you've got
	// a GOTO implemented!
	offset int

	// A stack for handling GOSUB/RETURN calls
	gstack *Stack

	// vars holds the variables set in the program, via LET.
	vars *Variables

	// STDIN is an input-reader used for the INPUT statement
	STDIN *bufio.Reader

	// Hack: Was the previous statement a GOTO?
	jump bool

	// lines is a lookup table - the key is the line-number of
	// the source program, and the value is the offset in our
	// program-array that this is located at.
	lines map[string]int

	// functions holds builtin-functions
	functions *Builtins
}

// New is our constructor.
//
// Given a lexer we store all the tokens it produced in our array, and
// initialise some other state.
func New(stream *tokenizer.Tokenizer) *Interpreter {
	t := &Interpreter{offset: 0}

	// setup a stack for holding line-numbers for GOSUB/RETURN
	t.gstack = NewStack()

	// setup storage for variable-contents
	t.vars = NewVars()

	// Built-in functions are stored here.
	t.functions = NewBuiltins()

	// Add in our builtins.
	//
	// These are implemented in golang in the file builtins.go
	//
	t.functions.Register("ABS", 1, ABS)
	t.functions.Register("ACS", 1, ACS)
	t.functions.Register("ASN", 1, ASN)
	t.functions.Register("ATN", 1, ATN)
	t.functions.Register("COS", 1, COS)
	t.functions.Register("EXP", 1, EXP)
	t.functions.Register("INT", 1, INT)
	t.functions.Register("LN", 1, LN)
	t.functions.Register("PI", 0, PI)
	t.functions.Register("RND", 1, RND)
	t.functions.Register("SGN", 1, SGN)
	t.functions.Register("SIN", 1, SIN)
	t.functions.Register("SQR", 1, SQR)
	t.functions.Register("TAN", 1, TAN)

	// Primitives that operate upon strings
	t.functions.Register("LEFT$", 3, LEFT)
	t.functions.Register("LEN", 1, LEN)
	t.functions.Register("RIGHT$", 3, RIGHT)

	// allow reading from STDIN
	t.STDIN = bufio.NewReader(os.Stdin)

	//
	// Setup a map to hold our jump-targets
	//
	t.lines = make(map[string]int)

	//
	// Save the tokens that our program consists of, one by one,
	// until we hit the end.
	//
	// We also record the offset at which each line starts, which
	// means that the GOTO & GOSUB statements don't need to scan
	// the program from start to finish to find the destination
	// to jump to.
	//
	offset := 0
	for {
		tok := stream.NextToken()
		if tok.Type == token.EOF {
			break
		}

		// Did we find a line-number?
		if tok.Type == token.LINENO {

			// Save the offset in the map
			line := tok.Literal

			// Already an offset?  That means we
			// have duplicate line-numbers
			if t.lines[line] != 0 {
				fmt.Printf("WARN: Line %s is duplicated - GOTO/GOSUB behaviour is undefined\n", line)
			}
			t.lines[line] = offset
		}

		// Regardless append the token to our array
		t.program = append(t.program, tok)

		offset++
	}

	return t
}

////
//
// Helpers for stuff
//
////

// factor
func (e *Interpreter) factor() object.Object {

	tok := e.program[e.offset]
	switch tok.Type {
	case token.LBRACKET:
		// skip past the lbracket
		e.offset++

		// handle the expr
		ret := e.expr()

		// skip past the rbracket
		tok = e.program[e.offset]
		if tok.Type != token.RBRACKET {
			fmt.Printf("Unclosed bracket around expression!\n")
			os.Exit(1)
		}
		e.offset++

		// Return the result of the sub-expression
		return (ret)
	case token.INT:
		i, err := strconv.ParseFloat(tok.Literal, 64)
		if err == nil {
			e.offset++
			return &object.NumberObject{Value: i}
		}
		fmt.Printf("Failed to convert %s -> float64 %s\n", tok.Literal, err.Error())
		os.Exit(3)

	case token.STRING:
		e.offset++
		return &object.StringObject{Value: tok.Literal}

	case token.IDENT:

		//
		// Get the contents of the variable.
		//
		val := e.GetVariable(tok.Literal)
		e.offset++
		return val
	}

	fmt.Printf("factor() - unhandled token: %v\n", tok)
	os.Exit(33)

	// never-reached
	return &object.NumberObject{Value: 1}
}

// terminal
func (e *Interpreter) term() object.Object {

	f1 := e.factor()

	tok := e.program[e.offset]

	for tok.Type == token.ASTERISK ||
		tok.Type == token.SLASH ||
		tok.Type == token.MOD {

		// skip the operator
		e.offset++

		// get the next factor
		f2 := e.factor()

		//
		// Type-check
		//
		if f1.Type() != object.NUMBER {
			fmt.Printf("term() only handles integers")
			os.Exit(3)
		}
		if f2.Type() != object.NUMBER {
			fmt.Printf("term() only handles integers")
			os.Exit(3)
		}

		if tok.Type == token.ASTERISK {
			f1 = &object.NumberObject{Value: f1.(*object.NumberObject).Value * f2.(*object.NumberObject).Value}
		}
		if tok.Type == token.SLASH {
			f1 = &object.NumberObject{Value: f1.(*object.NumberObject).Value / f2.(*object.NumberObject).Value}
		}
		if tok.Type == token.MOD {
			f1 = &object.NumberObject{Value: float64(int(f1.(*object.NumberObject).Value) % int(f2.(*object.NumberObject).Value))}
		}

		// repeat?
		tok = e.program[e.offset]
	}
	return f1
}

// expression
func (e *Interpreter) expr() object.Object {

	t1 := e.term()

	tok := e.program[e.offset]

	for tok.Type == token.PLUS ||
		tok.Type == token.MINUS {

		// skip the operator
		e.offset++

		t2 := e.term()

		//
		// Strings can be joined
		//
		if t1.Type() != t2.Type() {
			fmt.Printf("expr() - type mismatch\n")
			os.Exit(1)
		}

		//
		// Strings?
		//
		if t1.Type() == object.STRING {

			s1 := t1.(*object.StringObject).Value
			s2 := t2.(*object.StringObject).Value

			if tok.Type == token.PLUS {
				t1 = &object.StringObject{Value: s1 + s2}
			} else {
				fmt.Printf("Token not handled for two strings: %s\n", tok.Literal)
				os.Exit(2)
			}

		} else {

			//
			// Working with numbers.
			//
			n1 := t1.(*object.NumberObject).Value
			n2 := t2.(*object.NumberObject).Value

			if tok.Type == token.PLUS {
				t1 = &object.NumberObject{Value: n1 + n2}
			} else if tok.Type == token.MINUS {
				t1 = &object.NumberObject{Value: n1 - n2}
			} else {
				fmt.Printf("Token not handled for two numbers: %s\n", tok.Literal)
				os.Exit(2)

			}
		}
		// repeat?
		tok = e.program[e.offset]
	}

	return t1
}

// compare runs a comparison function (!)
func (e *Interpreter) compare() bool {

	// Get the first statement
	t1 := e.expr()

	// Get the comparison function
	op := e.program[e.offset]
	e.offset++

	// Get the second expression
	t2 := e.expr()

	//
	// We'll handle string equality testing here.
	//
	if t1.Type() == object.STRING && t2.Type() == object.STRING {

		v1 := t1.(*object.StringObject).Value
		v2 := t2.(*object.StringObject).Value

		switch op.Type {
		case token.ASSIGN:
			if v1 == v2 {
				return true
			}
		case token.NOT_EQUALS:
			if v1 == v2 {
				return true
			}
		}
		return false
	}

	//
	// Type-checks because our comparision function only works
	// on integers.
	//
	if t1.Type() != object.NUMBER {
		fmt.Printf("compare() only accepts integers right now")
		os.Exit(2)
	}
	if t2.Type() != object.NUMBER {
		fmt.Printf("compare() only accepts integers right now")
		os.Exit(2)
	}

	// Cast to int.
	v1 := t1.(*object.NumberObject).Value
	v2 := t2.(*object.NumberObject).Value

	switch op.Type {
	case token.ASSIGN:
		if v1 == v2 {
			return true
		}

	case token.GT:
		if v1 > v2 {
			return true
		}
	case token.GT_EQUALS:
		if v1 >= v2 {
			return true
		}
	case token.LT:
		if v1 < v2 {
			return true
		}

	case token.LT_EQUALS:
		if v1 <= v2 {
			return true
		}

	case token.NOT_EQUALS:
		if v1 != v2 {
			return true
		}
	default:
		fmt.Printf("Unknown comparison: %v\n", op)
		os.Exit(32)
	}
	return false
}

// Call the built-in with the given name if we can.
func (e *Interpreter) callBuiltin(name string) (object.Object, error) {

	if e.functions.Exists(name) {

		//
		// OK the function exists.
		//
		// Fetch it, so we know how many arguments
		// it should expect.
		//
		n, fun := e.functions.Get(name)

		//
		// skip past the function-call itself
		//
		e.offset++

		var args []token.Token
		for i := 0; i < n; i++ {
			args = append(args, e.program[e.offset])
			e.offset++
		}

		//
		// Call the function
		//
		out, err := fun(*e, args)
		return out, err
	}
	return nil, nil
}

////
//
// Statement-handlers
//
////

// runForLoop handles a FOR loop
func (e *Interpreter) runForLoop() error {
	// we expect "ID = NUM to NUM [STEP NUM]"

	// Bump past the FOR token
	e.offset++

	// We now expect a token
	target := e.program[e.offset]
	e.offset++
	if target.Type != token.IDENT {
		return fmt.Errorf("Expected IDENT after FOR, got %v", target)
	}

	// Now an EQUALS
	eq := e.program[e.offset]
	e.offset++
	if eq.Type != token.ASSIGN {
		return fmt.Errorf("Expected = after 'FOR %s' , got %v", target.Literal, eq)
	}

	// Now an integer
	startI := e.program[e.offset]
	e.offset++
	if startI.Type != token.INT {
		return fmt.Errorf("Expected INT after 'FOR %s=', got %v", target.Literal, startI)
	}

	start, err := strconv.ParseFloat(startI.Literal, 64)
	if err != nil {
		return fmt.Errorf("Failed to convert %s to an int %s", startI.Literal, err.Error())
	}

	// Now TO
	to := e.program[e.offset]
	e.offset++
	if to.Type != token.TO {
		return fmt.Errorf("Expected TO after 'FOR %s=%s', got %v", target.Literal, startI, to)
	}

	// Now an integer/variable
	endI := e.program[e.offset]
	e.offset++

	var end int

	if endI.Type == token.INT {
		v, err := strconv.ParseFloat(endI.Literal, 64)
		if err != nil {
			return fmt.Errorf("Failed to convert %s to an int %s", endI.Literal, err.Error())
		}

		end = int(v)
	} else if endI.Type == token.IDENT {

		x := e.GetVariable(endI.Literal)
		if x.Type() != object.NUMBER {
			return fmt.Errorf("End-variable must be an integer!")
		}
		end = int(x.(*object.NumberObject).Value)
	} else {
		return fmt.Errorf("Expected INT/VARIABLE after 'FOR %s=%s TO', got %v", target.Literal, startI, endI)
	}

	// Default step is 1.
	stepI := "1"

	// Is the next token a step?
	if e.program[e.offset].Type == token.STEP {
		e.offset++

		s := e.program[e.offset]
		e.offset++
		if s.Type != token.INT {
			return fmt.Errorf("Expected INT after 'FOR %s=%s TO %s STEP', got %v", target.Literal, startI, endI, s)
		}
		stepI = s.Literal
	}

	step, err := strconv.ParseFloat(stepI, 64)
	if err != nil {
		return fmt.Errorf("Failed to convert %s to an int %s", stepI, err.Error())
	}

	//
	// Now we can record the important details of the for-loop
	// in a hash.
	//
	// The key observersions here are that all the magic
	// really involved in the FOR-loop happens at the point
	// you interpret the "NEXT X" section.
	//
	// Handling the NEXT statement involves:
	//
	//  Incrementing the step-variable
	//  Looking for termination
	//  If not over then "jumping back".
	//
	// So for a for-loop we just record the start/end conditions
	// and the address of the body of the loop - ie. the next
	// token - so that the next-handler can GOTO there.
	//
	// It is almost beautifully elegent.
	//
	f := ForLoop{id: target.Literal,
		offset: e.offset,
		start:  int(start),
		end:    int(end),
		step:   int(step)}

	//
	// Set the variable to the starting-value
	//
	e.vars.Set(target.Literal, &object.NumberObject{Value: start})

	//
	// And record our loop - keyed on the name of the variable
	// which is used as the index.  This allows easy and natural
	// nested-loops.
	//
	// Did I say this is elegent?
	//
	AddForLoop(f)
	return nil
}

// runGOSUB handles a control-flow change
func (e *Interpreter) runGOSUB() error {

	// Skip the GOSUB-instruction itself
	e.offset++

	// Get the target
	target := e.program[e.offset]

	// We expect the target to be an int
	if target.Type != token.INT {
		return (fmt.Errorf("ERROR: GOSUB should be followed by an integer"))
	}

	//
	// We want to store the return address on our GOSUB-stack,
	// so that the next RETURN will continue execution at the
	// next instruction.
	//
	// Because we only support one statement per-line we can
	// handle that by bumping forward.  That should put us on the
	// LINENO of the following-line.
	//
	e.offset++
	e.gstack.Push(e.offset)

	//
	// Lookup the offset of the given line-number in our program/
	//
	offset := e.lines[target.Literal]

	//
	// If we found it then use it.
	//
	if offset > 0 {
		e.offset = offset
		return nil
	}

	return fmt.Errorf("Failed to GOSUB %s", target.Literal)
}

// runGOTO handles a control-flow change
func (e *Interpreter) runGOTO() error {

	// Skip the GOTO-instruction
	e.offset++

	// Get the GOTO-target
	target := e.program[e.offset]

	// We expect the target to be an int
	if target.Type != token.INT {
		return fmt.Errorf("ERROR: GOTO should be followed by an integer")
	}

	//
	// Lookup the offset of the given line-number in our program/
	//
	offset := e.lines[target.Literal]

	//
	// If we found it then use it.
	//
	if offset > 0 {
		e.offset = offset
		return nil
	}

	return fmt.Errorf("Failed to GOTO %s", target.Literal)
}

// runINPUT handles input of numbers from the user.
//
// NOTE:
//   INPUT "Foo", a   -> Reads an integer
//   INPUT "Foo", a$  -> Reads a string
func (e *Interpreter) runINPUT() error {

	// Skip the INPUT-instruction
	e.offset++

	// Get the prompt
	prompt := e.program[e.offset]
	e.offset++

	// We expect a comma
	comma := e.program[e.offset]
	e.offset++
	if comma.Type != token.COMMA {
		return fmt.Errorf("ERROR: INPUT should be : INPUT \"prompt\",var")
	}

	// Now the ID
	ident := e.program[e.offset]
	e.offset++
	if ident.Type != token.IDENT {
		return fmt.Errorf("ERROR: INPUT should be : INPUT \"prompt\",var")
	}

	//
	// Print the prompt
	//
	fmt.Printf(prompt.Literal)

	//
	// Read the input from the user.
	//
	input, _ := e.STDIN.ReadString('\n')
	input = strings.TrimRight(input, "\n")

	//
	// Now we handle the type-conversion.
	//
	if strings.HasSuffix(ident.Literal, "$") {
		// We set a string
		e.vars.Set(ident.Literal, &object.StringObject{Value: input})
		return nil
	}

	// We set an int
	i, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return err
	}

	//
	// Set the value
	//
	e.vars.Set(ident.Literal, &object.NumberObject{Value: i})
	return nil
}

// runIF handles conditional testing.
//
// There are a lot of choices to be made when it comes to IF, such as
// whether to support an ELSE section or not.  And what to allow
// inside the matching section generally:
//
// A single statement?
// A block?
//
// Here we _only_ allow:
//
//  IF $EXPR THEN $STATEMENT ELSE $STATEMENT NEWLINE
//
// $STATEMENT will only be a single expression
//
func (e *Interpreter) runIF() error {

	// Bump past the IF token
	e.offset++

	// Get the result of the comparison-function
	// against the two arguments.
	result := e.compare()

	// We now expect THEN
	target := e.program[e.offset]
	e.offset++
	if target.Type != token.THEN {
		return fmt.Errorf("Expected THEN after IF EXPR, got %v", target)
	}

	//
	// OK so if our comparison succeeded we can execute the single
	// statement between THEN + ELSE
	//
	// Otherwise between ELSE + Newline
	//
	if result {

		//
		// Execute single statement
		//
		e.RunOnce()

		//
		// If the user made a jump then we'll
		// abort here, because if the single-statement modified our
		// flow control we're screwed.
		//
		// (Because we'll start searching from the NEW location.)
		//
		//
		if e.jump {
			return nil
		}

		//
		// We get the next token, it should either be ELSE
		// or newline.  Handle the newline first.
		//
		tmp := e.program[e.offset]
		e.offset++
		if tmp.Type == token.NEWLINE {
			return nil
		}

		//
		// OK then we hit else so we skip forward until we
		// hit the newline.
		//
		for tmp.Type != token.NEWLINE {
			tmp = e.program[e.offset]
			e.offset++
		}
	} else {

		//
		// Here the test failed.
		//
		// Skip over the truthy-condition until we either
		// hit ELSE, or the newline that will terminate our
		// IF-statement.
		//
		//
		for {
			tmp := e.program[e.offset]
			e.offset++

			// If we hit the newline then we're done
			if tmp.Type == token.NEWLINE {
				return nil
			}

			// Otherwise did we hit the else?
			if tmp.Type == token.ELSE {

				// Execute the single statement
				e.RunOnce()

				// Then return.
				return nil
			}
		}
	}

	return nil
}

// runLET handles variable creation/updating.
func (e *Interpreter) runLET() error {

	// Bump past the LET token
	e.offset++

	// We now expect an ID
	target := e.program[e.offset]
	e.offset++
	if target.Type != token.IDENT {
		return fmt.Errorf("Expected IDENT after LET, got %v", target)
	}

	// Now "="
	assign := e.program[e.offset]
	if assign.Type != token.ASSIGN {
		return fmt.Errorf("Expected assignment after LET x, got %v", assign)
	}
	e.offset++

	// now we're at the expression/value/whatever
	res := e.expr()

	// Store the result
	e.vars.Set(target.Literal, res)
	return nil
}

// runNEXT handles the NEXT statement
func (e *Interpreter) runNEXT() error {
	// Bump past the NEXT token
	e.offset++

	// Get the identifier
	target := e.program[e.offset]
	e.offset++
	if target.Type != token.IDENT {
		return fmt.Errorf("Expected IDENT after NEXT in FOR loop, got %v", target)
	}

	// OK we've found the tail of a loop
	//
	// We need to bump the value of the given variable by the offset
	// and compare it against the max.
	//
	// If the max hasn't finished we loop around again.
	//
	// If it has we remove the for-loop
	//
	data := GetForLoop(target.Literal)

	//
	// Get the variable value, and increase it.
	//
	cur := e.vars.Get(target.Literal)
	if cur.Type() != object.NUMBER {
		return fmt.Errorf("NEXT variable %s is not a number!", target.Literal)
	}
	iVal := cur.(*object.NumberObject).Value

	//
	// Increment the number.
	//
	iVal += float64(data.step)

	//
	// Set it
	//
	e.vars.Set(target.Literal, &object.NumberObject{Value: iVal})

	//
	// Have we finnished?
	//
	if data.finished {
		RemoveForLoop(target.Literal)
		return nil
	}

	//
	// If we've reached our limit we mark this as complete,
	// but note that we dont' terminate to allow the actual
	// end-number to be inclusive.
	//
	if iVal == float64(data.end) {
		data.finished = true

		// updates-in-place.  bad name
		AddForLoop(data)
	}

	//
	// Otherwise loop again
	//
	e.offset = data.offset
	return nil
}

// runPRINT handles a print!
// NOTE:
//  Print basically swallows input up to the next newline.
//  However it also stops at ":" to cope with the case of printing in an IF
func (e *Interpreter) runPRINT() error {

	// Bump past the PRINT token
	e.offset++

	// Now keep lookin for things to print until we hit a newline.
	for e.offset < len(e.program) {

		// Get the token
		tok := e.program[e.offset]

		// End of the line, or statement?
		if tok.Type == token.NEWLINE || tok.Type == token.COLON {
			return nil
		}

		// Printing a literal?
		if tok.Type == token.INT || tok.Type == token.STRING {
			fmt.Printf("%s", tok.Literal)
		} else if tok.Type == token.COMMA {
			fmt.Printf(" ")
		} else if tok.Type == token.IDENT {

			//
			// This might be a variable, or a function-call.
			//
			// GetVariable handles both.
			//
			val := e.GetVariable(tok.Literal)

			if val.Type() == object.STRING {
				fmt.Printf("%s", val.(*object.StringObject).Value)
			}
			if val.Type() == object.NUMBER {
				n := val.(*object.NumberObject).Value

				// If the value is basically an
				// int then cast it to avoid
				// 3 looking like 3.0000
				if n == float64(int(n)) {
					fmt.Printf("%d", int(n))
				} else {
					fmt.Printf("%f", n)
				}
			}
		} else {
			// OK we're not printing:
			//
			//  an int
			//  a string
			//  a variable
			//
			// As a fall-back we'll assume we've been given
			// an expression, and print the result.
			//
			out := e.expr()

			if out.Type() == object.STRING {
				fmt.Printf("%s", out.(*object.StringObject).Value)
			}
			if out.Type() == object.NUMBER {
				n := out.(*object.NumberObject).Value

				// If the value is basically an
				// int then cast it to avoid
				// 3 looking like 3.0000
				if n == float64(int(n)) {
					fmt.Printf("%d", int(n))
				} else {
					fmt.Printf("%f", n)
				}
			}
		}
		e.offset++
	}

	return nil
}

// REM handles a REM statement
// This merely swallows input until the following newline / EOF.
func (e *Interpreter) runREM() error {

	for e.offset < len(e.program) {
		tok := e.program[e.offset]
		if tok.Type == token.NEWLINE {
			return nil
		}
		e.offset++
	}

	return nil
}

// RETURN handles a control-flow operation
func (e *Interpreter) runRETURN() error {

	// Stack can't be empty
	if e.gstack.Empty() {
		return fmt.Errorf("RETURN without GOSUB")
	}

	// Get the return address
	ret, err := e.gstack.Pop()
	if err != nil {
		return fmt.Errorf("Error handling RETURN: %s", err.Error())
	}

	// Return execution where we left off.
	e.offset = ret
	return nil
}

////
//
// Our core public API
//
////

// RunOnce executes a single statement.
func (e *Interpreter) RunOnce() error {

	//
	// Get the current token
	//
	tok := e.program[e.offset]
	var err error

	e.jump = false

	//
	// Handle this token
	//
	switch tok.Type {
	case token.NEWLINE:
		// NOP
	case token.LINENO:
		// NOP
	case token.END:
		e.finished = true
		return nil
	case token.FOR:
		err = e.runForLoop()
	case token.GOSUB:
		err = e.runGOSUB()
	case token.GOTO:
		err = e.runGOTO()
		e.jump = true
	case token.INPUT:
		err = e.runINPUT()
	case token.IF:
		err = e.runIF()
		e.offset--
	case token.LET:
		err = e.runLET()
		//
		// NOTE:
		//
		//   The LET statement bumps past itself
		//   So we need to ensure that the increment
		//   at the end of this case-statement doesn't
		//   run twice and get us out of sync.
		//
		//   This is annoying.
		//
		e.offset--
	case token.NEXT:
		err = e.runNEXT()
	case token.PRINT:
		err = e.runPRINT()
	case token.REM:
		err = e.runREM()
	case token.RETURN:
		err = e.runRETURN()
	case token.IDENT:
		_, err = e.callBuiltin(tok.Literal)
		e.offset--
	default:
		err = fmt.Errorf("Token not handled: %v", tok)
	}

	//
	// Ready for the next instruction
	//
	e.offset++

	//
	// Error?
	//
	if err != nil {
		return err
	}
	return nil
}

// Run launches the program, and does not return until it is over.
//
// A program will terminate when the control reaches the end of the
// final-line, or when the "END" token is encountered.
func (e *Interpreter) Run() error {

	//
	// We walk our series of tokens.
	//
	for e.offset < len(e.program) {

		err := e.RunOnce()

		if err != nil {
			return err
		}

		if e.finished {
			return nil
		}
	}

	return nil
}

// SetVariable sets the contents of a variable in the interpreterr environment.
//
// Useful for testing/embedding.
//
func (e *Interpreter) SetVariable(id string, val object.Object) {
	e.vars.Set(id, val)
}

// GetVariable returns the contents of the given variable.
//
// Useful for testing/embedding.
//
func (e *Interpreter) GetVariable(id string) object.Object {

	//
	// Before we lookup the value of a variable
	// we'll look for a built-in functin.
	//
	if e.functions.Exists(id) {

		out, err := e.callBuiltin(id)

		if err != nil {
			fmt.Printf("Error calling builtin %s - %s\n",
				id, err.Error())
			os.Exit(1)
		}
		e.offset--
		return out
	}

	return (e.vars.Get(id))
}

// RegisterBuiltin registers a function as a built-in, so that it can
// be called from the users' BASIC program.
//
// Useful for embedding.
//
func (e *Interpreter) RegisterBuiltin(name string, nArgs int, ft BuiltinSig) {
	e.functions.Register(name, nArgs, ft)
}
