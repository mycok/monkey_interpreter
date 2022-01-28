package ast

import (
	"github.com/mycok/monkey_interpreter/token"
)

// Node interface is implemented by node types. 
type Node interface {
	TokenLiteral() string
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

// Program type serves as the root node of every AST our parser produces.
// Every valid monkey program is a series / collection of valid statements / statement nodes.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement type represents an entire let statement in a program including the expression part.
type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

// TokenLiteral returns a string literal value of the token.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode() {}

// Identifier type represents identifier expressions in a program.
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral returns a string literal value of the token.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode() {}

