package lexer

import "github.com/hugoleodev/hug/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tk = newToken(token.ASSIGN, l.ch)
	case ';':
		tk = newToken(token.SEMICOLON, l.ch)
	case ',':
		tk = newToken(token.COMMA, l.ch)
	case '+':
		tk = newToken(token.PLUS, l.ch)
	case '-':
		tk = newToken(token.MINUS, l.ch)
	case '!':
		tk = newToken(token.BANG, l.ch)
	case '/':
		tk = newToken(token.SLASH, l.ch)
	case '*':
		tk = newToken(token.ASTERISK, l.ch)
	case '<':
		tk = newToken(token.LT, l.ch)
	case '>':
		tk = newToken(token.GT, l.ch)
	case '{':
		tk = newToken(token.LBRACE, l.ch)
	case '}':
		tk = newToken(token.RBRACE, l.ch)
	case '(':
		tk = newToken(token.LPAREN, l.ch)
	case ')':
		tk = newToken(token.RPAREN, l.ch)
	case 0:
		tk.Literal = ""
		tk.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tk.Literal = l.readIdentifier()
			tk.Type = token.LookupIdent(tk.Literal)
			return tk
		} else if isDigit(l.ch) {
			tk.Literal = l.readNumber()
			tk.Type = token.INT
			return tk
		} else {
			tk = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tk
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
