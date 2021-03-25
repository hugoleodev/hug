package parser

import (
	"testing"

	"github.com/hugoleodev/hug/ast"
	"github.com/hugoleodev/hug/lexer"
)

func TestLetStatements(t *testing.T) {

	t.Run("parse let statements", func(t *testing.T) {
		input := `
		let x = 13;
		let y = 9;
		let foobar = 131313;
		`

		l := lexer.New(input)
		p := New(l)

		program := p.ParseProgram()
		checkParseErrors(t, p)

		if program == nil {
			t.Fatalf("ParseProgram() returned nil")
		}

		if len(program.Statements) != 3 {
			t.Fatalf("program.Statements does not contain 3 statements. got=%d",
				len(program.Statements))
		}

		tests := []struct {
			expectedIdentifier string
		}{
			{"x"},
			{"y"},
			{"foobar"},
		}

		for i, tt := range tests {
			stmt := program.Statements[i]
			if !testLetStatement(t, stmt, tt.expectedIdentifier) {
				return
			}
		}
	})

	t.Run("check errors for bad let statement", func(t *testing.T) {
		input := `
		let x 13;
		let = 9;
		let 131313;
		`

		l := lexer.New(input)
		p := New(l)

		_ = p.ParseProgram()

		errors := p.errors

		if len(errors) != 3 {
			t.Errorf("p.errors does not contain 3 errors. got=%d", len(errors))
		}

		if len(p.Errors()) != len(errors) {
			t.Errorf("Mismatch between p.Errors() and errors messages")
		}
	})

}

func checkParseErrors(t *testing.T, p *Parser) {
	t.Helper()

	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parse error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got='%q'", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true

}
