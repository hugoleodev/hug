package lexer

import (
	"testing"

	"github.com/hugoleodev/hug/token"
)

func TestNextToken(t *testing.T) {
	t.Run("basic tokens support", func(t *testing.T) {
		input := `=+(){},;`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
		}

		l := New(input)

		for i, tt := range tests {
			tk := l.NextToken()

			assertNthTokenType(t, i, tt.expectedType, tk.Type)
			assertNthTokenLiteral(t, i, tt.expectedLiteral, tk.Literal)
		}
	})

	t.Run("support to identifiers, keywwords and numbers", func(t *testing.T) {
		input := `let thirteen = 13;
		let nine = 9;
		
		let add = fn(x, y) {
			x + y;
		};

		let result = add(thirteen, nine);
		!-/*9;
		9 < 13 > 9;
		`

		tests := []struct {
			expectedType    token.TokenType
			expectedLiteral string
		}{
			{token.LET, "let"},
			{token.IDENT, "thirteen"},
			{token.ASSIGN, "="},
			{token.INT, "13"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "nine"},
			{token.ASSIGN, "="},
			{token.INT, "9"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "thirteen"},
			{token.COMMA, ","},
			{token.IDENT, "nine"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "9"},
			{token.SEMICOLON, ";"},
			{token.INT, "9"},
			{token.LT, "<"},
			{token.INT, "13"},
			{token.GT, ">"},
			{token.INT, "9"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		l := New(input)

		for i, tt := range tests {
			tk := l.NextToken()

			assertNthTokenType(t, i, tt.expectedType, tk.Type)
			assertNthTokenLiteral(t, i, tt.expectedLiteral, tk.Literal)
		}
	})

}

func assertNthTokenType(t *testing.T, nth int, expected, got token.TokenType) {
	t.Helper()

	if got != expected {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			nth, expected, got)
	}
}

func assertNthTokenLiteral(t *testing.T, nth int, expected, got string) {
	t.Helper()

	if got != expected {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			nth, expected, got)
	}
}
