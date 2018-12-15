// Package tokenizer contains the tokenizer we use for parsing BASIC programs.
//
// Given a string containing a complete BASIC program this package allows
// that to be iterated over as a series of tokens.
//
// Our interpeter is intentionally naive, and executes tokens directly, without
// any intermediary representation.
//
package tokenizer

import (
	"github.com/skx/gobasic/token"
)

// Tokenizer holds our state.
type Tokenizer struct {
	// current character position.
	position int

	// next character position.
	readPosition int

	// current character.
	ch rune

	// rune slice of input string.
	characters []rune

	// The previous token.
	prevToken token.Token
}

// New returns a Tokenizer instance from the specified string input.
func New(input string) *Tokenizer {

	//
	// NOTE: We parse line-numbers by looking for:
	//
	//  1. NEWLINE
	//  2. INT
	//
	// To ensure that we can find the line-number of the first line
	// we also setup a fake "previous" character of a newline.  This
	// means we don't actually need to prefix our input with such a thing.
	//
	l := &Tokenizer{characters: []rune(input)}
	l.prevToken.Type = token.NEWLINE
	l.readChar()
	return l
}

// readChar reads forward one character.
func (l *Tokenizer) readChar() {
	if l.readPosition >= len(l.characters) {
		l.ch = rune(0)
	} else {
		l.ch = l.characters[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken reads and returns the next available token, skipping any
// white space which might be present.
func (l *Tokenizer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case rune('='):
		tok = newToken(token.ASSIGN, l.ch)
	case rune(':'):
		tok = newToken(token.COLON, l.ch)
	case rune(';'):
		tok = newToken(token.SEMICOLON, l.ch)
	case rune(','):
		tok = newToken(token.COMMA, l.ch)
	case rune('+'):
		tok = newToken(token.PLUS, l.ch)
	case rune('-'):
		// -3 is "-3".  "3 - 4" is -1.
		if isDigit(l.peekChar()) {
			// swallow the -
			l.readChar()

			// read the number
			tok.Literal = l.readNumber()
			tok.Type = token.INT

			tok.Literal = "-" + tok.Literal

		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case rune('/'):
		tok = newToken(token.SLASH, l.ch)
	case rune('^'):
		tok = newToken(token.POW, l.ch)
	case rune('%'):
		tok = newToken(token.MOD, l.ch)
	case rune('*'):
		tok = newToken(token.ASTERISK, l.ch)
	case rune('('):
		tok = newToken(token.LBRACKET, l.ch)
	case rune(')'):
		tok = newToken(token.RBRACKET, l.ch)
	case rune('['):
		tok = newToken(token.LINDEX, l.ch)
	case rune(']'):
		tok = newToken(token.RINDEX, l.ch)
	case rune('<'):
		if l.peekChar() == rune('>') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOTEQUALS, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTEQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case rune('>'):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTEQUALS, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case rune('"'):
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case rune('\n'):
		tok.Type = token.NEWLINE
		tok.Literal = "\\n"
	case rune(0):
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
		} else {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
		}
	}
	l.readChar()

	//
	// Hack: A number that follows a newline is a line-number,
	// not an integer.
	//
	if l.prevToken.Type == token.NEWLINE && tok.Type == token.INT {
		tok.Type = token.LINENO
	}

	//
	// Store the previous token - which is used solely for our
	// line-number hack.
	//
	l.prevToken = tok

	return tok
}

// newToken is a simple helper for returning a new token.
func newToken(tokenType token.Type, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier is designed to read an identifier (name of variable,
// function, etc).
func (l *Tokenizer) readIdentifier() string {

	id := ""

	for isIdentifier(l.peekChar()) {
		id += string(l.ch)
		l.readChar()
	}
	id += string(l.ch)
	return id
}

// skip white space
func (l *Tokenizer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

// read a number, note that this only handles integers.
func (l *Tokenizer) readNumber() string {
	str := ""

	for isDigit(l.peekChar()) || l.peekChar() == rune('.') {
		str += string(l.ch)
		l.readChar()
	}
	str += string(l.ch)
	return str
}

// read a string, handling "\t", "\n", etc.
func (l *Tokenizer) readString() string {
	out := ""

	for {
		l.readChar()
		if l.ch == '"' {
			break
		}
		if l.ch == rune(0) {
			break
		}

		//
		// Handle \n, \r, \t, \", etc.
		//
		if l.ch == '\\' {
			l.readChar()

			if l.ch == rune('n') {
				l.ch = '\n'
			}
			if l.ch == rune('r') {
				l.ch = '\r'
			}
			if l.ch == rune('t') {
				l.ch = '\t'
			}
			if l.ch == rune('"') {
				l.ch = '"'
			}
			if l.ch == rune('\\') {
				l.ch = '\\'
			}
		}
		out = out + string(l.ch)
	}

	return out
}

// peek character looks at the next character which is available for consumption
func (l *Tokenizer) peekChar() rune {
	if l.readPosition >= len(l.characters) {
		return rune(0)
	}
	return l.characters[l.readPosition]
}

// determinate ch is identifier or not
func isIdentifier(ch rune) bool {
	return !isWhitespace(ch) && !isBrace(ch) && !isOperator(ch) && !isComparison(ch) && !isCompound(ch) && !isBrace(ch) && !isParen(ch) && !isBracket(ch) && !isEmpty(ch) && (ch != rune('\n'))
}

// is white space: note that a newline is NOT considered whitespace
// as we need that in our evaluator.
func isWhitespace(ch rune) bool {
	return ch == rune(' ') || ch == rune('\t') || ch == rune('\r')
}

// is operators
func isOperator(ch rune) bool {
	return ch == rune('+') || ch == rune('-') || ch == rune('/') || ch == rune('*')
}

// is comparison
func isComparison(ch rune) bool {
	return ch == rune('=') || ch == rune('!') || ch == rune('>') || ch == rune('<')
}

// is compound
func isCompound(ch rune) bool {
	return ch == rune(',') || ch == rune(':') || ch == rune('"') || ch == rune(';')
}

// is brace
func isBrace(ch rune) bool {
	return ch == rune('{') || ch == rune('}')
}

// is bracket
func isBracket(ch rune) bool {
	return ch == rune('[') || ch == rune(']')
}

// is parenthesis
func isParen(ch rune) bool {
	return ch == rune('(') || ch == rune(')')
}

// is empty
func isEmpty(ch rune) bool {
	return rune(0) == ch
}

// is Digit
func isDigit(ch rune) bool {
	return rune('0') <= ch && ch <= rune('9')
}
