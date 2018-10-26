// Package token contains the tokens we understand when it comes
// to parsing our BASIC input.
package token

import (
	"fmt"
	"strings"
)

// Type is a string
type Type string

// Token contains a single token.
type Token struct {
	// Type holds the type of the token
	Type Type

	// Literal holds the literal value.
	Literal string
}

// pre-defined token-types
const (
	// Core
	EOF     = "EOF"     // End of file
	NEWLINE = "NEWLINE" // Newlines are kept in our lexer-stream
	LINENO  = "LINENO"  // Line-number of each input.

	// Types
	IDENT   = "IDENT"   // Identifier (i.e. variable name)
	INT     = "INT"     // integer literal
	STRING  = "STRING"  // string literal
	BUILTIN = "BUILTIN" // builtin-function

	// Implemented keywords.
	END    = "END"
	GOSUB  = "GOSUB"
	GOTO   = "GOTO"
	INPUT  = "INPUT"
	LET    = "LET"
	PRINT  = "PRINT"
	REM    = "REM"
	RETURN = "RETURN"

	// Did I mention that for-loops work?  :D
	FOR  = "FOR"
	NEXT = "NEXT"
	STEP = "STEP"
	TO   = "TO"

	// And conditionals?
	IF   = "IF"
	THEN = "THEN"
	ELSE = "ELSE"

	// Binary operators
	AND = "AND"
	OR  = "OR"

	// Woo-operators
	ASSIGN   = "=" // LET x = 3
	ASTERISK = "*" // integer multiplication
	COMMA    = "," // PRINT 3, 54
	MINUS    = "-" // integer subtraction
	MOD      = "%" // integer modulus
	PLUS     = "+" // integer addition
	SLASH    = "/" // integer division

	COLON    = ":"
	LBRACKET = "("
	RBRACKET = ")"

	// Comparison functions.
	GT         = ">"
	GT_EQUALS  = ">="
	LT         = "<"
	LT_EQUALS  = "<="
	NOT_EQUALS = "<>"
)

// reversed keywords
var keywords = map[string]Type{
	"and":    AND,
	"else":   ELSE,
	"end":    END,
	"for":    FOR,
	"gosub":  GOSUB,
	"goto":   GOTO,
	"if":     IF,
	"input":  INPUT,
	"let":    LET,
	"next":   NEXT,
	"or":     OR,
	"print":  PRINT,
	"rem":    REM,
	"return": RETURN,
	"step":   STEP,
	"then":   THEN,
	"to":     TO,
}

// LookupIdentifier used to determine whether identifier is keyword nor not.
// We handle both upper-case and lower-cased keywords, for example both
// "print" and "PRINT" are considered identical.
func LookupIdentifier(identifier string) Type {
	id := strings.ToLower(identifier)
	if tok, ok := keywords[id]; ok {
		return tok
	}
	return IDENT
}

// String creates a string-representation of a token
func (t Token) String() string {

	//
	// Special-case newline-token doesn't need an embedded newline.
	//
	lit := t.Literal
	if t.Type == NEWLINE {
		lit = "\\n"
	}
	return (fmt.Sprintf("Token{Type:%s Value:%s}", t.Type, lit))
}
