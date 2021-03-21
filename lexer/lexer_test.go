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
