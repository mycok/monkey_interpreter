package parser

import (
	"testing"

	"github.com/mycok/monkey_interpreter/ast"
	"github.com/mycok/monkey_interpreter/lexer"
)

func TestParseLetStatements(t *testing.T) {
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

func TestParseLetStatementsErrors(t *testing.T) {
	input := `
	let 67;
	let me 87;
	let = 4;
	let z == 10;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	// Checks for parse errors and stops the test execution by marking the entire test
	// function as Failed. This stops execution at this point in the function.
	checkParserErrors(t, p)
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.errors
	if len(errors) == 0 {
		return
	}

	// Indicate to the user that the parser encountered errors while parsing
	// the provided tokens.
	t.Errorf("parser had %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
