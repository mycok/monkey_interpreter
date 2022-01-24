package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// User defined identifiers + literals
	IDENT = "IDENT" // add, foobar, x,y, ....etc
	INT = "INT" // 1234567890

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"	
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

func LookupIndentifier(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}

	return IDENT
}

