package parser

import (
	"fmt"
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
		t.Fatalf("program.Statements expected to contain 3 statements. but got: %d instead", len(program.Statements))
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
		t.Fatalf("program.Statements expected to contain 3 statements. but got: %d instead", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt is not of a valid *ast.ReturnStatement type. got: %T instead", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral is not 'return', got: %T instead", returnStmt.TokenLiteral())
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
		t.Fatalf("program.Statements expected to contain 1 statement. but got: %d instead", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not a valid *ast.ExpressionStatement type. got: %T instead", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not a valid *ast.Identifier type. got: %T instead", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s. got: %s instead", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not %s. got: %s instead", "foobar", ident.TokenLiteral())
	}

}

func TestParseIntegerLiteralExpressions(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statement expected to contain 1 statement. but got: %d instead", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not a valid *ast.ExpressionStatement type. got: %T instead", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression is not a valid *ast.Identifier type. got: %T instead", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got: %d instead", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral() not %s. got: %s instead", "foobar", literal.TokenLiteral())
	}

}

func TestParsePrefixExpressions(t *testing.T) {
	tests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"+23;", "+", 23},
		{"-15;", "-", 15},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {

			t.Fatalf("program.Statements expected to contain 1 statement. but got: %d instead", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not a valid *ast.ExpressionStatement type. got: %T instead", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expression is not a valid *ast.PrefixExpression type. got: %T instead", stmt.Expression)
		}

		if exp.Operator != tc.operator {
			t.Fatalf("exp.Operator is not '%s'. got: %s instead", tc.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tc.integerValue) {
			return
		}
	}
}

func TestParseInfixExpressions(t *testing.T) {
	tests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{
			input:      "2 + 1;",
			leftValue:  2,
			operator:   "+",
			rightValue: 1,
		},
		{
			input:      "6 - 4;",
			leftValue:  6,
			operator:   "-",
			rightValue: 4,
		},
		{
			input:      "9 * 6;",
			leftValue:  9,
			operator:   "*",
			rightValue: 6,
		},
		{
			input:      "9 / 6;",
			leftValue:  9,
			operator:   "/",
			rightValue: 6,
		},
		{
			input:      "9 > 6;",
			leftValue:  9,
			operator:   ">",
			rightValue: 6,
		},
		{
			input:      "9 < 6;",
			leftValue:  9,
			operator:   "<",
			rightValue: 6,
		},
		{
			input:      "9 * 6;",
			leftValue:  9,
			operator:   "*",
			rightValue: 6,
		},
		{
			input:      "9 == 6;",
			leftValue:  9,
			operator:   "==",
			rightValue: 6,
		},
		{
			input:      "9 != 6;",
			leftValue:  9,
			operator:   "!=",
			rightValue: 6,
		},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements expected to contain 1 statement. but got: %d instead", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not a valid *ast.ExpressionStatement type. got: %T instead", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("expression is not a valid *ast.InfixExpression type. got: %T instead", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tc.leftValue) {
			return
		}

		if exp.Operator != tc.operator {
			t.Errorf("exp.operator is not '%s', got: '%s' instead", tc.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tc.rightValue) {
			return
		}
	}

}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "-a * b;",
			expected: "((-a) * b)",
		},
		{
			input:    "!-a;",
			expected: "(!(-a))",
		},
		{
			input:    "1 + 2 + 3;",
			expected: "((1 + 2) + 3)",
		},
		{
			input:    "a + b - c;",
			expected: "((a + b) - c)",
		},
		{
			input:    "a * b * c;",
			expected: "((a * b) * c)",
		},
		{
			input:    "a * b / c;",
			expected: "((a * b) / c)",
		},
		{
			input:    "a + b / c;",
			expected: "(a + (b / c))",
		},
		{
			input:    "a + b * c + d / e - f;",
			expected: "(((a + (b * c)) + (d / e)) - f)",
		},
		{
			input:    "3 + 4; -5 * 5;",
			expected: "(3 + 4)((-5) * 5)",
		},
		{
			input:    "5 > 4 == 3 < 4;",
			expected: "((5 > 4) == (3 < 4))",
		}, {
			input:    "5 < 4 != 3 > 4;",
			expected: "((5 < 4) != (3 > 4))",
		},
		{
			input:    "3 + 4 * 5 == 3 * 1 + 4 * 5;",
			expected: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		}, {
			input:    "3 + 4 * 5 == 3 * 1 + 4 * 5;",
			expected: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tc := range tests {
		l := lexer.New(tc.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)

		output := program.String()

		if tc.expected != output {
			t.Errorf("expected %s, got %s instead", tc.expected, output)
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() is not a 'let'. got: %q instead", s.TokenLiteral())

		return false
	}

	letSmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not a valid *ast.LetStatement type. got: %T instead", s)

		return false
	}

	if letSmt.Name.Value != name {
		t.Errorf("letSmt.Name.Value is not '%s'. got: %s instead", name, letSmt.Name.Value)

		return false
	}

	if letSmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name is not '%s'. got: %s instead", name, letSmt.Name.Value)

		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	intLit, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not *ast.IntegerLiteral. got: %T instead", il)

		return false
	}

	if intLit.Value != value {
		t.Errorf("intV.Value is not %d. got: %d instead", value, intLit.Value)

		return false
	}

	if intLit.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("intV.TokenLiteral() is not %d. got: %s instead", value, intLit.TokenLiteral())

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
