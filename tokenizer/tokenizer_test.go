package tokenizer

import (
	"testing"

	"github.com/skx/gobasic/token"
)

// TestMathOperators enures we can recognise "mathematical" operators.
func TestMathOperators(t *testing.T) {
	input := `+-/*%=`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "N"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.MOD, "%"},
		{token.ASSIGN, "="},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestMiscTokens just tests the tokens we've not otherwise covered.
func TestMiscTokens(t *testing.T) {
	input := `(),:`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "N"},
		{token.LBRACKET, "("},
		{token.RBRACKET, ")"},
		{token.COMMA, ","},
		{token.COLON, ":"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestLineNot is a trivial test of line-number parsing.
func TestLineNo(t *testing.T) {
	input := `10 PRINT
20 PRINT`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "N"},
		{token.LINENO, "10"},
		{token.PRINT, "PRINT"},
		{token.NEWLINE, "N"},
		{token.LINENO, "20"},
		{token.PRINT, "PRINT"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestStringParse tests we can cope with control-characters inside strings.
func TestStringParse(t *testing.T) {
	input := `10 LET a="\n\r\t\\\""
20 PRINT`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "N"},
		{token.LINENO, "10"},
		{token.LET, "LET"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.STRING, "\n\r\t\\\""},
		{token.NEWLINE, "N"},
		{token.LINENO, "20"},
		{token.PRINT, "PRINT"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestComparisons tests that we parse the different comparison operators.
func TestComparisons(t *testing.T) {
	input := `10 IF A < B
20 IF A <= B
30 IF A > B
40 IF A >= B
50 IF A <> B
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "N"},

		{token.LINENO, "10"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.LT, "<"},
		{token.IDENT, "B"},
		{token.NEWLINE, "N"},

		{token.LINENO, "20"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.LT_EQUALS, "<="},
		{token.IDENT, "B"},
		{token.NEWLINE, "N"},

		{token.LINENO, "30"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.GT, ">"},
		{token.IDENT, "B"},
		{token.NEWLINE, "N"},

		{token.LINENO, "40"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.GT_EQUALS, ">="},
		{token.IDENT, "B"},
		{token.NEWLINE, "N"},

		{token.LINENO, "50"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.NOT_EQUALS, "<>"},
		{token.IDENT, "B"},
		{token.NEWLINE, "N"},

		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

// TestNumber tests that positive and negative numbers are OK.
func TestNumber(t *testing.T) {
	input := `10 PRINT -4
20 PRINT 5 - 3`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "N"},
		{token.LINENO, "10"},
		{token.PRINT, "PRINT"},
		{token.INT, "-4"},
		{token.NEWLINE, "N"},
		{token.LINENO, "20"},
		{token.PRINT, "PRINT"},
		{token.INT, "5"},
		{token.MINUS, "-"},
		{token.INT, "3"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%v", i, tt.expectedType, tok)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
