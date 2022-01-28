package parser

import (
	"testing"

	"github.com/mycok/monkey_interpreter/lexer"
	"github.com/mycok/monkey_interpreter/ast"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements expected to contain 3 statements. but got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tc := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tc.expectedIdentifier) {
			return
		}

	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() is not a 'let'. got:%q", s.TokenLiteral())

		return false
	}

	letSmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not a valid *ast.LetStatement type. got:%T", s)

		return false
	}

	if letSmt.Name.Value != name {
		t.Errorf("letSmt.Name.Value not '%s'. got:%s", name, letSmt.Name.Value)

		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got:%s", name, letSmt.Name.Value)

		return false
	}

	return true
}