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

// TokenLiteral returns a token literal value of the token.
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

// TokenLiteral returns a token literal value of the token.
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

// ReturnStatement represents a return statement in a program including the expression part.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral returns a token literal value of the token.
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

// TokenLiteral returns a token literal value of the token.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String returns the string representation of the ExpressionStatement type.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (es *ExpressionStatement) statementNode() {}

// Identifier represents identifier expressions in a program. ie ("foobar;").
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral returns a token literal value of the token.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String returns the Identifier.Value.
func (i *Identifier) String() string { return i.Value }

func (i *Identifier) expressionNode() {}

// IntegerLiteral represents a single int literal token. ie (5; or "5;").
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral returns a token literal value of the token.
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// String returns a string representation of the IntegerLiteral type.
func (il *IntegerLiteral) String() string { return il.Token.Literal }

func (il *IntegerLiteral) expressionNode() {}

// PrefixExpression represents an expression such as (!4, +3, -19).
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// TokenLiteral returns a token literal value of the token.
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String returns a string representation of the PrefixExpression type.
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) expressionNode() {}

// InfixExpression represents an expression such as (4 * 19).
type InfixExpression struct {
	Token    token.Token // Refers to Operator token such as (*, +, >)
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral returns a token literal value of the token.
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String returns a string representation of the InfixExpression type.
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) expressionNode() {}
