package lexer

import (
	"testing"

	"github.com/hugoleodev/hug/token"
)

func TestNextToken(t *testing.T) {
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

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("testes[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tk.Literal)
		}
	}
}
