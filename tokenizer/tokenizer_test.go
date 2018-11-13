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
		{token.NEWLINE, "\\n"},
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
	input := `(),:;`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LBRACKET, "("},
		{token.RBRACKET, ")"},
		{token.COMMA, ","},
		{token.COLON, ":"},
		{token.SEMICOLON, ";"},
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

	// Attempt to read past the end of our input-stream
	for i := 0; i < 10; i++ {
		tok := l.NextToken()
		if tok.Type != token.EOF {
			t.Errorf("EOF wasn't hit properly")
		}
	}

}

// TestLineNot is a trivial test of line-number parsing.
func TestLineNo(t *testing.T) {
	input := `10 REM
20 REM`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LINENO, "10"},
		{token.REM, "REM"},
		{token.NEWLINE, "\\n"},
		{token.LINENO, "20"},
		{token.REM, "REM"},
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
20 REM OK`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LINENO, "10"},
		{token.LET, "LET"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.STRING, "\n\r\t\\\""},
		{token.NEWLINE, "\\n"},
		{token.LINENO, "20"},
		{token.REM, "REM"},
		{token.IDENT, "OK"},
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
		{token.NEWLINE, "\\n"},

		{token.LINENO, "10"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.LT, "<"},
		{token.IDENT, "B"},
		{token.NEWLINE, "\\n"},

		{token.LINENO, "20"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.LTEQUALS, "<="},
		{token.IDENT, "B"},
		{token.NEWLINE, "\\n"},

		{token.LINENO, "30"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.GT, ">"},
		{token.IDENT, "B"},
		{token.NEWLINE, "\\n"},

		{token.LINENO, "40"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.GTEQUALS, ">="},
		{token.IDENT, "B"},
		{token.NEWLINE, "\\n"},

		{token.LINENO, "50"},
		{token.IF, "IF"},
		{token.IDENT, "A"},
		{token.NOTEQUALS, "<>"},
		{token.IDENT, "B"},
		{token.NEWLINE, "\\n"},
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
	input := `10 REM -4.3
20 REM 5 - 3`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LINENO, "10"},
		{token.REM, "REM"},
		{token.INT, "-4.3"},
		{token.NEWLINE, "\\n"},
		{token.LINENO, "20"},
		{token.REM, "REM"},
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

// TestPow tests that we can parse "^".
func TestPow(t *testing.T) {
	input := `10 PRINT 2 ^ 3`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LINENO, "10"},
		{token.IDENT, "PRINT"},
		{token.INT, "2"},
		{token.POW, "^"},
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

// TestIdent ensures we parse a variable-name such as "A3" correctly.
func TestIdent(t *testing.T) {
	input := `10 LET a3 = 6`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LINENO, "10"},
		{token.LET, "LET"},
		{token.IDENT, "a3"},
		{token.ASSIGN, "="},
		{token.INT, "6"},
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

// Test null terminates a string
func TestNullString(t *testing.T) {
	input := "10 LET a = \"steve\000\n20 PRINT a"

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		// implicit newline which is a pain.
		{token.NEWLINE, "\\n"},
		{token.LINENO, "10"},
		{token.LET, "LET"},
		{token.IDENT, "a"},
		{token.ASSIGN, "="},
		{token.STRING, "steve"},
		{token.NEWLINE, "\\n"},
		{token.LINENO, "20"},
		{token.IDENT, "PRINT"},
		{token.IDENT, "a"},
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
