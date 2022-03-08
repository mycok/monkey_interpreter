package ast

import (
	"bytes"

	"github.com/mycok/monkey_interpreter/token"
)

// Node interface is implemented by node types.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement interface is implemented by statement types.
type Statement interface {
	Node
	statementNode()
}

// Expression interface is implemented by expression types.
type Expression interface {
	Node
	expressionNode()
}

// Program serves as the root node of every AST our parser produces.
// Every valid monkey program is a series / collection of valid statements / statement nodes.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns a string literal value of the statement.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// String returns the string representation of p.statements slice.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents an entire let statement in a program including the expression part.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral returns a string literal value of the token.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String returns the string representation of the LetStatement type.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() {}

// Identifier represents identifier expressions in a program.
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral returns a string literal value of the token.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String returns the Identifier.Value.
func (i *Identifier) String() string { return i.Value }

func (i *Identifier) expressionNode() {}

// ReturnStatement represents a return statement in a program including the expression part.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral returns a string literal value of the token.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String returns the string representation of the ReturnStatement type.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement represents an expression statement in a program. such as (x + 55).
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral returns a string literal value of the token.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String returns the string representation of the ExpressionStatement type.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (es *ExpressionStatement) statementNode() {}
