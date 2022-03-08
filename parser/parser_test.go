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
	checkParserErrors(t, p)

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

func TestParseReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 15;
	return add(5, 2);
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements expected to contain 3 statements. but got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not of a valid *ast.ReturnStatement type. got:%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral is not 'return', got:%T", returnStmt.TokenLiteral())
		}
	}
}

func TestParseIdentifierExpressions(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements expected to contain 1 statement. but got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not a valid *ast.ExpressionStatement type. got:%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not a valid *ast.Identifier type. got:%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got:%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not %s. got:%s", "foobar", ident.TokenLiteral())
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
		t.Errorf("letSmt.Name.Value is not '%s'. got:%s", name, letSmt.Name.Value)

		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name is not '%s'. got:%s", name, letSmt.Name.Value)

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
